package v1

import (
	"context"

	"github.com/dannyota/greennode-community-sdk/v2/greennode/client"
	"github.com/dannyota/greennode-community-sdk/v2/greennode/entity"
	sdkerror "github.com/dannyota/greennode-community-sdk/v2/greennode/sdkerror"
	"github.com/dannyota/greennode-community-sdk/v2/greennode/services/common"
)

type NetworkServiceV1 struct {
	VNetworkClient *client.ServiceClient
}

func (s *NetworkServiceV1) getProjectID() string {
	return s.VNetworkClient.GetProjectID()
}

type NetworkServiceInternalV1 struct {
	VNetworkClient *client.ServiceClient
}

func (s *NetworkServiceInternalV1) getProjectID() string {
	return s.VNetworkClient.GetProjectID()
}

func (s *NetworkServiceV1) GetEndpointByID(ctx context.Context, opts *GetEndpointByIDRequest) (*entity.Endpoint, error) {
	url := getEndpointByIDURL(s.VNetworkClient, opts)
	resp := new(GetEndpointByIDResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.NetworkGatewayErrorType)
	req := client.NewRequest().
		WithOkCodes(200).
		// WithUserId(s.getUserId()).
		WithJSONResponse(resp).
		WithJSONError(errResp)

	if _, sdkErr := s.VNetworkClient.Get(ctx, url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp).
			WithKVparameters(
				"endpointId", opts.GetEndpointID(),
				"projectId", s.getProjectID()).
			WithErrorCategories(sdkerror.ErrCatProductVNetwork)
	}

	return resp.ToEntityEndpoint(), nil
}

func (s *NetworkServiceV1) CreateEndpoint(ctx context.Context, opts *CreateEndpointRequest) (*entity.Endpoint, error) {
	url := createEndpointURL(s.VNetworkClient)
	resp := new(CreateEndpointResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.NetworkGatewayErrorType)
	req := client.NewRequest().
		WithOkCodes(201).
		WithUserID(opts.GetPortalUserID()).
		WithJSONBody(opts.ToRequestBody(s.VNetworkClient)).
		WithJSONResponse(resp).
		WithJSONError(errResp)

	if _, sdkErr := s.VNetworkClient.Post(ctx, url, req); sdkErr != nil {
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
			WithErrorCategories(sdkerror.ErrCatProductVNetwork)
	}

	return resp.ToEntityEndpoint(), nil
}

func (s *NetworkServiceV1) DeleteEndpointByID(ctx context.Context, opts *DeleteEndpointByIDRequest) error {
	url := deleteEndpointByIDURL(s.VNetworkClient, opts)
	errResp := sdkerror.NewErrorResponse(sdkerror.NetworkGatewayErrorType)
	req := client.NewRequest().
		WithOkCodes(200).
		WithJSONBody(opts.ToRequestBody(s.VNetworkClient)).
		// WithUserId(s.getUserId()).
		WithJSONError(errResp)

	if _, sdkErr := s.VNetworkClient.Delete(ctx, url, req); sdkErr != nil {
		return sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.EcVNetworkEndpointStatusInvalid,
			sdkerror.EcVServerNetworkNotFound,
			sdkerror.EcVServerSubnetNotFound).
			WithParameters(common.StructToMap(opts)).
			WithErrorCategories(sdkerror.ErrCatProductVNetwork)
	}

	return nil
}

func (s *NetworkServiceV1) ListEndpoints(ctx context.Context, opts *ListEndpointsRequest) (*entity.ListEndpoints, error) {
	url := listEndpointsURL(s.VNetworkClient, opts)
	resp := new(ListEndpointsResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.NetworkGatewayErrorType)
	req := client.NewRequest().
		WithOkCodes(200).
		// WithUserId(s.getUserId()).
		WithJSONResponse(resp).
		WithJSONError(errResp)

	if _, sdkErr := s.VNetworkClient.Get(ctx, url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp).
			WithKVparameters("projectId", s.getProjectID()).
			WithParameters(common.StructToMap(opts)).
			WithErrorCategories(sdkerror.ErrCatProductVNetwork)
	}

	return resp.ToEntityListEndpoints(), nil
}

// ________________________________________________________________________ NetworkServiceInternalV1

func (s *NetworkServiceInternalV1) ListTagsByEndpointID(ctx context.Context, opts *ListTagsByEndpointIDRequest) (*entity.ListTags, error) {
	url := listTagsByEndpointIDURL(s.VNetworkClient, opts)
	resp := new(ListTagsByEndpointIDResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.NetworkGatewayErrorType)
	req := client.NewRequest().
		WithMapHeaders(opts.GetMapHeaders()).
		WithOkCodes(200).
		WithJSONResponse(resp).
		WithJSONError(errResp)

	if _, sdkErr := s.VNetworkClient.Get(ctx, url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp).
			WithKVparameters("projectId", s.getProjectID()).
			WithParameters(common.StructToMap(opts)).
			WithErrorCategories(sdkerror.ErrCatProductVNetwork)
	}

	return resp.ToEntityListTags(), nil
}

func (s *NetworkServiceInternalV1) CreateTagsWithEndpointID(ctx context.Context, opts *CreateTagsWithEndpointIDRequest) error {
	url := createTagsWithEndpointIDURL(s.VNetworkClient, opts)
	errResp := sdkerror.NewErrorResponse(sdkerror.NetworkGatewayErrorType)
	req := client.NewRequest().
		WithMapHeaders(opts.GetMapHeaders()).
		WithOkCodes(200).
		WithJSONBody(opts).
		WithJSONError(errResp)

	if _, sdkErr := s.VNetworkClient.Post(ctx, url, req); sdkErr != nil {
		return sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.EcVNetworkEndpointTagExisted,
			sdkerror.EcVNetworkEndpointTagNotFound).
			WithKVparameters("projectId", s.getProjectID()).
			WithParameters(common.StructToMap(opts)).
			WithErrorCategories(sdkerror.ErrCatProductVNetwork)
	}

	return nil
}

func (s *NetworkServiceInternalV1) DeleteTagOfEndpoint(ctx context.Context, opts *DeleteTagOfEndpointRequest) error {
	url := deleteTagOfEndpointURL(s.VNetworkClient, opts)
	errResp := sdkerror.NewErrorResponse(sdkerror.NetworkGatewayErrorType)
	req := client.NewRequest().
		WithMapHeaders(opts.GetMapHeaders()).
		WithOkCodes(200).
		WithJSONError(errResp)

	if _, sdkErr := s.VNetworkClient.Delete(ctx, url, req); sdkErr != nil {
		return sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.EcVNetworkEndpointTagNotFound).
			WithKVparameters("projectId", s.getProjectID()).
			WithParameters(common.StructToMap(opts)).
			WithErrorCategories(sdkerror.ErrCatProductVNetwork)
	}

	return nil
}

func (s *NetworkServiceInternalV1) UpdateTagValueOfEndpoint(ctx context.Context, opts *UpdateTagValueOfEndpointRequest) error {
	url := updateTagValueOfEndpointURL(s.VNetworkClient, opts)
	errResp := sdkerror.NewErrorResponse(sdkerror.NetworkGatewayErrorType)
	req := client.NewRequest().
		WithMapHeaders(opts.GetMapHeaders()).
		WithOkCodes(200).
		WithJSONBody(opts).
		WithJSONError(errResp)

	if _, sdkErr := s.VNetworkClient.Put(ctx, url, req); sdkErr != nil {
		return sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.EcVNetworkEndpointTagNotFound).
			WithKVparameters("projectId", s.getProjectID()).
			WithParameters(common.StructToMap(opts)).
			WithErrorCategories(sdkerror.ErrCatProductVNetwork)
	}

	return nil
}

func (s *NetworkServiceInternalV1) CreateEndpoint(ctx context.Context, opts *CreateEndpointRequest) (*entity.Endpoint, error) {
	url := createEndpointURL(s.VNetworkClient)
	resp := new(CreateEndpointResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.NetworkGatewayErrorType)
	req := client.NewRequest().
		WithOkCodes(201).
		// WithUserId(s.getUserId()).
		WithJSONBody(opts.ToRequestBody(s.VNetworkClient)).
		WithJSONResponse(resp).
		WithJSONError(errResp)

	if _, sdkErr := s.VNetworkClient.Post(ctx, url, req); sdkErr != nil {
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
			WithErrorCategories(sdkerror.ErrCatProductVNetwork)
	}

	return resp.ToEntityEndpoint(), nil
}
