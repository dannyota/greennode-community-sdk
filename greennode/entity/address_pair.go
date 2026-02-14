package entity

type AddressPair struct {
	Id                 string
	VirtualIpAddressId string
	VirtualSubnetId    string
	NetworkInterfaceIp string
	NetworkInterfaceId string
	CIDR               string
}

func (s *AddressPair) GetId() string {
	return s.Id
}

type ListAddressPairs struct {
	Items []*AddressPair
}

func (s *ListAddressPairs) Len() int {
	return len(s.Items)
}

func (s *ListAddressPairs) At(idx int) *AddressPair {
	if idx < 0 || idx >= len(s.Items) {
		return nil
	}

	return s.Items[idx]
}
