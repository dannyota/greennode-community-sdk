package entity

type AddressPair struct {
	ID                 string `json:"uuid"`
	VirtualIPAddressID string `json:"virtualIpAddressId"`
	VirtualSubnetID    string `json:"virtualSubnetId"`
	NetworkInterfaceIP string `json:"networkInterfaceIp"`
	NetworkInterfaceID string `json:"networkInterfaceId"`
	CIDR               string `json:"cidr"`
}

type ListAddressPairs struct {
	Items []*AddressPair
}

func (l ListAddressPairs) Len() int {
	return len(l.Items)
}

func (l ListAddressPairs) At(idx int) *AddressPair {
	if idx < 0 || idx >= len(l.Items) {
		return nil
	}

	return l.Items[idx]
}
