package v2

import "github.com/dannyota/greennode-community-sdk/v2/greennode/services/common"

func NewListAllServersBySecgroupIDRequest(secgroupID string) *ListAllServersBySecgroupIDRequest {
	opt := new(ListAllServersBySecgroupIDRequest)
	opt.SecgroupID = secgroupID
	return opt
}

type ListAllServersBySecgroupIDRequest struct {
	common.SecgroupCommon
}
