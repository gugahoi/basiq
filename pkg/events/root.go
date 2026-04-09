package events

import (
	"github.com/gugahoi/basiq/pkg/events/getcmd"
	"github.com/gugahoi/basiq/pkg/events/listallcmd"
	"github.com/gugahoi/basiq/pkg/events/listcmd"
	"github.com/gugahoi/basiq/pkg/events/testcmd"
	"github.com/gugahoi/basiq/pkg/events/typescmd"
	"github.com/urfave/cli/v2"
)

func NewRootCmd() *cli.Command {
	return &cli.Command{
		Name:  "events",
		Usage: "commands to manage events in Basiq",
		Subcommands: []*cli.Command{
			listcmd.New(),
			getcmd.New(),
			typescmd.New(),
			listallcmd.New(),
			testcmd.New(),
		},
	}
}
