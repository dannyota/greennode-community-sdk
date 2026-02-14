package inter

import lsentity "github.com/dannyota/greennode-community-sdk/v2/greennode/entity"

type CreateLoadBalancerResponse struct {
	UUID string `json:"uuid"`
}

func (s *CreateLoadBalancerResponse) ToEntityLoadBalancer() *lsentity.LoadBalancer {
	return &lsentity.LoadBalancer{
		UUID: s.UUID,
	}
}
