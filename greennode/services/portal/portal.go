package portal

import (
	"github.com/dannyota/greennode-community-sdk/v2/greennode/client"
	portalv1 "github.com/dannyota/greennode-community-sdk/v2/greennode/services/portal/v1"
	portalv2 "github.com/dannyota/greennode-community-sdk/v2/greennode/services/portal/v2"
)

func NewPortalServiceV1(psvcClient client.ServiceClient) PortalServiceV1 {
	return &portalv1.PortalServiceV1{
		PortalClient: psvcClient,
	}
}

func NewPortalServiceV2(psvcClient client.ServiceClient) PortalServiceV2 {
	return &portalv2.PortalServiceV2{
		PortalClient: psvcClient,
	}
}
