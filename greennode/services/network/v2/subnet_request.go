package v2

import "github.com/dannyota/greennode-community-sdk/v2/greennode/services/common"

func NewGetSubnetByIDRequest(networkID, subnetID string) *GetSubnetByIDRequest {
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

func (r *GetSubnetByIDRequest) AddUserAgent(agent ...string) *GetSubnetByIDRequest {
	r.UserAgent.AddUserAgent(agent...)
	return r
}

type SecondarySubnetUpdateBody struct {
	Name string `json:"name"`
	CIDR string `json:"cidr"`
}
type UpdateSubnetBody struct {
	Name                    string                      `json:"name"`
	CIDR                    string                      `json:"cidr"`
	SecondarySubnetRequests []SecondarySubnetUpdateBody `json:"secondarySubnetRequests"`
}

func NewUpdateSubnetByIDRequest(networkID, subnetID string, updateBody *UpdateSubnetBody) *UpdateSubnetByIDRequest {
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

func (r *UpdateSubnetByIDRequest) ToRequestBody() any {
	return r.UpdateSubnetBody
}

func (r *UpdateSubnetByIDRequest) AddUserAgent(agent ...string) *UpdateSubnetByIDRequest {
	r.UserAgent.AddUserAgent(agent...)
	return r
}
