package loadBalancer

import (
	"log"
	"net"
	"sync"
	"time"
)

type LoadBalancerRoundRobin struct {
	ip       string
	port     string
	name     string
	nAddrs   int
	addrs    []string
	addrsMap map[string]bool
	mu       sync.Mutex
}

func (loadbalancer *LoadBalancerRoundRobin) NewLoadBalancer(name string, ip string, port string) {
	loadbalancer.ip = ip
	loadbalancer.port = port
	loadbalancer.name = name
	log.SetPrefix(loadbalancer.name + " ")
	loadbalancer.addrsMap = make(map[string]bool)

}

func (loadbalancer *LoadBalancerRoundRobin) AddNode(addr string) *LoadBalancerError {
	if !livenessProbe(addr) {
		return &LoadBalancerError{500, "Connection error"}
	}

	loadbalancer.mu.Lock()
	defer loadbalancer.mu.Unlock()

	loadbalancer.addrsMap[addr] = true
	loadbalancer.addrs = append(loadbalancer.addrs, addr)

	loadbalancer.nAddrs = loadbalancer.nAddrs + 1

	return nil

}

func (loadbalancer *LoadBalancerRoundRobin) Forward(conn net.Conn) error {
	b := make([]byte, 1024)
	addr := loadbalancer.GetNode()
	log.Println("[FORWARDING] TO ", addr)
	nodeConn, err := net.Dial("tcp", addr)
	if err != nil {
		panic(err)
	}

	//forward
	conn.Read(b)
	nodeConn.Write(b)

	//backward
	nodeConn.Read(b)
	conn.Write(b)

	defer conn.Close()
	defer nodeConn.Close()

	return nil
}

func (loadbalancer *LoadBalancerRoundRobin) GetNode() string {
	for i := 0; i < loadbalancer.nAddrs; i++ {
		index := i % loadbalancer.nAddrs
		if loadbalancer.addrsMap[loadbalancer.addrs[index]] {
			loadbalancer.mu.Lock()
			loadbalancer.addrsMap[loadbalancer.addrs[index]] = false

			defer loadbalancer.mu.Unlock()

			return loadbalancer.addrs[index]
		}
	}

	return ""

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

	log.Println("[LIVENESS PROBE] Successfully connected to ", addr)
	return true

}

func (loadbalancer *LoadBalancerRoundRobin) Start() {

	l, err := net.Listen("tcp", loadbalancer.ip+":"+loadbalancer.port)
	if err != nil {
		log.Fatalln(err)
	}

	defer l.Close()
	for {

		conn, err := l.Accept()
		if err != nil {
			log.Fatalln(err)
		}
		go loadbalancer.Forward(conn)
	}

}
