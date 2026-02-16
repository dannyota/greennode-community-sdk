package v1

import (
	"context"

	"github.com/dannyota/greennode-community-sdk/v2/greennode/client"
	sdkerror "github.com/dannyota/greennode-community-sdk/v2/greennode/sdkerror"
)

func (s *PortalServiceV1) ListZones(ctx context.Context) (*ListZones, error) {
	url := listZonesURL(s.Client)
	resp := new(ListZoneResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithOKCodes(200).
		WithJSONResponse(resp).
		WithJSONError(errResp)

	if _, sdkErr := s.Client.Get(ctx, url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp).
			WithKVparameters("projectId", s.Client.ProjectID)
	}

	return resp.ToEntityListZones(), nil
}
