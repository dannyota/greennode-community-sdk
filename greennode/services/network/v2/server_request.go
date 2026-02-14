package v2

import "github.com/dannyota/greennode-community-sdk/v2/greennode/services/common"

func NewListAllServersBySecgroupIdRequest(secgroupId string) IListAllServersBySecgroupIdRequest {
	opt := new(ListAllServersBySecgroupIdRequest)
	opt.SecgroupId = secgroupId
	return opt
}

type ListAllServersBySecgroupIdRequest struct {
	common.SecgroupCommon
	common.UserAgent
}

func (s *ListAllServersBySecgroupIdRequest) AddUserAgent(agent ...string) IListAllServersBySecgroupIdRequest {
	s.UserAgent.AddUserAgent(agent...)
	return s
}
