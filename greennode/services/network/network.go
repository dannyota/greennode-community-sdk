package network

import (
	lsclient "github.com/dannyota/greennode-community-sdk/v2/greennode/client"
	lsnetworkSvcV1 "github.com/dannyota/greennode-community-sdk/v2/greennode/services/network/v1"
	lsnetworkSvcV2 "github.com/dannyota/greennode-community-sdk/v2/greennode/services/network/v2"
)

func NewNetworkServiceV2(psvcClient lsclient.IServiceClient) INetworkServiceV2 {
	return &lsnetworkSvcV2.NetworkServiceV2{
		VserverClient: psvcClient,
	}
}

func NewNetworkServiceV1(psvcClient lsclient.IServiceClient) INetworkServiceV1 {
	return &lsnetworkSvcV1.NetworkServiceV1{
		VNetworkClient: psvcClient,
	}
}

func NewNetworkServiceInternalV1(psvcClient lsclient.IServiceClient) INetworkServiceInternalV1 {
	return &lsnetworkSvcV1.NetworkServiceInternalV1{
		VNetworkClient: psvcClient,
	}
}
