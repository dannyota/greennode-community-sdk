package v1

type EndpointCategory struct {
	ID        string
	Name      string
	IsDefault bool
}

type EndpointServiceDetail struct {
	EndpointAuthURL    string
	EndpointURL        string
	TargetCIDR         string
	EndpointEncryptURL string
}

type EndpointService struct {
	ID           string
	Name         string
	EndpointURL  string
	EndpointType string
	Detail       *EndpointServiceDetail
}

type EndpointVPC struct {
	ID        string
	Name      string
	CIDR      string
	Status    string
	DnsStatus string
}

type EndpointSubnet struct {
	ID     string
	Name   string
	Status string
	CIDR   string
	ZoneID string
}

type EndpointPackage struct {
	ID          string
	Name        string
	Description string
}

type EndpointProject struct {
	ID               string
	BackendProjectID string
	PortalUserID     int
	VServerProjectID string
}

type Endpoint struct {
	ID                string
	Name              string
	IPv4Address       string
	EndpointURL       string
	EndpointAuthURL   string
	Status            string
	VpcID             string
	EndpointServiceID string
	BillingStatus     string
	EndpointType      string
	Version           string
	Description       string
	CreatedAt         string
	UpdatedAt         string
	ZoneUuid          string
	EnableDnsName     bool
	EndpointDomains   []string
	Category          *EndpointCategory
	Service           *EndpointService
	VPC               *EndpointVPC
	Subnet            *EndpointSubnet
	Package           *EndpointPackage
	Project           *EndpointProject
}

func (e Endpoint) IsUsable() bool {
	return e.Status == "ACTIVE"
}

func (e Endpoint) IsError() bool {
	return e.Status == "ERROR"
}

type ListEndpoints struct {
	Items     []*Endpoint
	Page      int
	PageSize  int
	TotalPage int
	TotalItem int
}

func (l ListEndpoints) Len() int {
	return len(l.Items)
}

func (l ListEndpoints) At(idx int) *Endpoint {
	if idx < 0 || idx >= l.Len() {
		return nil
	}

	return l.Items[idx]
}

type VNetworkRegion struct {
	ID              string
	Name            string
	Code            string
	GatewayURL      string
	DashboardURL    string
	VServerEndpoint string
}

type ListVNetworkRegions struct {
	Items []*VNetworkRegion
}
