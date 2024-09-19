package listcmd

import (
	"context"
	"log"

	"github.com/gugahoi/basiq/internal/api"
	"github.com/mitchellh/mapstructure"
)

type ListWebhooksResponse struct {
	Id               string
	Name             *string
	Description      *string
	Url              string
	Status           api.WebhookStatus
	SubscribedEvents *[]string
	Type             string
}

// list lists all webhooks.
// https://api.basiq.io/reference/listappwebhooks
func List(c *api.ClientWithResponses) {
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
