package v1

import "github.com/dannyota/greennode-community-sdk/v2/greennode/services/common"

type GetPortalInfoRequest struct {
	BackEndProjectID string
}

func (s *GetPortalInfoRequest) GetBackEndProjectID() string {
	return s.BackEndProjectID
}

type ListProjectsRequest struct {
	common.UserAgent
}

func (s *ListProjectsRequest) AddUserAgent(agent ...string) IListProjectsRequest {
	s.UserAgent.AddUserAgent(agent...)
	return s
}
