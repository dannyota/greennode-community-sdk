package v2

import (
	"github.com/dannyota/greennode-community-sdk/v2/greennode/client"
	"github.com/dannyota/greennode-community-sdk/v2/greennode/entity"
	sdkerror "github.com/dannyota/greennode-community-sdk/v2/greennode/sdkerror"
)

func (s *LoadBalancerServiceV2) CreateLoadBalancer(opts ICreateLoadBalancerRequest) (*entity.LoadBalancer, sdkerror.Error) {
	url := createLoadBalancerURL(s.VLBClient)
	resp := new(CreateLoadBalancerResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithHeader("User-Agent", opts.ParseUserAgent()).
		WithOkCodes(202).
		WithJSONBody(opts.ToRequestBody()).
		WithJSONResponse(resp).
		WithJSONError(errResp)

	if _, sdkErr := s.VLBClient.Post(url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.WithErrorLoadBalancerExceedQuota(errResp)).
			WithParameters(opts.ToMap()).
			AppendCategories(sdkerror.ErrCatProductVlb)
	}

	return resp.ToEntityLoadBalancer(), nil
}

func (s *LoadBalancerServiceV2) ResizeLoadBalancer(opts IResizeLoadBalancerRequest) (*entity.LoadBalancer, sdkerror.Error) {
	url := resizeLoadBalancerURL(s.VLBClient, opts)
	resp := new(ResizeLoadBalancerResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithHeader("User-Agent", opts.ParseUserAgent()).
		WithOkCodes(202).
		WithJSONBody(opts.ToRequestBody()).
		WithJSONResponse(resp).
		WithJSONError(errResp)

	if _, sdkErr := s.VLBClient.Put(url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.WithErrorLoadBalancerNotFound(errResp),
			sdkerror.WithErrorLoadBalancerNotReady(errResp))
	}

	return resp.ToEntityLoadBalancer(), nil
}

func (s *LoadBalancerServiceV2) ListLoadBalancerPackages(opts IListLoadBalancerPackagesRequest) (*entity.ListLoadBalancerPackages, sdkerror.Error) {
	url := listLoadBalancerPackagesURL(s.VLBClient, opts)
	resp := new(ListLoadBalancerPackagesResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithHeader("User-Agent", opts.ParseUserAgent()).
		WithOkCodes(200).
		WithJSONResponse(resp).
		WithJSONError(errResp)

	if _, sdkErr := s.VLBClient.Get(url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp)
	}

	return resp.ToEntityListLoadBalancerPackages(), nil
}

func (s *LoadBalancerServiceV2) GetLoadBalancerByID(opts IGetLoadBalancerByIDRequest) (*entity.LoadBalancer, sdkerror.Error) {
	url := getLoadBalancerByIDURL(s.VLBClient, opts)
	resp := new(GetLoadBalancerByIDResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithHeader("User-Agent", opts.ParseUserAgent()).
		WithOkCodes(200).
		WithJSONResponse(resp).
		WithJSONError(errResp)

	if _, sdkErr := s.VLBClient.Get(url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.WithErrorLoadBalancerNotFound(errResp)).
			WithKVparameters("loadBalancerId", opts.GetLoadBalancerID()).
			AppendCategories(sdkerror.ErrCatProductVlb)
	}

	return resp.ToEntityLoadBalancer(), nil
}

func (s *LoadBalancerServiceV2) ListLoadBalancers(opts IListLoadBalancersRequest) (*entity.ListLoadBalancers, sdkerror.Error) {
	url := listLoadBalancersURL(s.VLBClient, opts)
	resp := new(ListLoadBalancersResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithHeader("User-Agent", opts.ParseUserAgent()).
		WithOkCodes(200).
		WithJSONResponse(resp).
		WithJSONError(errResp)

	if _, sdkErr := s.VLBClient.Get(url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp).
			AppendCategories(sdkerror.ErrCatProductVlb)
	}

	return resp.ToEntityListLoadBalancers(), nil
}

func (s *LoadBalancerServiceV2) GetPoolHealthMonitorByID(opts IGetPoolHealthMonitorByIDRequest) (*entity.HealthMonitor, sdkerror.Error) {
	url := getPoolHealthMonitorByIDURL(s.VLBClient, opts)
	resp := new(GetPoolHealthMonitorByIDResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithHeader("User-Agent", opts.ParseUserAgent()).
		WithOkCodes(200).
		WithJSONResponse(resp).
		WithJSONError(errResp)

	if _, sdkErr := s.VLBClient.Get(url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.WithErrorLoadBalancerNotFound(errResp),
			sdkerror.WithErrorPoolNotFound(errResp))
	}

	return resp.ToEntityHealthMonitor(), nil
}

func (s *LoadBalancerServiceV2) CreatePool(opts ICreatePoolRequest) (*entity.Pool, sdkerror.Error) {
	url := createPoolURL(s.VLBClient, opts)
	resp := new(CreatePoolResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithHeader("User-Agent", opts.ParseUserAgent()).
		WithOkCodes(202).
		WithJSONBody(opts.ToRequestBody()).
		WithJSONResponse(resp).
		WithJSONError(errResp)

	if _, sdkErr := s.VLBClient.Post(url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.WithErrorLoadBalancerNotFound(errResp),
			sdkerror.WithErrorLoadBalancerNotFound2(errResp),
			sdkerror.WithErrorLoadBalancerNotReady(errResp),
			sdkerror.WithErrorLoadBalancerDuplicatePoolName(errResp)).
			WithParameters(opts.ToMap()).
			AppendCategories(sdkerror.ErrCatProductVlb)
	}

	return resp.ToEntityPool(), nil
}

func (s *LoadBalancerServiceV2) UpdatePool(opts IUpdatePoolRequest) sdkerror.Error {
	url := updatePoolURL(s.VLBClient, opts)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithHeader("User-Agent", opts.ParseUserAgent()).
		WithOkCodes(202).
		WithJSONBody(opts.ToRequestBody()).
		WithJSONError(errResp)

	if _, sdkErr := s.VLBClient.Put(url, req); sdkErr != nil {
		return sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.WithErrorLoadBalancerNotFound2(errResp),
			sdkerror.WithErrorLoadBalancerNotReady(errResp),
			sdkerror.WithErrorListenerNotFound(errResp)).
			AppendCategories(sdkerror.ErrCatProductVlb)
	}
	return nil
}

func (s *LoadBalancerServiceV2) CreateListener(opts ICreateListenerRequest) (*entity.Listener, sdkerror.Error) {
	url := createListenerURL(s.VLBClient, opts)
	resp := new(CreateListenerResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithHeader("User-Agent", opts.ParseUserAgent()).
		WithOkCodes(202).
		WithJSONBody(opts.ToRequestBody()).
		WithJSONResponse(resp).
		WithJSONError(errResp)

	if _, sdkErr := s.VLBClient.Post(url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.WithErrorLoadBalancerNotFound2(errResp),
			sdkerror.WithErrorLoadBalancerNotReady(errResp),
			sdkerror.WithErrorListenerDuplicateName(errResp),
			sdkerror.WithErrorPoolNotFound(errResp),
			sdkerror.WithErrorListenerDuplicateProtocolOrPort(errResp)).
			WithParameters(opts.ToMap()).
			AppendCategories(sdkerror.ErrCatProductVlb)
	}

	return resp.ToEntityListener(), nil
}

func (s *LoadBalancerServiceV2) UpdateListener(opts IUpdateListenerRequest) sdkerror.Error {
	url := updateListenerURL(s.VLBClient, opts)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithHeader("User-Agent", opts.ParseUserAgent()).
		WithOkCodes(202).
		WithJSONBody(opts.ToRequestBody()).
		WithJSONError(errResp)

	if _, sdkErr := s.VLBClient.Put(url, req); sdkErr != nil {
		return sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.WithErrorLoadBalancerNotFound2(errResp),
			sdkerror.WithErrorLoadBalancerNotReady(errResp),
			sdkerror.WithErrorListenerNotFound(errResp)).
			AppendCategories(sdkerror.ErrCatProductVlb)
	}

	return nil
}

func (s *LoadBalancerServiceV2) ListListenersByLoadBalancerID(opts IListListenersByLoadBalancerIDRequest) (*entity.ListListeners, sdkerror.Error) {
	url := listListenersByLoadBalancerIDURL(s.VLBClient, opts)
	resp := new(ListListenersByLoadBalancerIDResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithHeader("User-Agent", opts.ParseUserAgent()).
		WithOkCodes(200).
		WithJSONResponse(resp).
		WithJSONError(errResp)

	if _, sdkErr := s.VLBClient.Get(url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.WithErrorLoadBalancerNotFound(errResp)).
			WithKVparameters("loadBalancerId", opts.GetLoadBalancerID()).
			AppendCategories(sdkerror.ErrCatProductVlb)
	}

	return resp.ToEntityListListeners(), nil
}

func (s *LoadBalancerServiceV2) ListPoolsByLoadBalancerID(opts IListPoolsByLoadBalancerIDRequest) (*entity.ListPools, sdkerror.Error) {
	url := listPoolsByLoadBalancerIDURL(s.VLBClient, opts)
	resp := new(ListPoolsByLoadBalancerIDResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithHeader("User-Agent", opts.ParseUserAgent()).
		WithOkCodes(200).
		WithJSONResponse(resp).
		WithJSONError(errResp)

	if _, sdkErr := s.VLBClient.Get(url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.WithErrorLoadBalancerNotFound(errResp)).
			WithKVparameters("loadBalancerId", opts.GetLoadBalancerID()).
			AppendCategories(sdkerror.ErrCatProductVlb)
	}

	return resp.ToEntityListPools(), nil
}

func (s *LoadBalancerServiceV2) UpdatePoolMembers(opts IUpdatePoolMembersRequest) sdkerror.Error {
	url := updatePoolMembersURL(s.VLBClient, opts)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithHeader("User-Agent", opts.ParseUserAgent()).
		WithOkCodes(202).
		WithJSONBody(opts.ToRequestBody()).
		WithJSONError(errResp)

	if _, sdkErr := s.VLBClient.Put(url, req); sdkErr != nil {
		return sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.WithErrorLoadBalancerNotFound2(errResp),
			sdkerror.WithErrorLoadBalancerNotReady(errResp),
			sdkerror.WithErrorPoolNotFound(errResp),
			sdkerror.WithErrorMemberMustIdentical(errResp)).
			AppendCategories(sdkerror.ErrCatProductVlb)
	}

	return nil
}

func (s *LoadBalancerServiceV2) ListPoolMembers(opts IListPoolMembersRequest) (*entity.ListMembers, sdkerror.Error) {
	url := listPoolMembersURL(s.VLBClient, opts)
	resp := new(ListPoolMembersResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithHeader("User-Agent", opts.ParseUserAgent()).
		WithOkCodes(200).
		WithJSONResponse(resp).
		WithJSONError(errResp)

	if _, sdkErr := s.VLBClient.Get(url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.WithErrorLoadBalancerNotFound(errResp),
			sdkerror.WithErrorPoolNotFound(errResp)).
			WithKVparameters("loadBalancerId", opts.GetLoadBalancerID(), "poolId", opts.GetPoolID()).
			AppendCategories(sdkerror.ErrCatProductVlb)
	}

	return resp.ToEntityListMembers(), nil
}

func (s *LoadBalancerServiceV2) DeletePoolByID(opts IDeletePoolByIDRequest) sdkerror.Error {
	url := deletePoolByIDURL(s.VLBClient, opts)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithHeader("User-Agent", opts.ParseUserAgent()).
		WithOkCodes(202).
		WithJSONError(errResp)

	if _, sdkErr := s.VLBClient.Delete(url, req); sdkErr != nil {
		return sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.WithErrorPoolInUse(errResp),
			sdkerror.WithErrorLoadBalancerNotReady(errResp),
			sdkerror.WithErrorPoolNotFound(errResp)).
			AppendCategories(sdkerror.ErrCatProductVlb)
	}

	return nil
}

func (s *LoadBalancerServiceV2) DeleteListenerByID(opts IDeleteListenerByIDRequest) sdkerror.Error {
	url := deleteListenerByIDURL(s.VLBClient, opts)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithHeader("User-Agent", opts.ParseUserAgent()).
		WithOkCodes(202).
		WithJSONError(errResp)

	if _, sdkErr := s.VLBClient.Delete(url, req); sdkErr != nil {
		return sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.WithErrorLoadBalancerNotFound2(errResp),
			sdkerror.WithErrorListenerNotFound(errResp),
			sdkerror.WithErrorLoadBalancerNotReady(errResp)).
			AppendCategories(sdkerror.ErrCatProductVlb)
	}

	return nil
}

func (s *LoadBalancerServiceV2) DeleteLoadBalancerByID(opts IDeleteLoadBalancerByIDRequest) sdkerror.Error {
	url := deleteLoadBalancerByIDURL(s.VLBClient, opts)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithHeader("User-Agent", opts.ParseUserAgent()).
		WithOkCodes(202).
		WithJSONError(errResp)

	if _, sdkErr := s.VLBClient.Delete(url, req); sdkErr != nil {
		return sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.WithErrorLoadBalancerNotFound(errResp),
			sdkerror.WithErrorLoadBalancerNotReady(errResp),
			sdkerror.WithErrorLoadBalancerIsCreating(errResp),
			sdkerror.WithErrorLoadBalancerIsDeleting(errResp)).
			WithKVparameters(
				"loadBalancerId", opts.GetLoadBalancerID(),
				"projectId", s.getProjectID()).
			AppendCategories(sdkerror.ErrCatProductVlb)
	}

	return nil
}

func (s *LoadBalancerServiceV2) GetPoolByID(opts IGetPoolByIDRequest) (*entity.Pool, sdkerror.Error) {
	url := getPoolByIDURL(s.VLBClient, opts)
	resp := new(GetPoolByIDResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithHeader("User-Agent", opts.ParseUserAgent()).
		WithOkCodes(200).
		WithJSONResponse(resp).
		WithJSONError(errResp)

	if _, sdkErr := s.VLBClient.Get(url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.WithErrorLoadBalancerNotFound(errResp),
			sdkerror.WithErrorPoolNotFound(errResp)).
			WithKVparameters(
				"loadBalancerId", opts.GetLoadBalancerID(),
				"poolId", opts.GetPoolID()).
			AppendCategories(sdkerror.ErrCatProductVlb)
	}

	return resp.ToEntityPool(), nil
}

func (s *LoadBalancerServiceV2) GetListenerByID(opts IGetListenerByIDRequest) (*entity.Listener, sdkerror.Error) {
	url := getListenerByIDURL(s.VLBClient, opts)
	resp := new(GetListenerByIDResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithHeader("User-Agent", opts.ParseUserAgent()).
		WithOkCodes(200).
		WithJSONResponse(resp).
		WithJSONError(errResp)

	if _, sdkErr := s.VLBClient.Get(url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.WithErrorLoadBalancerNotFound(errResp),
			sdkerror.WithErrorListenerNotFound(errResp)).
			WithKVparameters(
				"loadBalancerId", opts.GetLoadBalancerID(),
				"listenerId", opts.GetListenerID()).
			AppendCategories(sdkerror.ErrCatProductVlb)
	}

	return resp.ToEntityListener(), nil
}

func (s *LoadBalancerServiceV2) ResizeLoadBalancerByID(opts IResizeLoadBalancerByIDRequest) sdkerror.Error {
	url := resizeLoadBalancerByIDURL(s.VLBClient, opts)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithHeader("User-Agent", opts.ParseUserAgent()).
		WithOkCodes(202).
		WithJSONBody(opts.ToRequestBody()).
		WithJSONError(errResp)

	if _, sdkErr := s.VLBClient.Put(url, req); sdkErr != nil {
		return sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.WithErrorLoadBalancerPackageNotFound(errResp),
			sdkerror.WithErrorLoadBalancerNotFound(errResp),
			sdkerror.WithErrorLoadBalancerNotFound2(errResp),
			sdkerror.WithErrorLoadBalancerNotReady(errResp),
			sdkerror.WithErrorLoadBalancerResizeSamePackage(errResp)).
			WithParameters(opts.ToMap()).
			AppendCategories(sdkerror.ErrCatProductVlb)
	}

	return nil
}

func (s *LoadBalancerServiceV2) ScaleLoadBalancer(opts IScaleLoadBalancerRequest) (*entity.LoadBalancer, sdkerror.Error) {
	url := scaleLoadBalancerURL(s.VLBClient, opts)
	resp := new(ScaleLoadBalancerResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithHeader("User-Agent", opts.ParseUserAgent()).
		WithOkCodes(202).
		WithJSONBody(opts.ToRequestBody()).
		WithJSONResponse(resp).
		WithJSONError(errResp)

	if _, sdkErr := s.VLBClient.Put(url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.WithErrorLoadBalancerNotFound(errResp),
			sdkerror.WithErrorLoadBalancerNotFound2(errResp),
			sdkerror.WithErrorLoadBalancerNotReady(errResp)).
			WithParameters(opts.ToMap()).
			AppendCategories(sdkerror.ErrCatProductVlb)
	}

	return resp.ToEntityLoadBalancer(), nil
}

// policy

func (s *LoadBalancerServiceV2) ListPolicies(opts IListPoliciesRequest) (*entity.ListPolicies, sdkerror.Error) {
	url := listPoliciesURL(s.VLBClient, opts)
	resp := new(ListPoliciesResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithHeader("User-Agent", opts.ParseUserAgent()).
		WithOkCodes(200).
		WithJSONResponse(resp).
		WithJSONError(errResp)

	if _, sdkErr := s.VLBClient.Get(url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp)
	}

	return resp.ToEntityListPolicies(), nil
}

func (s *LoadBalancerServiceV2) CreatePolicy(opts ICreatePolicyRequest) (*entity.Policy, sdkerror.Error) {
	url := createPolicyURL(s.VLBClient, opts)
	resp := new(CreatePolicyResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithHeader("User-Agent", opts.ParseUserAgent()).
		WithOkCodes(202).
		WithJSONBody(opts.ToRequestBody()).
		WithJSONResponse(resp).
		WithJSONError(errResp)

	if _, sdkErr := s.VLBClient.Post(url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.WithErrorLoadBalancerNotFound(errResp),
			sdkerror.WithErrorListenerNotFound(errResp),
		).WithParameters(opts.ToMap())
	}

	return resp.ToEntityPolicy(), nil
}

func (s *LoadBalancerServiceV2) GetPolicyByID(opts IGetPolicyByIDRequest) (*entity.Policy, sdkerror.Error) {
	url := getPolicyByIDURL(s.VLBClient, opts)
	resp := new(GetPolicyResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithHeader("User-Agent", opts.ParseUserAgent()).
		WithOkCodes(200).
		WithJSONResponse(resp).
		WithJSONError(errResp)

	if _, sdkErr := s.VLBClient.Get(url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.WithErrorLoadBalancerNotFound(errResp),
			sdkerror.WithErrorListenerNotFound(errResp),
		).WithKVparameters("policyId", opts.GetPolicyID())
	}

	return resp.ToEntityPolicy(), nil
}

func (s *LoadBalancerServiceV2) UpdatePolicy(opts IUpdatePolicyRequest) sdkerror.Error {
	url := updatePolicyURL(s.VLBClient, opts)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithHeader("User-Agent", opts.ParseUserAgent()).
		WithOkCodes(202).
		WithJSONBody(opts.ToRequestBody()).
		WithJSONError(errResp)

	if _, sdkErr := s.VLBClient.Put(url, req); sdkErr != nil {
		return sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.WithErrorLoadBalancerNotFound(errResp),
			sdkerror.WithErrorListenerNotFound(errResp),
		)
	}

	return nil
}

func (s *LoadBalancerServiceV2) DeletePolicyByID(opts IDeletePolicyByIDRequest) sdkerror.Error {
	url := deletePolicyByIDURL(s.VLBClient, opts)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithHeader("User-Agent", opts.ParseUserAgent()).
		WithOkCodes(202).
		WithJSONError(errResp)

	if _, sdkErr := s.VLBClient.Delete(url, req); sdkErr != nil {
		return sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.WithErrorLoadBalancerNotFound(errResp),
			sdkerror.WithErrorListenerNotFound(errResp),
		)
	}

	return nil
}

func (s *LoadBalancerServiceV2) ReorderPolicies(opts IReorderPoliciesRequest) sdkerror.Error {
	url := reorderPoliciesURL(s.VLBClient, opts)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithHeader("User-Agent", opts.ParseUserAgent()).
		WithOkCodes(202).
		WithJSONBody(opts.ToRequestBody()).
		WithJSONError(errResp)
	if _, sdkErr := s.VLBClient.Put(url, req); sdkErr != nil {
		return sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.WithErrorLoadBalancerNotFound(errResp),
			sdkerror.WithErrorListenerNotFound(errResp),
		)
	}
	return nil
}


func (s *LoadBalancerServiceV2) ListCertificates(opts IListCertificatesRequest) (*entity.ListCertificates, sdkerror.Error) {
	url := listCertificatesURL(s.VLBClient)
	resp := new(ListCertificatesResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithHeader("User-Agent", opts.ParseUserAgent()).
		WithOkCodes(200).
		WithJSONResponse(resp).
		WithJSONError(errResp)

	if _, sdkErr := s.VLBClient.Get(url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp)
	}

	return resp.ToEntityListCertificates(), nil
}

func (s *LoadBalancerServiceV2) GetCertificateByID(opts IGetCertificateByIDRequest) (*entity.Certificate, sdkerror.Error) {
	url := getCertificateByIDURL(s.VLBClient, opts)
	resp := new(GetCertificateByIDResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithHeader("User-Agent", opts.ParseUserAgent()).
		WithOkCodes(200).
		WithJSONResponse(resp).
		WithJSONError(errResp)

	if _, sdkErr := s.VLBClient.Get(url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp)
	}

	return resp.ToEntityCertificate(), nil
}

func (s *LoadBalancerServiceV2) CreateCertificate(opts ICreateCertificateRequest) (*entity.Certificate, sdkerror.Error) {
	url := createCertificateURL(s.VLBClient)
	resp := new(CreateCertificateResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithHeader("User-Agent", opts.ParseUserAgent()).
		WithOkCodes(201).
		WithJSONBody(opts.ToRequestBody()).
		WithJSONResponse(resp).
		WithJSONError(errResp)

	if _, sdkErr := s.VLBClient.Post(url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp)
	}

	return resp.ToEntityCertificate(), nil
}

func (s *LoadBalancerServiceV2) DeleteCertificateByID(opts IDeleteCertificateByIDRequest) sdkerror.Error {
	url := deleteCertificateByIDURL(s.VLBClient, opts)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithHeader("User-Agent", opts.ParseUserAgent()).
		WithOkCodes(204).
		WithJSONError(errResp)

	if _, sdkErr := s.VLBClient.Delete(url, req); sdkErr != nil {
		return sdkerror.SdkErrorHandler(sdkErr, errResp)
	}

	return nil
}
