package listcmd

import (
	"context"
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/gugahoi/basiq/internal/api"
	"github.com/mitchellh/mapstructure"
	"github.com/urfave/cli/v2"
)

type Webhook struct {
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
func New() *cli.Command {
	return &cli.Command{
		Name:        "list",
		Aliases:     []string{"ls"},
		Usage:       "list all webhooks",
		Description: "`list` lists all webhooks",
		Action: func(ctx *cli.Context) error {
			client := ctx.App.Metadata["client"].(*api.ClientWithResponses)
			return exec(client)
		},
	}
}

func exec(c *api.ClientWithResponses) error {
	webhooks, err := c.ListAppWebhooksWithResponse(context.Background())
	if err != nil {
		return fmt.Errorf("failed to list webhooks: %w", err)
	}
	if webhooks.StatusCode() != 200 {
		return fmt.Errorf("failed to list webhooks: [%d] %s", webhooks.StatusCode(), string(webhooks.Body))
	}

	var result []Webhook
	err = mapstructure.Decode(*webhooks.JSON200.Data, &result)
	if err != nil {
		return fmt.Errorf("failed to decode webhook list: %w", err)
	}

	w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', 0)
	for _, webhook := range result {
		fmt.Fprintf(w, "%s\t%s\t%s\t%s\n", webhook.Id, *webhook.Name, *webhook.Description, webhook.Url)
	}
	w.Flush()
	return nil
}
