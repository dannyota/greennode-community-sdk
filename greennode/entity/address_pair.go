package entity

type AddressPair struct {
	ID                 string
	VirtualIpAddressID string
	VirtualSubnetID    string
	NetworkInterfaceIp string
	NetworkInterfaceID string
	CIDR               string
}

func (s *AddressPair) GetID() string {
	return s.ID
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
