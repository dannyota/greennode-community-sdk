package gateway

import (
	lsclient "github.com/dannyota/greennode-community-sdk/v2/greennode/client"
	lsidentitySvc "github.com/dannyota/greennode-community-sdk/v2/greennode/services/identity"
)

type iamGatewayV2 struct {
	identityService lsidentitySvc.IIdentityServiceV2
}

func NewIamGatewayV2(psvcClient lsclient.IServiceClient) IIamGatewayV2 {
	return &iamGatewayV2{
		identityService: lsidentitySvc.NewIdentityService(psvcClient),
	}
}

func (s *iamGatewayV2) IdentityService() lsidentitySvc.IIdentityServiceV2 {
	return s.identityService
}
