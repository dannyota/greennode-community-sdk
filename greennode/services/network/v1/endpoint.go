package v1

import (
	"github.com/dannyota/greennode-community-sdk/v2/greennode/client"
	"github.com/dannyota/greennode-community-sdk/v2/greennode/entity"
	sdkerror "github.com/dannyota/greennode-community-sdk/v2/greennode/sdkerror"
)

func (s *NetworkServiceV1) GetEndpointById(popts IGetEndpointByIdRequest) (*entity.Endpoint, sdkerror.Error) {
	url := getEndpointByIdUrl(s.VNetworkClient, popts)
	resp := new(GetEndpointByIdResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.NetworkGatewayErrorType)
	req := client.NewRequest().
		WithOkCodes(200).
		// WithUserId(s.getUserId()).
		WithJsonResponse(resp).
		WithJsonError(errResp)

	if _, sdkErr := s.VNetworkClient.Get(url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp).
			WithKVparameters(
				"endpointId", popts.GetEndpointId(),
				"projectId", s.getProjectId()).
			WithErrorCategories(sdkerror.ErrCatProductVNetwork)
	}

	return resp.ToEntityEndpoint(), nil
}

func (s *NetworkServiceV1) CreateEndpoint(popts ICreateEndpointRequest) (*entity.Endpoint, sdkerror.Error) {
	url := createEndpointUrl(s.VNetworkClient)
	resp := new(CreateEndpointResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.NetworkGatewayErrorType)
	req := client.NewRequest().
		WithOkCodes(201).
		WithUserId(popts.GetPortalUserId()).
		WithJsonBody(popts.ToRequestBody(s.VNetworkClient)).
		WithJsonResponse(resp).
		WithJsonError(errResp)

	if _, sdkErr := s.VNetworkClient.Post(url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.WithErrorEndpointOfVpcExists(errResp),
			sdkerror.WithErrorLockOnProcess(errResp),
			sdkerror.WithErrorNetworkNotFound(errResp),
			sdkerror.WithErrorPurchaseIssue(errResp),
			sdkerror.WithErrorPaymentMethodNotAllow(errResp),
			sdkerror.WithErrorCreditNotEnough(errResp),
			sdkerror.WithErrorSubnetNotFound(errResp),
			sdkerror.WithErrorEndpointPackageNotBelongToEndpointService(errResp),
			sdkerror.WithErrorContainInvalidCharacter(errResp)).
			WithParameters(popts.ToMap()).
			WithErrorCategories(sdkerror.ErrCatProductVNetwork)
	}

	return resp.ToEntityEndpoint(), nil
}

func (s *NetworkServiceV1) DeleteEndpointById(popts IDeleteEndpointByIdRequest) sdkerror.Error {
	url := deleteEndpointByIdUrl(s.VNetworkClient, popts)
	errResp := sdkerror.NewErrorResponse(sdkerror.NetworkGatewayErrorType)
	req := client.NewRequest().
		WithHeader("User-Agent", popts.ParseUserAgent()).
		WithOkCodes(200).
		WithJsonBody(popts.ToRequestBody(s.VNetworkClient)).
		// WithUserId(s.getUserId()).
		WithJsonError(errResp)

	if _, sdkErr := s.VNetworkClient.Delete(url, req); sdkErr != nil {
		return sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.WithErrorEndpointStatusInvalid(errResp),
			sdkerror.WithErrorNetworkNotFound(errResp),
			sdkerror.WithErrorSubnetNotFound(errResp)).
			WithParameters(popts.ToMap()).
			WithErrorCategories(sdkerror.ErrCatProductVNetwork)
	}

	return nil
}

func (s *NetworkServiceV1) ListEndpoints(popts IListEndpointsRequest) (*entity.ListEndpoints, sdkerror.Error) {
	url := listEndpointsUrl(s.VNetworkClient, popts)
	resp := new(ListEndpointsResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.NetworkGatewayErrorType)
	req := client.NewRequest().
		WithHeader("User-Agent", popts.ParseUserAgent()).
		WithOkCodes(200).
		// WithUserId(s.getUserId()).
		WithJsonResponse(resp).
		WithJsonError(errResp)

	if _, sdkErr := s.VNetworkClient.Get(url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp).
			WithKVparameters("projectId", s.getProjectId()).
			WithParameters(popts.ToMap()).
			WithErrorCategories(sdkerror.ErrCatProductVNetwork)
	}

	return resp.ToEntityListEndpoints(), nil
}

// ________________________________________________________________________ NetworkServiceInternalV1

func (s *NetworkServiceInternalV1) ListTagsByEndpointId(popts IListTagsByEndpointIdRequest) (*entity.ListTags, sdkerror.Error) {
	url := listTagsByEndpointIdUrl(s.VNetworkClient, popts)
	resp := new(ListTagsByEndpointIdResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.NetworkGatewayErrorType)
	req := client.NewRequest().
		WithMapHeaders(popts.GetMapHeaders()).
		WithHeader("User-Agent", popts.ParseUserAgent()).
		WithOkCodes(200).
		WithJsonResponse(resp).
		WithJsonError(errResp)

	if _, sdkErr := s.VNetworkClient.Get(url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp).
			WithKVparameters("projectId", s.getProjectId()).
			WithParameters(popts.ToMap()).
			WithErrorCategories(sdkerror.ErrCatProductVNetwork)
	}

	return resp.ToEntityListTags(), nil
}

func (s *NetworkServiceInternalV1) CreateTagsWithEndpointId(popts ICreateTagsWithEndpointIdRequest) sdkerror.Error {
	url := createTagsWithEndpointIdUrl(s.VNetworkClient, popts)
	errResp := sdkerror.NewErrorResponse(sdkerror.NetworkGatewayErrorType)
	req := client.NewRequest().
		WithMapHeaders(popts.GetMapHeaders()).
		WithHeader("User-Agent", popts.ParseUserAgent()).
		WithOkCodes(200).
		WithJsonBody(popts.ToRequestBody()).
		WithJsonError(errResp)

	if _, sdkErr := s.VNetworkClient.Post(url, req); sdkErr != nil {
		return sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.WithErrorEndpointTagExisted(errResp),
			sdkerror.WithErrorEndpointTagNotFound(errResp)).
			WithKVparameters("projectId", s.getProjectId()).
			WithParameters(popts.ToMap()).
			WithErrorCategories(sdkerror.ErrCatProductVNetwork)
	}

	return nil
}

func (s *NetworkServiceInternalV1) DeleteTagOfEndpoint(popts IDeleteTagOfEndpointRequest) sdkerror.Error {
	url := deleteTagOfEndpointUrl(s.VNetworkClient, popts)
	errResp := sdkerror.NewErrorResponse(sdkerror.NetworkGatewayErrorType)
	req := client.NewRequest().
		WithMapHeaders(popts.GetMapHeaders()).
		WithHeader("User-Agent", popts.ParseUserAgent()).
		WithOkCodes(200).
		WithJsonError(errResp)

	if _, sdkErr := s.VNetworkClient.Delete(url, req); sdkErr != nil {
		return sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.WithErrorEndpointTagNotFound(errResp)).
			WithKVparameters("projectId", s.getProjectId()).
			WithParameters(popts.ToMap()).
			WithErrorCategories(sdkerror.ErrCatProductVNetwork)
	}

	return nil
}

func (s *NetworkServiceInternalV1) UpdateTagValueOfEndpoint(popts IUpdateTagValueOfEndpointRequest) sdkerror.Error {
	url := updateTagValueOfEndpointUrl(s.VNetworkClient, popts)
	errResp := sdkerror.NewErrorResponse(sdkerror.NetworkGatewayErrorType)
	req := client.NewRequest().
		WithMapHeaders(popts.GetMapHeaders()).
		WithHeader("User-Agent", popts.ParseUserAgent()).
		WithOkCodes(200).
		WithJsonBody(popts.ToRequestBody()).
		WithJsonError(errResp)

	if _, sdkErr := s.VNetworkClient.Put(url, req); sdkErr != nil {
		return sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.WithErrorEndpointTagNotFound(errResp)).
			WithKVparameters("projectId", s.getProjectId()).
			WithParameters(popts.ToMap()).
			WithErrorCategories(sdkerror.ErrCatProductVNetwork)
	}

	return nil
}

func (s *NetworkServiceInternalV1) CreateEndpoint(popts ICreateEndpointRequest) (*entity.Endpoint, sdkerror.Error) {
	url := createEndpointUrl(s.VNetworkClient)
	resp := new(CreateEndpointResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.NetworkGatewayErrorType)
	req := client.NewRequest().
		WithOkCodes(201).
		// WithUserId(s.getUserId()).
		WithJsonBody(popts.ToRequestBody(s.VNetworkClient)).
		WithJsonResponse(resp).
		WithJsonError(errResp)

	if _, sdkErr := s.VNetworkClient.Post(url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.WithErrorEndpointOfVpcExists(errResp),
			sdkerror.WithErrorLockOnProcess(errResp),
			sdkerror.WithErrorNetworkNotFound(errResp),
			sdkerror.WithErrorPurchaseIssue(errResp),
			sdkerror.WithErrorPaymentMethodNotAllow(errResp),
			sdkerror.WithErrorCreditNotEnough(errResp),
			sdkerror.WithErrorSubnetNotFound(errResp),
			sdkerror.WithErrorEndpointPackageNotBelongToEndpointService(errResp),
			sdkerror.WithErrorContainInvalidCharacter(errResp)).
			WithParameters(popts.ToMap()).
			WithErrorCategories(sdkerror.ErrCatProductVNetwork)
	}

	return resp.ToEntityEndpoint(), nil
}
