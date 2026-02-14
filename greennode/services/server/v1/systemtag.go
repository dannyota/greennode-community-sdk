package v1

import (
	"github.com/dannyota/greennode-community-sdk/v2/greennode/client"
	"github.com/dannyota/greennode-community-sdk/v2/greennode/entity"
	sdkerror "github.com/dannyota/greennode-community-sdk/v2/greennode/sdk_error"
)

func (s *ServerServiceInternalV1) CreateSystemTags(popts ICreateSystemTagRequest) (*[]entity.SystemTag, sdkerror.IError) {

	url := createSystemTagUrl(s.VServerClient)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)

	rawResp := new([]SystemTagResponse)

	req := client.NewRequest().
		WithHeader("User-Agent", popts.ParseUserAgent()).
		WithOkCodes(200).
		WithJsonBody(popts.ToRequestBody()).
		WithJsonResponse(rawResp).
		WithJsonError(errResp)

	if _, sdkErr := s.VServerClient.Post(url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp)
	}

	result := make([]entity.SystemTag, 0, len(*rawResp))

	for _, r := range *rawResp {
		result = append(result, r.toSystemTag())
	}

	return &result, nil
}
