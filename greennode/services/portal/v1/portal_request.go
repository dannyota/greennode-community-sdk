package v1

import "github.com/dannyota/greennode-community-sdk/v2/greennode/services/common"

type GetPortalInfoRequest struct {
	BackEndProjectID string
}

func (r *GetPortalInfoRequest) GetBackEndProjectID() string {
	return r.BackEndProjectID
}

type ListProjectsRequest struct {
	common.UserAgent
}

func (r *ListProjectsRequest) AddUserAgent(agent ...string) *ListProjectsRequest {
	r.UserAgent.AddUserAgent(agent...)
	return r
}
