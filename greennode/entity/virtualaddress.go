package entity

type VirtualAddress struct {
	ID              string
	Name            string
	EndpointAddress string
	VpcID           string
	SubnetID        string
	Description     string
	SubnetCidr      string
	VpcCidr         string
	AddressPairIps  []string
}

func (va *VirtualAddress) GetID() string {
	return va.ID
}

func (va *VirtualAddress) GetName() string {
	return va.Name
}

func (va *VirtualAddress) GetEndpointAddress() string {
	return va.EndpointAddress
}

func (va *VirtualAddress) GetVpcID() string {
	return va.VpcID
}

func (va *VirtualAddress) GetSubnetID() string {
	return va.SubnetID
}

func (va *VirtualAddress) GetDescription() string {
	return va.Description
}

func (va *VirtualAddress) GetSubnetCidr() string {
	return va.SubnetCidr
}

func (va *VirtualAddress) GetVpcCidr() string {
	return va.VpcCidr
}

func (va *VirtualAddress) GetAddressPairIps() []string {
	return va.AddressPairIps
}
