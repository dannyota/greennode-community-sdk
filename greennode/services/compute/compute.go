package compute

import (
	"github.com/dannyota/greennode-community-sdk/v2/greennode/client"
	computev2 "github.com/dannyota/greennode-community-sdk/v2/greennode/services/compute/v2"
)

func NewComputeServiceV2(svcClient client.ServiceClient) ComputeServiceV2 {
	return &computev2.ComputeServiceV2{
		VServerClient: svcClient,
	}
}
