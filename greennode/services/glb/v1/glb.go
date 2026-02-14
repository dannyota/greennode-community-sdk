package v1

import (
	"fmt"

	"github.com/dannyota/greennode-community-sdk/v2/greennode/client"
	"github.com/dannyota/greennode-community-sdk/v2/greennode/entity"
	sdkerror "github.com/dannyota/greennode-community-sdk/v2/greennode/sdkerror"
)

func (s *GLBServiceV1) ListGlobalPools(opts IListGlobalPoolsRequest) (*entity.ListGlobalPools, sdkerror.Error) {
	url := listGlobalPoolsURL(s.VLBClient, opts)
	resp := new(ListGlobalPoolsResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.GlobalLoadBalancerErrorType)
	req := client.NewRequest().
		WithHeader("User-Agent", opts.ParseUserAgent()).
		WithOkCodes(200).
		WithJSONResponse(resp).
		WithJSONError(errResp)

	if _, sdkErr := s.VLBClient.Get(url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.WithErrorGlobalLoadBalancerNotFound(errResp)).
			WithKVparameters("loadBalancerId", opts.GetLoadBalancerID()).
			AppendCategories(sdkerror.ErrCatProductVlb)
	}

	return resp.ToEntityListGlobalPools(), nil
}


func (s *GLBServiceV1) CreateGlobalPool(opts ICreateGlobalPoolRequest) (*entity.GlobalPool, sdkerror.Error) {
	url := createGlobalPoolURL(s.VLBClient, opts)
	resp := new(CreateGlobalPoolResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.GlobalLoadBalancerErrorType)
	req := client.NewRequest().
		WithHeader("User-Agent", opts.ParseUserAgent()).
		WithOkCodes(202).
		WithJSONBody(opts.ToRequestBody()).
		WithJSONResponse(resp).
		WithJSONError(errResp)

	if _, sdkErr := s.VLBClient.Post(url, req); sdkErr != nil {
		fmt.Println("sdkErr: ", sdkErr)
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.WithErrorGlobalLoadBalancerNotFound(errResp)).
			WithParameters(opts.ToMap()).
			AppendCategories(sdkerror.ErrCatProductVlb)
	}

	return resp.ToEntityPool(), nil
}


func (s *GLBServiceV1) UpdateGlobalPool(opts IUpdateGlobalPoolRequest) (*entity.GlobalPool, sdkerror.Error) {
	url := updateGlobalPoolURL(s.VLBClient, opts)
	resp := new(UpdateGlobalPoolResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.GlobalLoadBalancerErrorType)
	req := client.NewRequest().
		WithHeader("User-Agent", opts.ParseUserAgent()).
		WithOkCodes(202).
		WithJSONBody(opts.ToRequestBody()).
		WithJSONResponse(resp).
		WithJSONError(errResp)

	if _, sdkErr := s.VLBClient.Put(url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.WithErrorGlobalLoadBalancerNotFound(errResp)).
			WithParameters(opts.ToMap()).
			AppendCategories(sdkerror.ErrCatProductVlb)
	}

	return resp.ToEntityPool(), nil
}


func (s *GLBServiceV1) DeleteGlobalPool(opts IDeleteGlobalPoolRequest) sdkerror.Error {
	url := deleteGlobalPoolURL(s.VLBClient, opts)
	errResp := sdkerror.NewErrorResponse(sdkerror.GlobalLoadBalancerErrorType)
	req := client.NewRequest().
		WithHeader("User-Agent", opts.ParseUserAgent()).
		WithOkCodes(202).
		WithJSONError(errResp)

	if _, sdkErr := s.VLBClient.Delete(url, req); sdkErr != nil {
		return sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.WithErrorGlobalLoadBalancerNotFound(errResp)).
			AppendCategories(sdkerror.ErrCatProductVlb)
	}

	return nil
}


func (s *GLBServiceV1) ListGlobalPoolMembers(opts IListGlobalPoolMembersRequest) (*entity.ListGlobalPoolMembers, sdkerror.Error) {
	url := listGlobalPoolMembersURL(s.VLBClient, opts)
	resp := new(ListGlobalPoolMembersResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.GlobalLoadBalancerErrorType)
	req := client.NewRequest().
		WithHeader("User-Agent", opts.ParseUserAgent()).
		WithOkCodes(200).
		WithJSONResponse(resp).
		WithJSONError(errResp)

	if _, sdkErr := s.VLBClient.Get(url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.WithErrorGlobalLoadBalancerNotFound(errResp)).
			AppendCategories(sdkerror.ErrCatProductVlb)
	}

	return resp.ToEntityListGlobalPoolMembers(), nil
}


func (s *GLBServiceV1) GetGlobalPoolMember(opts IGetGlobalPoolMemberRequest) (*entity.GlobalPoolMember, sdkerror.Error) {
	url := getGlobalPoolMemberURL(s.VLBClient, opts)
	resp := new(GetGlobalPoolMemberResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.GlobalLoadBalancerErrorType)
	req := client.NewRequest().
		WithHeader("User-Agent", opts.ParseUserAgent()).
		WithOkCodes(200).
		WithJSONResponse(resp).
		WithJSONError(errResp)

	if _, sdkErr := s.VLBClient.Get(url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.WithErrorGlobalLoadBalancerNotFound(errResp)).
			WithKVparameters("loadBalancerId", opts.GetLoadBalancerID()).
			WithKVparameters("poolId", opts.GetPoolID()).
			WithKVparameters("poolMemberId", opts.GetPoolMemberID()).
			AppendCategories(sdkerror.ErrCatProductVlb)
	}

	return resp.ToEntityGlobalPoolMember(), nil
}


func (s *GLBServiceV1) DeleteGlobalPoolMember(opts IDeleteGlobalPoolMemberRequest) sdkerror.Error {
	url := deleteGlobalPoolMemberURL(s.VLBClient, opts)
	errResp := sdkerror.NewErrorResponse(sdkerror.GlobalLoadBalancerErrorType)
	req := client.NewRequest().
		WithHeader("User-Agent", opts.ParseUserAgent()).
		WithOkCodes(202).
		WithJSONError(errResp)

	if _, sdkErr := s.VLBClient.Delete(url, req); sdkErr != nil {
		return sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.WithErrorGlobalLoadBalancerNotFound(errResp)).
			WithKVparameters("loadBalancerId", opts.GetLoadBalancerID()).
			WithKVparameters("poolId", opts.GetPoolID()).
			WithKVparameters("poolMemberId", opts.GetPoolMemberID()).
			AppendCategories(sdkerror.ErrCatProductVlb)
	}

	return nil
}


func (s *GLBServiceV1) UpdateGlobalPoolMember(opts IUpdateGlobalPoolMemberRequest) (*entity.GlobalPoolMember, sdkerror.Error) {
	url := updateGlobalPoolMemberURL(s.VLBClient, opts)
	resp := new(UpdateGlobalPoolMemberResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.GlobalLoadBalancerErrorType)
	req := client.NewRequest().
		WithHeader("User-Agent", opts.ParseUserAgent()).
		WithOkCodes(202).
		WithJSONBody(opts.ToRequestBody()).
		WithJSONResponse(resp).
		WithJSONError(errResp)

	if _, sdkErr := s.VLBClient.Put(url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.WithErrorGlobalLoadBalancerNotFound(errResp)).
			WithParameters(opts.ToMap()).
			AppendCategories(sdkerror.ErrCatProductVlb)
	}

	return resp.ToEntityGlobalPoolMember(), nil
}


func (s *GLBServiceV1) PatchGlobalPoolMembers(opts IPatchGlobalPoolMembersRequest) sdkerror.Error {
	url := patchGlobalPoolMembersURL(s.VLBClient, opts)
	errResp := sdkerror.NewErrorResponse(sdkerror.GlobalLoadBalancerErrorType)
	req := client.NewRequest().
		WithHeader("User-Agent", opts.ParseUserAgent()).
		WithOkCodes(202).
		WithJSONBody(opts.ToRequestBody()).
		WithJSONError(errResp)

	if _, sdkErr := s.VLBClient.Patch(url, req); sdkErr != nil {
		return sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.WithErrorGlobalLoadBalancerNotFound(errResp)).
			WithParameters(opts.ToMap()).
			AppendCategories(sdkerror.ErrCatProductVlb)
	}

	return nil
}


func (s *GLBServiceV1) ListGlobalListeners(opts IListGlobalListenersRequest) (*entity.ListGlobalListeners, sdkerror.Error) {
	url := listGlobalListenersURL(s.VLBClient, opts)
	resp := new(ListGlobalListenersResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.GlobalLoadBalancerErrorType)
	req := client.NewRequest().
		WithHeader("User-Agent", opts.ParseUserAgent()).
		WithOkCodes(200).
		WithJSONResponse(resp).
		WithJSONError(errResp)

	if _, sdkErr := s.VLBClient.Get(url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.WithErrorGlobalLoadBalancerNotFound(errResp)).
			AppendCategories(sdkerror.ErrCatProductVlb)
	}

	return resp.ToEntityListGlobalListeners(), nil
}


func (s *GLBServiceV1) CreateGlobalListener(opts ICreateGlobalListenerRequest) (*entity.GlobalListener, sdkerror.Error) {
	url := createGlobalListenerURL(s.VLBClient, opts)
	resp := new(CreateGlobalListenerResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.GlobalLoadBalancerErrorType)
	req := client.NewRequest().
		WithHeader("User-Agent", opts.ParseUserAgent()).
		WithOkCodes(202).
		WithJSONBody(opts.ToRequestBody()).
		WithJSONResponse(resp).
		WithJSONError(errResp)

	if _, sdkErr := s.VLBClient.Post(url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.WithErrorGlobalLoadBalancerNotFound(errResp)).
			WithParameters(opts.ToMap()).
			AppendCategories(sdkerror.ErrCatProductVlb)
	}

	return resp.ToEntityGlobalListener(), nil
}


func (s *GLBServiceV1) UpdateGlobalListener(opts IUpdateGlobalListenerRequest) (*entity.GlobalListener, sdkerror.Error) {
	url := updateGlobalListenerURL(s.VLBClient, opts)
	resp := new(UpdateGlobalListenerResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.GlobalLoadBalancerErrorType)
	req := client.NewRequest().
		WithHeader("User-Agent", opts.ParseUserAgent()).
		WithOkCodes(202).
		WithJSONBody(opts.ToRequestBody()).
		WithJSONResponse(resp).
		WithJSONError(errResp)

	if _, sdkErr := s.VLBClient.Put(url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.WithErrorGlobalLoadBalancerNotFound(errResp)).
			WithParameters(opts.ToMap()).
			AppendCategories(sdkerror.ErrCatProductVlb)

	}

	return resp.ToEntityGlobalListener(), nil
}


func (s *GLBServiceV1) GetGlobalListener(opts IGetGlobalListenerRequest) (*entity.GlobalListener, sdkerror.Error) {
	url := getGlobalListenerURL(s.VLBClient, opts)
	resp := new(GetGlobalListenerResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.GlobalLoadBalancerErrorType)
	req := client.NewRequest().
		WithHeader("User-Agent", opts.ParseUserAgent()).
		WithOkCodes(200).
		WithJSONResponse(resp).
		WithJSONError(errResp)

	if _, sdkErr := s.VLBClient.Get(url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.WithErrorGlobalLoadBalancerNotFound(errResp)).
			WithKVparameters("loadBalancerId", opts.GetLoadBalancerID()).
			WithKVparameters("listenerId", opts.GetListenerID()).
			AppendCategories(sdkerror.ErrCatProductVlb)
	}

	return resp.ToEntityGlobalListener(), nil
}


func (s *GLBServiceV1) DeleteGlobalListener(opts IDeleteGlobalListenerRequest) sdkerror.Error {
	url := deleteGlobalListenerURL(s.VLBClient, opts)
	errResp := sdkerror.NewErrorResponse(sdkerror.GlobalLoadBalancerErrorType)
	req := client.NewRequest().
		WithHeader("User-Agent", opts.ParseUserAgent()).
		WithOkCodes(202).
		WithJSONError(errResp)

	if _, sdkErr := s.VLBClient.Delete(url, req); sdkErr != nil {
		return sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.WithErrorGlobalLoadBalancerNotFound(errResp)).
			AppendCategories(sdkerror.ErrCatProductVlb)
	}

	return nil
}


func (s *GLBServiceV1) ListGlobalLoadBalancers(opts IListGlobalLoadBalancersRequest) (*entity.ListGlobalLoadBalancers, sdkerror.Error) {
	url := listGlobalLoadBalancersURL(s.VLBClient, opts)
	resp := new(ListGlobalLoadBalancersResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.GlobalLoadBalancerErrorType)
	req := client.NewRequest().
		WithHeader("User-Agent", opts.ParseUserAgent()).
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
	opts ICreateGlobalLoadBalancerRequest,
) (*entity.GlobalLoadBalancer, sdkerror.Error) {
	url := createGlobalLoadBalancerURL(s.VLBClient, opts)
	resp := new(CreateGlobalLoadBalancerResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.GlobalLoadBalancerErrorType)
	req := client.NewRequest().
		WithHeader("User-Agent", opts.ParseUserAgent()).
		WithOkCodes(202).
		WithJSONBody(opts.ToRequestBody()).
		WithJSONResponse(resp).
		WithJSONError(errResp)

	if _, sdkErr := s.VLBClient.Post(url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp).
			WithParameters(opts.ToMap()).
			AppendCategories(sdkerror.ErrCatProductVlb)
	}

	return resp.ToEntityGlobalLoadBalancer(), nil
}


func (s *GLBServiceV1) DeleteGlobalLoadBalancer(opts IDeleteGlobalLoadBalancerRequest) sdkerror.Error {
	url := deleteGlobalLoadBalancerURL(s.VLBClient, opts)
	errResp := sdkerror.NewErrorResponse(sdkerror.GlobalLoadBalancerErrorType)
	req := client.NewRequest().
		WithHeader("User-Agent", opts.ParseUserAgent()).
		WithOkCodes(202).
		WithJSONError(errResp)

	if _, sdkErr := s.VLBClient.Delete(url, req); sdkErr != nil {
		return sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.WithErrorGlobalLoadBalancerNotFound(errResp)).
			AppendCategories(sdkerror.ErrCatProductVlb)
	}

	return nil
}


func (s *GLBServiceV1) GetGlobalLoadBalancerByID(
	opts IGetGlobalLoadBalancerByIDRequest,
) (*entity.GlobalLoadBalancer, sdkerror.Error) {
	url := getGlobalLoadBalancerByIDURL(s.VLBClient, opts)
	resp := new(GetGlobalLoadBalancerByIDResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.GlobalLoadBalancerErrorType)
	req := client.NewRequest().
		WithHeader("User-Agent", opts.ParseUserAgent()).
		WithOkCodes(200).
		WithJSONResponse(resp).
		WithJSONError(errResp)

	if _, sdkErr := s.VLBClient.Get(url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.WithErrorGlobalLoadBalancerNotFound(errResp)).
			AppendCategories(sdkerror.ErrCatProductVlb)
	}

	return resp.ToEntityGlobalLoadBalancer(), nil
}

func (s *GLBServiceV1) ListGlobalPackages(opts IListGlobalPackagesRequest) (*entity.ListGlobalPackages, sdkerror.Error) {
	url := listGlobalPackagesURL(s.VLBClient, opts)
	resp := new(ListGlobalPackagesResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.GlobalLoadBalancerErrorType)
	req := client.NewRequest().
		WithHeader("User-Agent", opts.ParseUserAgent()).
		WithOkCodes(200).
		WithJSONResponse(resp).
		WithJSONError(errResp)

	if _, sdkErr := s.VLBClient.Get(url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp).
			AppendCategories(sdkerror.ErrCatProductVlb)
	}

	return resp.ToEntityListGlobalPackages(), nil
}

func (s *GLBServiceV1) ListGlobalRegions(opts IListGlobalRegionsRequest) (*entity.ListGlobalRegions, sdkerror.Error) {
	url := listGlobalRegionsURL(s.VLBClient, opts)
	resp := new(ListGlobalRegionsResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.GlobalLoadBalancerErrorType)
	req := client.NewRequest().
		WithHeader("User-Agent", opts.ParseUserAgent()).
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
	opts IGetGlobalLoadBalancerUsageHistoriesRequest,
) (*entity.ListGlobalLoadBalancerUsageHistories, sdkerror.Error) {
	url := getGlobalLoadBalancerUsageHistoriesURL(s.VLBClient, opts)
	resp := new(GetGlobalLoadBalancerUsageHistoriesResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.GlobalLoadBalancerErrorType)
	req := client.NewRequest().
		WithHeader("User-Agent", opts.ParseUserAgent()).
		WithOkCodes(200).
		WithJSONResponse(resp).
		WithJSONError(errResp)

	if _, sdkErr := s.VLBClient.Get(url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.WithErrorGlobalLoadBalancerNotFound(errResp)).
			AppendCategories(sdkerror.ErrCatProductVlb)
	}

	return resp.ToEntityGlobalLoadBalancerUsageHistories(), nil
}
