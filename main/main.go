package main

import (
	"LoadBalancer/server"
)

func StartServer(name string, ip string, port string) {
	server := new(server.Server)
	server.NewServer(name, ip, port)
	server.Start()
}

func main() {

	//create a new handler

	go StartServer("S1", "127.0.0.1", "8080")

	go StartServer("S2", "127.0.0.2", "8080")

	select {}

}
