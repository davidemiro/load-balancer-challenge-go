package loadBalancer

import (
	"LoadBalancer/server"
	"fmt"
	"net/http"
)

type LoadBalancerRoundRobin struct {
	handler *server.HttpHandler
	ip      string
	port    string
	name    string
	next    int
	addrs   []string
	n_addrs int
}

func (loadBalancerRR *LoadBalancerRoundRobin) NewLoadBalancer(name string, ip string, port string) {
	loadBalancerRR.next = 0
	loadBalancerRR.ip = ip
	loadBalancerRR.port = port

}

func (loadBalancerRR *LoadBalancerRoundRobin) AddNode(addr string) {
	loadBalancerRR.addrs = append(loadBalancerRR.addrs, addr)
	loadBalancerRR.n_addrs += 1
}

func (loadBalancerRR *LoadBalancerRoundRobin) forward(w http.ResponseWriter, req *http.Request) *http.Response {

	// Get the addr of the node to forward the request
	if len(loadBalancerRR.addrs) == 0 {
		fmt.Println("Internal Server Error.\n No backend nodes available.\n")
		return nil
	}

	//Implementation of Round Robin stragegy
	selected_addr := loadBalancerRR.addrs[loadBalancerRR.next]
	loadBalancerRR.next = (loadBalancerRR.next + 1) % loadBalancerRR.n_addrs

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

func (LoadBalancerRR *LoadBalancerRoundRobin) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	resp := LoadBalancerRR.forward(w, req)

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

func (LoadBalancerRR *LoadBalancerRoundRobin) Start() {
	LoadBalancerRR.handler = new(server.HttpHandler)
	LoadBalancerRR.handler.NewHttpHandler(LoadBalancerRR.name)
	fmt.Printf("Starting server at port %s and address %s\n", LoadBalancerRR.port, LoadBalancerRR.ip)
	if err := http.ListenAndServe(fmt.Sprintf("%s:%s", LoadBalancerRR.ip, LoadBalancerRR.port), LoadBalancerRR.handler); err != nil {
		fmt.Println(err)
	}

}
