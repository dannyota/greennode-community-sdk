package inter

import lsclient "github.com/dannyota/greennode-community-sdk/v2/greennode/client"

func createLoadBalancerUrl(psc lsclient.IServiceClient) string {
	return psc.ServiceURL("loadBalancers")
}
