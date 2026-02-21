package v2

import (
	"context"

	"github.com/dannyota/greennode-community-sdk/greennode/client"
	sdkerror "github.com/dannyota/greennode-community-sdk/greennode/sdkerror"
	computev2 "github.com/dannyota/greennode-community-sdk/greennode/services/compute/v2"
)

func (s *NetworkServiceV2) ListAllServersBySecgroupID(ctx context.Context, opts *ListAllServersBySecgroupIDRequest) (*computev2.ListServers, error) {
	url := listAllServersBySecgroupIDURL(s.Client, opts)
	resp := new(ListAllServersBySecgroupIDResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithOKCodes(200).
		WithJSONResponse(resp).
		WithJSONError(errResp)

	if _, sdkErr := s.Client.Get(ctx, url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.EcVServerSecgroupNotFound).
			WithKVparameters(
				"secgroupId", opts.SecgroupID,
				"projectId", s.Client.ProjectID)
	}

	return resp.ToEntityListServers(), nil
}
