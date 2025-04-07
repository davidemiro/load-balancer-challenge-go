package loadBalancer

type LoadBalancer interface {
	Start(balancer LoadBalancer)
	Forward()
	GetNode() string
	AddNode(address string)
}

type LoadBalancerError struct {
	Code    int
	Message string
}
