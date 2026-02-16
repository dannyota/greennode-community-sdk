package v2

import "github.com/dannyota/greennode-community-sdk/v2/greennode/services/common"

func NewListAllServersBySecgroupIDRequest(secgroupID string) *ListAllServersBySecgroupIDRequest {
	return &ListAllServersBySecgroupIDRequest{
		SecgroupCommon: common.SecgroupCommon{
			SecgroupID: secgroupID,
		},
	}
}

type ListAllServersBySecgroupIDRequest struct {
	common.SecgroupCommon
}
