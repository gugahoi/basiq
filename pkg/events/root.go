package events

import (
	"github.com/gugahoi/basiq/pkg/events/getcmd"
	"github.com/gugahoi/basiq/pkg/events/listcmd"
	"github.com/gugahoi/basiq/pkg/events/testcmd"
	"github.com/gugahoi/basiq/pkg/events/typescmd"
	"github.com/urfave/cli/v2"
)

// NewRootCmd returns the top-level "events" command with subcommands for managing events and event types.
func NewRootCmd() *cli.Command {
	return &cli.Command{
		Name:  "events",
		Usage: "commands to manage events in Basiq",
		Subcommands: []*cli.Command{
			typescmd.New(),
			listcmd.New(),
			getcmd.New(),
			testcmd.New(),
		},
	}
}
