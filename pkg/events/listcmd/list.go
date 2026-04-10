package listcmd

import (
	"context"
	"fmt"
	"os"
	"strings"
	"text/tabwriter"

	"github.com/gugahoi/basiq/internal/api/events"
	"github.com/gugahoi/basiq/tools"
	"github.com/urfave/cli/v3"
)

// New returns a cli.Command that lists all events with optional filters.
func New() *cli.Command {
	return &cli.Command{
		Name:      "list",
		Usage:     "list all events",
		UsageText: "basiq events list [user_id=<user_id>] [type=<type>] [entity=<entity>]",
		Action: func(ctx context.Context, cmd *cli.Command) error {
			return exec(tools.GetEventsClient(cmd), cmd.Args().Slice()...)
		},
	}
}

// parseArgs parses key=value filter arguments into a ListAllFilters struct.
// Supported keys: entity, type, user_id.
func parseArgs(args []string) events.ListAllFilters {
	var filters events.ListAllFilters
	for _, arg := range args {
		parts := strings.Split(arg, "=")
		key := parts[0]
		value := parts[1]

		switch key {
		case "entity":
			filters.Entity = &value
		case "type":
			filters.Type = &value
		case "user_id":
			filters.UserId = &value
		}
	}
	return filters
}

// exec fetches and displays all events, applying any provided filters.
func exec(client *events.Client, args ...string) error {
	filters := parseArgs(args)

	res, err := client.ListAll(context.Background(), filters)
	if err != nil {
		return fmt.Errorf("failed to list events: %w", err)
	}

	w := tabwriter.NewWriter(os.Stdout, 0, 0, 4, ' ', 0)
	for _, event := range res.Data {
		fmt.Fprintf(w, "%s\t%s\t%s\t%s\n", event.Id, event.Entity, event.EventType, event.Data)
	}
	w.Flush()

	return nil
}
