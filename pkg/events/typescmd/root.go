package typescmd

import (
	"github.com/gugahoi/basiq/pkg/events/typescmd/getcmd"
	"github.com/gugahoi/basiq/pkg/events/typescmd/listcmd"
	"github.com/urfave/cli/v2"
)

// New returns a cli.Command group for event type operations.
func New() *cli.Command {
	return &cli.Command{
		Name:  "types",
		Usage: "commands to manage event types",
		Subcommands: []*cli.Command{
			listcmd.New(),
			getcmd.New(),
		},
	}
}
