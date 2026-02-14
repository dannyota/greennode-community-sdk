package v2

import "github.com/dannyota/greennode-community-sdk/v2/greennode/services/common"

func NewGetSubnetByIDRequest(networkID, subnetID string) IGetSubnetByIDRequest {
	opt := new(GetSubnetByIDRequest)
	opt.NetworkID = networkID
	opt.SubnetID = subnetID
	return opt
}

type GetSubnetByIDRequest struct {
	common.UserAgent
	common.SubnetCommon
	common.NetworkCommon
}

func (s *GetSubnetByIDRequest) AddUserAgent(agent ...string) IGetSubnetByIDRequest {
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

func NewUpdateSubnetByIDRequest(networkID, subnetID string, updateBody *UpdateSubnetBody) IUpdateSubnetByIDRequest {
	opt := new(UpdateSubnetByIDRequest)
	opt.NetworkID = networkID
	opt.SubnetID = subnetID
	opt.UpdateSubnetBody = updateBody
	return opt
}

type UpdateSubnetByIDRequest struct {
	UpdateSubnetBody *UpdateSubnetBody `json:"subnet"`
	common.UserAgent
	common.SubnetCommon
	common.NetworkCommon
}

func (s *UpdateSubnetByIDRequest) ToRequestBody() interface{} {
	return s.UpdateSubnetBody
}

func (s *UpdateSubnetByIDRequest) AddUserAgent(agent ...string) IUpdateSubnetByIDRequest {
	s.UserAgent.AddUserAgent(agent...)
	return s
}
