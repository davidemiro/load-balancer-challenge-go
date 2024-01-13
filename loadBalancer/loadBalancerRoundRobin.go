package loadBalancer

import (
	"log"
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"
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

	// Create a handler function that serves as the entry point for incoming requests
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Modify the request before it is sent to the destination server if needed
		// For example, you might want to change headers or add authentication

		//Implementation of Round Robin stragegy
		selected_addr := loadbalancer.addrs[loadbalancer.next]
		loadbalancer.next = (loadbalancer.next + 1) % loadbalancer.n_addrs

		log.Println("Selected address " + selected_addr)

		targetURL, error := url.Parse("http://" + selected_addr)
		if error != nil {
			log.Fatal(error)
		}

		// Create a reverse proxy
		proxy := httputil.NewSingleHostReverseProxy(targetURL)

		// Forward the request to the target server
		proxy.ServeHTTP(w, r)
	})

	// Start the HTTP server on a specific port
	if err := http.ListenAndServe(loadbalancer.ip+":"+loadbalancer.port, nil); err != nil {
		log.Fatalln(err)
		panic(err)
	}

}
