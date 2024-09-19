package main

import (
	"flag"
	"log"
	"os"

	"github.com/gugahoi/basiq/pkg/createcmd"
	"github.com/gugahoi/basiq/pkg/deletecmd"
	"github.com/gugahoi/basiq/pkg/getcmd"
	"github.com/gugahoi/basiq/pkg/listcmd"
	"github.com/gugahoi/basiq/tools"
)

func main() {
	// remove timestamp from log lines
	log.SetFlags(0)
	log.SetOutput(os.Stdout)

	var apikey string
	flag.StringVar(&apikey, "apikey", "", "Basiq API key")
	flag.Parse()

	if apikey == "" {
		log.Fatalln("apikey is required")
	}

	c := tools.CreateClient(apikey)

	args := flag.Args()

	cmd := args[0]

	switch cmd {
	case "list":
		listcmd.List(c)
	case "delete":
		deletecmd.Delete(c, args[2])
	case "create":
		createcmd.Create(c, args[2:])
	case "get":
		getcmd.Get(c, args[2])
	default:
		usage()
	}
}

func usage() {
	log.Fatalln("available commands: list, get")
}
