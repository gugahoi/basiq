package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gugahoi/basiq/pkg/createcmd"
	"github.com/gugahoi/basiq/pkg/deletecmd"
	"github.com/gugahoi/basiq/pkg/getcmd"
	"github.com/gugahoi/basiq/pkg/listcmd"
	"github.com/gugahoi/basiq/tools"
	"github.com/urfave/cli/v2"
)

func main() {
	// remove timestamp from log lines
	log.SetFlags(0)
	log.SetOutput(os.Stdout)

	app := cli.NewApp()

	app.Authors = []*cli.Author{
		&cli.Author{
			Name:  "Gustavo Hoirisch",
			Email: "github@gustavo.com.au",
		},
	}
	app.Usage = "Basiq CLI client"
	app.Before = func(ctx *cli.Context) error {
		apikey := ctx.String("apikey")
		if apikey == "" {
			return fmt.Errorf("apikey is required")
		}
		ctx.App.Metadata["client"] = tools.CreateClient(apikey)
		return nil
	}
	app.Commands = []*cli.Command{
		listcmd.New(),
		createcmd.New(),
		deletecmd.New(),
		getcmd.New(),
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
