package intervpc

import (
	lbv2 "danny.vn/greennode/greennode/services/loadbalancer/v2"
)

type CreateLoadBalancerResponse struct {
	UUID string `json:"uuid"`
}

func (r *CreateLoadBalancerResponse) ToEntityLoadBalancer() *lbv2.LoadBalancer {
	return &lbv2.LoadBalancer{
		UUID: r.UUID,
	}
}
