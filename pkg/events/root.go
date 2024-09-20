package events

import (
	"fmt"

	"github.com/gugahoi/basiq/pkg/events/getcmd"
	"github.com/gugahoi/basiq/pkg/events/listcmd"
	"github.com/gugahoi/basiq/tools"
	"github.com/urfave/cli/v2"
)

func NewRootCmd() *cli.Command {
	return &cli.Command{
		Name:  "events",
		Usage: "commands to manage events in Basiq",
		Before: func(ctx *cli.Context) error {
			apikey := ctx.String("apikey")
			if apikey == "" {
				return fmt.Errorf("apikey is required")
			}
			ctx.App.Metadata["client"] = tools.CreateEventsClient(apikey)
			return nil
		},
		Subcommands: []*cli.Command{
			listcmd.New(),
			getcmd.New(),
		},
	}
}
