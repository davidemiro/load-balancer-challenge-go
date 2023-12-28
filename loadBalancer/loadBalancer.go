package loadBalancer

import "net/http"

type LoadBalancer interface {
	forward(balancer LoadBalancer) (w http.ResponseWriter, req *http.Request)
}

type LoadBalancerError struct {
	Code    int
	Message string
}
