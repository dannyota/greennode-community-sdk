package v2

import "github.com/dannyota/greennode-community-sdk/v2/greennode/services/common"

type ListSecgroupRequest struct {
	common.UserAgent
}

func (s *ListSecgroupRequest) AddUserAgent(pagent ...string) IListSecgroupRequest {
	s.UserAgent.AddUserAgent(pagent...)
	return s
}

type DeleteSecgroupByIdRequest struct { //__________________________________________________________________________________
	common.UserAgent
	common.SecgroupCommon
}

func (s *DeleteSecgroupByIdRequest) AddUserAgent(pagent ...string) IDeleteSecgroupByIdRequest {
	s.UserAgent.AddUserAgent(pagent...)
	return s
}

type CreateSecgroupRequest struct { //__________________________________________________________________________________
	Name        string `json:"name"`
	Description string `json:"description"`

	common.UserAgent
}

func (s *CreateSecgroupRequest) ToRequestBody() interface{} {
	return s
}

func (s *CreateSecgroupRequest) AddUserAgent(pagent ...string) ICreateSecgroupRequest {
	s.UserAgent.AddUserAgent(pagent...)
	return s
}

func (s *CreateSecgroupRequest) GetSecgroupName() string {
	return s.Name
}

type GetSecgroupByIdRequest struct { //_________________________________________________________________________________
	common.SecgroupCommon
	common.UserAgent
}

func (s *GetSecgroupByIdRequest) AddUserAgent(pagent ...string) IGetSecgroupByIdRequest {
	s.UserAgent.AddUserAgent(pagent...)
	return s
}
