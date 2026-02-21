package v2

type GetNetworkByIDResponse struct {
	Status      string   `json:"status"`
	ElasticIps  []string `json:"elasticIps"`
	DisplayName string   `json:"displayName"`
	ID          string   `json:"id"`
	CreatedAt   string   `json:"createdAt"`
	Cidr        string   `json:"cidr"`
}

func (r *GetNetworkByIDResponse) ToEntityNetwork() *Network {
	return &Network{
		Status:     r.Status,
		ElasticIps: r.ElasticIps,
		Name:       r.DisplayName,
		ID:         r.ID,
		CreatedAt:  r.CreatedAt,
		Cidr:       r.Cidr,
	}
}
