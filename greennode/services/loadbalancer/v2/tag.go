package v2

import (
	"context"

	"github.com/dannyota/greennode-community-sdk/v2/greennode/client"
	"github.com/dannyota/greennode-community-sdk/v2/greennode/entity"
	sdkerror "github.com/dannyota/greennode-community-sdk/v2/greennode/sdkerror"
	"github.com/dannyota/greennode-community-sdk/v2/greennode/services/common"
)

func (s *LoadBalancerServiceV2) ListTags(ctx context.Context, opts *ListTagsRequest) (*entity.ListTags, error) {
	url := listTagsURL(s.ServerClient, opts)
	resp := new(ListTagsResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithOkCodes(200).
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
		WithOkCodes(200).
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
	tags := new(entity.ListTags)
	for _, tag := range tmpTags.Items {
		if !tag.SystemTag {
			tags.Items = append(tags.Items, tag)
		}
	}

	url := updateTagsURL(s.ServerClient, opts)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	opts.prepare(tags)
	req := client.NewRequest().
		WithOkCodes(200).
		WithJSONBody(opts).
		WithJSONError(errResp)

	if _, sdkErr = s.ServerClient.Put(ctx, url, req); sdkErr != nil {
		return sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.EcTagKeyInvalid).WithParameters(common.StructToMap(opts))
	}

	return nil

}
