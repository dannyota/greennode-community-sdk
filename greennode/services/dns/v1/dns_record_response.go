package v1

type ListRecordsResponse struct {
	ListData  []*DnsRecord `json:"listData"`
	Page      int          `json:"page"`
	PageSize  int          `json:"pageSize"`
	TotalPage int          `json:"totalPage"`
	TotalItem int          `json:"totalItem"`
}

func (r *ListRecordsResponse) ToEntityListRecords() *ListDnsRecords {
	return &ListDnsRecords{
		ListData:  r.ListData,
		Page:      r.Page,
		PageSize:  r.PageSize,
		TotalPage: r.TotalPage,
		TotalItem: r.TotalItem,
	}
}

type GetRecordResponse struct {
	Data *DnsRecord `json:"data"`
}

func (r *GetRecordResponse) ToEntityDnsRecord() *DnsRecord {
	return r.Data
}

type CreateDnsRecordResponse struct {
	Data *DnsRecord `json:"data"`
}

func (r *CreateDnsRecordResponse) ToEntityDnsRecord() *DnsRecord {
	return r.Data
}
