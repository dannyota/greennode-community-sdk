package v2

import lsclient "github.com/dannyota/greennode-community-sdk/v2/greennode/client"

type ComputeServiceV2 struct {
	VServerClient lsclient.IServiceClient
}

func (s *ComputeServiceV2) getProjectId() string {
	return s.VServerClient.GetProjectId()
}

const (
	defaultOffsetListServerGroups = 0
	defaultLimitListServerGroups  = 10
)
