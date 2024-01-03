package main

import (
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

	go StartServer("S1", "127.0.0.2", "8081")

	go StartServer("S2", "127.0.0.3", "8081")

	loadBalancer := new(loadBalancer.LoadBalancerRoundRobin)
	loadBalancer.NewLoadBalancer("LB", "127.0.0.1", "8080")
	loadBalancer.AddNode("127.0.0.2:8081")
	loadBalancer.AddNode("127.0.0.3:8081")

	loadBalancer.Start()

	select {}

}
