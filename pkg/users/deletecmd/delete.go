package deletecmd

import (
	"context"
	"fmt"

	"github.com/gugahoi/basiq/internal/api/users"
	"github.com/gugahoi/basiq/tools"
	"github.com/urfave/cli/v3"
)

// New returns a cli.Command that permanently deletes a user by ID.
func New() *cli.Command {
	return &cli.Command{
		Name:    "delete",
		Usage:   "permanently delete a user by ID",
		Aliases: []string{"rm"},
		Before: func(ctx context.Context, cmd *cli.Command) (context.Context, error) {
			if cmd.Args().Len() == 0 {
				return ctx, fmt.Errorf("missing user ID")
			}
			return ctx, nil
		},
		Action: func(ctx context.Context, cmd *cli.Command) error {
			client := tools.GetUsersClient(cmd)
			return exec(ctx, client, cmd.Args().First())
		},
	}
}

// exec deletes a user by ID.
func exec(ctx context.Context, client *users.Client, userID string) error {
	if err := client.DeleteUser(ctx, userID); err != nil {
		return err
	}
	fmt.Printf("user %s deleted\n", userID)
	return nil
}
