package inter

import "github.com/dannyota/greennode-community-sdk/v2/greennode/client"

func createLoadBalancerUrl(sc client.ServiceClient) string {
	return sc.ServiceURL("loadBalancers")
}
