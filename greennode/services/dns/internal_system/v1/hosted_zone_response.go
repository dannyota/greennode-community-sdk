package v1

import (
	"github.com/dannyota/greennode-community-sdk/v2/greennode/entity"
)

type GetHostedZoneByIDResponse struct {
	Data *entity.HostedZone `json:"data"`
}

func (r *GetHostedZoneByIDResponse) ToEntityHostedZone() *entity.HostedZone {
	return r.Data
}

type ListHostedZonesResponse struct {
	ListData  []*entity.HostedZone `json:"listData"`
	Page      int                  `json:"page"`
	PageSize  int                  `json:"pageSize"`
	TotalPage int                  `json:"totalPage"`
	TotalItem int                  `json:"totalItem"`
}

func (r *ListHostedZonesResponse) ToEntityListHostedZones() *entity.ListHostedZones {
	return &entity.ListHostedZones{
		ListData:  r.ListData,
		Page:      r.Page,
		PageSize:  r.PageSize,
		TotalPage: r.TotalPage,
		TotalItem: r.TotalItem,
	}
}

type CreateHostedZoneResponse struct {
	Data *entity.HostedZone `json:"data"`
}

func (r *CreateHostedZoneResponse) ToEntityHostedZone() *entity.HostedZone {
	return r.Data
}
