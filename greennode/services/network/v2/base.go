package v2

import "github.com/dannyota/greennode-community-sdk/v2/greennode/client"

type NetworkServiceV2 struct {
	VServerClient *client.ServiceClient
}

func (s *NetworkServiceV2) getProjectID() string {
	return s.VServerClient.GetProjectID()
}
