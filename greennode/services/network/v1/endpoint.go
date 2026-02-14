package v1

import (
	"github.com/dannyota/greennode-community-sdk/v2/greennode/client"
	"github.com/dannyota/greennode-community-sdk/v2/greennode/entity"
	sdkerror "github.com/dannyota/greennode-community-sdk/v2/greennode/sdkerror"
)

func (s *NetworkServiceV1) GetEndpointByID(opts *GetEndpointByIDRequest) (*entity.Endpoint, sdkerror.Error) {
	url := getEndpointByIDURL(s.VNetworkClient, opts)
	resp := new(GetEndpointByIDResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.NetworkGatewayErrorType)
	req := client.NewRequest().
		WithOkCodes(200).
		// WithUserId(s.getUserId()).
		WithJSONResponse(resp).
		WithJSONError(errResp)

	if _, sdkErr := s.VNetworkClient.Get(url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp).
			WithKVparameters(
				"endpointId", opts.GetEndpointID(),
				"projectId", s.getProjectID()).
			WithErrorCategories(sdkerror.ErrCatProductVNetwork)
	}

	return resp.ToEntityEndpoint(), nil
}

func (s *NetworkServiceV1) CreateEndpoint(opts *CreateEndpointRequest) (*entity.Endpoint, sdkerror.Error) {
	url := createEndpointURL(s.VNetworkClient)
	resp := new(CreateEndpointResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.NetworkGatewayErrorType)
	req := client.NewRequest().
		WithOkCodes(201).
		WithUserID(opts.GetPortalUserID()).
		WithJSONBody(opts.ToRequestBody(s.VNetworkClient)).
		WithJSONResponse(resp).
		WithJSONError(errResp)

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

func (s *NetworkServiceV1) DeleteEndpointByID(opts *DeleteEndpointByIDRequest) sdkerror.Error {
	url := deleteEndpointByIDURL(s.VNetworkClient, opts)
	errResp := sdkerror.NewErrorResponse(sdkerror.NetworkGatewayErrorType)
	req := client.NewRequest().
		WithHeader("User-Agent", opts.ParseUserAgent()).
		WithOkCodes(200).
		WithJSONBody(opts.ToRequestBody(s.VNetworkClient)).
		// WithUserId(s.getUserId()).
		WithJSONError(errResp)

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

func (s *NetworkServiceV1) ListEndpoints(opts *ListEndpointsRequest) (*entity.ListEndpoints, sdkerror.Error) {
	url := listEndpointsURL(s.VNetworkClient, opts)
	resp := new(ListEndpointsResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.NetworkGatewayErrorType)
	req := client.NewRequest().
		WithHeader("User-Agent", opts.ParseUserAgent()).
		WithOkCodes(200).
		// WithUserId(s.getUserId()).
		WithJSONResponse(resp).
		WithJSONError(errResp)

	if _, sdkErr := s.VNetworkClient.Get(url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp).
			WithKVparameters("projectId", s.getProjectID()).
			WithParameters(opts.ToMap()).
			WithErrorCategories(sdkerror.ErrCatProductVNetwork)
	}

	return resp.ToEntityListEndpoints(), nil
}

// ________________________________________________________________________ NetworkServiceInternalV1

func (s *NetworkServiceInternalV1) ListTagsByEndpointID(opts *ListTagsByEndpointIDRequest) (*entity.ListTags, sdkerror.Error) {
	url := listTagsByEndpointIDURL(s.VNetworkClient, opts)
	resp := new(ListTagsByEndpointIDResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.NetworkGatewayErrorType)
	req := client.NewRequest().
		WithMapHeaders(opts.GetMapHeaders()).
		WithHeader("User-Agent", opts.ParseUserAgent()).
		WithOkCodes(200).
		WithJSONResponse(resp).
		WithJSONError(errResp)

	if _, sdkErr := s.VNetworkClient.Get(url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp).
			WithKVparameters("projectId", s.getProjectID()).
			WithParameters(opts.ToMap()).
			WithErrorCategories(sdkerror.ErrCatProductVNetwork)
	}

	return resp.ToEntityListTags(), nil
}

func (s *NetworkServiceInternalV1) CreateTagsWithEndpointID(opts *CreateTagsWithEndpointIDRequest) sdkerror.Error {
	url := createTagsWithEndpointIDURL(s.VNetworkClient, opts)
	errResp := sdkerror.NewErrorResponse(sdkerror.NetworkGatewayErrorType)
	req := client.NewRequest().
		WithMapHeaders(opts.GetMapHeaders()).
		WithHeader("User-Agent", opts.ParseUserAgent()).
		WithOkCodes(200).
		WithJSONBody(opts.ToRequestBody()).
		WithJSONError(errResp)

	if _, sdkErr := s.VNetworkClient.Post(url, req); sdkErr != nil {
		return sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.WithErrorEndpointTagExisted(errResp),
			sdkerror.WithErrorEndpointTagNotFound(errResp)).
			WithKVparameters("projectId", s.getProjectID()).
			WithParameters(opts.ToMap()).
			WithErrorCategories(sdkerror.ErrCatProductVNetwork)
	}

	return nil
}

func (s *NetworkServiceInternalV1) DeleteTagOfEndpoint(opts *DeleteTagOfEndpointRequest) sdkerror.Error {
	url := deleteTagOfEndpointURL(s.VNetworkClient, opts)
	errResp := sdkerror.NewErrorResponse(sdkerror.NetworkGatewayErrorType)
	req := client.NewRequest().
		WithMapHeaders(opts.GetMapHeaders()).
		WithHeader("User-Agent", opts.ParseUserAgent()).
		WithOkCodes(200).
		WithJSONError(errResp)

	if _, sdkErr := s.VNetworkClient.Delete(url, req); sdkErr != nil {
		return sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.WithErrorEndpointTagNotFound(errResp)).
			WithKVparameters("projectId", s.getProjectID()).
			WithParameters(opts.ToMap()).
			WithErrorCategories(sdkerror.ErrCatProductVNetwork)
	}

	return nil
}

func (s *NetworkServiceInternalV1) UpdateTagValueOfEndpoint(opts *UpdateTagValueOfEndpointRequest) sdkerror.Error {
	url := updateTagValueOfEndpointURL(s.VNetworkClient, opts)
	errResp := sdkerror.NewErrorResponse(sdkerror.NetworkGatewayErrorType)
	req := client.NewRequest().
		WithMapHeaders(opts.GetMapHeaders()).
		WithHeader("User-Agent", opts.ParseUserAgent()).
		WithOkCodes(200).
		WithJSONBody(opts.ToRequestBody()).
		WithJSONError(errResp)

	if _, sdkErr := s.VNetworkClient.Put(url, req); sdkErr != nil {
		return sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.WithErrorEndpointTagNotFound(errResp)).
			WithKVparameters("projectId", s.getProjectID()).
			WithParameters(opts.ToMap()).
			WithErrorCategories(sdkerror.ErrCatProductVNetwork)
	}

	return nil
}

func (s *NetworkServiceInternalV1) CreateEndpoint(opts *CreateEndpointRequest) (*entity.Endpoint, sdkerror.Error) {
	url := createEndpointURL(s.VNetworkClient)
	resp := new(CreateEndpointResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.NetworkGatewayErrorType)
	req := client.NewRequest().
		WithOkCodes(201).
		// WithUserId(s.getUserId()).
		WithJSONBody(opts.ToRequestBody(s.VNetworkClient)).
		WithJSONResponse(resp).
		WithJSONError(errResp)

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
