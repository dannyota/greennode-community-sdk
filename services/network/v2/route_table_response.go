package v2

type routeResp struct {
	ID                   string `json:"id"`
	RouteTableID         string `json:"routeTableId"`
	RoutingType          string `json:"routingType"`
	DestinationCidrBlock string `json:"destinationCidrBlock"`
	Target               string `json:"target"`
	Status               string `json:"status"`
}

type routeTableResp struct {
	ID        string      `json:"id"`
	Name      string      `json:"name"`
	Status    string      `json:"status"`
	NetworkID string      `json:"networkId"`
	CreatedAt string      `json:"createdAt"`
	Routes    []routeResp `json:"routes"`
}

func (r *routeTableResp) toEntity() *RouteTable {
	routes := make([]*Route, 0, len(r.Routes))
	for _, rt := range r.Routes {
		routes = append(routes, &Route{
			ID:                   rt.ID,
			RouteTableID:         rt.RouteTableID,
			RoutingType:          rt.RoutingType,
			DestinationCidrBlock: rt.DestinationCidrBlock,
			Target:               rt.Target,
			Status:               rt.Status,
		})
	}
	return &RouteTable{
		ID:        r.ID,
		Name:      r.Name,
		Status:    r.Status,
		NetworkID: r.NetworkID,
		CreatedAt: r.CreatedAt,
		Routes:    routes,
	}
}

type ListRouteTablesResponse struct {
	ListData  []routeTableResp `json:"listData"`
	Page      int              `json:"page"`
	PageSize  int              `json:"pageSize"`
	TotalPage int              `json:"totalPage"`
	TotalItem int              `json:"totalItem"`
}

func (r *ListRouteTablesResponse) ToEntityListRouteTables() *ListRouteTables {
	result := &ListRouteTables{
		Page:      r.Page,
		PageSize:  r.PageSize,
		TotalPage: r.TotalPage,
		TotalItem: r.TotalItem,
	}
	for i := range r.ListData {
		result.Items = append(result.Items, r.ListData[i].toEntity())
	}
	return result
}
