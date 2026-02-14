package v2

import (
	"github.com/dannyota/greennode-community-sdk/v2/greennode/client"
	"github.com/dannyota/greennode-community-sdk/v2/greennode/entity"
	sdkerror "github.com/dannyota/greennode-community-sdk/v2/greennode/sdk_error"
)

func (s *NetworkServiceV2) ListAllServersBySecgroupId(popts IListAllServersBySecgroupIdRequest) (*entity.ListServers, sdkerror.IError) {
	url := listAllServersBySecgroupIdUrl(s.VserverClient, popts)
	resp := new(ListAllServersBySecgroupIdResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithHeader("User-Agent", popts.ParseUserAgent()).
		WithOkCodes(200).
		WithJsonResponse(resp).
		WithJsonError(errResp)

	if _, sdkErr := s.VserverClient.Get(url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.WithErrorSecgroupNotFound(errResp)).
			WithKVparameters(
				"secgroupId", popts.GetSecgroupId(),
				"projectId", s.getProjectId())
	}

	return resp.ToEntityListServers(), nil
}
