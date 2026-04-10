package main

import (
	"context"
	"log"
	"net/mail"
	"os"

	"github.com/urfave/cli/v3"

	"github.com/gugahoi/basiq/pkg/events"
	"github.com/gugahoi/basiq/pkg/webhooks"
)

func main() {
	// remove timestamp from log lines
	log.SetFlags(0)
	log.SetOutput(os.Stdout)

	app := &cli.Command{
		Authors: []any{
			&mail.Address{
				Name:    "Gustavo Hoirisch",
				Address: "github@gustavo.com.au",
			},
		},
		EnableShellCompletion: true,
		Usage:                 "Basiq CLI client",
		Commands: []*cli.Command{
			webhooks.NewRootCmd(),
			events.NewRootCmd(),
		},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "apikey",
				Sources:  cli.EnvVars("BASIQ_APIKEY"),
				Usage:    "Basiq API key",
				Required: true,
			},
		},
	}

	if err := app.Run(context.Background(), os.Args); err != nil {
		log.Fatalln(err)
	}
}
