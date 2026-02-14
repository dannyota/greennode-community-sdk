package inter

import lsclient "github.com/dannyota/greennode-community-sdk/v2/greennode/client"

type LoadBalancerServiceInternal struct {
	VLBClient lsclient.IServiceClient
}
