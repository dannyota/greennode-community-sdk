package v2

import "github.com/dannyota/greennode-community-sdk/v2/greennode/client"

type NetworkServiceV2 struct {
	VserverClient client.ServiceClient
}

func (s *NetworkServiceV2) getProjectID() string {
	return s.VserverClient.GetProjectID()
}
