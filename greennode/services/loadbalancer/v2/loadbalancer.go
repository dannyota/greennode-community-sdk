package v2

import (
	"github.com/dannyota/greennode-community-sdk/v2/greennode/client"
	"github.com/dannyota/greennode-community-sdk/v2/greennode/entity"
	sdkerror "github.com/dannyota/greennode-community-sdk/v2/greennode/sdkerror"
)

func (s *LoadBalancerServiceV2) CreateLoadBalancer(opts *CreateLoadBalancerRequest) (*entity.LoadBalancer, error) {
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

func (s *LoadBalancerServiceV2) ResizeLoadBalancer(opts *ResizeLoadBalancerRequest) (*entity.LoadBalancer, error) {
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

func (s *LoadBalancerServiceV2) ListLoadBalancerPackages(opts *ListLoadBalancerPackagesRequest) (*entity.ListLoadBalancerPackages, error) {
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

func (s *LoadBalancerServiceV2) GetLoadBalancerByID(opts *GetLoadBalancerByIDRequest) (*entity.LoadBalancer, error) {
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

func (s *LoadBalancerServiceV2) ListLoadBalancers(opts *ListLoadBalancersRequest) (*entity.ListLoadBalancers, error) {
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

func (s *LoadBalancerServiceV2) GetPoolHealthMonitorByID(opts *GetPoolHealthMonitorByIDRequest) (*entity.HealthMonitor, error) {
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

func (s *LoadBalancerServiceV2) CreatePool(opts *CreatePoolRequest) (*entity.Pool, error) {
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

func (s *LoadBalancerServiceV2) UpdatePool(opts *UpdatePoolRequest) error {
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

func (s *LoadBalancerServiceV2) CreateListener(opts *CreateListenerRequest) (*entity.Listener, error) {
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

func (s *LoadBalancerServiceV2) UpdateListener(opts *UpdateListenerRequest) error {
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

func (s *LoadBalancerServiceV2) ListListenersByLoadBalancerID(opts *ListListenersByLoadBalancerIDRequest) (*entity.ListListeners, error) {
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

func (s *LoadBalancerServiceV2) ListPoolsByLoadBalancerID(opts *ListPoolsByLoadBalancerIDRequest) (*entity.ListPools, error) {
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

func (s *LoadBalancerServiceV2) UpdatePoolMembers(opts *UpdatePoolMembersRequest) error {
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

func (s *LoadBalancerServiceV2) ListPoolMembers(opts *ListPoolMembersRequest) (*entity.ListMembers, error) {
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

func (s *LoadBalancerServiceV2) DeletePoolByID(opts *DeletePoolByIDRequest) error {
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

func (s *LoadBalancerServiceV2) DeleteListenerByID(opts *DeleteListenerByIDRequest) error {
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

func (s *LoadBalancerServiceV2) DeleteLoadBalancerByID(opts *DeleteLoadBalancerByIDRequest) error {
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

func (s *LoadBalancerServiceV2) GetPoolByID(opts *GetPoolByIDRequest) (*entity.Pool, error) {
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

func (s *LoadBalancerServiceV2) GetListenerByID(opts *GetListenerByIDRequest) (*entity.Listener, error) {
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

func (s *LoadBalancerServiceV2) ResizeLoadBalancerByID(opts *ResizeLoadBalancerByIDRequest) error {
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

func (s *LoadBalancerServiceV2) ScaleLoadBalancer(opts *ScaleLoadBalancerRequest) (*entity.LoadBalancer, error) {
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

func (s *LoadBalancerServiceV2) ListPolicies(opts *ListPoliciesRequest) (*entity.ListPolicies, error) {
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

func (s *LoadBalancerServiceV2) CreatePolicy(opts *CreatePolicyRequest) (*entity.Policy, error) {
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

func (s *LoadBalancerServiceV2) GetPolicyByID(opts *GetPolicyByIDRequest) (*entity.Policy, error) {
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

func (s *LoadBalancerServiceV2) UpdatePolicy(opts *UpdatePolicyRequest) error {
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

func (s *LoadBalancerServiceV2) DeletePolicyByID(opts *DeletePolicyByIDRequest) error {
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

func (s *LoadBalancerServiceV2) ReorderPolicies(opts *ReorderPoliciesRequest) error {
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


func (s *LoadBalancerServiceV2) ListCertificates(opts *ListCertificatesRequest) (*entity.ListCertificates, error) {
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

func (s *LoadBalancerServiceV2) GetCertificateByID(opts *GetCertificateByIDRequest) (*entity.Certificate, error) {
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

func (s *LoadBalancerServiceV2) CreateCertificate(opts *CreateCertificateRequest) (*entity.Certificate, error) {
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

func (s *LoadBalancerServiceV2) DeleteCertificateByID(opts *DeleteCertificateByIDRequest) error {
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
