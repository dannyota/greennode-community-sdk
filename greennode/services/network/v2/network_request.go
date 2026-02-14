package v2

import "github.com/dannyota/greennode-community-sdk/v2/greennode/services/common"

func NewGetNetworkByIDRequest(networkID string) IGetNetworkByIDRequest {
	opt := new(GetNetworkByIDRequest)
	opt.NetworkID = networkID
	return opt
}

type GetNetworkByIDRequest struct {
	common.NetworkCommon
	common.UserAgent
}

func (s *GetNetworkByIDRequest) AddUserAgent(agent ...string) IGetNetworkByIDRequest {
	s.UserAgent.AddUserAgent(agent...)
	return s
}
