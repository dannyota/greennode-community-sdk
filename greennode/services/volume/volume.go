package volume

import (
	"github.com/dannyota/greennode-community-sdk/v2/greennode/client"
	volumev1 "github.com/dannyota/greennode-community-sdk/v2/greennode/services/volume/v1"
	volumev2 "github.com/dannyota/greennode-community-sdk/v2/greennode/services/volume/v2"
)

func NewVolumeServiceV2(psvcClient client.ServiceClient) VolumeServiceV2 {
	return &volumev2.VolumeServiceV2{
		VServerClient: psvcClient,
	}
}

func NewVolumeServiceV1(psvcClient client.ServiceClient) VolumeServiceV1 {
	return &volumev1.VolumeServiceV1{
		VServerClient: psvcClient,
	}
}
