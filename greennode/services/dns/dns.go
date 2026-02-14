package dns

import (
	lsclient "github.com/dannyota/greennode-community-sdk/v2/greennode/client"
	lsdnsSvcInternal "github.com/dannyota/greennode-community-sdk/v2/greennode/services/dns/internal_system/v1"
	lsdnsSvcV1 "github.com/dannyota/greennode-community-sdk/v2/greennode/services/dns/v1"
)

func NewVDnsServiceV1(psvcClient lsclient.IServiceClient) IVDnsServiceV1 {
	return &lsdnsSvcV1.VDnsServiceV1{
		DnsClient: psvcClient,
	}
}

func NewVDnsServiceInternal(psvcClient lsclient.IServiceClient) IVDnsServiceInternal {
	return &lsdnsSvcInternal.VDnsServiceInternal{
		DnsClient: psvcClient,
	}
}
