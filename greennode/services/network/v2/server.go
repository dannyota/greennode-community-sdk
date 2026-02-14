package v2

import (
	"github.com/dannyota/greennode-community-sdk/v2/greennode/client"
	"github.com/dannyota/greennode-community-sdk/v2/greennode/entity"
	sdkerror "github.com/dannyota/greennode-community-sdk/v2/greennode/sdkerror"
)

func (s *NetworkServiceV2) ListAllServersBySecgroupID(opts *ListAllServersBySecgroupIDRequest) (*entity.ListServers, sdkerror.Error) {
	url := listAllServersBySecgroupIDURL(s.VserverClient, opts)
	resp := new(ListAllServersBySecgroupIDResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithHeader("User-Agent", opts.ParseUserAgent()).
		WithOkCodes(200).
		WithJSONResponse(resp).
		WithJSONError(errResp)

	if _, sdkErr := s.VserverClient.Get(url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.WithErrorSecgroupNotFound(errResp)).
			WithKVparameters(
				"secgroupId", opts.GetSecgroupID(),
				"projectId", s.getProjectID())
	}

	return resp.ToEntityListServers(), nil
}
