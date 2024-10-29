package loadBalancer

import (
	"log"
	"net"
	"time"
)

type LoadBalancerRoundRobin struct {
	ip      string
	port    string
	name    string
	next    int
	addrs   []string
	n_addrs int
}

func (loadbalancer *LoadBalancerRoundRobin) NewLoadBalancer(name string, ip string, port string) {
	loadbalancer.next = 0
	loadbalancer.ip = ip
	loadbalancer.port = port
	loadbalancer.name = name
	log.SetPrefix(loadbalancer.name + " ")

}

func (loadbalancer *LoadBalancerRoundRobin) AddNode(addr string) *LoadBalancerError {
	if !livenessProbe(addr) {
		return &LoadBalancerError{500, "Connection error"}
	}
	loadbalancer.addrs = append(loadbalancer.addrs, addr)
	loadbalancer.n_addrs += 1

	return nil

}

func (loadbalancer *LoadBalancerRoundRobin) Forward(conn net.Conn) error {
	return nil
}

func (loadbalancer *LoadBalancerRoundRobin) Backward(addr string) error {
	return nil
}

func livenessProbe(addr string) bool {

	// Use DialTimeout to set a timeout for the connection attempt
	conn, err := net.DialTimeout("tcp", addr, 5*time.Second)
	if err != nil {
		log.Fatalln(err)
		return false
	}

	// Close the connection after establishing it
	defer conn.Close()

	log.Println("Successfully connected to " + addr + "\n")
	return true

}

func (loadbalancer *LoadBalancerRoundRobin) Start() {

	l, err := net.Listen("tcp", loadbalancer.ip+":"+loadbalancer.port)
	if err != nil {
		log.Fatalln(err)
	}
	defer l.Close()

	conn, err := l.Accept()
	if err != nil {
		log.Fatalln(err)
	}
	go loadbalancer.Forward(conn)

}
