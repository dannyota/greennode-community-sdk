package dns

import (
	"github.com/dannyota/greennode-community-sdk/v2/greennode/client"
	dnsinternalv1 "github.com/dannyota/greennode-community-sdk/v2/greennode/services/dns/internal_system/v1"
	dnsv1 "github.com/dannyota/greennode-community-sdk/v2/greennode/services/dns/v1"
)

func NewVDnsServiceV1(svcClient client.ServiceClient) VDnsServiceV1 {
	return &dnsv1.VDnsServiceV1{
		DnsClient: svcClient,
	}
}

func NewVDnsServiceInternal(svcClient client.ServiceClient) VDnsServiceInternal {
	return &dnsinternalv1.VDnsServiceInternal{
		DnsClient: svcClient,
	}
}
