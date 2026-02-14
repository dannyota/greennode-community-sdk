package v2

import (
	"github.com/dannyota/greennode-community-sdk/v2/greennode/client"
	"github.com/dannyota/greennode-community-sdk/v2/greennode/entity"
	sdkerror "github.com/dannyota/greennode-community-sdk/v2/greennode/sdkerror"
)

func (s *NetworkServiceV2) ListAllServersBySecgroupId(opts IListAllServersBySecgroupIdRequest) (*entity.ListServers, sdkerror.Error) {
	url := listAllServersBySecgroupIdUrl(s.VserverClient, opts)
	resp := new(ListAllServersBySecgroupIdResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithHeader("User-Agent", opts.ParseUserAgent()).
		WithOkCodes(200).
		WithJsonResponse(resp).
		WithJsonError(errResp)

	if _, sdkErr := s.VserverClient.Get(url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.WithErrorSecgroupNotFound(errResp)).
			WithKVparameters(
				"secgroupId", opts.GetSecgroupId(),
				"projectId", s.getProjectId())
	}

	return resp.ToEntityListServers(), nil
}
