package v2

import (
	"context"

	"danny.vn/greennode/greennode/client"
	sdkerror "danny.vn/greennode/greennode/sdkerror"
)

func (s *NetworkServiceV2) GetSubnetByID(ctx context.Context, opts *GetSubnetByIDRequest) (*Subnet, error) {
	url := getSubnetByIDURL(s.Client, opts)
	resp := new(GetSubnetByIDResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithOKCodes(200).
		WithJSONResponse(resp).
		WithJSONError(errResp)

	if _, sdkErr := s.Client.Get(ctx, url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.EcVServerSubnetNotBelongNetwork,
			sdkerror.EcVServerSubnetNotFound).
			WithKVparameters(
				"subnetId", opts.SubnetID,
				"projectId", s.Client.ProjectID)
	}

	return resp.ToEntitySubnet(), nil
}

func (s *NetworkServiceV2) UpdateSubnetByID(ctx context.Context, opts *UpdateSubnetByIDRequest) (*Subnet, error) {
	url := updateSubnetByIDURL(s.Client, opts)
	resp := new(UpdateSubnetByIDResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithOKCodes(200, 201, 202, 203, 204).
		WithJSONBody(opts.UpdateSubnetBody).
		WithJSONResponse(resp).
		WithJSONError(errResp)

	_, sdkErr := s.Client.Patch(ctx, url, req)
	if sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp)
	}

	return resp.ToEntitySubnet(), nil
}
