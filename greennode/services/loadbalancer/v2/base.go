package v2

import "github.com/dannyota/greennode-community-sdk/v2/greennode/client"

type LoadBalancerServiceV2 struct {
	VLBClient     client.IServiceClient
	VServerClient client.IServiceClient
}

func (s *LoadBalancerServiceV2) getProjectId() string {
	return s.VLBClient.GetProjectId()
}

const (
	defaultPageListLoadBalancer = 1
	defaultSizeListLoadBalancer = 10
)
