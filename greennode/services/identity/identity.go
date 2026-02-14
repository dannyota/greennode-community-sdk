package identity

import (
	lsclient "github.com/dannyota/greennode-community-sdk/v2/greennode/client"
	lsidentitySvcV2 "github.com/dannyota/greennode-community-sdk/v2/greennode/services/identity/v2"
)

func NewIdentityService(psvcClient lsclient.IServiceClient) IIdentityServiceV2 {
	return &lsidentitySvcV2.IdentityServiceV2{
		IamClient: psvcClient,
	}
}
