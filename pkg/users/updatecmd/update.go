package updatecmd

import (
	"context"
	"fmt"
	"os"
	"strings"
	"text/tabwriter"

	"github.com/gugahoi/basiq/internal/api/users"
	"github.com/gugahoi/basiq/tools"
	"github.com/urfave/cli/v3"
)

// New returns a cli.Command that updates an existing user by ID.
func New() *cli.Command {
	return &cli.Command{
		Name:      "update",
		Usage:     "update a user",
		UsageText: "update <id> [email=<email>] [mobile=<mobile>] [firstName=<firstName>] [lastName=<lastName>] [businessName=<businessName>]",
		Before: func(ctx context.Context, cmd *cli.Command) (context.Context, error) {
			if cmd.Args().Len() < 2 {
				return ctx, fmt.Errorf("missing user ID or fields to update")
			}
			return ctx, nil
		},
		Action: func(ctx context.Context, cmd *cli.Command) error {
			client := tools.GetUsersClient(cmd)
			return exec(ctx, client, cmd.Args().First(), cmd.Args().Tail())
		},
	}
}

// parseArgs parses key=value arguments into an UpdateUserBody.
// Supported keys: email, mobile, firstName, lastName, businessName, businessIdNo, businessIdNoType, verificationDate.
func parseArgs(args []string) *users.UpdateUserBody {
	var payload users.UpdateUserBody

	for _, arg := range args {
		if !strings.Contains(arg, "=") {
			continue
		}
		kv := strings.SplitN(arg, "=", 2)
		key, value := kv[0], kv[1]

		switch key {
		case "email":
			payload.Email = &value
		case "mobile":
			payload.Mobile = &value
		case "firstName":
			payload.FirstName = &value
		case "lastName":
			payload.LastName = &value
		case "businessName":
			payload.BusinessName = &value
		case "businessIdNo":
			payload.BusinessIdNo = &value
		case "businessIdNoType":
			payload.BusinessIdNoType = &value
		case "verificationDate":
			payload.VerificationDate = &value
		}
	}

	return &payload
}

// exec updates a user and prints the updated details in tabular form.
func exec(ctx context.Context, client *users.Client, userID string, args []string) error {
	payload := parseArgs(args)

	user, err := client.UpdateUser(ctx, userID, payload)
	if err != nil {
		return err
	}

	w := tabwriter.NewWriter(os.Stdout, 0, 0, 4, ' ', 0)
	fmt.Fprintf(w, "ID\tEMAIL\tMOBILE\tFIRST NAME\tLAST NAME\n")
	fmt.Fprintf(w, "%s\t%s\t%s\t%s\t%s\n", user.Id, user.Email, user.Mobile, user.FirstName, user.LastName)
	w.Flush()

	return nil
}
