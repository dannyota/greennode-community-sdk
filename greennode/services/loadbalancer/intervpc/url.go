package intervpc

import "github.com/dannyota/greennode-community-sdk/greennode/client"

func createLoadBalancerURL(sc *client.ServiceClient) string {
	return sc.ServiceURL("loadBalancers")
}
