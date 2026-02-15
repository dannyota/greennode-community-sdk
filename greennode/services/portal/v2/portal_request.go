package v2

type GetQuotaByNameRequest struct {
	Name QuotaName
}

func NewGetQuotaByNameRequest(name QuotaName) *GetQuotaByNameRequest {
	return &GetQuotaByNameRequest{
		Name: name,
	}
}

func (r *GetQuotaByNameRequest) GetName() QuotaName {
	return r.Name
}
