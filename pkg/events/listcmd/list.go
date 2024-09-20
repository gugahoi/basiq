package listcmd

import (
	"context"
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/gugahoi/basiq/internal/api/events"
	"github.com/urfave/cli/v2"
)

func New() *cli.Command {
	return &cli.Command{
		Name:  "list",
		Usage: "list all events types",
		Action: func(ctx *cli.Context) error {
			client := ctx.App.Metadata["client"].(*events.Client)
			return exec(client)
		},
	}
}

func exec(client *events.Client) error {
	types, err := client.ListAllTypes(context.Background())
	if err != nil {
		return fmt.Errorf("failed to list events: %w", err)
	}
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', 0)
	for _, t := range types.Data {
		fmt.Fprintf(w, "%s\t%s\n", t.Id, t.Description)
	}
	w.Flush()
	return nil
}
