package gateway

import (
	"github.com/dannyota/greennode-community-sdk/v2/greennode/client"
	identityv2 "github.com/dannyota/greennode-community-sdk/v2/greennode/services/identity/v2"
)

type IAMGatewayV2 struct {
	identityService *identityv2.IdentityServiceV2
}

func NewIAMGatewayV2(svcClient client.ServiceClient) *IAMGatewayV2 {
	return &IAMGatewayV2{
		identityService: &identityv2.IdentityServiceV2{IAMClient: svcClient},
	}
}

func (g *IAMGatewayV2) IdentityService() *identityv2.IdentityServiceV2 {
	return g.identityService
}
