package v1

type vnetworkRegionResp struct {
	Uuid              string `json:"uuid,omitempty"`
	Name              string `json:"name,omitempty"`
	Code              string `json:"code,omitempty"`
	GatewayURL        string `json:"gatewayUrl,omitempty"`
	VnetworkDashboard string `json:"vnetworkDashboard,omitempty"`
	VServerEndpoint   string `json:"vserverEndpoint,omitempty"`
}

type ListVNetworkRegionsResponse struct {
	Data []vnetworkRegionResp `json:"data"`
}

func (r *ListVNetworkRegionsResponse) ToEntityListVNetworkRegions() *ListVNetworkRegions {
	items := make([]*VNetworkRegion, 0, len(r.Data))
	for _, item := range r.Data {
		items = append(items, &VNetworkRegion{
			ID:              item.Uuid,
			Name:            item.Name,
			Code:            item.Code,
			GatewayURL:      item.GatewayURL,
			DashboardURL:    item.VnetworkDashboard,
			VServerEndpoint: item.VServerEndpoint,
		})
	}
	return &ListVNetworkRegions{
		Items: items,
	}
}
