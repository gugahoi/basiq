package getcmd

import (
	"context"
	"log"

	"github.com/gugahoi/basiq/internal/api"
)

// get retrieves a webhook.
// https://api.basiq.io/reference/getwebhook
func Get(c *api.ClientWithResponses, webhookID string) {
	webhook, err := c.GetWebhookWithResponse(context.Background(), webhookID)
	if err != nil {
		log.Fatalln("failed to get webhook", err)
	}
	if webhook.StatusCode() != 200 {
		log.Fatalln("failed to get webhook", webhook.StatusCode(), string(webhook.Body))
	}
	log.Printf("%v\n", webhook.JSON200.Description)
}
