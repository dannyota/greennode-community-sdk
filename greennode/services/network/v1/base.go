package v1

import "github.com/dannyota/greennode-community-sdk/v2/greennode/client"

type NetworkServiceV1 struct {
	VNetworkClient client.ServiceClient
}

func (s *NetworkServiceV1) getProjectId() string {
	return s.VNetworkClient.GetProjectId()
}

type NetworkServiceInternalV1 struct {
	VNetworkClient client.ServiceClient
}

func (s *NetworkServiceInternalV1) getProjectId() string {
	return s.VNetworkClient.GetProjectId()
}
