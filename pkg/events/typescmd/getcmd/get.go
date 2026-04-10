package getcmd

import (
	"context"
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/gugahoi/basiq/internal/api/events"
	"github.com/gugahoi/basiq/tools"
	"github.com/urfave/cli/v3"
)

// New returns a cli.Command that retrieves an event type by ID.
func New() *cli.Command {
	return &cli.Command{
		Name:  "get",
		Usage: "get an event type by ID",
		Before: func(ctx context.Context, cmd *cli.Command) (context.Context, error) {
			if cmd.Args().Len() == 0 {
				return ctx, fmt.Errorf("missing event type ID")
			}
			return ctx, nil
		},
		Action: func(ctx context.Context, cmd *cli.Command) error {
			client := tools.GetEventsClient(cmd)
			return exec(client, cmd.Args().First())
		},
	}
}

// exec fetches and displays a single event type by ID.
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
