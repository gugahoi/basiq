package users

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

// HttpRequestDoer performs HTTP requests.
//
// The standard http.Client implements this interface.
type HttpRequestDoer interface {
	Do(req *http.Request) (*http.Response, error)
}

// RequestEditorFn is the function signature for the RequestEditor callback function.
type RequestEditorFn func(ctx context.Context, req *http.Request) error

// ClientOption allows setting custom parameters during construction.
type ClientOption func(*Client) error

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
	// Server is the endpoint of the server conforming to this interface, with scheme,
	// https://au-api.basiq.io for example.
	Server string

	// Client is used for performing requests, typically a *http.Client with any
	// customised settings, such as certificate chains.
	Client HttpRequestDoer

	// RequestEditors is a list of callbacks for modifying requests before sending over the network.
	RequestEditors []RequestEditorFn
}

// NewClient creates a new Client with the given server URL and options.
func NewClient(server string, opts ...ClientOption) (*Client, error) {
	client := Client{Server: server}

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

// User represents a Basiq user.
type User struct {
	Id                 string  `json:"id"`
	Email              string  `json:"email"`
	Mobile             string  `json:"mobile"`
	FirstName          string  `json:"firstName"`
	LastName           string  `json:"lastName"`
	BusinessName       string  `json:"businessName"`
	BusinessIdNo       string  `json:"businessIdNo"`
	BusinessIdNoType   string  `json:"businessIdNoType"`
	VerificationStatus bool    `json:"verificationStatus"`
	VerificationDate   string  `json:"verificationDate"`
}

// UpdateUserBody contains the fields that can be updated on a user.
type UpdateUserBody struct {
	Email              *string `json:"email,omitempty"`
	Mobile             *string `json:"mobile,omitempty"`
	FirstName          *string `json:"firstName,omitempty"`
	LastName           *string `json:"lastName,omitempty"`
	BusinessName       *string `json:"businessName,omitempty"`
	BusinessIdNo       *string `json:"businessIdNo,omitempty"`
	BusinessIdNoType   *string `json:"businessIdNoType,omitempty"`
	VerificationStatus *bool   `json:"verificationStatus,omitempty"`
	VerificationDate   *string `json:"verificationDate,omitempty"`
}

func (c *Client) applyEditors(ctx context.Context, req *http.Request) error {
	for _, r := range c.RequestEditors {
		if err := r(ctx, req); err != nil {
			return err
		}
	}
	return nil
}

func (c *Client) newRequest(ctx context.Context, method, path string, body io.Reader) (*http.Request, error) {
	url := c.Server + strings.TrimPrefix(path, "/")

	req, err := http.NewRequestWithContext(ctx, method, url, body)
	if err != nil {
		return nil, err
	}

	req.Header.Set("accept", "application/json")
	req.Header.Set("content-type", "application/json")

	if err := c.applyEditors(ctx, req); err != nil {
		return nil, err
	}

	return req, nil
}

// GetUser retrieves a user by their ID.
// https://api.basiq.io/reference/getuser
func (c *Client) GetUser(ctx context.Context, userID string) (*User, error) {
	req, err := c.newRequest(ctx, "GET", fmt.Sprintf("/users/%s", userID), nil)
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
		return nil, fmt.Errorf("failed to get user: [%d] %s", res.StatusCode, string(body))
	}

	var user User
	if err := json.Unmarshal(body, &user); err != nil {
		return nil, fmt.Errorf("failed to parse user: %w", err)
	}

	return &user, nil
}

// UpdateUser updates a user by their ID.
// https://api.basiq.io/reference/updateuser
func (c *Client) UpdateUser(ctx context.Context, userID string, payload *UpdateUserBody) (*User, error) {
	jsonBody, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("failed to serialize user: %w", err)
	}

	req, err := c.newRequest(ctx, "POST", fmt.Sprintf("/users/%s", userID), bytes.NewReader(jsonBody))
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
		return nil, fmt.Errorf("failed to update user: [%d] %s", res.StatusCode, string(body))
	}

	var user User
	if err := json.Unmarshal(body, &user); err != nil {
		return nil, fmt.Errorf("failed to parse user: %w", err)
	}

	return &user, nil
}

// DeleteUser permanently deletes a user and all associated data.
// https://api.basiq.io/reference/deleteuser
func (c *Client) DeleteUser(ctx context.Context, userID string) error {
	req, err := c.newRequest(ctx, "DELETE", fmt.Sprintf("/users/%s", userID), nil)
	if err != nil {
		return err
	}

	res, err := c.Client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != 204 {
		body, _ := io.ReadAll(res.Body)
		return fmt.Errorf("failed to delete user: [%d] %s", res.StatusCode, string(body))
	}

	return nil
}
