package v2

import (
	"encoding/base64"

	"github.com/dannyota/greennode-community-sdk/v2/greennode/client"
	"github.com/dannyota/greennode-community-sdk/v2/greennode/entity"
	sdkerror "github.com/dannyota/greennode-community-sdk/v2/greennode/sdkerror"
)

func (s *IdentityServiceV2) GetAccessToken(opts IGetAccessTokenRequest) (*entity.AccessToken, sdkerror.Error) {
	url := getAccessTokenURL(s.IamClient)
	resp := new(GetAccessTokenResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.IamErrorType)
	req := client.NewRequest().
		WithOkCodes(200).
		WithJSONResponse(resp).
		WithSkipAuth(true).
		WithJSONError(errResp).
		WithJSONBody(opts.ToRequestBody()).
		WithHeader("Content-Type", "application/x-www-form-urlencoded").
		WithHeader("Authorization", "Basic "+base64.StdEncoding.EncodeToString([]byte(opts.GetClientID()+":"+opts.GetClientSecret())))

	if _, sdkErr := s.IamClient.Post(url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.WithErrorTooManyFailedLogin(errResp),
			sdkerror.WithErrorAuthenticationFailed(errResp),
			sdkerror.WithErrorUnknownAuthFailure(errResp)). // Always put this handler at the end
			WithKVparameters("clientId", opts.GetClientID())
	}

	return resp.ToEntityAccessToken(), nil
}
