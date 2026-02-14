package v2

import "github.com/dannyota/greennode-community-sdk/v2/greennode/entity"

type GetNetworkByIDResponse struct {
	Status      string   `json:"status"`
	ElasticIps  []string `json:"elasticIps"`
	DisplayName string   `json:"displayName"`
	ID          string   `json:"id"`
	CreatedAt   string   `json:"createdAt"`
	Cidr        string   `json:"cidr"`
}

func (r *GetNetworkByIDResponse) ToEntityNetwork() *entity.Network {
	return &entity.Network{
		Status:     r.Status,
		ElasticIps: r.ElasticIps,
		Name:       r.DisplayName,
		ID:         r.ID,
		CreatedAt:  r.CreatedAt,
		Cidr:       r.Cidr,
	}
}
