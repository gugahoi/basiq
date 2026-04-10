package testcmd

import (
	"context"
	"fmt"

	"github.com/gugahoi/basiq/internal/api/events"
	"github.com/gugahoi/basiq/tools"
	"github.com/urfave/cli/v3"
)

// New returns a cli.Command that posts a test message for an event type.
func New() *cli.Command {
	return &cli.Command{
		Name:  "test",
		Usage: "post a test message",
		Action: func(ctx context.Context, cmd *cli.Command) error {
			client := tools.GetEventsClient(cmd)
			return exec(client, cmd.Args().First())
		},
	}
}

// exec posts a test message for the given event type ID.
func exec(client *events.Client, eventTypeID string) error {
	err := client.TestMessage(context.Background(), eventTypeID)
	if err != nil {
		return fmt.Errorf("failed to list events: %w", err)
	}
	return nil
}
