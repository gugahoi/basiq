package events

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

// HttpRequestDoer performs HTTP requests.
//
// The standard http.Client implements this interface.
type HttpRequestDoer interface {
	Do(req *http.Request) (*http.Response, error)
}

// RequestEditorFn  is the function signature for the RequestEditor callback function
type RequestEditorFn func(ctx context.Context, req *http.Request) error

// WithRequestEditorFn allows setting up a callback function, which will be
// called right before sending the request. This can be used to mutate the request.
func WithRequestEditorFn(fn RequestEditorFn) ClientOption {
	return func(c *Client) error {
		c.RequestEditors = append(c.RequestEditors, fn)
		return nil
	}
}

// Client which conforms to the OpenAPI3 specification for this service.
type Client struct {
	// The endpoint of the server conforming to this interface, with scheme,
	// https://api.deepmap.com for example. This can contain a path relative
	// to the server, such as https://api.deepmap.com/dev-test, and all the
	// paths in the swagger spec will be appended to the server.
	Server string

	// Doer for performing requests, typically a *http.Client with any
	// customized settings, such as certificate chains.
	Client HttpRequestDoer

	// A list of callbacks for modifying requests which are generated before sending over
	// the network.
	RequestEditors []RequestEditorFn
}

// ClientOption allows setting custom parameters during construction
type ClientOption func(*Client) error

func NewClient(server string, opts ...ClientOption) (*Client, error) {
	client := Client{
		Server: server,
	}

	// mutate client and add all optional params
	for _, o := range opts {
		if err := o(&client); err != nil {
			return nil, err
		}
	}

	if !strings.HasSuffix(client.Server, "/") {
		client.Server += "/"
	}

	if client.Client == nil {
		client.Client = &http.Client{}
	}

	return &client, nil
}

type EventType struct {
	Type        string
	Id          string
	Description string
	Links       struct {
		Self string
	}
}

type ListAllTypesResponse struct {
	Data  []EventType
	Count int
	Size  int
	Type  string
	Links struct {
		Self string
		Next string
	}
}

func (c *Client) applyEditors(ctx context.Context, req *http.Request) error {
	for _, r := range c.RequestEditors {
		if err := r(ctx, req); err != nil {
			return err
		}
	}
	return nil
}

func (c *Client) buildUrl(path string) (string, error) {
	var err error

	serverURL, err := url.Parse(c.Server)
	if err != nil {
		return "", err
	}

	if path[0] == '/' {
		path = "." + path
	}

	queryURL, err := serverURL.Parse(path)
	if err != nil {
		return "", err
	}

	return queryURL.String(), nil
}

func (c *Client) createRequest(ctx context.Context, method string, url string) (*http.Request, error) {
	url, err := c.buildUrl(url)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req); err != nil {
		return nil, err
	}
	return req, nil
}

func (c *Client) ListAllTypes(ctx context.Context) (*ListAllTypesResponse, error) {
	req, err := c.createRequest(ctx, "GET", "/events/types")
	if err != nil {
		return nil, err
	}

	res, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != 200 {
		return nil, fmt.Errorf("failed to list events: [%d] %s", res.StatusCode, res.Status)
	}

	body, _ := io.ReadAll(res.Body)

	var result ListAllTypesResponse
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to list events: %w", err)
	}

	return &result, nil
}

type GetTypeResponse = EventType

func (c *Client) GetType(ctx context.Context, id string) (*GetTypeResponse, error) {
	req, err := c.createRequest(ctx, "GET", fmt.Sprintf("/events/types/%s", id))
	if err != nil {
		return nil, err
	}

	res, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != 200 {
		return nil, fmt.Errorf("failed to get event type: [%d] %s", res.StatusCode, res.Status)
	}

	body, _ := io.ReadAll(res.Body)

	var result GetTypeResponse
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to get event: %w", err)
	}

	return &result, nil
}

type Event struct {
	CreatedDate string
	Data        json.RawMessage
	DataRef     string
	Entity      string
	EventType   string
	Id          string
	Links       struct{ Self string }
	Type        string
	UserId      string
}

type ListAllResponse struct {
	Data  []Event
	Type  string
	Links struct {
		Self string
	}
}

type ListAllFilters struct {
	UserId *string
	Entity *string
	Type   *string
}

// ListAll lists all events
// Filters: userId, entity, type
func (c *Client) ListAll(ctx context.Context, filters ListAllFilters) (*ListAllResponse, error) {
	queryStr := []string{}

	if filters.UserId != nil {
		queryStr = append(queryStr, fmt.Sprintf("user.id.eq('%s')", *filters.UserId))
	}
	if filters.Entity != nil {
		queryStr = append(queryStr, fmt.Sprintf("event.entity.eq('%s')", *filters.Entity))
	}
	if filters.Type != nil {
		queryStr = append(queryStr, fmt.Sprintf("event.type.eq('%s')", *filters.Type))
	}

	query := url.Values{}
	if len(queryStr) > 0 {
		query.Set("filter", strings.Join(queryStr, ","))
	}

	url := fmt.Sprintf("/events?%s", query.Encode())
	req, err := c.createRequest(ctx, "GET", url)
	if err != nil {
		return nil, err
	}

	res, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	if res.StatusCode != 200 {
		return nil, fmt.Errorf("failed to get events: [%d] %s", res.StatusCode, string(body))
	}

	var result ListAllResponse
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to parse events: %w", err)
	}

	return &result, nil
}
