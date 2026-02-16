package v1

import (
	"context"

	"github.com/dannyota/greennode-community-sdk/v2/greennode/client"
	"github.com/dannyota/greennode-community-sdk/v2/greennode/entity"
	sdkerror "github.com/dannyota/greennode-community-sdk/v2/greennode/sdkerror"
)

type PortalServiceV1 struct {
	Client *client.ServiceClient
}

func (s *PortalServiceV1) GetPortalInfo(ctx context.Context, opts *GetPortalInfoRequest) (*entity.Portal, error) {
	url := getPortalInfoURL(s.Client, opts)
	resp := new(entity.Portal)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithOkCodes(200).
		WithJSONResponse(resp).
		WithJSONError(errResp)

	if _, sdkErr := s.Client.Get(ctx, url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp).
			WithKVparameters("backendProjectId", opts.BackEndProjectID)
	}

	return resp, nil
}

func (s *PortalServiceV1) ListProjects(ctx context.Context, opts *ListProjectsRequest) (*entity.ListPortals, error) {
	url := listProjectsURL(s.Client)
	resp := new(ListProjectsResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithOkCodes(200).
		WithJSONResponse(resp).
		WithJSONError(errResp)

	if _, sdkErr := s.Client.Get(ctx, url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp)
	}

	return resp.ToEntityListPortals(), nil
}
