package main

import (
	"flag"
	"log"
	"os"

	"github.com/gugahoi/basiq/internal/api"
	"github.com/gugahoi/basiq/pkg/createcmd"
	"github.com/gugahoi/basiq/pkg/deletecmd"
	"github.com/gugahoi/basiq/pkg/getcmd"
	"github.com/gugahoi/basiq/pkg/listcmd"
	"github.com/oapi-codegen/oapi-codegen/v2/pkg/securityprovider"
)

func createClient(apikey string) *api.ClientWithResponses {
	token := getAuthToken(apikey)
	auth, err := securityprovider.NewSecurityProviderBearerToken(token)

	client, err := api.NewClientWithResponses(ServerURL, api.WithRequestEditorFn(auth.Intercept))
	if err != nil {
		log.Fatalln("failed to generate Basiq client", err)
	}
	return client

}

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

	c := createClient(apikey)

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
