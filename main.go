package main

import (
	"log"
	"os"

	"github.com/gugahoi/basiq/pkg/events"
	"github.com/gugahoi/basiq/pkg/webhooks"
	"github.com/urfave/cli/v2"
)

func main() {
	// remove timestamp from log lines
	log.SetFlags(0)
	log.SetOutput(os.Stdout)

	app := cli.NewApp()

	app.Authors = []*cli.Author{
		{
			Name:  "Gustavo Hoirisch",
			Email: "github@gustavo.com.au",
		},
	}
	app.Usage = "Basiq CLI client"
	app.Commands = []*cli.Command{
		webhooks.NewRootCmd(),
		events.NewRootCmd(),
	}
	app.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:     "apikey",
			EnvVars:  []string{"BASIQ_APIKEY"},
			Usage:    "Basiq API key",
			Required: true,
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatalln(err)
	}
}
