package v2

import "github.com/dannyota/greennode-community-sdk/v2/greennode/entity"

type GetSubnetByIDResponse struct {
	UUID                   string `json:"uuid"`
	CreatedAt              string `json:"createdAt"`
	UpdatedAt              string `json:"updatedAt,omitempty"`
	Status                 string `json:"status"`
	Cidr                   string `json:"cidr"`
	NetworkUuid            string `json:"networkUuid"`
	RouteTableUuid         string `json:"routeTableUuid,omitempty"`
	Name                   string `json:"name"`
	InterfaceAclPolicyUuid string `json:"interfaceAclPolicyUuid,omitempty"`
	InterfaceAclPolicyName string `json:"interfaceAclPolicyName,omitempty"`
	SecondarySubnets       []struct {
		UUID string `json:"uuid"`
		Name string `json:"name"`
		Cidr string `json:"cidr"`
	} `json:"secondarySubnets,omitempty"`
	Zone struct {
		UUID string `json:"uuid"`
	} `json:"zone"`
}

func (r *GetSubnetByIDResponse) ToEntitySubnet() *entity.Subnet {
	secondaryRange := make([]entity.SubnetSecondaryRange, 0, len(r.SecondarySubnets))
	for _, sr := range r.SecondarySubnets {
		secondaryRange = append(secondaryRange, entity.SubnetSecondaryRange{
			ID:   sr.UUID,
			Name: sr.Name,
			Cidr: sr.Cidr,
		})
	}
	return &entity.Subnet{
		ID:                     r.UUID,
		NetworkID:              r.NetworkUuid,
		Name:                   r.Name,
		Status:                 r.Status,
		Cidr:                   r.Cidr,
		RouteTableID:           r.RouteTableUuid,
		InterfaceAclPolicyID:   r.InterfaceAclPolicyUuid,
		InterfaceAclPolicyName: r.InterfaceAclPolicyName,
		SecondarySubnets:       secondaryRange,
		ZoneID:                 r.Zone.UUID,
	}
}

type UpdateSubnetByIDResponse struct {
	Data GetSubnetByIDResponse `json:"data"`
}

func (r *UpdateSubnetByIDResponse) ToEntitySubnet() *entity.Subnet {
	return r.Data.ToEntitySubnet()
}
