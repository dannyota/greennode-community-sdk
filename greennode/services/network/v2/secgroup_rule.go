package v2

import (
	"github.com/dannyota/greennode-community-sdk/v2/greennode/client"
	"github.com/dannyota/greennode-community-sdk/v2/greennode/entity"
	sdkerror "github.com/dannyota/greennode-community-sdk/v2/greennode/sdkerror"
)

func (s *NetworkServiceV2) CreateSecgroupRule(opts *CreateSecgroupRuleRequest) (*entity.SecgroupRule, sdkerror.Error) {
	url := createSecgroupRuleURL(s.VserverClient, opts)
	resp := new(CreateSecgroupRuleResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithHeader("User-Agent", opts.ParseUserAgent()).
		WithOkCodes(201).
		WithJSONBody(opts.ToRequestBody()).
		WithJSONResponse(resp).
		WithJSONError(errResp)

	if _, sdkErr := s.VserverClient.Post(url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.WithErrorSecgroupNotFound(errResp),
			sdkerror.WithErrorSecgroupRuleExceedQuota(errResp),
			sdkerror.WithErrorSecgroupRuleAlreadyExists(errResp)).
			WithParameters(opts.ToMap()).
			WithKVparameters("projectId", s.getProjectID())
	}

	return resp.ToEntitySecgroupRule(), nil
}

func (s *NetworkServiceV2) DeleteSecgroupRuleByID(opts *DeleteSecgroupRuleByIDRequest) sdkerror.Error {
	url := deleteSecgroupRuleByIDURL(s.VserverClient, opts)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithHeader("User-Agent", opts.ParseUserAgent()).
		WithOkCodes(204).
		WithJSONError(errResp)

	if _, sdkErr := s.VserverClient.Delete(url, req); sdkErr != nil {
		return sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.WithErrorSecgroupRuleNotFound(errResp),
			sdkerror.WithErrorSecgroupNotFound(errResp)).
			WithKVparameters(
				"secgroupId", opts.GetSecgroupID(),
				"secgroupRuleId", opts.GetSecgroupRuleID(),
				"projectId", s.getProjectID())
	}

	return nil
}

func (s *NetworkServiceV2) ListSecgroupRulesBySecgroupID(opts *ListSecgroupRulesBySecgroupIDRequest) (*entity.ListSecgroupRules, sdkerror.Error) {
	url := listSecgroupRulesBySecgroupIDURL(s.VserverClient, opts)
	resp := new(ListSecgroupRulesBySecgroupIDResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithHeader("User-Agent", opts.ParseUserAgent()).
		WithOkCodes(200).
		WithJSONResponse(resp).
		WithJSONError(errResp)

	if _, sdkErr := s.VserverClient.Get(url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.WithErrorSecgroupNotFound(errResp)).
			WithKVparameters("projectId", s.getProjectID(), "secgroupId", opts.GetSecgroupID())
	}

	return resp.ToEntityListSecgroupRules(), nil
}
