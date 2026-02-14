package dns

import (
	"github.com/dannyota/greennode-community-sdk/v2/greennode/client"
	dnsinternalv1 "github.com/dannyota/greennode-community-sdk/v2/greennode/services/dns/internal_system/v1"
	dnsv1 "github.com/dannyota/greennode-community-sdk/v2/greennode/services/dns/v1"
)

func NewVDnsServiceV1(psvcClient client.IServiceClient) IVDnsServiceV1 {
	return &dnsv1.VDnsServiceV1{
		DnsClient: psvcClient,
	}
}

func NewVDnsServiceInternal(psvcClient client.IServiceClient) IVDnsServiceInternal {
	return &dnsinternalv1.VDnsServiceInternal{
		DnsClient: psvcClient,
	}
}
