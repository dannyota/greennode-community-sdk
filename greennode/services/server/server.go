package server

import (
	lsclient "github.com/dannyota/greennode-community-sdk/v2/greennode/client"
	lsserverSvcV1 "github.com/dannyota/greennode-community-sdk/v2/greennode/services/server/v1"
)

func NewServerServiceInternalV1(psvcClient lsclient.IServiceClient) IServerServiceInternalV1 {
	return &lsserverSvcV1.ServerServiceInternalV1{
		VServerClient: psvcClient,
	}
}
