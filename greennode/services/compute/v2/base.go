package v2

import "github.com/dannyota/greennode-community-sdk/v2/greennode/client"

type ComputeServiceV2 struct {
	VServerClient *client.ServiceClient
}

func (s *ComputeServiceV2) getProjectID() string {
	return s.VServerClient.GetProjectID()
}

const (
	defaultOffsetListServerGroups = 0
	defaultLimitListServerGroups  = 10
)
