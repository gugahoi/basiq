package deletecmd

import (
	"context"
	"log"

	"github.com/gugahoi/basiq/internal/api"
)

// delete deletes a webhook.
// https://api.basiq.io/reference/deletewebhook
func Delete(c *api.ClientWithResponses, webhookID string) {
	_, err := c.DeleteWebhook(context.Background(), webhookID)
	if err != nil {
		log.Fatalln("failed to delete webhook", err)
	}
}
