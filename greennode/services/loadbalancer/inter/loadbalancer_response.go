package inter

import "github.com/dannyota/greennode-community-sdk/v2/greennode/entity"

type CreateLoadBalancerResponse struct {
	UUID string `json:"uuid"`
}

func (s *CreateLoadBalancerResponse) ToEntityLoadBalancer() *entity.LoadBalancer {
	return &entity.LoadBalancer{
		UUID: s.UUID,
	}
}
