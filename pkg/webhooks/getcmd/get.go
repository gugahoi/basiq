package getcmd

import (
	"context"
	"fmt"
	"log"

	"github.com/gugahoi/basiq/internal/api"
	"github.com/urfave/cli/v2"
)

func New() *cli.Command {
	return &cli.Command{
		Name:  "get",
		Usage: "retrieve a webhook by ID",
		Action: func(ctx *cli.Context) error {
			client := ctx.App.Metadata["client"].(*api.ClientWithResponses)
			return exec(client, ctx.Args().First())
		},
	}
}

// get retrieves a webhook.
// https://api.basiq.io/reference/getwebhook
func exec(c *api.ClientWithResponses, webhookID string) error {
	webhook, err := c.GetWebhookWithResponse(context.Background(), webhookID)
	if err != nil {
		return fmt.Errorf("failed to get webhook: %w", err)
	}
	if webhook.StatusCode() != 200 {
		return fmt.Errorf("failed to get webhook: [%d] %s", webhook.StatusCode(), string(webhook.Body))
	}
	log.Printf("%v\n", webhook.JSON200.Description)

	return nil
}
