package intervpc

import "danny.vn/greennode/client"

func createLoadBalancerURL(sc *client.ServiceClient) string {
	return sc.ServiceURL("loadBalancers")
}
