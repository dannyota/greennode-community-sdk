package v2

import "github.com/dannyota/greennode-community-sdk/v2/greennode/client"

type NetworkServiceV2 struct {
	VserverClient client.IServiceClient
}

func (s *NetworkServiceV2) getProjectId() string {
	return s.VserverClient.GetProjectId()
}
