package getcmd

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
		Name:  "get",
		Usage: "get an event type",
		Before: func(ctx *cli.Context) error {
			if ctx.Args().Len() == 0 {
				return fmt.Errorf("missing event type ID")
			}
			return nil
		},
		Action: func(ctx *cli.Context) error {
			client := ctx.App.Metadata["client"].(*events.Client)
			return exec(client, ctx.Args().First())
		},
	}
}

func exec(client *events.Client, id string) error {
	data, err := client.GetType(context.Background(), id)
	if err != nil {
		return fmt.Errorf("failed to get event type: %w", err)
	}
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 4, ' ', 0)
	fmt.Fprintf(w, "%s\t%s\n", data.Id, data.Description)
	w.Flush()
	return nil
}
