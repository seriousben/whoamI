package main

import (
	"flag"

	"github.com/seriousben/whoamI/server"
)

var port string

func init() {
	flag.StringVar(&port, "port", "80", "give me a port number")
}

func main() {
	flag.Parse()
	server.Start(port)
}
