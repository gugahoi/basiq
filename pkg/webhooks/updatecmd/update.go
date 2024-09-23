package updatecmd

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"strings"

	"github.com/gugahoi/basiq/internal/api"
	"github.com/urfave/cli/v2"
)

// create creates a webhook.
// https://api.basiq.io/reference/addwebhook
func New() *cli.Command {
	return &cli.Command{
		Name:      "update",
		Usage:     "update a webhook",
		UsageText: `update <id> url=<url> description=<description> name=<name> events=<event1,event2,...>`,
		Before: func(ctx *cli.Context) error {
			if ctx.Args().Len() == 0 {
				return fmt.Errorf("invalid number of arguments")
			}
			return nil
		},

		Action: func(ctx *cli.Context) error {
			client := ctx.App.Metadata["client"].(*api.ClientWithResponses)
			return exec(client, ctx.Args().First(), ctx.Args().Slice())
		},
	}
}

// parseArgs parses the arguments to the update command. It accepts the following arguments in key=value format:
// url=<url> description=<description> name=<name> events=<event1,event2,...>
func parseArgs(args []string) *api.WebhookBody {
	var payload api.WebhookBody

	for _, arg := range args {
		if strings.Contains(arg, "=") {
			kv := strings.Split(arg, "=")
			key := kv[0]
			value := kv[1]

			switch key {
			case "url":
				payload.Url = value
			case "description":
				payload.Description = &value
			case "name":
				payload.Name = &value
			case "events":
				subscribedEvents := strings.Split(value, ",")
				payload.SubscribedEvents = &subscribedEvents
			}
		}
	}

	return &payload
}

func exec(c *api.ClientWithResponses, id string, args []string) error {
	payload := parseArgs(args)

	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to serialize webhook: %w", err)
	}
	response, err := c.UpdateWebhookWithBody(context.Background(), id, "application/json", bytes.NewReader(jsonPayload))
	if err != nil {
		return fmt.Errorf("failed to create webhook: %w", err)
	}
	if response.StatusCode != 200 {
		body, err := io.ReadAll(response.Body)
		defer response.Body.Close()
		if err != nil {
			return fmt.Errorf("failed to update webhook: %w", err)
		}
		return fmt.Errorf("failed to update webhook: [%d] %s", response.StatusCode, string(body))
	}

	return nil
}
