package v2

import "github.com/dannyota/greennode-community-sdk/v2/greennode/services/common"

func NewGetNetworkByIdRequest(pnetworkId string) IGetNetworkByIdRequest {
	opt := new(GetNetworkByIdRequest)
	opt.NetworkId = pnetworkId
	return opt
}

type GetNetworkByIdRequest struct {
	common.NetworkCommon
	common.UserAgent
}

func (s *GetNetworkByIdRequest) AddUserAgent(pagent ...string) IGetNetworkByIdRequest {
	s.UserAgent.AddUserAgent(pagent...)
	return s
}
