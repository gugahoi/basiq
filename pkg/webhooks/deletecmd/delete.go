package deletecmd

import (
	"context"
	"fmt"

	"github.com/gugahoi/basiq/internal/api"
	"github.com/gugahoi/basiq/tools"
	"github.com/urfave/cli/v2"
)

func New() *cli.Command {
	return &cli.Command{
		Name:    "delete",
		Usage:   "delete a webhook by ID",
		Aliases: []string{"rm"},
		Action: func(ctx *cli.Context) error {
			client := tools.GetClient(ctx)
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
