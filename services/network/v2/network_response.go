package v2

type GetNetworkByIDResponse struct {
	Status      string   `json:"status"`
	ElasticIps  []string `json:"elasticIps"`
	DisplayName string   `json:"displayName"`
	ID          string   `json:"id"`
	CreatedAt   string   `json:"createdAt"`
	Cidr        string   `json:"cidr"`
}

func (r *GetNetworkByIDResponse) ToEntityNetwork() *Network {
	return &Network{
		Status:     r.Status,
		ElasticIps: r.ElasticIps,
		Name:       r.DisplayName,
		UUID:       r.ID,
		CreatedAt:  r.CreatedAt,
		Cidr:       r.Cidr,
	}
}

type networkZoneResp struct {
	Uuid          string   `json:"uuid"`
	Name          string   `json:"name"`
	ZoneType      string   `json:"zoneType"`
	IsDefault     bool     `json:"isDefault"`
	Description   string   `json:"description"`
	IsEnabled     bool     `json:"isEnabled"`
	OpenstackZone string   `json:"openstackZone"`
	IpRanges      []string `json:"ipRanges"`
}

type networkResp struct {
	Status         string          `json:"status"`
	ElasticIps     []string        `json:"elasticIps"`
	DisplayName    string          `json:"displayName"`
	ID             string          `json:"id"`
	CreatedAt      string          `json:"createdAt"`
	Cidr           string          `json:"cidr"`
	DhcpOptionName string          `json:"dhcpOptionName"`
	DhcpOptionID   string          `json:"dhcpOptionId"`
	RouteTableName string          `json:"routeTableName"`
	RouteTableID   string          `json:"routeTableId"`
	Zone           networkZoneResp `json:"zone"`
	DnsStatus      string          `json:"dnsStatus"`
	DnsID          string          `json:"dnsId"`
}

func (r *networkResp) toEntityNetwork() *Network {
	return &Network{
		Status:         r.Status,
		ElasticIps:     r.ElasticIps,
		Name:           r.DisplayName,
		UUID:           r.ID,
		CreatedAt:      r.CreatedAt,
		Cidr:           r.Cidr,
		DhcpOptionName: r.DhcpOptionName,
		DhcpOptionID:   r.DhcpOptionID,
		RouteTableName: r.RouteTableName,
		RouteTableID:   r.RouteTableID,
		Zone: NetworkZone{
			Uuid:          r.Zone.Uuid,
			Name:          r.Zone.Name,
			ZoneType:      r.Zone.ZoneType,
			IsDefault:     r.Zone.IsDefault,
			Description:   r.Zone.Description,
			IsEnabled:     r.Zone.IsEnabled,
			OpenstackZone: r.Zone.OpenstackZone,
			IpRanges:      r.Zone.IpRanges,
		},
		DnsStatus: r.DnsStatus,
		DnsID:     r.DnsID,
	}
}

type ListNetworksResponse struct {
	ListData  []networkResp `json:"listData"`
	Page      int           `json:"page"`
	PageSize  int           `json:"pageSize"`
	TotalPage int           `json:"totalPage"`
	TotalItem int           `json:"totalItem"`
}

func (r *ListNetworksResponse) ToEntityListNetworks() *ListNetworks {
	result := &ListNetworks{
		Page:      r.Page,
		PageSize:  r.PageSize,
		TotalPage: r.TotalPage,
		TotalItem: r.TotalItem,
	}
	for i := range r.ListData {
		result.Items = append(result.Items, r.ListData[i].toEntityNetwork())
	}
	return result
}
