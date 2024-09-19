package main

import (
	"bytes"
	"context"
	"encoding/json"
	"flag"
	"log"
	"os"
	"strings"

	basiq "github.com/gugahoi/basiq/internal/api"
	"github.com/mitchellh/mapstructure"
	"github.com/oapi-codegen/oapi-codegen/v2/pkg/securityprovider"
)

func createClient(apikey string) *basiq.ClientWithResponses {
	token := getAuthToken(apikey)
	auth, err := securityprovider.NewSecurityProviderBearerToken(token)

	client, err := basiq.NewClientWithResponses(ServerURL, basiq.WithRequestEditorFn(auth.Intercept))
	if err != nil {
		log.Fatalln("failed to generate Basiq client", err)
	}
	return client

}

func main() {
	// remove timestamp from log lines
	log.SetFlags(0)
	log.SetOutput(os.Stdout)

	var apikey string
	flag.StringVar(&apikey, "apikey", "", "Basiq API key")
	flag.Parse()

	if apikey == "" {
		log.Fatalln("apikey is required")
	}

	c := createClient(apikey)

	args := flag.Args()
	cmd := args[0]

	switch cmd {
	case "list":
		list(c)
	case "delete":
		delete(c, args[2])
	case "create":
		create(c, args[2:])
	case "get":
		get(c, args[2])
	default:
		usage()
	}
}

func usage() {
	log.Fatalln("available commands: list, get")
}

// get retrieves a webhook.
// https://api.basiq.io/reference/getwebhook
func get(c *basiq.ClientWithResponses, webhookID string) {
	webhook, err := c.GetWebhookWithResponse(context.Background(), webhookID)
	if err != nil {
		log.Fatalln("failed to get webhook", err)
	}
	if webhook.StatusCode() != 200 {
		log.Fatalln("failed to get webhook", webhook.StatusCode(), string(webhook.Body))
	}
	log.Printf("%v\n", webhook.JSON200.Description)
}

type ListWebhooksResponse struct {
	Id               string
	Name             *string
	Description      *string
	Url              string
	Status           basiq.WebhookStatus
	SubscribedEvents *[]string
	Type             string
}

// list lists all webhooks.
// https://api.basiq.io/reference/listappwebhooks
func list(c *basiq.ClientWithResponses) {
	webhooks, err := c.ListAppWebhooksWithResponse(context.Background())
	if err != nil {
		log.Fatalln("failed to list webhooks", err)
	}
	if webhooks.StatusCode() != 200 {
		log.Fatalln("failed to list webhooks", webhooks.StatusCode(), string(webhooks.Body))
	}

	var result []ListWebhooksResponse
	err = mapstructure.Decode(*webhooks.JSON200.Data, &result)
	if err != nil {
		log.Fatalln("failed to decode webhook list", err)
	}

	for _, webhook := range result {
		log.Printf("%s\t%s\t%s\t%s\n", webhook.Id, *webhook.Name, *webhook.Description, webhook.Url)
	}

}

// delete deletes a webhook.
// https://api.basiq.io/reference/deletewebhook
func delete(c *basiq.ClientWithResponses, webhookID string) {
	_, err := c.DeleteWebhook(context.Background(), webhookID)
	if err != nil {
		log.Fatalln("failed to delete webhook", err)
	}
}

// create creates a webhook.
// https://api.basiq.io/reference/addwebhook
func create(c *basiq.ClientWithResponses, args []string) {
	description := args[0]
	url := args[1]
	name := args[2]
	events := args[3]

	subscribedEvents := strings.Split(events, ",")

	webhook := basiq.WebhookBody{
		Description:      &description,
		Url:              url,
		Name:             &name,
		SubscribedEvents: &subscribedEvents,
	}

	webhookJSON, err := json.Marshal(webhook)
	if err != nil {
		log.Fatalln("failed to serialize webhook", err)
	}
	response, err := c.AddWebhookWithBodyWithResponse(context.Background(), "application/json", bytes.NewReader(webhookJSON))
	if err != nil {
		log.Fatalln("failed to create webhook", err)
	}
	if response.StatusCode() != 201 {
		log.Fatalln("failed to create webhook", response.StatusCode(), string(response.Body))
	}
	log.Printf("%v\n", response.JSON201.Id)
}
