package main

import (
	"LoadBalancer/client"
	"LoadBalancer/loadBalancer"
	"LoadBalancer/server"
)

func StartServer(name string, ip string, port string) {
	server := new(server.Server)
	server.NewServer(name, ip, port)
	server.Start()
}

func main() {

	//create a new handler

	loadBalancer := new(loadBalancer.LoadBalancerRoundRobin)
	loadBalancer.NewLoadBalancer("LB", "127.0.10.2", "8082")
	go loadBalancer.Start()

	go StartServer("S1", "127.0.0.2", "8082")

	go StartServer("S2", "127.0.0.3", "8082")

	loadBalancer.AddNode("127.0.0.2:8082")
	loadBalancer.AddNode("127.0.0.3:8082")

	client.StartClient("C1", "127.0.10.2:8082")
	client.StartClient("C2", "127.0.10.2:8082")

}
