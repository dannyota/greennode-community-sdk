package v1

import lsclient "github.com/dannyota/greennode-community-sdk/v2/greennode/client"

type GLBServiceV1 struct {
	VLBClient     lsclient.IServiceClient
	VServerClient lsclient.IServiceClient
}

const (
	defaultOffsetListGlobalLoadBalancer = 0
	defaultLimitListGlobalLoadBalancer  = 10
)
