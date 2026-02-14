package compute

import (
	lsclient "github.com/dannyota/greennode-community-sdk/v2/greennode/client"
	lscomputeSvcV2 "github.com/dannyota/greennode-community-sdk/v2/greennode/services/compute/v2"
)

func NewComputeServiceV2(psvcClient lsclient.IServiceClient) IComputeServiceV2 {
	return &lscomputeSvcV2.ComputeServiceV2{
		VServerClient: psvcClient,
	}
}
