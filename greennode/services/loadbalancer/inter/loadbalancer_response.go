package inter

import (
	lbv2 "github.com/dannyota/greennode-community-sdk/v2/greennode/services/loadbalancer/v2"
)

type CreateLoadBalancerResponse struct {
	UUID string `json:"uuid"`
}

func (r *CreateLoadBalancerResponse) ToEntityLoadBalancer() *lbv2.LoadBalancer {
	return &lbv2.LoadBalancer{
		UUID: r.UUID,
	}
}
