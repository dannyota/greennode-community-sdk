package v2

import "github.com/dannyota/greennode-community-sdk/v2/greennode/services/common"

type ListSecgroupRequest struct {
	common.UserAgent
}

func (s *ListSecgroupRequest) AddUserAgent(agent ...string) IListSecgroupRequest {
	s.UserAgent.AddUserAgent(agent...)
	return s
}

type DeleteSecgroupByIDRequest struct { //__________________________________________________________________________________
	common.UserAgent
	common.SecgroupCommon
}

func (s *DeleteSecgroupByIDRequest) AddUserAgent(agent ...string) IDeleteSecgroupByIDRequest {
	s.UserAgent.AddUserAgent(agent...)
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

func (s *CreateSecgroupRequest) AddUserAgent(agent ...string) ICreateSecgroupRequest {
	s.UserAgent.AddUserAgent(agent...)
	return s
}

func (s *CreateSecgroupRequest) GetSecgroupName() string {
	return s.Name
}

type GetSecgroupByIDRequest struct { //_________________________________________________________________________________
	common.SecgroupCommon
	common.UserAgent
}

func (s *GetSecgroupByIDRequest) AddUserAgent(agent ...string) IGetSecgroupByIDRequest {
	s.UserAgent.AddUserAgent(agent...)
	return s
}
