package v2

import (
	"context"
	"encoding/base64"

	"github.com/dannyota/greennode-community-sdk/v2/greennode/client"
	sdkerror "github.com/dannyota/greennode-community-sdk/v2/greennode/sdkerror"
)

type IdentityServiceV2 struct {
	Client *client.ServiceClient
}

func (s *IdentityServiceV2) GetAccessToken(ctx context.Context, opts *GetAccessTokenRequest) (*AccessToken, error) {
	url := getAccessTokenURL(s.Client)
	resp := new(GetAccessTokenResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.IAMErrorType)
	req := client.NewRequest().
		WithOKCodes(200).
		WithJSONResponse(resp).
		WithSkipAuth(true).
		WithJSONError(errResp).
		WithJSONBody(opts).
		WithHeader("Content-Type", "application/x-www-form-urlencoded").
		WithHeader("Authorization", "Basic "+base64.StdEncoding.EncodeToString([]byte(opts.ClientID+":"+opts.ClientSecret)))

	if _, sdkErr := s.Client.Post(ctx, url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.EcTooManyFailedLogins,
			sdkerror.EcAuthenticationFailed,
			sdkerror.EcUnknownAuthFailure). // Always put this handler at the end
			WithKVparameters("clientId", opts.ClientID)
	}

	return resp.ToEntityAccessToken(), nil
}
