package v1

import (
	"github.com/dannyota/greennode-community-sdk/v2/greennode/client"
	"github.com/dannyota/greennode-community-sdk/v2/greennode/entity"
	sdkerror "github.com/dannyota/greennode-community-sdk/v2/greennode/sdkerror"
)

func (s *NetworkServiceV1) GetEndpointById(opts IGetEndpointByIdRequest) (*entity.Endpoint, sdkerror.Error) {
	url := getEndpointByIdUrl(s.VNetworkClient, opts)
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
				"endpointId", opts.GetEndpointId(),
				"projectId", s.getProjectId()).
			WithErrorCategories(sdkerror.ErrCatProductVNetwork)
	}

	return resp.ToEntityEndpoint(), nil
}

func (s *NetworkServiceV1) CreateEndpoint(opts ICreateEndpointRequest) (*entity.Endpoint, sdkerror.Error) {
	url := createEndpointUrl(s.VNetworkClient)
	resp := new(CreateEndpointResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.NetworkGatewayErrorType)
	req := client.NewRequest().
		WithOkCodes(201).
		WithUserId(opts.GetPortalUserId()).
		WithJsonBody(opts.ToRequestBody(s.VNetworkClient)).
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
			WithParameters(opts.ToMap()).
			WithErrorCategories(sdkerror.ErrCatProductVNetwork)
	}

	return resp.ToEntityEndpoint(), nil
}

func (s *NetworkServiceV1) DeleteEndpointById(opts IDeleteEndpointByIdRequest) sdkerror.Error {
	url := deleteEndpointByIdUrl(s.VNetworkClient, opts)
	errResp := sdkerror.NewErrorResponse(sdkerror.NetworkGatewayErrorType)
	req := client.NewRequest().
		WithHeader("User-Agent", opts.ParseUserAgent()).
		WithOkCodes(200).
		WithJsonBody(opts.ToRequestBody(s.VNetworkClient)).
		// WithUserId(s.getUserId()).
		WithJsonError(errResp)

	if _, sdkErr := s.VNetworkClient.Delete(url, req); sdkErr != nil {
		return sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.WithErrorEndpointStatusInvalid(errResp),
			sdkerror.WithErrorNetworkNotFound(errResp),
			sdkerror.WithErrorSubnetNotFound(errResp)).
			WithParameters(opts.ToMap()).
			WithErrorCategories(sdkerror.ErrCatProductVNetwork)
	}

	return nil
}

func (s *NetworkServiceV1) ListEndpoints(opts IListEndpointsRequest) (*entity.ListEndpoints, sdkerror.Error) {
	url := listEndpointsUrl(s.VNetworkClient, opts)
	resp := new(ListEndpointsResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.NetworkGatewayErrorType)
	req := client.NewRequest().
		WithHeader("User-Agent", opts.ParseUserAgent()).
		WithOkCodes(200).
		// WithUserId(s.getUserId()).
		WithJsonResponse(resp).
		WithJsonError(errResp)

	if _, sdkErr := s.VNetworkClient.Get(url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp).
			WithKVparameters("projectId", s.getProjectId()).
			WithParameters(opts.ToMap()).
			WithErrorCategories(sdkerror.ErrCatProductVNetwork)
	}

	return resp.ToEntityListEndpoints(), nil
}

// ________________________________________________________________________ NetworkServiceInternalV1

func (s *NetworkServiceInternalV1) ListTagsByEndpointId(opts IListTagsByEndpointIdRequest) (*entity.ListTags, sdkerror.Error) {
	url := listTagsByEndpointIdUrl(s.VNetworkClient, opts)
	resp := new(ListTagsByEndpointIdResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.NetworkGatewayErrorType)
	req := client.NewRequest().
		WithMapHeaders(opts.GetMapHeaders()).
		WithHeader("User-Agent", opts.ParseUserAgent()).
		WithOkCodes(200).
		WithJsonResponse(resp).
		WithJsonError(errResp)

	if _, sdkErr := s.VNetworkClient.Get(url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp).
			WithKVparameters("projectId", s.getProjectId()).
			WithParameters(opts.ToMap()).
			WithErrorCategories(sdkerror.ErrCatProductVNetwork)
	}

	return resp.ToEntityListTags(), nil
}

func (s *NetworkServiceInternalV1) CreateTagsWithEndpointId(opts ICreateTagsWithEndpointIdRequest) sdkerror.Error {
	url := createTagsWithEndpointIdUrl(s.VNetworkClient, opts)
	errResp := sdkerror.NewErrorResponse(sdkerror.NetworkGatewayErrorType)
	req := client.NewRequest().
		WithMapHeaders(opts.GetMapHeaders()).
		WithHeader("User-Agent", opts.ParseUserAgent()).
		WithOkCodes(200).
		WithJsonBody(opts.ToRequestBody()).
		WithJsonError(errResp)

	if _, sdkErr := s.VNetworkClient.Post(url, req); sdkErr != nil {
		return sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.WithErrorEndpointTagExisted(errResp),
			sdkerror.WithErrorEndpointTagNotFound(errResp)).
			WithKVparameters("projectId", s.getProjectId()).
			WithParameters(opts.ToMap()).
			WithErrorCategories(sdkerror.ErrCatProductVNetwork)
	}

	return nil
}

func (s *NetworkServiceInternalV1) DeleteTagOfEndpoint(opts IDeleteTagOfEndpointRequest) sdkerror.Error {
	url := deleteTagOfEndpointUrl(s.VNetworkClient, opts)
	errResp := sdkerror.NewErrorResponse(sdkerror.NetworkGatewayErrorType)
	req := client.NewRequest().
		WithMapHeaders(opts.GetMapHeaders()).
		WithHeader("User-Agent", opts.ParseUserAgent()).
		WithOkCodes(200).
		WithJsonError(errResp)

	if _, sdkErr := s.VNetworkClient.Delete(url, req); sdkErr != nil {
		return sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.WithErrorEndpointTagNotFound(errResp)).
			WithKVparameters("projectId", s.getProjectId()).
			WithParameters(opts.ToMap()).
			WithErrorCategories(sdkerror.ErrCatProductVNetwork)
	}

	return nil
}

func (s *NetworkServiceInternalV1) UpdateTagValueOfEndpoint(opts IUpdateTagValueOfEndpointRequest) sdkerror.Error {
	url := updateTagValueOfEndpointUrl(s.VNetworkClient, opts)
	errResp := sdkerror.NewErrorResponse(sdkerror.NetworkGatewayErrorType)
	req := client.NewRequest().
		WithMapHeaders(opts.GetMapHeaders()).
		WithHeader("User-Agent", opts.ParseUserAgent()).
		WithOkCodes(200).
		WithJsonBody(opts.ToRequestBody()).
		WithJsonError(errResp)

	if _, sdkErr := s.VNetworkClient.Put(url, req); sdkErr != nil {
		return sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.WithErrorEndpointTagNotFound(errResp)).
			WithKVparameters("projectId", s.getProjectId()).
			WithParameters(opts.ToMap()).
			WithErrorCategories(sdkerror.ErrCatProductVNetwork)
	}

	return nil
}

func (s *NetworkServiceInternalV1) CreateEndpoint(opts ICreateEndpointRequest) (*entity.Endpoint, sdkerror.Error) {
	url := createEndpointUrl(s.VNetworkClient)
	resp := new(CreateEndpointResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.NetworkGatewayErrorType)
	req := client.NewRequest().
		WithOkCodes(201).
		// WithUserId(s.getUserId()).
		WithJsonBody(opts.ToRequestBody(s.VNetworkClient)).
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
			WithParameters(opts.ToMap()).
			WithErrorCategories(sdkerror.ErrCatProductVNetwork)
	}

	return resp.ToEntityEndpoint(), nil
}
