package entity

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
