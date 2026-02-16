package entity

type VirtualAddress struct {
	ID              string   `json:"uuid"`
	Name            string   `json:"name"`
	EndpointAddress string   `json:"ipAddress"`
	VpcID           string   `json:"networkId"`
	SubnetID        string   `json:"subnetId"`
	Description     string   `json:"description"`
	SubnetCidr      string   `json:"subnetCIDR"`
	VpcCidr         string   `json:"networkCIDR"`
	AddressPairIps  []string `json:"addressPairIps"`
}

