package v2

import (
	"context"

	"danny.vn/greennode/client"
	sdkerror "danny.vn/greennode/sdkerror"
	"danny.vn/greennode/services/common"
)

func (s *NetworkServiceV2) ListInterconnects(ctx context.Context, opts *ListInterconnectsRequest) (*ListInterconnects, error) {
	url := listInterconnectsURL(s.Client, opts)
	resp := new(ListInterconnectsResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithOKCodes(200).
		WithJSONResponse(resp).
		WithJSONError(errResp)

	if _, sdkErr := s.Client.Get(ctx, url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp).
			WithParameters(common.StructToMap(opts)).
			WithKVparameters("projectId", s.Client.ProjectID)
	}

	return resp.ToEntityListInterconnects(), nil
}
