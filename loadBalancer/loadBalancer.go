package loadBalancer

type LoadBalancer struct {
	next          int
	servers_addrs []string
}

func (server *LoadBalancer) NewLoadBalancer(next int, servers_addrs []string) {

}
