package main

import (
	"flag"
	"log"
	"os"

	"github.com/seriousben/whoamI/client"
	"github.com/seriousben/whoamI/server"
)

func main() {
	configCommand := flag.NewFlagSet("config", flag.ExitOnError)
	port := configCommand.String("port", "8080", "give me a port number")

	command := "server"
	if len(os.Args) >= 2 {
		command = os.Args[1]
	}

	switch command {
	case "healthcheck":
		configCommand.Parse(os.Args[2:])
		client.Healthcheck(*port)
	case "serve":
		configCommand.Parse(os.Args[2:])
		server.Start(*port)
	default:
		log.Fatalf("unknown command %s", command)
	}
}
