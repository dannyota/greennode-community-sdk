package v2

import "github.com/dannyota/greennode-community-sdk/v2/greennode/entity"

type GetAllAddressPairByVirtualSubnetIDResponse struct {
	Data []*AddressPairResponse `json:"data"`
}

type AddressPairResponse struct {
	UUID               string `json:"uuid"`
	VirtualIpAddressID string `json:"virtualIpAddressId"`
	VirtualSubnetID    string `json:"virtualSubnetId"`
	NetworkInterfaceIp string `json:"networkInterfaceIp"`
	NetworkInterfaceID string `json:"networkInterfaceId"`

	CIDR string `json:"cidr"`
	// ID              string `json:"id"`
	// BackendSubnetId int    `json:"backendSubnetId"`
	// ProjectId       string `json:"projectId"`
	// CreatedAt       string `json:"createdAt"`
	// DeletedAt       string `json:"deletedAt"`
}

func (s *AddressPairResponse) toAddressPair() *entity.AddressPair {
	return &entity.AddressPair{
		ID:                 s.UUID,
		VirtualIpAddressID: s.VirtualIpAddressID,
		VirtualSubnetID:    s.VirtualSubnetID,
		NetworkInterfaceIp: s.NetworkInterfaceIp,
		NetworkInterfaceID: s.NetworkInterfaceID,
		CIDR:               s.CIDR,
	}
}

func (s *GetAllAddressPairByVirtualSubnetIDResponse) ToListAddressPair() []*entity.AddressPair {
	addressPairs := make([]*entity.AddressPair, 0, len(s.Data))
	for _, addressPair := range s.Data {
		addressPairs = append(addressPairs, addressPair.toAddressPair())
	}
	return addressPairs
}

type SetAddressPairInVirtualSubnetResponse struct {
	Data *AddressPairResponse `json:"data"`
}

func (s *SetAddressPairInVirtualSubnetResponse) ToAddressPair() *entity.AddressPair {
	return s.Data.toAddressPair()
}

type CreateAddressPairResponse struct {
	Data AddressPairResponse `json:"data"`
}

func (s *CreateAddressPairResponse) ToAddressPair() *entity.AddressPair {
	return s.Data.toAddressPair()
}
