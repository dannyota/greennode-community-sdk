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

func (s *GetNetworkByIDResponse) ToEntityNetwork() *entity.Network {
	return &entity.Network{
		Status:     s.Status,
		ElasticIps: s.ElasticIps,
		Name:       s.DisplayName,
		ID:         s.ID,
		CreatedAt:  s.CreatedAt,
		Cidr:       s.Cidr,
	}
}
