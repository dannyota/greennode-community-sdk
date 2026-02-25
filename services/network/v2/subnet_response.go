package v2

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

func (r *GetSubnetByIDResponse) ToEntitySubnet() *Subnet {
	secondaryRange := make([]SubnetSecondaryRange, 0, len(r.SecondarySubnets))
	for _, sr := range r.SecondarySubnets {
		secondaryRange = append(secondaryRange, SubnetSecondaryRange{
			UUID: sr.UUID,
			Name: sr.Name,
			Cidr: sr.Cidr,
		})
	}
	return &Subnet{
		UUID:                   r.UUID,
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

type ListSubnetsByNetworkIDResponse []GetSubnetByIDResponse

func (r ListSubnetsByNetworkIDResponse) ToEntityListSubnets() *ListSubnets {
	result := &ListSubnets{}
	for i := range r {
		result.Items = append(result.Items, r[i].ToEntitySubnet())
	}
	return result
}

type UpdateSubnetByIDResponse struct {
	Data GetSubnetByIDResponse `json:"data"`
}

func (r *UpdateSubnetByIDResponse) ToEntitySubnet() *Subnet {
	return r.Data.ToEntitySubnet()
}
