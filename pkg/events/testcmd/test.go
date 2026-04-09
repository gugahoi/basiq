package testcmd

import (
	"context"
	"fmt"

	"github.com/gugahoi/basiq/internal/api/events"
	"github.com/gugahoi/basiq/tools"
	"github.com/urfave/cli/v2"
)

func New() *cli.Command {
	return &cli.Command{
		Name:  "test",
		Usage: "post a test message",
		Action: func(ctx *cli.Context) error {
			client := tools.GetEventsClient(ctx)
			return exec(client, ctx.Args().First())
		},
	}
}

func exec(client *events.Client, eventTypeID string) error {
	err := client.TestMessage(context.Background(), eventTypeID)
	if err != nil {
		return fmt.Errorf("failed to list events: %w", err)
	}
	return nil
}
