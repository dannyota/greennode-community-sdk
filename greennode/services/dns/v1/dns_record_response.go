package v1

import (
	"github.com/dannyota/greennode-community-sdk/v2/greennode/entity"
)

type ListRecordsResponse struct {
	ListData  []*entity.DnsRecord `json:"listData"`
	Page      int                 `json:"page"`
	PageSize  int                 `json:"pageSize"`
	TotalPage int                 `json:"totalPage"`
	TotalItem int                 `json:"totalItem"`
}

func (r *ListRecordsResponse) ToEntityListRecords() *entity.ListDnsRecords {
	return &entity.ListDnsRecords{
		ListData:  r.ListData,
		Page:      r.Page,
		PageSize:  r.PageSize,
		TotalPage: r.TotalPage,
		TotalItem: r.TotalItem,
	}
}

type GetRecordResponse struct {
	Data *entity.DnsRecord `json:"data"`
}

func (r *GetRecordResponse) ToEntityDnsRecord() *entity.DnsRecord {
	return r.Data
}

type CreateDnsRecordResponse struct {
	Data *entity.DnsRecord `json:"data"`
}

func (r *CreateDnsRecordResponse) ToEntityDnsRecord() *entity.DnsRecord {
	return r.Data
}
