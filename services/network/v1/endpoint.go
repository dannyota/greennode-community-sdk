package v1

import (
	"context"

	"danny.vn/greennode/client"
	sdkerror "danny.vn/greennode/sdkerror"
	"danny.vn/greennode/services/common"
	"danny.vn/greennode/types"
)

type NetworkServiceV1 struct {
	Client *client.ServiceClient
}

type NetworkServiceInternalV1 struct {
	Client *client.ServiceClient
}

func (s *NetworkServiceV1) GetEndpointByID(ctx context.Context, opts *GetEndpointByIDRequest) (*Endpoint, error) {
	url := getEndpointByIDURL(s.Client, opts)
	resp := new(GetEndpointByIDResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.NetworkGatewayErrorType)
	req := client.NewRequest().
		WithOKCodes(200).
		// WithUserId(s.getUserId()).
		WithJSONResponse(resp).
		WithJSONError(errResp)

	if _, sdkErr := s.Client.Get(ctx, url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp).
			WithKVparameters(
				"endpointId", opts.EndpointID,
				"projectId", s.Client.ProjectID).
			AppendCategories(sdkerror.ErrCatProductVNetwork)
	}

	return resp.ToEntityEndpoint(), nil
}

func (s *NetworkServiceV1) CreateEndpoint(ctx context.Context, opts *CreateEndpointRequest) (*Endpoint, error) {
	url := createEndpointURL(s.Client)
	resp := new(CreateEndpointResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.NetworkGatewayErrorType)
	req := client.NewRequest().
		WithOKCodes(201).
		WithUserID(opts.ResourceInfo.PortalUserID).
		WithJSONBody(opts.ToRequestBody(s.Client)).
		WithJSONResponse(resp).
		WithJSONError(errResp)

	if _, sdkErr := s.Client.Post(ctx, url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.EcVNetworkEndpointOfVpcExists,
			sdkerror.EcVNetworkLockOnProcess,
			sdkerror.EcVServerNetworkNotFound,
			sdkerror.EcPurchaseIssue,
			sdkerror.EcPaymentMethodNotAllow,
			sdkerror.EcCreditNotEnough,
			sdkerror.EcVServerSubnetNotFound,
			sdkerror.EcVNetworkEndpointPackageNotBelongToEndpointService,
			sdkerror.EcVNetworkContainInvalidCharacter).
			WithParameters(common.StructToMap(opts)).
			AppendCategories(sdkerror.ErrCatProductVNetwork)
	}

	return resp.ToEntityEndpoint(), nil
}

func (s *NetworkServiceV1) DeleteEndpointByID(ctx context.Context, opts *DeleteEndpointByIDRequest) error {
	url := deleteEndpointByIDURL(s.Client, opts)
	errResp := sdkerror.NewErrorResponse(sdkerror.NetworkGatewayErrorType)
	req := client.NewRequest().
		WithOKCodes(200).
		WithJSONBody(opts.ToRequestBody(s.Client)).
		// WithUserId(s.getUserId()).
		WithJSONError(errResp)

	if _, sdkErr := s.Client.Delete(ctx, url, req); sdkErr != nil {
		return sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.EcVNetworkEndpointStatusInvalid,
			sdkerror.EcVServerNetworkNotFound,
			sdkerror.EcVServerSubnetNotFound).
			WithParameters(common.StructToMap(opts)).
			AppendCategories(sdkerror.ErrCatProductVNetwork)
	}

	return nil
}

func (s *NetworkServiceV1) ListVNetworkRegions(ctx context.Context, opts *ListVNetworkRegionsRequest) (*ListVNetworkRegions, error) {
	url := listVNetworkRegionsURL(s.Client)
	resp := new(ListVNetworkRegionsResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.NetworkGatewayErrorType)
	req := client.NewRequest().
		WithOKCodes(200).
		WithJSONResponse(resp).
		WithJSONError(errResp)

	if _, sdkErr := s.Client.Get(ctx, url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp).
			AppendCategories(sdkerror.ErrCatProductVNetwork)
	}

	return resp.ToEntityListVNetworkRegions(), nil
}

func (s *NetworkServiceV1) ListEndpoints(ctx context.Context, opts *ListEndpointsRequest) (*ListEndpoints, error) {
	url := listEndpointsURL(s.Client, opts)
	resp := new(ListEndpointsResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.NetworkGatewayErrorType)
	req := client.NewRequest().
		WithOKCodes(200).
		// WithUserId(s.getUserId()).
		WithJSONResponse(resp).
		WithJSONError(errResp)

	if _, sdkErr := s.Client.Get(ctx, url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp).
			WithKVparameters("projectId", s.Client.ProjectID).
			WithParameters(common.StructToMap(opts)).
			AppendCategories(sdkerror.ErrCatProductVNetwork)
	}

	return resp.ToEntityListEndpoints(), nil
}

// ________________________________________________________________________ NetworkServiceInternalV1

func (s *NetworkServiceInternalV1) ListTagsByEndpointID(ctx context.Context, opts *ListTagsByEndpointIDRequest) (*types.ListTags, error) {
	url := listTagsByEndpointIDURL(s.Client, opts)
	resp := new(ListTagsByEndpointIDResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.NetworkGatewayErrorType)
	req := client.NewRequest().
		WithUserID(opts.PortalUser.ID).
		WithOKCodes(200).
		WithJSONResponse(resp).
		WithJSONError(errResp)

	if _, sdkErr := s.Client.Get(ctx, url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp).
			WithKVparameters("projectId", s.Client.ProjectID).
			WithParameters(common.StructToMap(opts)).
			AppendCategories(sdkerror.ErrCatProductVNetwork)
	}

	return resp.ToEntityListTags(), nil
}

func (s *NetworkServiceInternalV1) CreateTagsWithEndpointID(ctx context.Context, opts *CreateTagsWithEndpointIDRequest) error {
	url := createTagsWithEndpointIDURL(s.Client, opts)
	errResp := sdkerror.NewErrorResponse(sdkerror.NetworkGatewayErrorType)
	req := client.NewRequest().
		WithUserID(opts.PortalUser.ID).
		WithOKCodes(200).
		WithJSONBody(opts).
		WithJSONError(errResp)

	if _, sdkErr := s.Client.Post(ctx, url, req); sdkErr != nil {
		return sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.EcVNetworkEndpointTagExisted,
			sdkerror.EcVNetworkEndpointTagNotFound).
			WithKVparameters("projectId", s.Client.ProjectID).
			WithParameters(common.StructToMap(opts)).
			AppendCategories(sdkerror.ErrCatProductVNetwork)
	}

	return nil
}

func (s *NetworkServiceInternalV1) DeleteTagOfEndpoint(ctx context.Context, opts *DeleteTagOfEndpointRequest) error {
	url := deleteTagOfEndpointURL(s.Client, opts)
	errResp := sdkerror.NewErrorResponse(sdkerror.NetworkGatewayErrorType)
	req := client.NewRequest().
		WithUserID(opts.PortalUser.ID).
		WithOKCodes(200).
		WithJSONError(errResp)

	if _, sdkErr := s.Client.Delete(ctx, url, req); sdkErr != nil {
		return sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.EcVNetworkEndpointTagNotFound).
			WithKVparameters("projectId", s.Client.ProjectID).
			WithParameters(common.StructToMap(opts)).
			AppendCategories(sdkerror.ErrCatProductVNetwork)
	}

	return nil
}

func (s *NetworkServiceInternalV1) UpdateTagValueOfEndpoint(ctx context.Context, opts *UpdateTagValueOfEndpointRequest) error {
	url := updateTagValueOfEndpointURL(s.Client, opts)
	errResp := sdkerror.NewErrorResponse(sdkerror.NetworkGatewayErrorType)
	req := client.NewRequest().
		WithUserID(opts.PortalUser.ID).
		WithOKCodes(200).
		WithJSONBody(opts).
		WithJSONError(errResp)

	if _, sdkErr := s.Client.Put(ctx, url, req); sdkErr != nil {
		return sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.EcVNetworkEndpointTagNotFound).
			WithKVparameters("projectId", s.Client.ProjectID).
			WithParameters(common.StructToMap(opts)).
			AppendCategories(sdkerror.ErrCatProductVNetwork)
	}

	return nil
}

func (s *NetworkServiceInternalV1) CreateEndpoint(ctx context.Context, opts *CreateEndpointRequest) (*Endpoint, error) {
	url := createEndpointURL(s.Client)
	resp := new(CreateEndpointResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.NetworkGatewayErrorType)
	req := client.NewRequest().
		WithOKCodes(201).
		// WithUserId(s.getUserId()).
		WithJSONBody(opts.ToRequestBody(s.Client)).
		WithJSONResponse(resp).
		WithJSONError(errResp)

	if _, sdkErr := s.Client.Post(ctx, url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.EcVNetworkEndpointOfVpcExists,
			sdkerror.EcVNetworkLockOnProcess,
			sdkerror.EcVServerNetworkNotFound,
			sdkerror.EcPurchaseIssue,
			sdkerror.EcPaymentMethodNotAllow,
			sdkerror.EcCreditNotEnough,
			sdkerror.EcVServerSubnetNotFound,
			sdkerror.EcVNetworkEndpointPackageNotBelongToEndpointService,
			sdkerror.EcVNetworkContainInvalidCharacter).
			WithParameters(common.StructToMap(opts)).
			AppendCategories(sdkerror.ErrCatProductVNetwork)
	}

	return resp.ToEntityEndpoint(), nil
}
