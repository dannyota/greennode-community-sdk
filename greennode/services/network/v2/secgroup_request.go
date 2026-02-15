package v2

import "github.com/dannyota/greennode-community-sdk/v2/greennode/services/common"

type ListSecgroupRequest struct {
	common.UserAgent
}

func (r *ListSecgroupRequest) AddUserAgent(agent ...string) *ListSecgroupRequest {
	r.UserAgent.AddUserAgent(agent...)
	return r
}

type DeleteSecgroupByIDRequest struct { //__________________________________________________________________________________
	common.UserAgent
	common.SecgroupCommon
}

func (r *DeleteSecgroupByIDRequest) AddUserAgent(agent ...string) *DeleteSecgroupByIDRequest {
	r.UserAgent.AddUserAgent(agent...)
	return r
}

type CreateSecgroupRequest struct { //__________________________________________________________________________________
	Name        string `json:"name"`
	Description string `json:"description"`

	common.UserAgent
}

func (r *CreateSecgroupRequest) AddUserAgent(agent ...string) *CreateSecgroupRequest {
	r.UserAgent.AddUserAgent(agent...)
	return r
}

func (r *CreateSecgroupRequest) GetSecgroupName() string {
	return r.Name
}

type GetSecgroupByIDRequest struct { //_________________________________________________________________________________
	common.SecgroupCommon
	common.UserAgent
}

func (r *GetSecgroupByIDRequest) AddUserAgent(agent ...string) *GetSecgroupByIDRequest {
	r.UserAgent.AddUserAgent(agent...)
	return r
}
