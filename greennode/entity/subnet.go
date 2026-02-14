package entity

type Subnet struct {
	ID                     string
	NetworkID              string
	Name                   string
	Status                 string
	Cidr                   string
	RouteTableID           string
	InterfaceAclPolicyID   string
	InterfaceAclPolicyName string
	SecondarySubnets       []SubnetSecondaryRange
	ZoneID                 string
}

type SubnetSecondaryRange struct {
	ID   string
	Name string
	Cidr string
}
