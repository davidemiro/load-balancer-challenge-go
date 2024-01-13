package loadBalancer

type LoadBalancer interface {
	Start(balancer LoadBalancer)
}

type LoadBalancerError struct {
	Code    int
	Message string
}
