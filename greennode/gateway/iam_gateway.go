package gateway

import (
	"github.com/dannyota/greennode-community-sdk/v2/greennode/client"
	"github.com/dannyota/greennode-community-sdk/v2/greennode/services/identity"
)

type iamGatewayV2 struct {
	identityService identity.IdentityServiceV2
}

func NewIamGatewayV2(svcClient client.ServiceClient) IamGatewayV2 {
	return &iamGatewayV2{
		identityService: identity.NewIdentityService(svcClient),
	}
}

func (s *iamGatewayV2) IdentityService() identity.IdentityServiceV2 {
	return s.identityService
}
