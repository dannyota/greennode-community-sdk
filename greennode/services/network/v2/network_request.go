package v2

import lscommon "github.com/dannyota/greennode-community-sdk/v2/greennode/services/common"

func NewGetNetworkByIdRequest(pnetworkId string) IGetNetworkByIdRequest {
	opt := new(GetNetworkByIdRequest)
	opt.NetworkId = pnetworkId
	return opt
}

type GetNetworkByIdRequest struct {
	lscommon.NetworkCommon
	lscommon.UserAgent
}

func (s *GetNetworkByIdRequest) AddUserAgent(pagent ...string) IGetNetworkByIdRequest {
	s.UserAgent.AddUserAgent(pagent...)
	return s
}
