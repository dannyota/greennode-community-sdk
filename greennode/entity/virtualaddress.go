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

