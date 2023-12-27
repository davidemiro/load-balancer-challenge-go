package main

import (
	"LoadBalancer/server"
)

func main() {

	//create a new handler
	server := new(server.Server)
	server.NewServer("S1", "127.0.0.1", "8080")
	server.Start()

}
