package v2

type Network struct {
	Status          string      `json:"status"`
	ElasticIps      []string    `json:"elasticIps"`
	Name            string      `json:"displayName"`
	ID              string      `json:"id"`
	CreatedAt       string      `json:"createdAt"`
	Cidr            string      `json:"cidr"`
	DhcpOptionName  string      `json:"dhcpOptionName"`
	DhcpOptionID    string      `json:"dhcpOptionId"`
	RouteTableName  string      `json:"routeTableName"`
	RouteTableID    string      `json:"routeTableId"`
	Zone            NetworkZone `json:"zone"`
	DnsStatus       string      `json:"dnsStatus"`
	DnsID           string      `json:"dnsId"`
}

type NetworkZone struct {
	Uuid          string   `json:"uuid"`
	Name          string   `json:"name"`
	ZoneType      string   `json:"zoneType"`
	IsDefault     bool     `json:"isDefault"`
	Description   string   `json:"description"`
	IsEnabled     bool     `json:"isEnabled"`
	OpenstackZone string   `json:"openstackZone"`
	IpRanges      []string `json:"ipRanges"`
}

type ListNetworks struct {
	Items     []*Network
	Page      int
	PageSize  int
	TotalPage int
	TotalItem int
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

type ListSubnets struct {
	Items []*Subnet
}

type Secgroup struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Status      string `json:"status"`
	CreatedAt   string `json:"createdAt"`
	IsSystem    bool   `json:"isSystem"`
}

type ListSecgroups struct {
	Items     []*Secgroup
	Page      int
	PageSize  int
	TotalPage int
	TotalItem int
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

type Route struct {
	ID                    string `json:"id"`
	RouteTableID          string `json:"routeTableId"`
	RoutingType           string `json:"routingType"`
	DestinationCidrBlock  string `json:"destinationCidrBlock"`
	Target                string `json:"target"`
	Status                string `json:"status"`
}

type RouteTable struct {
	ID        string   `json:"id"`
	Name      string   `json:"name"`
	Status    string   `json:"status"`
	NetworkID string   `json:"networkId"`
	CreatedAt string   `json:"createdAt"`
	Routes    []*Route `json:"routes"`
}

type ListRouteTables struct {
	Items     []*RouteTable
	Page      int
	PageSize  int
	TotalPage int
	TotalItem int
}

type Peering struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Status    string `json:"status"`
	FromVpcID string `json:"fromVpcId"`
	FromCidr  string `json:"fromCidr"`
	EndVpcID  string `json:"endVpcId"`
	EndCidr   string `json:"endCidr"`
	CreatedAt string `json:"createdAt"`
}

type ListPeerings struct {
	Items     []*Peering
	Page      int
	PageSize  int
	TotalPage int
	TotalItem int
}

type Interconnect struct {
	ID           string `json:"id"`
	ProjectID    string `json:"projectId"`
	Name         string `json:"name"`
	Description  string `json:"description"`
	EnableGw2    bool   `json:"enableGw2"`
	CircuitID    int    `json:"circuitId"`
	Gw01IP       string `json:"gw01Ip"`
	Gw02IP       string `json:"gw02Ip"`
	GwVIP        string `json:"gwVip"`
	RemoteGw01IP string `json:"remoteGw01Ip"`
	RemoteGw02IP string `json:"remoteGw02Ip"`
	PackageID    string `json:"packageId"`
	Status       string `json:"status"`
	TypeID       string `json:"typeId"`
	TypeName     string `json:"typeName"`
	CreatedAt    string `json:"createdAt"`
}

type ListInterconnects struct {
	Items     []*Interconnect
	Page      int
	PageSize  int
	TotalPage int
	TotalItem int
}
