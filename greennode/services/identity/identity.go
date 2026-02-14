package identity

import (
	"github.com/dannyota/greennode-community-sdk/v2/greennode/client"
	"github.com/dannyota/greennode-community-sdk/v2/greennode/entity"
	sdkerror "github.com/dannyota/greennode-community-sdk/v2/greennode/sdkerror"
	identityv2 "github.com/dannyota/greennode-community-sdk/v2/greennode/services/identity/v2"
)

type IdentityServiceV2 interface {
	GetAccessToken(opts identityv2.IGetAccessTokenRequest) (*entity.AccessToken, sdkerror.Error)
}

func NewIdentityService(svcClient client.ServiceClient) IdentityServiceV2 {
	return &identityv2.IdentityServiceV2{
		IAMClient: svcClient,
	}
}
