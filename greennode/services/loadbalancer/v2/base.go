package v2

import "github.com/dannyota/greennode-community-sdk/v2/greennode/client"

type LoadBalancerServiceV2 struct {
	VLBClient     client.ServiceClient
	VServerClient client.ServiceClient
}

func (s *LoadBalancerServiceV2) getProjectID() string {
	return s.VLBClient.GetProjectID()
}

const (
	defaultPageListLoadBalancer = 1
	defaultSizeListLoadBalancer = 10
)
