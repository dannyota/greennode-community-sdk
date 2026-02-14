package v2

import lsentity "github.com/dannyota/greennode-community-sdk/v2/greennode/entity"

type GetNetworkByIdResponse struct {
	Status      string   `json:"status"`
	ElasticIps  []string `json:"elasticIps"`
	DisplayName string   `json:"displayName"`
	ID          string   `json:"id"`
	CreatedAt   string   `json:"createdAt"`
	Cidr        string   `json:"cidr"`
}

func (s *GetNetworkByIdResponse) ToEntityNetwork() *lsentity.Network {
	return &lsentity.Network{
		Status:     s.Status,
		ElasticIps: s.ElasticIps,
		Name:       s.DisplayName,
		Id:         s.ID,
		CreatedAt:  s.CreatedAt,
		Cidr:       s.Cidr,
	}
}
