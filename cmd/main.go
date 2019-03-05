package main

import (
	"flag"
	"github.com/itimofeev/http-speed-test"
)

func main() {
	isServerPtr := flag.String("mode", "server", "can be server or client (default is server)")

	flag.Parse()

	if *isServerPtr == "client" {
		speedt.RunClient()
		return
	}
	speedt.RunServer(":13579")
}
