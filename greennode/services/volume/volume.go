package volume

import (
	"github.com/dannyota/greennode-community-sdk/v2/greennode/client"
	volumev1 "github.com/dannyota/greennode-community-sdk/v2/greennode/services/volume/v1"
	volumev2 "github.com/dannyota/greennode-community-sdk/v2/greennode/services/volume/v2"
)

func NewVolumeServiceV2(svcClient client.ServiceClient) VolumeServiceV2 {
	return &volumev2.VolumeServiceV2{
		VServerClient: svcClient,
	}
}

func NewVolumeServiceV1(svcClient client.ServiceClient) VolumeServiceV1 {
	return &volumev1.VolumeServiceV1{
		VServerClient: svcClient,
	}
}
