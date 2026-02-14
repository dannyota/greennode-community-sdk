package v2

import "github.com/dannyota/greennode-community-sdk/v2/greennode/services/common"

type IGetSecgroupByIDRequest interface {
	AddUserAgent(agent ...string) IGetSecgroupByIDRequest
	ParseUserAgent() string
	GetSecgroupID() string
}

type ICreateSecgroupRequest interface {
	ToRequestBody() any
	GetSecgroupName() string
	AddUserAgent(agent ...string) ICreateSecgroupRequest
	ParseUserAgent() string
}

type IDeleteSecgroupByIDRequest interface {
	GetSecgroupID() string
	AddUserAgent(agent ...string) IDeleteSecgroupByIDRequest
	ParseUserAgent() string
}

type IListSecgroupRequest interface {
	AddUserAgent(agent ...string) IListSecgroupRequest
	ParseUserAgent() string
}

type ListSecgroupRequest struct {
	common.UserAgent
}

func (r *ListSecgroupRequest) AddUserAgent(agent ...string) IListSecgroupRequest {
	r.UserAgent.AddUserAgent(agent...)
	return r
}

type DeleteSecgroupByIDRequest struct { //__________________________________________________________________________________
	common.UserAgent
	common.SecgroupCommon
}

func (r *DeleteSecgroupByIDRequest) AddUserAgent(agent ...string) IDeleteSecgroupByIDRequest {
	r.UserAgent.AddUserAgent(agent...)
	return r
}

type CreateSecgroupRequest struct { //__________________________________________________________________________________
	Name        string `json:"name"`
	Description string `json:"description"`

	common.UserAgent
}

func (r *CreateSecgroupRequest) ToRequestBody() any {
	return r
}

func (r *CreateSecgroupRequest) AddUserAgent(agent ...string) ICreateSecgroupRequest {
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

func (r *GetSecgroupByIDRequest) AddUserAgent(agent ...string) IGetSecgroupByIDRequest {
	r.UserAgent.AddUserAgent(agent...)
	return r
}
