package v1

import (
	"context"

	"danny.vn/greennode/client"
	sdkerror "danny.vn/greennode/sdkerror"
)

type ComputeServiceV1 struct {
	Client *client.ServiceClient
}

func (s *ComputeServiceV1) ListOSImages(ctx context.Context, opts *ListOSImagesRequest) (*ListOSImages, error) {
	url := listOSImagesURL(s.Client, opts)
	resp := new(listImagesResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithOKCodes(200).
		WithJSONResponse(resp).
		WithJSONError(errResp)

	if _, sdkErr := s.Client.Get(ctx, url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp).
			WithKVparameters("projectId", s.Client.ProjectID)
	}

	return resp.ToEntityListOSImages(), nil
}

func (s *ComputeServiceV1) ListGPUImages(ctx context.Context) (*ListOSImages, error) {
	url := listGPUImagesURL(s.Client)
	resp := new(listImagesResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithOKCodes(200).
		WithJSONResponse(resp).
		WithJSONError(errResp)

	if _, sdkErr := s.Client.Get(ctx, url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp).
			WithKVparameters("projectId", s.Client.ProjectID)
	}

	return resp.ToEntityListOSImages(), nil
}
