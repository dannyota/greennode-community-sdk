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

func (r *AddressPairResponse) toAddressPair() *entity.AddressPair {
	return &entity.AddressPair{
		ID:                 r.UUID,
		VirtualIpAddressID: r.VirtualIpAddressID,
		VirtualSubnetID:    r.VirtualSubnetID,
		NetworkInterfaceIp: r.NetworkInterfaceIp,
		NetworkInterfaceID: r.NetworkInterfaceID,
		CIDR:               r.CIDR,
	}
}

func (r *GetAllAddressPairByVirtualSubnetIDResponse) ToListAddressPair() []*entity.AddressPair {
	addressPairs := make([]*entity.AddressPair, 0, len(r.Data))
	for _, addressPair := range r.Data {
		addressPairs = append(addressPairs, addressPair.toAddressPair())
	}
	return addressPairs
}

type SetAddressPairInVirtualSubnetResponse struct {
	Data *AddressPairResponse `json:"data"`
}

func (r *SetAddressPairInVirtualSubnetResponse) ToAddressPair() *entity.AddressPair {
	return r.Data.toAddressPair()
}

type CreateAddressPairResponse struct {
	Data AddressPairResponse `json:"data"`
}

func (r *CreateAddressPairResponse) ToAddressPair() *entity.AddressPair {
	return r.Data.toAddressPair()
}
