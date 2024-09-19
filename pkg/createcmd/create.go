package createcmd

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/gugahoi/basiq/internal/api"
	"github.com/urfave/cli/v2"
)

// create creates a webhook.
// https://api.basiq.io/reference/addwebhook
func New() *cli.Command {
	return &cli.Command{
		Name:  "create",
		Usage: "create a webhook",
		Action: func(ctx *cli.Context) error {
			client := ctx.App.Metadata["client"].(*api.ClientWithResponses)
			return exec(client, ctx.Args().Tail())
		},
	}
}

func exec(c *api.ClientWithResponses, args []string) error {
	description := args[0]
	url := args[1]
	name := args[2]
	events := args[3]

	subscribedEvents := strings.Split(events, ",")

	payload := api.WebhookBody{
		Description:      &description,
		Url:              url,
		Name:             &name,
		SubscribedEvents: &subscribedEvents,
	}

	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to serialize webhook: %w", err)
	}
	response, err := c.AddWebhookWithBodyWithResponse(context.Background(), "application/json", bytes.NewReader(jsonPayload))
	if err != nil {
		return fmt.Errorf("failed to create webhook: %w", err)
	}
	if response.StatusCode() != 201 {
		return fmt.Errorf("failed to create webhook: [%d] %s", response.StatusCode(), string(response.Body))
	}
	log.Printf("%v\n", response.JSON201.Id)

	return nil
}
