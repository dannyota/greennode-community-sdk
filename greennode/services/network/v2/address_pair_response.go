package v2

type GetAllAddressPairByVirtualSubnetIDResponse struct {
	Data []*AddressPairResponse `json:"data"`
}

type AddressPairResponse struct {
	UUID               string `json:"uuid"`
	VirtualIPAddressID string `json:"virtualIpAddressId"`
	VirtualSubnetID    string `json:"virtualSubnetId"`
	NetworkInterfaceIP string `json:"networkInterfaceIp"`
	NetworkInterfaceID string `json:"networkInterfaceId"`

	CIDR string `json:"cidr"`
}

func (r *AddressPairResponse) toEntityAddressPair() *AddressPair {
	return &AddressPair{
		ID:                 r.UUID,
		VirtualIPAddressID: r.VirtualIPAddressID,
		VirtualSubnetID:    r.VirtualSubnetID,
		NetworkInterfaceIP: r.NetworkInterfaceIP,
		NetworkInterfaceID: r.NetworkInterfaceID,
		CIDR:               r.CIDR,
	}
}

func (r *GetAllAddressPairByVirtualSubnetIDResponse) ToListAddressPair() []*AddressPair {
	addressPairs := make([]*AddressPair, 0, len(r.Data))
	for _, addressPair := range r.Data {
		addressPairs = append(addressPairs, addressPair.toEntityAddressPair())
	}
	return addressPairs
}

type SetAddressPairInVirtualSubnetResponse struct {
	Data *AddressPairResponse `json:"data"`
}

func (r *SetAddressPairInVirtualSubnetResponse) ToAddressPair() *AddressPair {
	return r.Data.toEntityAddressPair()
}

type CreateAddressPairResponse struct {
	Data AddressPairResponse `json:"data"`
}

func (r *CreateAddressPairResponse) ToAddressPair() *AddressPair {
	return r.Data.toEntityAddressPair()
}
