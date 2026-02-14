package identity

import (
	lsentity "github.com/dannyota/greennode-community-sdk/v2/greennode/entity"
	lsdkErr "github.com/dannyota/greennode-community-sdk/v2/greennode/sdk_error"
	lsidentitySvcV2 "github.com/dannyota/greennode-community-sdk/v2/greennode/services/identity/v2"
)

type IIdentityServiceV2 interface {
	GetAccessToken(popts lsidentitySvcV2.IGetAccessTokenRequest) (*lsentity.AccessToken, lsdkErr.IError)
}
