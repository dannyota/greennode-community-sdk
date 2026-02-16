package v2

func NewGetSubnetByIDRequest(networkID, subnetID string) *GetSubnetByIDRequest {
	return &GetSubnetByIDRequest{
		SubnetID:  subnetID,
		NetworkID: networkID,
	}
}

type GetSubnetByIDRequest struct {
	SubnetID  string
	NetworkID string
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
		SubnetID:         subnetID,
		NetworkID:        networkID,
	}
}

type UpdateSubnetByIDRequest struct {
	UpdateSubnetBody *UpdateSubnetBody `json:"subnet"`
	SubnetID         string
	NetworkID        string
}
