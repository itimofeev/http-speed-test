package main

import "flag"

func main() {
	isServerPtr := flag.String("mode", "server", "can be server or client (default is server)")

	flag.Parse()

	if *isServerPtr == "client" {
		runClient()
		return
	}
	runServer(":13579")
}
