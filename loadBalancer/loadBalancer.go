package loadBalancer

type LoadBalancer interface {
	Start(balancer LoadBalancer)
	Forward()
	Backward()
	AddNode(address string)
}

type LoadBalancerError struct {
	Code    int
	Message string
}
