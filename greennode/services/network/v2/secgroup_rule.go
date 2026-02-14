package v2

import (
	"github.com/dannyota/greennode-community-sdk/v2/greennode/client"
	"github.com/dannyota/greennode-community-sdk/v2/greennode/entity"
	sdkerror "github.com/dannyota/greennode-community-sdk/v2/greennode/sdkerror"
)

func (s *NetworkServiceV2) CreateSecgroupRule(opts ICreateSecgroupRuleRequest) (*entity.SecgroupRule, sdkerror.Error) {
	url := createSecgroupRuleUrl(s.VserverClient, opts)
	resp := new(CreateSecgroupRuleResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithHeader("User-Agent", opts.ParseUserAgent()).
		WithOkCodes(201).
		WithJsonBody(opts.ToRequestBody()).
		WithJsonResponse(resp).
		WithJsonError(errResp)

	if _, sdkErr := s.VserverClient.Post(url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.WithErrorSecgroupNotFound(errResp),
			sdkerror.WithErrorSecgroupRuleExceedQuota(errResp),
			sdkerror.WithErrorSecgroupRuleAlreadyExists(errResp)).
			WithParameters(opts.ToMap()).
			WithKVparameters("projectId", s.getProjectId())
	}

	return resp.ToEntitySecgroupRule(), nil
}

func (s *NetworkServiceV2) DeleteSecgroupRuleById(opts IDeleteSecgroupRuleByIdRequest) sdkerror.Error {
	url := deleteSecgroupRuleByIdUrl(s.VserverClient, opts)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithHeader("User-Agent", opts.ParseUserAgent()).
		WithOkCodes(204).
		WithJsonError(errResp)

	if _, sdkErr := s.VserverClient.Delete(url, req); sdkErr != nil {
		return sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.WithErrorSecgroupRuleNotFound(errResp),
			sdkerror.WithErrorSecgroupNotFound(errResp)).
			WithKVparameters(
				"secgroupId", opts.GetSecgroupId(),
				"secgroupRuleId", opts.GetSecgroupRuleId(),
				"projectId", s.getProjectId())
	}

	return nil
}

func (s *NetworkServiceV2) ListSecgroupRulesBySecgroupId(opts IListSecgroupRulesBySecgroupIdRequest) (*entity.ListSecgroupRules, sdkerror.Error) {
	url := listSecgroupRulesBySecgroupIdUrl(s.VserverClient, opts)
	resp := new(ListSecgroupRulesBySecgroupIdResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithHeader("User-Agent", opts.ParseUserAgent()).
		WithOkCodes(200).
		WithJsonResponse(resp).
		WithJsonError(errResp)

	if _, sdkErr := s.VserverClient.Get(url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.WithErrorSecgroupNotFound(errResp)).
			WithKVparameters("projectId", s.getProjectId(), "secgroupId", opts.GetSecgroupId())
	}

	return resp.ToEntityListSecgroupRules(), nil
}
