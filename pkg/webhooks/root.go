package webhooks

import (
	"github.com/gugahoi/basiq/pkg/webhooks/createcmd"
	"github.com/gugahoi/basiq/pkg/webhooks/deletecmd"
	"github.com/gugahoi/basiq/pkg/webhooks/getcmd"
	"github.com/gugahoi/basiq/pkg/webhooks/listcmd"
	"github.com/gugahoi/basiq/pkg/webhooks/updatecmd"

	"github.com/urfave/cli/v2"
)

func NewRootCmd() *cli.Command {
	return &cli.Command{
		Name:  "webhooks",
		Usage: "commands to manage webhooks in Basiq",
		Subcommands: []*cli.Command{
			createcmd.New(),
			deletecmd.New(),
			getcmd.New(),
			listcmd.New(),
			updatecmd.New(),
		},
	}
}
