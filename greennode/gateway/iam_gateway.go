package gateway

import (
	"github.com/dannyota/greennode-community-sdk/v2/greennode/client"
	"github.com/dannyota/greennode-community-sdk/v2/greennode/services/identity"
)

type iamGatewayV2 struct {
	identityService identity.IIdentityServiceV2
}

func NewIamGatewayV2(psvcClient client.IServiceClient) IIamGatewayV2 {
	return &iamGatewayV2{
		identityService: identity.NewIdentityService(psvcClient),
	}
}

func (s *iamGatewayV2) IdentityService() identity.IIdentityServiceV2 {
	return s.identityService
}
