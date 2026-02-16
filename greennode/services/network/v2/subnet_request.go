package v2

import "github.com/dannyota/greennode-community-sdk/v2/greennode/services/common"

func NewGetSubnetByIDRequest(networkID, subnetID string) *GetSubnetByIDRequest {
	return &GetSubnetByIDRequest{
		SubnetCommon: common.SubnetCommon{
			SubnetID: subnetID,
		},
		NetworkCommon: common.NetworkCommon{
			NetworkID: networkID,
		},
	}
}

type GetSubnetByIDRequest struct {
	common.SubnetCommon
	common.NetworkCommon
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
	return &UpdateSubnetByIDRequest{
		UpdateSubnetBody: updateBody,
		SubnetCommon: common.SubnetCommon{
			SubnetID: subnetID,
		},
		NetworkCommon: common.NetworkCommon{
			NetworkID: networkID,
		},
	}
}

type UpdateSubnetByIDRequest struct {
	UpdateSubnetBody *UpdateSubnetBody `json:"subnet"`
	common.SubnetCommon
	common.NetworkCommon
}
