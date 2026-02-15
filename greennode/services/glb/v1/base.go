package v1

import "github.com/dannyota/greennode-community-sdk/v2/greennode/client"

type GLBServiceV1 struct {
	VLBClient     *client.ServiceClient
	VServerClient *client.ServiceClient
}

const (
	defaultOffsetListGlobalLoadBalancer = 0
	defaultLimitListGlobalLoadBalancer  = 10
)
