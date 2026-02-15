package v2

import (
	"github.com/dannyota/greennode-community-sdk/v2/greennode/client"
	"github.com/dannyota/greennode-community-sdk/v2/greennode/entity"
	sdkerror "github.com/dannyota/greennode-community-sdk/v2/greennode/sdkerror"
)

func (s *NetworkServiceV2) GetSecgroupByID(opts *GetSecgroupByIDRequest) (*entity.Secgroup, error) {
	url := getSecgroupByIDURL(s.VServerClient, opts)
	resp := new(GetSecgroupByIDResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithHeader("User-Agent", opts.ParseUserAgent()).
		WithOkCodes(200).
		WithJSONResponse(resp).
		WithJSONError(errResp)

	if _, sdkErr := s.VServerClient.Get(url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.EcVServerSecgroupNotFound).
			WithKVparameters(
				"secgroupId", opts.GetSecgroupID(),
				"projectId", s.getProjectID())
	}

	return resp.ToEntitySecgroup(), nil
}

func (s *NetworkServiceV2) CreateSecgroup(opts *CreateSecgroupRequest) (*entity.Secgroup, error) {
	url := createSecgroupURL(s.VServerClient)
	resp := new(CreateSecgroupResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithHeader("User-Agent", opts.ParseUserAgent()).
		WithOkCodes(201).
		WithJSONBody(opts.ToRequestBody()).
		WithJSONResponse(resp).
		WithJSONError(errResp)

	if _, sdkErr := s.VServerClient.Post(url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.EcVServerSecgroupNameAlreadyExists,
			sdkerror.EcVServerSecgroupRuleExceedQuota,
			sdkerror.EcVServerSecgroupExceedQuota).
			WithKVparameters(
				"secgroupName", opts.GetSecgroupName(),
				"projectId", s.getProjectID())
	}

	return resp.ToEntitySecgroup(), nil
}

func (s *NetworkServiceV2) ListSecgroup(opts *ListSecgroupRequest) (*entity.ListSecgroups, error) {
	url := listSecgroupURL(s.VServerClient, opts)
	resp := new(ListSecgroupResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithHeader("User-Agent", opts.ParseUserAgent()).
		WithOkCodes(200).
		WithJSONResponse(resp).
		WithJSONError(errResp)

	if _, sdkErr := s.VServerClient.Get(url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp)
	}

	return resp.ToListEntitySecgroups(), nil
}

func (s *NetworkServiceV2) DeleteSecgroupByID(opts *DeleteSecgroupByIDRequest) error {
	url := deleteSecgroupByIDURL(s.VServerClient, opts)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithHeader("User-Agent", opts.ParseUserAgent()).
		WithOkCodes(204).
		WithJSONError(errResp)

	if _, sdkErr := s.VServerClient.Delete(url, req); sdkErr != nil {
		return sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.EcVServerSecgroupInUse,
			sdkerror.EcVServerSecgroupNotFound).
			WithKVparameters(
				"secgroupId", opts.GetSecgroupID(),
				"projectId", s.getProjectID())
	}

	return nil
}
