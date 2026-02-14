package v2

import (
	"github.com/dannyota/greennode-community-sdk/v2/greennode/client"
	"github.com/dannyota/greennode-community-sdk/v2/greennode/entity"
	sdkerror "github.com/dannyota/greennode-community-sdk/v2/greennode/sdkerror"
)

func (s *NetworkServiceV2) GetSecgroupByID(opts IGetSecgroupByIDRequest) (*entity.Secgroup, sdkerror.Error) {
	url := getSecgroupByIDURL(s.VserverClient, opts)
	resp := new(GetSecgroupByIDResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithHeader("User-Agent", opts.ParseUserAgent()).
		WithOkCodes(200).
		WithJSONResponse(resp).
		WithJSONError(errResp)

	if _, sdkErr := s.VserverClient.Get(url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.WithErrorSecgroupNotFound(errResp)).
			WithKVparameters(
				"secgroupId", opts.GetSecgroupID(),
				"projectId", s.getProjectID())
	}

	return resp.ToEntitySecgroup(), nil
}

func (s *NetworkServiceV2) CreateSecgroup(opts ICreateSecgroupRequest) (*entity.Secgroup, sdkerror.Error) {
	url := createSecgroupURL(s.VserverClient)
	resp := new(CreateSecgroupResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithHeader("User-Agent", opts.ParseUserAgent()).
		WithOkCodes(201).
		WithJSONBody(opts.ToRequestBody()).
		WithJSONResponse(resp).
		WithJSONError(errResp)

	if _, sdkErr := s.VserverClient.Post(url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.WithErrorSecgroupNameAlreadyExists(errResp),
			sdkerror.WithErrorSecgroupRuleExceedQuota(errResp),
			sdkerror.WithErrorSecgroupExceedQuota(errResp)).
			WithKVparameters(
				"secgroupName", opts.GetSecgroupName(),
				"projectId", s.getProjectID())
	}

	return resp.ToEntitySecgroup(), nil
}

func (s *NetworkServiceV2) ListSecgroup(opts IListSecgroupRequest) (*entity.ListSecgroups, sdkerror.Error) {
	url := listSecgroupURL(s.VserverClient, opts)
	resp := new(ListSecgroupResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithHeader("User-Agent", opts.ParseUserAgent()).
		WithOkCodes(200).
		WithJSONResponse(resp).
		WithJSONError(errResp)

	if _, sdkErr := s.VserverClient.Get(url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp)
	}

	return resp.ToListEntitySecgroups(), nil
}

func (s *NetworkServiceV2) DeleteSecgroupByID(opts IDeleteSecgroupByIDRequest) sdkerror.Error {
	url := deleteSecgroupByIDURL(s.VserverClient, opts)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithHeader("User-Agent", opts.ParseUserAgent()).
		WithOkCodes(204).
		WithJSONError(errResp)

	if _, sdkErr := s.VserverClient.Delete(url, req); sdkErr != nil {
		return sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.WithErrorSecgroupInUse(errResp),
			sdkerror.WithErrorSecgroupNotFound(errResp)).
			WithKVparameters(
				"secgroupId", opts.GetSecgroupID(),
				"projectId", s.getProjectID())
	}

	return nil
}
