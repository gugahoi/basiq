package deletecmd

import (
	"context"
	"fmt"

	"github.com/gugahoi/basiq/internal/api"
	"github.com/urfave/cli/v2"
)

func New() *cli.Command {
	return &cli.Command{
		Name:    "delete",
		Aliases: []string{"rm"},
		Action: func(ctx *cli.Context) error {
			client := ctx.App.Metadata["client"].(*api.ClientWithResponses)
			return exec(client, ctx.Args().First())
		},
	}
}

// delete deletes a webhook.
// https://api.basiq.io/reference/deletewebhook
func exec(c *api.ClientWithResponses, webhookID string) error {
	_, err := c.DeleteWebhook(context.Background(), webhookID)
	if err != nil {
		return fmt.Errorf("failed to delete webhook: %w", err)
	}
	return nil
}
