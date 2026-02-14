package v2

import "github.com/dannyota/greennode-community-sdk/v2/greennode/entity"

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

func (s *VirtualAddressDataResponse) toEntityVirtualAddress() *entity.VirtualAddress {
	return &entity.VirtualAddress{
		ID:              s.UUID,
		Name:            s.Name,
		EndpointAddress: s.IPAddress,
		VpcID:           s.NetworkID,
		SubnetID:        s.SubnetID,
		Description:     s.Description,
		SubnetCidr:      s.SubnetCIDR,
		VpcCidr:         s.NetworkCIDR,
		AddressPairIps:  s.AddressPairIps,
	}
}

// Response struct for API create virtual address cross project

type CreateVirtualAddressCrossProjectResponse struct {
	Data VirtualAddressDataResponse `json:"data"`
}

func (s *CreateVirtualAddressCrossProjectResponse) ToEntityVirtualAddress() *entity.VirtualAddress {
	return s.Data.toEntityVirtualAddress()
}

// Response struct for API get virtual address by ID
type GetVirtualAddressByIDResponse struct {
	Data VirtualAddressDataResponse `json:"data"`
}

func (s *GetVirtualAddressByIDResponse) ToEntityVirtualAddress() *entity.VirtualAddress {
	return s.Data.toEntityVirtualAddress()
}

// Response struct for API list address pair by virtual address ID
type ListAddressPairsByVirtualAddressIDResponse struct {
	Data []AddressPairResponse `json:"data"`
}

func (s *ListAddressPairsByVirtualAddressIDResponse) ToEntityListAddressPairs() *entity.ListAddressPairs {
	addressPairs := make([]*entity.AddressPair, 0, len(s.Data))
	for _, addressPair := range s.Data {
		addressPairs = append(addressPairs, &entity.AddressPair{
			ID:                 addressPair.UUID,
			VirtualIpAddressID: addressPair.VirtualIpAddressID,
			VirtualSubnetID:    addressPair.VirtualSubnetID,
			NetworkInterfaceIp: addressPair.NetworkInterfaceIp,
			NetworkInterfaceID: addressPair.NetworkInterfaceID,
			CIDR:               addressPair.CIDR,
		})
	}
	return &entity.ListAddressPairs{Items: addressPairs}
}
