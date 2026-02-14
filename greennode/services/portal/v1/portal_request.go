package v1

import "github.com/dannyota/greennode-community-sdk/v2/greennode/services/common"

type GetPortalInfoRequest struct {
	BackEndProjectId string
}

func (s *GetPortalInfoRequest) GetBackEndProjectId() string {
	return s.BackEndProjectId
}

type ListProjectsRequest struct {
	common.UserAgent
}

func (s *ListProjectsRequest) AddUserAgent(pagent ...string) IListProjectsRequest {
	s.UserAgent.AddUserAgent(pagent...)
	return s
}
