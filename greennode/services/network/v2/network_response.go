package v2

import "github.com/dannyota/greennode-community-sdk/v2/greennode/entity"

type GetNetworkByIdResponse struct {
	Status      string   `json:"status"`
	ElasticIps  []string `json:"elasticIps"`
	DisplayName string   `json:"displayName"`
	ID          string   `json:"id"`
	CreatedAt   string   `json:"createdAt"`
	Cidr        string   `json:"cidr"`
}

func (s *GetNetworkByIdResponse) ToEntityNetwork() *entity.Network {
	return &entity.Network{
		Status:     s.Status,
		ElasticIps: s.ElasticIps,
		Name:       s.DisplayName,
		Id:         s.ID,
		CreatedAt:  s.CreatedAt,
		Cidr:       s.Cidr,
	}
}
