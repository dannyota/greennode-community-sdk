package v2

import "github.com/dannyota/greennode-community-sdk/v2/greennode/services/common"

type IGetNetworkByIDRequest interface {
	GetNetworkID() string
	AddUserAgent(agent ...string) IGetNetworkByIDRequest
	ParseUserAgent() string
}

func NewGetNetworkByIDRequest(networkID string) IGetNetworkByIDRequest {
	opt := new(GetNetworkByIDRequest)
	opt.NetworkID = networkID
	return opt
}

type GetNetworkByIDRequest struct {
	common.NetworkCommon
	common.UserAgent
}

func (r *GetNetworkByIDRequest) AddUserAgent(agent ...string) IGetNetworkByIDRequest {
	r.UserAgent.AddUserAgent(agent...)
	return r
}
