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

func (s *VirtualAddress) GetID() string {
	return s.ID
}

func (s *VirtualAddress) GetName() string {
	return s.Name
}

func (s *VirtualAddress) GetEndpointAddress() string {
	return s.EndpointAddress
}

func (s *VirtualAddress) GetVpcID() string {
	return s.VpcID
}

func (s *VirtualAddress) GetSubnetID() string {
	return s.SubnetID
}

func (s *VirtualAddress) GetDescription() string {
	return s.Description
}

func (s *VirtualAddress) GetSubnetCidr() string {
	return s.SubnetCidr
}

func (s *VirtualAddress) GetVpcCidr() string {
	return s.VpcCidr
}

func (s *VirtualAddress) GetAddressPairIps() []string {
	return s.AddressPairIps
}
