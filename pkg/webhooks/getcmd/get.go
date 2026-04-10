package getcmd

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"

	"github.com/gugahoi/basiq/internal/api"
	"github.com/gugahoi/basiq/tools"
	"github.com/mitchellh/mapstructure"
	"github.com/urfave/cli/v3"
)

func New() *cli.Command {
	return &cli.Command{
		Name:  "get",
		Usage: "retrieve a webhook by ID",
		ShellComplete: func(ctx context.Context, cmd *cli.Command) {
			client := tools.GetClient(cmd)
			webhooks, err := client.ListAppWebhooksWithResponse(ctx)
			if err != nil || webhooks.StatusCode() != 200 || webhooks.JSON200 == nil {
				return
			}
			var result []struct {
				Id   string
				Name *string
			}
			_ = mapstructure.Decode(*webhooks.JSON200.Data, &result)
			for _, w := range result {
				if w.Name != nil {
					fmt.Printf("%s:%s\n", w.Id, *w.Name)
				} else {
					fmt.Println(w.Id)
				}
			}
		},
		Action: func(ctx context.Context, cmd *cli.Command) error {
			client := tools.GetClient(cmd)
			return exec(client, cmd.Args().First())
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
	var prettyJSON bytes.Buffer
	err = json.Indent(&prettyJSON, webhook.Body, "", "\t")
	fmt.Printf("%s\n", string(prettyJSON.Bytes()))

	return nil
}
