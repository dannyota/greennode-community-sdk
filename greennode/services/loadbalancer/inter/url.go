package inter

import "github.com/dannyota/greennode-community-sdk/v2/greennode/client"

func createLoadBalancerUrl(psc client.ServiceClient) string {
	return psc.ServiceURL("loadBalancers")
}
