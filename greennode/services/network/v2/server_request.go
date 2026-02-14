package v2

import "github.com/dannyota/greennode-community-sdk/v2/greennode/services/common"

func NewListAllServersBySecgroupIdRequest(psecgroupId string) IListAllServersBySecgroupIdRequest {
	opt := new(ListAllServersBySecgroupIdRequest)
	opt.SecgroupId = psecgroupId
	return opt
}

type ListAllServersBySecgroupIdRequest struct {
	common.SecgroupCommon
	common.UserAgent
}

func (s *ListAllServersBySecgroupIdRequest) AddUserAgent(pagent ...string) IListAllServersBySecgroupIdRequest {
	s.UserAgent.AddUserAgent(pagent...)
	return s
}
