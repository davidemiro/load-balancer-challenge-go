package loadBalancer

import (
	"fmt"
	"net/http"
)

type httpHandler struct {
	loadbalancer *LoadBalancerRoundRobin
}

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

}

func (loadbalancer *LoadBalancerRoundRobin) AddNode(addr string) {
	//TODO: add liveness probe to check addr correctness
	if livenessProbe(addr) {
		//return error("Connection error")
	}
	loadbalancer.addrs = append(loadbalancer.addrs, addr)
	loadbalancer.n_addrs += 1
}

func (loadbalancer *LoadBalancerRoundRobin) forward(w http.ResponseWriter, req *http.Request) *http.Response {

	// Get the addr of the node to forward the request
	if len(loadbalancer.addrs) == 0 {
		fmt.Println("Internal Server Error.\n No backend nodes available.\n")
		return nil
	}

	//Implementation of Round Robin stragegy
	selected_addr := loadbalancer.addrs[loadbalancer.next]
	loadbalancer.next = (loadbalancer.next + 1) % loadbalancer.n_addrs

	// Define the new connection to the backend
	newReq, err := http.NewRequest(req.Method, selected_addr, req.Body)
	fmt.Print(err)

	// Copy headers from the original request
	for key, value := range req.Header {

		newReq.Header.Set(key, value[0])
	}

	// Send the cloned request
	client := &http.Client{}
	nodeResp, err := client.Do(newReq)

	defer nodeResp.Body.Close()

	if err != nil {
		fmt.Println(err)
		return nil
	}

	defer nodeResp.Body.Close()

	return nodeResp

}

func livenessProbe(addr string) bool {
	return true

}

func (httpHandler *httpHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	resp := httpHandler.loadbalancer.forward(w, req)

	if resp == nil {
		return
	}

	// Copy headers from the original request
	for key, value := range req.Header {
		resp.Header.Set(key, value[0])
	}

	// Set the status code for the original response
	w.WriteHeader(resp.StatusCode)

	// Write the response body to the original client
	var bytes []byte
	resp.Body.Read(bytes)
	_, err := w.Write(bytes)

	if err != nil {
		fmt.Print(err)
	}

}

func (loadbalancer *LoadBalancerRoundRobin) Start() {
	handler := new(httpHandler)
	handler.loadbalancer = loadbalancer
	fmt.Printf("Starting server %s at port %s and address %s\n", loadbalancer.name, loadbalancer.port, loadbalancer.ip)
	if err := http.ListenAndServe(fmt.Sprintf("%s:%s", loadbalancer.ip, loadbalancer.port), handler); err != nil {
		fmt.Println(err)
	}

}
