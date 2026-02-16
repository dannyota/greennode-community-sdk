package intervpc

import "github.com/dannyota/greennode-community-sdk/v2/greennode/client"

func createLoadBalancerURL(sc *client.ServiceClient) string {
	return sc.ServiceURL("loadBalancers")
}
