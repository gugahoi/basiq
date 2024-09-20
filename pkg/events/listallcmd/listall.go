package listallcmd

import (
	"context"
	"fmt"
	"os"
	"strings"
	"text/tabwriter"

	"github.com/gugahoi/basiq/internal/api/events"
	"github.com/urfave/cli/v2"
)

func New() *cli.Command {
	return &cli.Command{
		Name:      "listall",
		Usage:     "list all events",
		UsageText: "basiq events listall [user_id=<user_id>] [type=<type>] [entity=<entity>]",
		Action: func(ctx *cli.Context) error {
			return exec(ctx.App.Metadata["client"].(*events.Client), ctx.Args().Slice()...)
		},
	}
}

// args can be:
// entity=<entity> type=<event_type> user_id=<user_id>
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
