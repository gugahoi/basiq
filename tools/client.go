package tools

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/gugahoi/basiq/internal/api"
	"github.com/gugahoi/basiq/internal/api/events"
	"github.com/oapi-codegen/oapi-codegen/v2/pkg/securityprovider"
)

const ServerURL = "https://au-api.basiq.io/"

// CreateClient creates a client.
func CreateClient(apikey string) *api.ClientWithResponses {
	token := getAuthToken(apikey)
	auth, err := securityprovider.NewSecurityProviderBearerToken(token)

	client, err := api.NewClientWithResponses(ServerURL, api.WithRequestEditorFn(auth.Intercept))
	if err != nil {
		log.Fatalln("failed to generate Basiq client", err)
	}
	return client
}

// CreateEventsClient creates an events client.
func CreateEventsClient(apikey string) *events.Client {
	token := getAuthToken(apikey)
	auth, err := securityprovider.NewSecurityProviderBearerToken(token)

	client, err := events.NewClient(ServerURL, events.WithRequestEditorFn(auth.Intercept))
	if err != nil {
		log.Fatalln("failed to generate Basiq client", err)
	}
	return client
}

type AuthResponse struct {
	AccessToken string `json:"access_token"`
}

func getAuthToken(apikey string) string {
	payload := strings.NewReader("scope=SERVER_ACCESS")
	req, _ := http.NewRequest("POST", ServerURL+"token", payload)

	req.Header.Add("accept", "application/json")
	req.Header.Add("authorization", "Basic "+apikey)
	req.Header.Add("basiq-version", "3.0")
	req.Header.Add("content-type", "application/x-www-form-urlencoded")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalln("unable to authenticate", err)
	}

	defer res.Body.Close()
	raw, _ := io.ReadAll(res.Body)
	body := string(raw)

	if res.StatusCode != 200 {
		log.Fatalln("unable to authenticate", body)
	}

	var authResponse AuthResponse
	err = json.Unmarshal(raw, &authResponse)
	if err != nil {
		log.Fatalln("unable to parse access_token", body)
	}

	return authResponse.AccessToken
}
