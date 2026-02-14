package identity

import (
	"github.com/dannyota/greennode-community-sdk/v2/greennode/client"
	identityv2 "github.com/dannyota/greennode-community-sdk/v2/greennode/services/identity/v2"
)

func NewIdentityService(svcClient client.ServiceClient) IdentityServiceV2 {
	return &identityv2.IdentityServiceV2{
		IamClient: svcClient,
	}
}
