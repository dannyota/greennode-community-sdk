package v2

import "github.com/dannyota/greennode-community-sdk/v2/greennode/services/common"

func NewGetNetworkByIDRequest(networkID string) *GetNetworkByIDRequest {
	return &GetNetworkByIDRequest{
		NetworkCommon: common.NetworkCommon{
			NetworkID: networkID,
		},
	}
}

type GetNetworkByIDRequest struct {
	common.NetworkCommon
}
