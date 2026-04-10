package getcmd

import (
	"context"
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/urfave/cli/v3"

	"github.com/gugahoi/basiq/internal/api/events"
	"github.com/gugahoi/basiq/tools"
)

// New returns a cli.Command that retrieves a single event by ID.
func New() *cli.Command {
	return &cli.Command{
		Name:  "get",
		Usage: "retrieve an event by ID",
		ShellComplete: func(ctx context.Context, cmd *cli.Command) {
			client := tools.GetEventsClient(cmd)
			res, err := client.ListAll(ctx, events.ListAllFilters{})
			if err != nil {
				return
			}
			for _, e := range res.Data {
				fmt.Printf("%s:%s.%s\n", e.Id, e.Entity, e.EventType)
			}
		},
		Before: func(ctx context.Context, cmd *cli.Command) (context.Context, error) {
			if cmd.Args().Len() == 0 {
				return ctx, fmt.Errorf("missing event ID")
			}
			return ctx, nil
		},
		Action: func(ctx context.Context, cmd *cli.Command) error {
			client := tools.GetEventsClient(cmd)
			return exec(client, cmd.Args().First())
		},
	}
}

func exec(client *events.Client, id string) error {
	data, err := client.GetEvent(context.Background(), id)
	if err != nil {
		return fmt.Errorf("failed to get event: %w", err)
	}
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 4, ' ', 0)
	fmt.Fprintf(w, "%s\t%s\t%s\t%s\n", data.Id, data.Entity, data.EventType, data.Data)
	w.Flush()
	return nil
}
