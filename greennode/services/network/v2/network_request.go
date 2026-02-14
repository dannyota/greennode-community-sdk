package v2

import "github.com/dannyota/greennode-community-sdk/v2/greennode/services/common"

func NewGetNetworkByIDRequest(networkID string) *GetNetworkByIDRequest {
	opt := new(GetNetworkByIDRequest)
	opt.NetworkID = networkID
	return opt
}

type GetNetworkByIDRequest struct {
	common.NetworkCommon
	common.UserAgent
}

func (r *GetNetworkByIDRequest) AddUserAgent(agent ...string) *GetNetworkByIDRequest {
	r.UserAgent.AddUserAgent(agent...)
	return r
}
