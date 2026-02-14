package v2

import "github.com/dannyota/greennode-community-sdk/v2/greennode/services/common"

func NewGetSubnetByIdRequest(networkId, subnetId string) IGetSubnetByIdRequest {
	opt := new(GetSubnetByIdRequest)
	opt.NetworkId = networkId
	opt.SubnetId = subnetId
	return opt
}

type GetSubnetByIdRequest struct {
	common.UserAgent
	common.SubnetCommon
	common.NetworkCommon
}

func (s *GetSubnetByIdRequest) AddUserAgent(agent ...string) IGetSubnetByIdRequest {
	s.UserAgent.AddUserAgent(agent...)
	return s
}

// --------------------------------------------------------
type SecondarySubnetUpdateBody struct {
	Name string `json:"name"`
	CIDR string `json:"cidr"`
}
type UpdateSubnetBody struct {
	Name                    string                      `json:"name"`
	CIDR                    string                      `json:"cidr"`
	SecondarySubnetRequests []SecondarySubnetUpdateBody `json:"secondarySubnetRequests"`
}

func NewUpdateSubnetByIdRequest(networkId, subnetId string, updateBody *UpdateSubnetBody) IUpdateSubnetByIdRequest {
	opt := new(UpdateSubnetByIdRequest)
	opt.NetworkId = networkId
	opt.SubnetId = subnetId
	opt.UpdateSubnetBody = updateBody
	return opt
}

type UpdateSubnetByIdRequest struct {
	UpdateSubnetBody *UpdateSubnetBody `json:"subnet"`
	common.UserAgent
	common.SubnetCommon
	common.NetworkCommon
}

func (s *UpdateSubnetByIdRequest) ToRequestBody() interface{} {
	return s.UpdateSubnetBody
}

func (s *UpdateSubnetByIdRequest) AddUserAgent(agent ...string) IUpdateSubnetByIdRequest {
	s.UserAgent.AddUserAgent(agent...)
	return s
}
