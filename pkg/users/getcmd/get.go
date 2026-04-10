package getcmd

import (
	"context"
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/gugahoi/basiq/internal/api/users"
	"github.com/gugahoi/basiq/tools"
	"github.com/urfave/cli/v3"
)

// New returns a cli.Command that retrieves a single user by ID.
func New() *cli.Command {
	return &cli.Command{
		Name:  "get",
		Usage: "retrieve a user by ID",
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

// exec retrieves a user and prints their details in tabular form.
func exec(ctx context.Context, client *users.Client, userID string) error {
	user, err := client.GetUser(ctx, userID)
	if err != nil {
		return err
	}

	w := tabwriter.NewWriter(os.Stdout, 0, 0, 4, ' ', 0)
	fmt.Fprintf(w, "ID\tEMAIL\tMOBILE\tFIRST NAME\tLAST NAME\n")
	fmt.Fprintf(w, "%s\t%s\t%s\t%s\t%s\n", user.Id, user.Email, user.Mobile, user.FirstName, user.LastName)
	w.Flush()

	return nil
}
