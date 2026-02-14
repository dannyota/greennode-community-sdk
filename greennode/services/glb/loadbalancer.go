package glb

import (
	"github.com/dannyota/greennode-community-sdk/v2/greennode/client"
	v1 "github.com/dannyota/greennode-community-sdk/v2/greennode/services/glb/v1"
)

func NewGLBServiceV1(svcClient client.ServiceClient) GLBServiceV1 {
	return &v1.GLBServiceV1{
		VLBClient: svcClient,
	}
}
