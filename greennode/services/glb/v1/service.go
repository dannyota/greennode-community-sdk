package v1

import (
	"context"

	"github.com/dannyota/greennode-community-sdk/v2/greennode/client"
	"github.com/dannyota/greennode-community-sdk/v2/greennode/entity"
	sdkerror "github.com/dannyota/greennode-community-sdk/v2/greennode/sdkerror"
	"github.com/dannyota/greennode-community-sdk/v2/greennode/services/common"
)

type GLBServiceV1 struct {
	VLBClient *client.ServiceClient
}

const (
	defaultOffsetListGlobalLoadBalancer = 0
	defaultLimitListGlobalLoadBalancer  = 10
)

func (s *GLBServiceV1) ListGlobalPools(ctx context.Context, opts *ListGlobalPoolsRequest) (*entity.ListGlobalPools, error) {
	url := listGlobalPoolsURL(s.VLBClient, opts)
	resp := new(ListGlobalPoolsResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.GlobalLoadBalancerErrorType)
	req := client.NewRequest().
		WithOkCodes(200).
		WithJSONResponse(resp).
		WithJSONError(errResp)

	if _, sdkErr := s.VLBClient.Get(ctx, url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.EcGlobalLoadBalancerNotFound).
			WithKVparameters("loadBalancerId", opts.LoadBalancerID).
			AppendCategories(sdkerror.ErrCatProductVlb)
	}

	return resp.ToEntityListGlobalPools(), nil
}

func (s *GLBServiceV1) CreateGlobalPool(ctx context.Context, opts *CreateGlobalPoolRequest) (*entity.GlobalPool, error) {
	url := createGlobalPoolURL(s.VLBClient, opts)
	resp := new(entity.GlobalPool)
	errResp := sdkerror.NewErrorResponse(sdkerror.GlobalLoadBalancerErrorType)
	req := client.NewRequest().
		WithOkCodes(202).
		WithJSONBody(opts).
		WithJSONResponse(resp).
		WithJSONError(errResp)

	if _, sdkErr := s.VLBClient.Post(ctx, url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.EcGlobalLoadBalancerNotFound).
			WithParameters(common.StructToMap(opts)).
			AppendCategories(sdkerror.ErrCatProductVlb)
	}

	return resp, nil
}

func (s *GLBServiceV1) UpdateGlobalPool(ctx context.Context, opts *UpdateGlobalPoolRequest) (*entity.GlobalPool, error) {
	url := updateGlobalPoolURL(s.VLBClient, opts)
	resp := new(UpdateGlobalPoolResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.GlobalLoadBalancerErrorType)
	req := client.NewRequest().
		WithOkCodes(202).
		WithJSONBody(opts).
		WithJSONResponse(resp).
		WithJSONError(errResp)

	if _, sdkErr := s.VLBClient.Put(ctx, url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.EcGlobalLoadBalancerNotFound).
			WithParameters(common.StructToMap(opts)).
			AppendCategories(sdkerror.ErrCatProductVlb)
	}

	return resp.ToEntityPool(), nil
}

func (s *GLBServiceV1) DeleteGlobalPool(ctx context.Context, opts *DeleteGlobalPoolRequest) error {
	url := deleteGlobalPoolURL(s.VLBClient, opts)
	errResp := sdkerror.NewErrorResponse(sdkerror.GlobalLoadBalancerErrorType)
	req := client.NewRequest().
		WithOkCodes(202).
		WithJSONError(errResp)

	if _, sdkErr := s.VLBClient.Delete(ctx, url, req); sdkErr != nil {
		return sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.EcGlobalLoadBalancerNotFound).
			AppendCategories(sdkerror.ErrCatProductVlb)
	}

	return nil
}

func (s *GLBServiceV1) ListGlobalPoolMembers(ctx context.Context, opts *ListGlobalPoolMembersRequest) (*entity.ListGlobalPoolMembers, error) {
	url := listGlobalPoolMembersURL(s.VLBClient, opts)
	resp := new(ListGlobalPoolMembersResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.GlobalLoadBalancerErrorType)
	req := client.NewRequest().
		WithOkCodes(200).
		WithJSONResponse(resp).
		WithJSONError(errResp)

	if _, sdkErr := s.VLBClient.Get(ctx, url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.EcGlobalLoadBalancerNotFound).
			AppendCategories(sdkerror.ErrCatProductVlb)
	}

	return resp.ToEntityListGlobalPoolMembers(), nil
}

func (s *GLBServiceV1) GetGlobalPoolMember(ctx context.Context, opts *GetGlobalPoolMemberRequest) (*entity.GlobalPoolMember, error) {
	url := getGlobalPoolMemberURL(s.VLBClient, opts)
	resp := new(GetGlobalPoolMemberResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.GlobalLoadBalancerErrorType)
	req := client.NewRequest().
		WithOkCodes(200).
		WithJSONResponse(resp).
		WithJSONError(errResp)

	if _, sdkErr := s.VLBClient.Get(ctx, url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.EcGlobalLoadBalancerNotFound).
			WithKVparameters("loadBalancerId", opts.LoadBalancerID).
			WithKVparameters("poolId", opts.PoolID).
			WithKVparameters("poolMemberId", opts.PoolMemberID).
			AppendCategories(sdkerror.ErrCatProductVlb)
	}

	return resp.ToEntityGlobalPoolMember(), nil
}

func (s *GLBServiceV1) DeleteGlobalPoolMember(ctx context.Context, opts *DeleteGlobalPoolMemberRequest) error {
	url := deleteGlobalPoolMemberURL(s.VLBClient, opts)
	errResp := sdkerror.NewErrorResponse(sdkerror.GlobalLoadBalancerErrorType)
	req := client.NewRequest().
		WithOkCodes(202).
		WithJSONError(errResp)

	if _, sdkErr := s.VLBClient.Delete(ctx, url, req); sdkErr != nil {
		return sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.EcGlobalLoadBalancerNotFound).
			WithKVparameters("loadBalancerId", opts.LoadBalancerID).
			WithKVparameters("poolId", opts.PoolID).
			WithKVparameters("poolMemberId", opts.PoolMemberID).
			AppendCategories(sdkerror.ErrCatProductVlb)
	}

	return nil
}

func (s *GLBServiceV1) UpdateGlobalPoolMember(ctx context.Context, opts *UpdateGlobalPoolMemberRequest) (*entity.GlobalPoolMember, error) {
	url := updateGlobalPoolMemberURL(s.VLBClient, opts)
	resp := new(UpdateGlobalPoolMemberResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.GlobalLoadBalancerErrorType)
	req := client.NewRequest().
		WithOkCodes(202).
		WithJSONBody(opts).
		WithJSONResponse(resp).
		WithJSONError(errResp)

	if _, sdkErr := s.VLBClient.Put(ctx, url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.EcGlobalLoadBalancerNotFound).
			WithParameters(common.StructToMap(opts)).
			AppendCategories(sdkerror.ErrCatProductVlb)
	}

	return resp.ToEntityGlobalPoolMember(), nil
}

func (s *GLBServiceV1) PatchGlobalPoolMembers(ctx context.Context, opts *PatchGlobalPoolMembersRequest) error {
	url := patchGlobalPoolMembersURL(s.VLBClient, opts)
	errResp := sdkerror.NewErrorResponse(sdkerror.GlobalLoadBalancerErrorType)
	req := client.NewRequest().
		WithOkCodes(202).
		WithJSONBody(opts).
		WithJSONError(errResp)

	if _, sdkErr := s.VLBClient.Patch(ctx, url, req); sdkErr != nil {
		return sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.EcGlobalLoadBalancerNotFound).
			WithParameters(common.StructToMap(opts)).
			AppendCategories(sdkerror.ErrCatProductVlb)
	}

	return nil
}

func (s *GLBServiceV1) ListGlobalListeners(ctx context.Context, opts *ListGlobalListenersRequest) (*entity.ListGlobalListeners, error) {
	url := listGlobalListenersURL(s.VLBClient, opts)
	resp := new(ListGlobalListenersResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.GlobalLoadBalancerErrorType)
	req := client.NewRequest().
		WithOkCodes(200).
		WithJSONResponse(resp).
		WithJSONError(errResp)

	if _, sdkErr := s.VLBClient.Get(ctx, url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.EcGlobalLoadBalancerNotFound).
			AppendCategories(sdkerror.ErrCatProductVlb)
	}

	return resp.ToEntityListGlobalListeners(), nil
}

func (s *GLBServiceV1) CreateGlobalListener(ctx context.Context, opts *CreateGlobalListenerRequest) (*entity.GlobalListener, error) {
	url := createGlobalListenerURL(s.VLBClient, opts)
	resp := new(entity.GlobalListener)
	errResp := sdkerror.NewErrorResponse(sdkerror.GlobalLoadBalancerErrorType)
	req := client.NewRequest().
		WithOkCodes(202).
		WithJSONBody(opts).
		WithJSONResponse(resp).
		WithJSONError(errResp)

	if _, sdkErr := s.VLBClient.Post(ctx, url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.EcGlobalLoadBalancerNotFound).
			WithParameters(common.StructToMap(opts)).
			AppendCategories(sdkerror.ErrCatProductVlb)
	}

	return resp, nil
}

func (s *GLBServiceV1) UpdateGlobalListener(ctx context.Context, opts *UpdateGlobalListenerRequest) (*entity.GlobalListener, error) {
	url := updateGlobalListenerURL(s.VLBClient, opts)
	resp := new(entity.GlobalListener)
	errResp := sdkerror.NewErrorResponse(sdkerror.GlobalLoadBalancerErrorType)
	req := client.NewRequest().
		WithOkCodes(202).
		WithJSONBody(opts).
		WithJSONResponse(resp).
		WithJSONError(errResp)

	if _, sdkErr := s.VLBClient.Put(ctx, url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.EcGlobalLoadBalancerNotFound).
			WithParameters(common.StructToMap(opts)).
			AppendCategories(sdkerror.ErrCatProductVlb)
	}

	return resp, nil
}

func (s *GLBServiceV1) GetGlobalListener(ctx context.Context, opts *GetGlobalListenerRequest) (*entity.GlobalListener, error) {
	url := getGlobalListenerURL(s.VLBClient, opts)
	resp := new(entity.GlobalListener)
	errResp := sdkerror.NewErrorResponse(sdkerror.GlobalLoadBalancerErrorType)
	req := client.NewRequest().
		WithOkCodes(200).
		WithJSONResponse(resp).
		WithJSONError(errResp)

	if _, sdkErr := s.VLBClient.Get(ctx, url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.EcGlobalLoadBalancerNotFound).
			WithKVparameters("loadBalancerId", opts.LoadBalancerID).
			WithKVparameters("listenerId", opts.ListenerID).
			AppendCategories(sdkerror.ErrCatProductVlb)
	}

	return resp, nil
}

func (s *GLBServiceV1) DeleteGlobalListener(ctx context.Context, opts *DeleteGlobalListenerRequest) error {
	url := deleteGlobalListenerURL(s.VLBClient, opts)
	errResp := sdkerror.NewErrorResponse(sdkerror.GlobalLoadBalancerErrorType)
	req := client.NewRequest().
		WithOkCodes(202).
		WithJSONError(errResp)

	if _, sdkErr := s.VLBClient.Delete(ctx, url, req); sdkErr != nil {
		return sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.EcGlobalLoadBalancerNotFound).
			AppendCategories(sdkerror.ErrCatProductVlb)
	}

	return nil
}

func (s *GLBServiceV1) ListGlobalLoadBalancers(ctx context.Context, opts *ListGlobalLoadBalancersRequest) (*entity.ListGlobalLoadBalancers, error) {
	url := listGlobalLoadBalancersURL(s.VLBClient, opts)
	resp := new(entity.ListGlobalLoadBalancers)
	errResp := sdkerror.NewErrorResponse(sdkerror.GlobalLoadBalancerErrorType)
	req := client.NewRequest().
		WithOkCodes(200).
		WithJSONResponse(resp).
		WithJSONError(errResp)

	if _, sdkErr := s.VLBClient.Get(ctx, url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp).
			AppendCategories(sdkerror.ErrCatProductVlb)
	}

	return resp, nil
}

func (s *GLBServiceV1) CreateGlobalLoadBalancer(
	ctx context.Context,
	opts *CreateGlobalLoadBalancerRequest,
) (*entity.GlobalLoadBalancer, error) {
	url := createGlobalLoadBalancerURL(s.VLBClient, opts)
	resp := new(CreateGlobalLoadBalancerResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.GlobalLoadBalancerErrorType)
	req := client.NewRequest().
		WithOkCodes(202).
		WithJSONBody(opts).
		WithJSONResponse(resp).
		WithJSONError(errResp)

	if _, sdkErr := s.VLBClient.Post(ctx, url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp).
			WithParameters(common.StructToMap(opts)).
			AppendCategories(sdkerror.ErrCatProductVlb)
	}

	return resp.ToEntityGlobalLoadBalancer(), nil
}

func (s *GLBServiceV1) DeleteGlobalLoadBalancer(ctx context.Context, opts *DeleteGlobalLoadBalancerRequest) error {
	url := deleteGlobalLoadBalancerURL(s.VLBClient, opts)
	errResp := sdkerror.NewErrorResponse(sdkerror.GlobalLoadBalancerErrorType)
	req := client.NewRequest().
		WithOkCodes(202).
		WithJSONError(errResp)

	if _, sdkErr := s.VLBClient.Delete(ctx, url, req); sdkErr != nil {
		return sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.EcGlobalLoadBalancerNotFound).
			AppendCategories(sdkerror.ErrCatProductVlb)
	}

	return nil
}

func (s *GLBServiceV1) GetGlobalLoadBalancerByID(
	ctx context.Context,
	opts *GetGlobalLoadBalancerByIDRequest,
) (*entity.GlobalLoadBalancer, error) {
	url := getGlobalLoadBalancerByIDURL(s.VLBClient, opts)
	resp := new(entity.GlobalLoadBalancer)
	errResp := sdkerror.NewErrorResponse(sdkerror.GlobalLoadBalancerErrorType)
	req := client.NewRequest().
		WithOkCodes(200).
		WithJSONResponse(resp).
		WithJSONError(errResp)

	if _, sdkErr := s.VLBClient.Get(ctx, url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.EcGlobalLoadBalancerNotFound).
			AppendCategories(sdkerror.ErrCatProductVlb)
	}

	return resp, nil
}

func (s *GLBServiceV1) ListGlobalPackages(ctx context.Context, opts *ListGlobalPackagesRequest) (*entity.ListGlobalPackages, error) {
	url := listGlobalPackagesURL(s.VLBClient, opts)
	resp := new(ListGlobalPackagesResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.GlobalLoadBalancerErrorType)
	req := client.NewRequest().
		WithOkCodes(200).
		WithJSONResponse(resp).
		WithJSONError(errResp)

	if _, sdkErr := s.VLBClient.Get(ctx, url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp).
			AppendCategories(sdkerror.ErrCatProductVlb)
	}

	return resp.ToEntityListGlobalPackages(), nil
}

func (s *GLBServiceV1) ListGlobalRegions(ctx context.Context, opts *ListGlobalRegionsRequest) (*entity.ListGlobalRegions, error) {
	url := listGlobalRegionsURL(s.VLBClient, opts)
	resp := new(ListGlobalRegionsResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.GlobalLoadBalancerErrorType)
	req := client.NewRequest().
		WithOkCodes(200).
		WithJSONResponse(resp).
		WithJSONError(errResp)

	if _, sdkErr := s.VLBClient.Get(ctx, url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp).
			AppendCategories(sdkerror.ErrCatProductVlb)
	}

	return resp.ToEntityListGlobalRegions(), nil
}

func (s *GLBServiceV1) GetGlobalLoadBalancerUsageHistories(
	ctx context.Context,
	opts *GetGlobalLoadBalancerUsageHistoriesRequest,
) (*entity.ListGlobalLoadBalancerUsageHistories, error) {
	url := getGlobalLoadBalancerUsageHistoriesURL(s.VLBClient, opts)
	resp := new(entity.ListGlobalLoadBalancerUsageHistories)
	errResp := sdkerror.NewErrorResponse(sdkerror.GlobalLoadBalancerErrorType)
	req := client.NewRequest().
		WithOkCodes(200).
		WithJSONResponse(resp).
		WithJSONError(errResp)

	if _, sdkErr := s.VLBClient.Get(ctx, url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.EcGlobalLoadBalancerNotFound).
			AppendCategories(sdkerror.ErrCatProductVlb)
	}

	return resp, nil
}
