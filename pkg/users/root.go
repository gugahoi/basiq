package users

import (
	"github.com/gugahoi/basiq/pkg/users/deletecmd"
	"github.com/gugahoi/basiq/pkg/users/getcmd"
	"github.com/gugahoi/basiq/pkg/users/updatecmd"
	"github.com/urfave/cli/v3"
)

// NewRootCmd returns the top-level "users" command with subcommands for managing users.
func NewRootCmd() *cli.Command {
	return &cli.Command{
		Name:  "users",
		Usage: "commands to manage users in Basiq",
		Commands: []*cli.Command{
			getcmd.New(),
			updatecmd.New(),
			deletecmd.New(),
		},
	}
}
