package gateway

import (
	"github.com/dannyota/greennode-community-sdk/v2/greennode/client"
	"github.com/dannyota/greennode-community-sdk/v2/greennode/services/identity"
)

type IAMGatewayV2 struct {
	identityService identity.IdentityServiceV2
}

func NewIAMGatewayV2(svcClient client.ServiceClient) *IAMGatewayV2 {
	return &IAMGatewayV2{
		identityService: identity.NewIdentityService(svcClient),
	}
}

func (g *IAMGatewayV2) IdentityService() identity.IdentityServiceV2 {
	return g.identityService
}
