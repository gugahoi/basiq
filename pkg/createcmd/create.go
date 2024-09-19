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
		Name:      "create",
		Usage:     "create a webhook",
		UsageText: `create url=<url> description=<description> name=<name> events=<event1,event2,...>`,
		Before: func(ctx *cli.Context) error {
			if ctx.Args().Len() == 0 {
				return fmt.Errorf("invalid number of arguments")
			}
			return nil
		},

		Action: func(ctx *cli.Context) error {
			client := ctx.App.Metadata["client"].(*api.ClientWithResponses)
			return exec(client, ctx.Args().Slice())
		},
	}
}

// parseArgs parses the arguments to the create command. It accepts the following arguments in key=value format:
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

	log.Printf("%#v\n", payload)
	return &payload
}

func exec(c *api.ClientWithResponses, args []string) error {
	payload := parseArgs(args)

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
