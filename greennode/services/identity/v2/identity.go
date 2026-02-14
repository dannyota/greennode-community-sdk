package v2

import (
	"encoding/base64"

	"github.com/dannyota/greennode-community-sdk/v2/greennode/client"
	"github.com/dannyota/greennode-community-sdk/v2/greennode/entity"
	sdkerror "github.com/dannyota/greennode-community-sdk/v2/greennode/sdkerror"
)

func (s *IdentityServiceV2) GetAccessToken(popts IGetAccessTokenRequest) (*entity.AccessToken, sdkerror.IError) {
	url := getAccessTokenUrl(s.IamClient)
	resp := new(GetAccessTokenResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.IamErrorType)
	req := client.NewRequest().
		WithOkCodes(200).
		WithJsonResponse(resp).
		WithSkipAuth(true).
		WithJsonError(errResp).
		WithJsonBody(popts.ToRequestBody()).
		WithHeader("Content-Type", "application/x-www-form-urlencoded").
		WithHeader("Authorization", "Basic "+base64.StdEncoding.EncodeToString([]byte(popts.GetClientId()+":"+popts.GetClientSecret())))

	if _, sdkErr := s.IamClient.Post(url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.WithErrorTooManyFailedLogin(errResp),
			sdkerror.WithErrorAuthenticationFailed(errResp),
			sdkerror.WithErrorUnknownAuthFailure(errResp)). // Always put this handler at the end
			WithKVparameters("clientId", popts.GetClientId())
	}

	return resp.ToEntityAccessToken(), nil
}
