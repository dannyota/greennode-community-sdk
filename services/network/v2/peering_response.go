package v2

type peeringResp struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Status    string `json:"status"`
	FromVpcID string `json:"fromVpcId"`
	FromCidr  string `json:"fromCidr"`
	EndVpcID  string `json:"endVpcId"`
	EndCidr   string `json:"endCidr"`
	CreatedAt string `json:"createdAt"`
}

func (r *peeringResp) toEntity() *Peering {
	return &Peering{
		ID:        r.ID,
		Name:      r.Name,
		Status:    r.Status,
		FromVpcID: r.FromVpcID,
		FromCidr:  r.FromCidr,
		EndVpcID:  r.EndVpcID,
		EndCidr:   r.EndCidr,
		CreatedAt: r.CreatedAt,
	}
}

type ListPeeringsResponse struct {
	ListData  []peeringResp `json:"listData"`
	Page      int           `json:"page"`
	PageSize  int           `json:"pageSize"`
	TotalPage int           `json:"totalPage"`
	TotalItem int           `json:"totalItem"`
}

func (r *ListPeeringsResponse) ToEntityListPeerings() *ListPeerings {
	result := &ListPeerings{
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
