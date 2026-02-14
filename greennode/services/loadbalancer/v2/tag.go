package v2

import (
	"github.com/dannyota/greennode-community-sdk/v2/greennode/client"
	"github.com/dannyota/greennode-community-sdk/v2/greennode/entity"
	sdkerror "github.com/dannyota/greennode-community-sdk/v2/greennode/sdkerror"
)

func (s *LoadBalancerServiceV2) ListTags(opts IListTagsRequest) (*entity.ListTags, sdkerror.Error) {
	url := listTagsURL(s.VServerClient, opts)
	resp := new(ListTagsResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithHeader("User-Agent", opts.ParseUserAgent()).
		WithOkCodes(200).
		WithJSONResponse(resp).
		WithJSONError(errResp)

	if _, sdkErr := s.VServerClient.Get(url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp)
	}

	return resp.ToEntityListTags(), nil
}

func (s *LoadBalancerServiceV2) CreateTags(opts ICreateTagsRequest) sdkerror.Error {
	url := createTagsURL(s.VServerClient, opts)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithHeader("User-Agent", opts.ParseUserAgent()).
		WithOkCodes(200).
		WithJSONBody(opts.ToRequestBody()).
		WithJSONError(errResp)

	if _, sdkErr := s.VServerClient.Put(url, req); sdkErr != nil {
		return sdkerror.SdkErrorHandler(sdkErr, errResp)
	}

	return nil
}

func (s *LoadBalancerServiceV2) UpdateTags(opts IUpdateTagsRequest) sdkerror.Error {
	tmpTags, sdkErr := s.ListTags(NewListTagsRequest(opts.GetLoadBalancerID()))
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

	url := updateTagsURL(s.VServerClient, opts)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithHeader("User-Agent", opts.ParseUserAgent()).
		WithOkCodes(200).
		WithJSONBody(opts.ToRequestBody(tags)).
		WithJSONError(errResp)

	if _, sdkErr = s.VServerClient.Put(url, req); sdkErr != nil {
		return sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.WithErrorTagKeyInvalid(errResp)).WithParameters(opts.ToMap())
	}

	return nil

}
