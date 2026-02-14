package v2

type GetQuotaByNameRequest struct {
	Name QuotaName
}

func NewGetQuotaByNameRequest(name QuotaName) IGetQuotaByNameRequest {
	return &GetQuotaByNameRequest{
		Name: name,
	}
}

func (s *GetQuotaByNameRequest) GetName() QuotaName {
	return s.Name
}
