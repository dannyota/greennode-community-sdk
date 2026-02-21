package v2

import (
	"context"

	"danny.vn/greennode/greennode/client"
	sdkerror "danny.vn/greennode/greennode/sdkerror"
	"danny.vn/greennode/greennode/services/common"
	types "danny.vn/greennode/greennode/types"
)

func (s *LoadBalancerServiceV2) ListTags(ctx context.Context, opts *ListTagsRequest) (*types.ListTags, error) {
	url := listTagsURL(s.ServerClient, opts)
	resp := new(ListTagsResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithOKCodes(200).
		WithJSONResponse(resp).
		WithJSONError(errResp)

	if _, sdkErr := s.ServerClient.Get(ctx, url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp)
	}

	return resp.ToEntityListTags(), nil
}

func (s *LoadBalancerServiceV2) CreateTags(ctx context.Context, opts *CreateTagsRequest) error {
	url := createTagsURL(s.ServerClient, opts)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithOKCodes(200).
		WithJSONBody(opts).
		WithJSONError(errResp)

	if _, sdkErr := s.ServerClient.Put(ctx, url, req); sdkErr != nil {
		return sdkerror.SdkErrorHandler(sdkErr, errResp)
	}

	return nil
}

func (s *LoadBalancerServiceV2) UpdateTags(ctx context.Context, opts *UpdateTagsRequest) error {
	tmpTags, sdkErr := s.ListTags(ctx, NewListTagsRequest(opts.LoadBalancerID))
	if sdkErr != nil {
		return sdkErr
	}

	// Do not update system tags
	tags := new(types.ListTags)
	for _, tag := range tmpTags.Items {
		if !tag.SystemTag {
			tags.Items = append(tags.Items, tag)
		}
	}

	url := updateTagsURL(s.ServerClient, opts)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	opts.prepare(tags)
	req := client.NewRequest().
		WithOKCodes(200).
		WithJSONBody(opts).
		WithJSONError(errResp)

	if _, sdkErr = s.ServerClient.Put(ctx, url, req); sdkErr != nil {
		return sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.EcTagKeyInvalid).WithParameters(common.StructToMap(opts))
	}

	return nil

}
