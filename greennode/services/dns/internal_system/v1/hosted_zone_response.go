package v1

import (
	dnsv1 "github.com/dannyota/greennode-community-sdk/v2/greennode/services/dns/v1"
)

type GetHostedZoneByIDResponse struct {
	Data *dnsv1.HostedZone `json:"data"`
}

func (r *GetHostedZoneByIDResponse) ToEntityHostedZone() *dnsv1.HostedZone {
	return r.Data
}

type ListHostedZonesResponse struct {
	ListData  []*dnsv1.HostedZone `json:"listData"`
	Page      int                 `json:"page"`
	PageSize  int                 `json:"pageSize"`
	TotalPage int                 `json:"totalPage"`
	TotalItem int                 `json:"totalItem"`
}

func (r *ListHostedZonesResponse) ToEntityListHostedZones() *dnsv1.ListHostedZones {
	return &dnsv1.ListHostedZones{
		ListData:  r.ListData,
		Page:      r.Page,
		PageSize:  r.PageSize,
		TotalPage: r.TotalPage,
		TotalItem: r.TotalItem,
	}
}

type CreateHostedZoneResponse struct {
	Data *dnsv1.HostedZone `json:"data"`
}

func (r *CreateHostedZoneResponse) ToEntityHostedZone() *dnsv1.HostedZone {
	return r.Data
}
