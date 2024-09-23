package webhooks

import (
	"fmt"

	"github.com/gugahoi/basiq/pkg/webhooks/createcmd"
	"github.com/gugahoi/basiq/pkg/webhooks/deletecmd"
	"github.com/gugahoi/basiq/pkg/webhooks/getcmd"
	"github.com/gugahoi/basiq/pkg/webhooks/listcmd"
	"github.com/gugahoi/basiq/pkg/webhooks/updatecmd"
	"github.com/gugahoi/basiq/tools"

	"github.com/urfave/cli/v2"
)

func NewRootCmd() *cli.Command {
	return &cli.Command{
		Name:  "webhooks",
		Usage: "commands to manage webhooks in Basiq",
		Before: func(ctx *cli.Context) error {
			apikey := ctx.String("apikey")
			if apikey == "" {
				return fmt.Errorf("apikey is required")
			}
			ctx.App.Metadata["client"] = tools.CreateClient(apikey)
			return nil
		},
		Subcommands: []*cli.Command{
			createcmd.New(),
			deletecmd.New(),
			getcmd.New(),
			listcmd.New(),
			updatecmd.New(),
		},
	}
}
