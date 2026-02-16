package v2

type Network struct {
	Status     string   `json:"status"`
	ElasticIps []string `json:"elasticIps"`
	Name       string   `json:"displayName"`
	ID         string   `json:"id"`
	CreatedAt  string   `json:"createdAt"`
	Cidr       string   `json:"cidr"`
}

type Subnet struct {
	ID                     string                 `json:"uuid"`
	NetworkID              string                 `json:"networkUuid"`
	Name                   string                 `json:"name"`
	Status                 string                 `json:"status"`
	Cidr                   string                 `json:"cidr"`
	RouteTableID           string                 `json:"routeTableUuid"`
	InterfaceAclPolicyID   string                 `json:"interfaceAclPolicyUuid"`
	InterfaceAclPolicyName string                 `json:"interfaceAclPolicyName"`
	SecondarySubnets       []SubnetSecondaryRange `json:"secondarySubnets"`
	ZoneID                 string                 `json:"zoneId"`
}

type SubnetSecondaryRange struct {
	ID   string `json:"uuid"`
	Name string `json:"name"`
	Cidr string `json:"cidr"`
}

type Secgroup struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Status      string `json:"status"`
}

type ListSecgroups struct {
	Items []*Secgroup
}

type SecgroupRule struct {
	ID             string `json:"id"`
	SecgroupID     string `json:"secgroupId"`
	Direction      string `json:"direction"`
	EtherType      string `json:"etherType"`
	Protocol       string `json:"protocol"`
	Description    string `json:"description"`
	RemoteIPPrefix string `json:"remoteIpPrefix"`
	PortRangeMax   int    `json:"portRangeMax"`
	PortRangeMin   int    `json:"portRangeMin"`
}

type ListSecgroupRules struct {
	Items []*SecgroupRule
}

func (l ListSecgroupRules) Len() int {
	return len(l.Items)
}

func (l ListSecgroupRules) Get(i int) *SecgroupRule {
	if i < 0 || i >= len(l.Items) {
		return nil
	}
	return l.Items[i]
}

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
