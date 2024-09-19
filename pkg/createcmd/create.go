package createcmd

import (
	"bytes"
	"context"
	"encoding/json"
	"log"
	"strings"

	"github.com/gugahoi/basiq/internal/api"
)

// create creates a webhook.
// https://api.basiq.io/reference/addwebhook
func Create(c *api.ClientWithResponses, args []string) {
	description := args[0]
	url := args[1]
	name := args[2]
	events := args[3]

	subscribedEvents := strings.Split(events, ",")

	webhook := api.WebhookBody{
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
