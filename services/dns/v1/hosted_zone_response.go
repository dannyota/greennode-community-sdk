package v1

type GetHostedZoneByIDResponse struct {
	Data *HostedZone `json:"data"`
}

func (r *GetHostedZoneByIDResponse) ToEntityHostedZone() *HostedZone {
	return r.Data
}

type ListHostedZonesResponse struct {
	ListData  []*HostedZone `json:"listData"`
	Page      int           `json:"page"`
	PageSize  int           `json:"pageSize"`
	TotalPage int           `json:"totalPage"`
	TotalItem int           `json:"totalItem"`
}

func (r *ListHostedZonesResponse) ToEntityListHostedZones() *ListHostedZones {
	return &ListHostedZones{
		ListData:  r.ListData,
		Page:      r.Page,
		PageSize:  r.PageSize,
		TotalPage: r.TotalPage,
		TotalItem: r.TotalItem,
	}
}

type CreateHostedZoneResponse struct {
	Data *HostedZone `json:"data"`
}

func (r *CreateHostedZoneResponse) ToEntityHostedZone() *HostedZone {
	return r.Data
}
