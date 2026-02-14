package identity

import (
	"github.com/dannyota/greennode-community-sdk/v2/greennode/entity"
	sdkerror "github.com/dannyota/greennode-community-sdk/v2/greennode/sdkerror"
	identityv2 "github.com/dannyota/greennode-community-sdk/v2/greennode/services/identity/v2"
)

type IdentityServiceV2 interface {
	GetAccessToken(popts identityv2.IGetAccessTokenRequest) (*entity.AccessToken, sdkerror.Error)
}
