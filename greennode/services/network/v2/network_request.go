package v2

import "github.com/dannyota/greennode-community-sdk/v2/greennode/services/common"

func NewGetNetworkByIdRequest(networkId string) IGetNetworkByIdRequest {
	opt := new(GetNetworkByIdRequest)
	opt.NetworkId = networkId
	return opt
}

type GetNetworkByIdRequest struct {
	common.NetworkCommon
	common.UserAgent
}

func (s *GetNetworkByIdRequest) AddUserAgent(agent ...string) IGetNetworkByIdRequest {
	s.UserAgent.AddUserAgent(agent...)
	return s
}
