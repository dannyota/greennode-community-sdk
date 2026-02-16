package v1

import (
	dnsv1 "github.com/dannyota/greennode-community-sdk/v2/greennode/services/dns/v1"
)

type ListRecordsResponse struct {
	ListData  []*dnsv1.DnsRecord `json:"listData"`
	Page      int                `json:"page"`
	PageSize  int                `json:"pageSize"`
	TotalPage int                `json:"totalPage"`
	TotalItem int                `json:"totalItem"`
}

func (r *ListRecordsResponse) ToEntityListRecords() *dnsv1.ListDnsRecords {
	return &dnsv1.ListDnsRecords{
		ListData:  r.ListData,
		Page:      r.Page,
		PageSize:  r.PageSize,
		TotalPage: r.TotalPage,
		TotalItem: r.TotalItem,
	}
}

type GetRecordResponse struct {
	Data *dnsv1.DnsRecord `json:"data"`
}

func (r *GetRecordResponse) ToEntityDnsRecord() *dnsv1.DnsRecord {
	return r.Data
}

type CreateDnsRecordResponse struct {
	Data *dnsv1.DnsRecord `json:"data"`
}

func (r *CreateDnsRecordResponse) ToEntityDnsRecord() *dnsv1.DnsRecord {
	return r.Data
}
