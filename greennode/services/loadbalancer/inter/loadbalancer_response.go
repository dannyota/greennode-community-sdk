package inter

import "github.com/dannyota/greennode-community-sdk/v2/greennode/entity"

type CreateLoadBalancerResponse struct {
	UUID string `json:"uuid"`
}

func (r *CreateLoadBalancerResponse) ToEntityLoadBalancer() *entity.LoadBalancer {
	return &entity.LoadBalancer{
		UUID: r.UUID,
	}
}
