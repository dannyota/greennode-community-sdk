package v2

import (
	"context"
	"encoding/base64"

	"github.com/dannyota/greennode-community-sdk/v2/greennode/client"
	"github.com/dannyota/greennode-community-sdk/v2/greennode/entity"
	sdkerror "github.com/dannyota/greennode-community-sdk/v2/greennode/sdkerror"
)

type IdentityServiceV2 struct {
	IAMClient *client.ServiceClient
}

func (s *IdentityServiceV2) GetAccessToken(ctx context.Context, opts *GetAccessTokenRequest) (*entity.AccessToken, error) {
	url := getAccessTokenURL(s.IAMClient)
	resp := new(GetAccessTokenResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.IAMErrorType)
	req := client.NewRequest().
		WithOkCodes(200).
		WithJSONResponse(resp).
		WithSkipAuth(true).
		WithJSONError(errResp).
		WithJSONBody(opts).
		WithHeader("Content-Type", "application/x-www-form-urlencoded").
		WithHeader("Authorization", "Basic "+base64.StdEncoding.EncodeToString([]byte(opts.ClientID+":"+opts.ClientSecret)))

	if _, sdkErr := s.IAMClient.Post(ctx, url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.EcTooManyFailedLogins,
			sdkerror.EcAuthenticationFailed,
			sdkerror.EcUnknownAuthFailure). // Always put this handler at the end
			WithKVparameters("clientId", opts.ClientID)
	}

	return resp.ToEntityAccessToken(), nil
}
