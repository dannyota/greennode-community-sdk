package v2

type VirtualAddressDataResponse struct {
	Name           string   `json:"name"`
	Description    string   `json:"description"`
	NetworkName    string   `json:"networkName"`
	NetworkCIDR    string   `json:"networkCIDR"`
	SubnetCIDR     string   `json:"subnetCIDR"`
	SubnetName     string   `json:"subnetName"`
	UUID           string   `json:"uuid"`
	NetworkID      string   `json:"networkId"`
	SubnetID       string   `json:"subnetId"`
	IPAddress      string   `json:"ipAddress"`
	CreatedAt      string   `json:"createdAt"`
	AddressPairIps []string `json:"addressPairIps"`
}

func (r *VirtualAddressDataResponse) toEntityVirtualAddress() *VirtualAddress {
	return &VirtualAddress{
		ID:              r.UUID,
		Name:            r.Name,
		EndpointAddress: r.IPAddress,
		VpcID:           r.NetworkID,
		SubnetID:        r.SubnetID,
		Description:     r.Description,
		SubnetCidr:      r.SubnetCIDR,
		VpcCidr:         r.NetworkCIDR,
		AddressPairIps:  r.AddressPairIps,
	}
}

// Response struct for API create virtual address cross project

type CreateVirtualAddressCrossProjectResponse struct {
	Data VirtualAddressDataResponse `json:"data"`
}

func (r *CreateVirtualAddressCrossProjectResponse) ToEntityVirtualAddress() *VirtualAddress {
	return r.Data.toEntityVirtualAddress()
}

// Response struct for API get virtual address by ID
type GetVirtualAddressByIDResponse struct {
	Data VirtualAddressDataResponse `json:"data"`
}

func (r *GetVirtualAddressByIDResponse) ToEntityVirtualAddress() *VirtualAddress {
	return r.Data.toEntityVirtualAddress()
}

// Response struct for API list address pair by virtual address ID
type ListAddressPairsByVirtualAddressIDResponse struct {
	Data []AddressPairResponse `json:"data"`
}

func (r *ListAddressPairsByVirtualAddressIDResponse) ToEntityListAddressPairs() *ListAddressPairs {
	addressPairs := make([]*AddressPair, 0, len(r.Data))
	for _, addressPair := range r.Data {
		addressPairs = append(addressPairs, &AddressPair{
			ID:                 addressPair.UUID,
			VirtualIPAddressID: addressPair.VirtualIPAddressID,
			VirtualSubnetID:    addressPair.VirtualSubnetID,
			NetworkInterfaceIP: addressPair.NetworkInterfaceIP,
			NetworkInterfaceID: addressPair.NetworkInterfaceID,
			CIDR:               addressPair.CIDR,
		})
	}
	return &ListAddressPairs{Items: addressPairs}
}
