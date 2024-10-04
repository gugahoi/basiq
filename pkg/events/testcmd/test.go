package testcmd

import (
	"context"
	"fmt"

	"github.com/gugahoi/basiq/internal/api/events"
	"github.com/urfave/cli/v2"
)

func New() *cli.Command {
	return &cli.Command{
		Name:  "test",
		Usage: "post a test message",
		Action: func(ctx *cli.Context) error {
			client := ctx.App.Metadata["client"].(*events.Client)
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
