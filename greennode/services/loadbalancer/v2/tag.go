package v2

import (
	"github.com/dannyota/greennode-community-sdk/v2/greennode/client"
	"github.com/dannyota/greennode-community-sdk/v2/greennode/entity"
	sdkerror "github.com/dannyota/greennode-community-sdk/v2/greennode/sdkerror"
)

func (s *LoadBalancerServiceV2) ListTags(popts IListTagsRequest) (*entity.ListTags, sdkerror.Error) {
	url := listTagsUrl(s.VServerClient, popts)
	resp := new(ListTagsResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithHeader("User-Agent", popts.ParseUserAgent()).
		WithOkCodes(200).
		WithJsonResponse(resp).
		WithJsonError(errResp)

	if _, sdkErr := s.VServerClient.Get(url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp)
	}

	return resp.ToEntityListTags(), nil
}

func (s *LoadBalancerServiceV2) CreateTags(popts ICreateTagsRequest) sdkerror.Error {
	url := createTagsUrl(s.VServerClient, popts)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithHeader("User-Agent", popts.ParseUserAgent()).
		WithOkCodes(200).
		WithJsonBody(popts.ToRequestBody()).
		WithJsonError(errResp)

	if _, sdkErr := s.VServerClient.Put(url, req); sdkErr != nil {
		return sdkerror.SdkErrorHandler(sdkErr, errResp)
	}

	return nil
}

func (s *LoadBalancerServiceV2) UpdateTags(popts IUpdateTagsRequest) sdkerror.Error {
	tmpTags, sdkErr := s.ListTags(NewListTagsRequest(popts.GetLoadBalancerId()))
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

	url := updateTagsUrl(s.VServerClient, popts)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithHeader("User-Agent", popts.ParseUserAgent()).
		WithOkCodes(200).
		WithJsonBody(popts.ToRequestBody(tags)).
		WithJsonError(errResp)

	if _, sdkErr = s.VServerClient.Put(url, req); sdkErr != nil {
		return sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.WithErrorTagKeyInvalid(errResp)).WithParameters(popts.ToMap())
	}

	return nil

}
