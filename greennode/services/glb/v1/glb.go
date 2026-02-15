package v1

import (
	"fmt"

	"github.com/dannyota/greennode-community-sdk/v2/greennode/client"
	"github.com/dannyota/greennode-community-sdk/v2/greennode/entity"
	sdkerror "github.com/dannyota/greennode-community-sdk/v2/greennode/sdkerror"
	"github.com/dannyota/greennode-community-sdk/v2/greennode/services/common"
)

func (s *GLBServiceV1) ListGlobalPools(opts *ListGlobalPoolsRequest) (*entity.ListGlobalPools, error) {
	url := listGlobalPoolsURL(s.VLBClient, opts)
	resp := new(ListGlobalPoolsResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.GlobalLoadBalancerErrorType)
	req := client.NewRequest().
		WithOkCodes(200).
		WithJSONResponse(resp).
		WithJSONError(errResp)

	if _, sdkErr := s.VLBClient.Get(url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.EcGlobalLoadBalancerNotFound).
			WithKVparameters("loadBalancerId", opts.GetLoadBalancerID()).
			AppendCategories(sdkerror.ErrCatProductVlb)
	}

	return resp.ToEntityListGlobalPools(), nil
}

func (s *GLBServiceV1) CreateGlobalPool(opts *CreateGlobalPoolRequest) (*entity.GlobalPool, error) {
	url := createGlobalPoolURL(s.VLBClient, opts)
	resp := new(CreateGlobalPoolResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.GlobalLoadBalancerErrorType)
	req := client.NewRequest().
		WithOkCodes(202).
		WithJSONBody(opts).
		WithJSONResponse(resp).
		WithJSONError(errResp)

	if _, sdkErr := s.VLBClient.Post(url, req); sdkErr != nil {
		fmt.Println("sdkErr: ", sdkErr)
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.EcGlobalLoadBalancerNotFound).
			WithParameters(common.StructToMap(opts)).
			AppendCategories(sdkerror.ErrCatProductVlb)
	}

	return resp.ToEntityPool(), nil
}

func (s *GLBServiceV1) UpdateGlobalPool(opts *UpdateGlobalPoolRequest) (*entity.GlobalPool, error) {
	url := updateGlobalPoolURL(s.VLBClient, opts)
	resp := new(UpdateGlobalPoolResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.GlobalLoadBalancerErrorType)
	req := client.NewRequest().
		WithOkCodes(202).
		WithJSONBody(opts).
		WithJSONResponse(resp).
		WithJSONError(errResp)

	if _, sdkErr := s.VLBClient.Put(url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.EcGlobalLoadBalancerNotFound).
			WithParameters(common.StructToMap(opts)).
			AppendCategories(sdkerror.ErrCatProductVlb)
	}

	return resp.ToEntityPool(), nil
}

func (s *GLBServiceV1) DeleteGlobalPool(opts *DeleteGlobalPoolRequest) error {
	url := deleteGlobalPoolURL(s.VLBClient, opts)
	errResp := sdkerror.NewErrorResponse(sdkerror.GlobalLoadBalancerErrorType)
	req := client.NewRequest().
		WithOkCodes(202).
		WithJSONError(errResp)

	if _, sdkErr := s.VLBClient.Delete(url, req); sdkErr != nil {
		return sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.EcGlobalLoadBalancerNotFound).
			AppendCategories(sdkerror.ErrCatProductVlb)
	}

	return nil
}

func (s *GLBServiceV1) ListGlobalPoolMembers(opts *ListGlobalPoolMembersRequest) (*entity.ListGlobalPoolMembers, error) {
	url := listGlobalPoolMembersURL(s.VLBClient, opts)
	resp := new(ListGlobalPoolMembersResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.GlobalLoadBalancerErrorType)
	req := client.NewRequest().
		WithOkCodes(200).
		WithJSONResponse(resp).
		WithJSONError(errResp)

	if _, sdkErr := s.VLBClient.Get(url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.EcGlobalLoadBalancerNotFound).
			AppendCategories(sdkerror.ErrCatProductVlb)
	}

	return resp.ToEntityListGlobalPoolMembers(), nil
}

func (s *GLBServiceV1) GetGlobalPoolMember(opts *GetGlobalPoolMemberRequest) (*entity.GlobalPoolMember, error) {
	url := getGlobalPoolMemberURL(s.VLBClient, opts)
	resp := new(GetGlobalPoolMemberResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.GlobalLoadBalancerErrorType)
	req := client.NewRequest().
		WithOkCodes(200).
		WithJSONResponse(resp).
		WithJSONError(errResp)

	if _, sdkErr := s.VLBClient.Get(url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.EcGlobalLoadBalancerNotFound).
			WithKVparameters("loadBalancerId", opts.GetLoadBalancerID()).
			WithKVparameters("poolId", opts.GetPoolID()).
			WithKVparameters("poolMemberId", opts.GetPoolMemberID()).
			AppendCategories(sdkerror.ErrCatProductVlb)
	}

	return resp.ToEntityGlobalPoolMember(), nil
}

func (s *GLBServiceV1) DeleteGlobalPoolMember(opts *DeleteGlobalPoolMemberRequest) error {
	url := deleteGlobalPoolMemberURL(s.VLBClient, opts)
	errResp := sdkerror.NewErrorResponse(sdkerror.GlobalLoadBalancerErrorType)
	req := client.NewRequest().
		WithOkCodes(202).
		WithJSONError(errResp)

	if _, sdkErr := s.VLBClient.Delete(url, req); sdkErr != nil {
		return sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.EcGlobalLoadBalancerNotFound).
			WithKVparameters("loadBalancerId", opts.GetLoadBalancerID()).
			WithKVparameters("poolId", opts.GetPoolID()).
			WithKVparameters("poolMemberId", opts.GetPoolMemberID()).
			AppendCategories(sdkerror.ErrCatProductVlb)
	}

	return nil
}

func (s *GLBServiceV1) UpdateGlobalPoolMember(opts *UpdateGlobalPoolMemberRequest) (*entity.GlobalPoolMember, error) {
	url := updateGlobalPoolMemberURL(s.VLBClient, opts)
	resp := new(UpdateGlobalPoolMemberResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.GlobalLoadBalancerErrorType)
	req := client.NewRequest().
		WithOkCodes(202).
		WithJSONBody(opts).
		WithJSONResponse(resp).
		WithJSONError(errResp)

	if _, sdkErr := s.VLBClient.Put(url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.EcGlobalLoadBalancerNotFound).
			WithParameters(common.StructToMap(opts)).
			AppendCategories(sdkerror.ErrCatProductVlb)
	}

	return resp.ToEntityGlobalPoolMember(), nil
}

func (s *GLBServiceV1) PatchGlobalPoolMembers(opts *PatchGlobalPoolMembersRequest) error {
	url := patchGlobalPoolMembersURL(s.VLBClient, opts)
	errResp := sdkerror.NewErrorResponse(sdkerror.GlobalLoadBalancerErrorType)
	req := client.NewRequest().
		WithOkCodes(202).
		WithJSONBody(opts).
		WithJSONError(errResp)

	if _, sdkErr := s.VLBClient.Patch(url, req); sdkErr != nil {
		return sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.EcGlobalLoadBalancerNotFound).
			WithParameters(common.StructToMap(opts)).
			AppendCategories(sdkerror.ErrCatProductVlb)
	}

	return nil
}

func (s *GLBServiceV1) ListGlobalListeners(opts *ListGlobalListenersRequest) (*entity.ListGlobalListeners, error) {
	url := listGlobalListenersURL(s.VLBClient, opts)
	resp := new(ListGlobalListenersResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.GlobalLoadBalancerErrorType)
	req := client.NewRequest().
		WithOkCodes(200).
		WithJSONResponse(resp).
		WithJSONError(errResp)

	if _, sdkErr := s.VLBClient.Get(url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.EcGlobalLoadBalancerNotFound).
			AppendCategories(sdkerror.ErrCatProductVlb)
	}

	return resp.ToEntityListGlobalListeners(), nil
}

func (s *GLBServiceV1) CreateGlobalListener(opts *CreateGlobalListenerRequest) (*entity.GlobalListener, error) {
	url := createGlobalListenerURL(s.VLBClient, opts)
	resp := new(CreateGlobalListenerResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.GlobalLoadBalancerErrorType)
	req := client.NewRequest().
		WithOkCodes(202).
		WithJSONBody(opts).
		WithJSONResponse(resp).
		WithJSONError(errResp)

	if _, sdkErr := s.VLBClient.Post(url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.EcGlobalLoadBalancerNotFound).
			WithParameters(common.StructToMap(opts)).
			AppendCategories(sdkerror.ErrCatProductVlb)
	}

	return resp.ToEntityGlobalListener(), nil
}

func (s *GLBServiceV1) UpdateGlobalListener(opts *UpdateGlobalListenerRequest) (*entity.GlobalListener, error) {
	url := updateGlobalListenerURL(s.VLBClient, opts)
	resp := new(UpdateGlobalListenerResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.GlobalLoadBalancerErrorType)
	req := client.NewRequest().
		WithOkCodes(202).
		WithJSONBody(opts).
		WithJSONResponse(resp).
		WithJSONError(errResp)

	if _, sdkErr := s.VLBClient.Put(url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.EcGlobalLoadBalancerNotFound).
			WithParameters(common.StructToMap(opts)).
			AppendCategories(sdkerror.ErrCatProductVlb)

	}

	return resp.ToEntityGlobalListener(), nil
}

func (s *GLBServiceV1) GetGlobalListener(opts *GetGlobalListenerRequest) (*entity.GlobalListener, error) {
	url := getGlobalListenerURL(s.VLBClient, opts)
	resp := new(GetGlobalListenerResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.GlobalLoadBalancerErrorType)
	req := client.NewRequest().
		WithOkCodes(200).
		WithJSONResponse(resp).
		WithJSONError(errResp)

	if _, sdkErr := s.VLBClient.Get(url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.EcGlobalLoadBalancerNotFound).
			WithKVparameters("loadBalancerId", opts.GetLoadBalancerID()).
			WithKVparameters("listenerId", opts.GetListenerID()).
			AppendCategories(sdkerror.ErrCatProductVlb)
	}

	return resp.ToEntityGlobalListener(), nil
}

func (s *GLBServiceV1) DeleteGlobalListener(opts *DeleteGlobalListenerRequest) error {
	url := deleteGlobalListenerURL(s.VLBClient, opts)
	errResp := sdkerror.NewErrorResponse(sdkerror.GlobalLoadBalancerErrorType)
	req := client.NewRequest().
		WithOkCodes(202).
		WithJSONError(errResp)

	if _, sdkErr := s.VLBClient.Delete(url, req); sdkErr != nil {
		return sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.EcGlobalLoadBalancerNotFound).
			AppendCategories(sdkerror.ErrCatProductVlb)
	}

	return nil
}

func (s *GLBServiceV1) ListGlobalLoadBalancers(opts *ListGlobalLoadBalancersRequest) (*entity.ListGlobalLoadBalancers, error) {
	url := listGlobalLoadBalancersURL(s.VLBClient, opts)
	resp := new(ListGlobalLoadBalancersResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.GlobalLoadBalancerErrorType)
	req := client.NewRequest().
		WithOkCodes(200).
		WithJSONResponse(resp).
		WithJSONError(errResp)

	if _, sdkErr := s.VLBClient.Get(url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp).
			AppendCategories(sdkerror.ErrCatProductVlb)
	}

	return resp.ToEntityListGlobalLoadBalancers(), nil
}

func (s *GLBServiceV1) CreateGlobalLoadBalancer(
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

	if _, sdkErr := s.VLBClient.Post(url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp).
			WithParameters(common.StructToMap(opts)).
			AppendCategories(sdkerror.ErrCatProductVlb)
	}

	return resp.ToEntityGlobalLoadBalancer(), nil
}

func (s *GLBServiceV1) DeleteGlobalLoadBalancer(opts *DeleteGlobalLoadBalancerRequest) error {
	url := deleteGlobalLoadBalancerURL(s.VLBClient, opts)
	errResp := sdkerror.NewErrorResponse(sdkerror.GlobalLoadBalancerErrorType)
	req := client.NewRequest().
		WithOkCodes(202).
		WithJSONError(errResp)

	if _, sdkErr := s.VLBClient.Delete(url, req); sdkErr != nil {
		return sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.EcGlobalLoadBalancerNotFound).
			AppendCategories(sdkerror.ErrCatProductVlb)
	}

	return nil
}

func (s *GLBServiceV1) GetGlobalLoadBalancerByID(
	opts *GetGlobalLoadBalancerByIDRequest,
) (*entity.GlobalLoadBalancer, error) {
	url := getGlobalLoadBalancerByIDURL(s.VLBClient, opts)
	resp := new(GetGlobalLoadBalancerByIDResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.GlobalLoadBalancerErrorType)
	req := client.NewRequest().
		WithOkCodes(200).
		WithJSONResponse(resp).
		WithJSONError(errResp)

	if _, sdkErr := s.VLBClient.Get(url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.EcGlobalLoadBalancerNotFound).
			AppendCategories(sdkerror.ErrCatProductVlb)
	}

	return resp.ToEntityGlobalLoadBalancer(), nil
}

func (s *GLBServiceV1) ListGlobalPackages(opts *ListGlobalPackagesRequest) (*entity.ListGlobalPackages, error) {
	url := listGlobalPackagesURL(s.VLBClient, opts)
	resp := new(ListGlobalPackagesResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.GlobalLoadBalancerErrorType)
	req := client.NewRequest().
		WithOkCodes(200).
		WithJSONResponse(resp).
		WithJSONError(errResp)

	if _, sdkErr := s.VLBClient.Get(url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp).
			AppendCategories(sdkerror.ErrCatProductVlb)
	}

	return resp.ToEntityListGlobalPackages(), nil
}

func (s *GLBServiceV1) ListGlobalRegions(opts *ListGlobalRegionsRequest) (*entity.ListGlobalRegions, error) {
	url := listGlobalRegionsURL(s.VLBClient, opts)
	resp := new(ListGlobalRegionsResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.GlobalLoadBalancerErrorType)
	req := client.NewRequest().
		WithOkCodes(200).
		WithJSONResponse(resp).
		WithJSONError(errResp)

	if _, sdkErr := s.VLBClient.Get(url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp).
			AppendCategories(sdkerror.ErrCatProductVlb)
	}

	return resp.ToEntityListGlobalRegions(), nil
}

func (s *GLBServiceV1) GetGlobalLoadBalancerUsageHistories(
	opts *GetGlobalLoadBalancerUsageHistoriesRequest,
) (*entity.ListGlobalLoadBalancerUsageHistories, error) {
	url := getGlobalLoadBalancerUsageHistoriesURL(s.VLBClient, opts)
	resp := new(GetGlobalLoadBalancerUsageHistoriesResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.GlobalLoadBalancerErrorType)
	req := client.NewRequest().
		WithOkCodes(200).
		WithJSONResponse(resp).
		WithJSONError(errResp)

	if _, sdkErr := s.VLBClient.Get(url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.EcGlobalLoadBalancerNotFound).
			AppendCategories(sdkerror.ErrCatProductVlb)
	}

	return resp.ToEntityGlobalLoadBalancerUsageHistories(), nil
}
