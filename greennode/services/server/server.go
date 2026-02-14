package server

import (
	"github.com/dannyota/greennode-community-sdk/v2/greennode/client"
	serverv1 "github.com/dannyota/greennode-community-sdk/v2/greennode/services/server/v1"
)

func NewServerServiceInternalV1(psvcClient client.IServiceClient) IServerServiceInternalV1 {
	return &serverv1.ServerServiceInternalV1{
		VServerClient: psvcClient,
	}
}
