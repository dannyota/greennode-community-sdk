package v2

import (
	"github.com/dannyota/greennode-community-sdk/v2/greennode/client"
	"github.com/dannyota/greennode-community-sdk/v2/greennode/entity"
	sdkerror "github.com/dannyota/greennode-community-sdk/v2/greennode/sdkerror"
)

func (s *ComputeServiceV2) CreateServer(opts ICreateServerRequest) (*entity.Server, sdkerror.Error) {
	url := createServerUrl(s.VServerClient)
	resp := new(CreateServerResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithHeader("User-Agent", opts.ParseUserAgent()).
		WithOkCodes(202).
		WithJsonBody(opts.ToRequestBody()).
		WithJsonResponse(resp).
		WithJsonError(errResp)

	if _, sdkErr := s.VServerClient.Post(url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.WithErrorPurchaseIssue(errResp),
			sdkerror.WithErrorSubnetNotFound(errResp),
			sdkerror.WithErrorImageNotFound(errResp),
			sdkerror.WithErrorServerExceedQuota(errResp),
			sdkerror.WithErrorServerExceedCpuQuota(errResp),
			sdkerror.WithErrorServerFlavorSystemExceedQuota(errResp),
			sdkerror.WithErrorVolumeTypeNotFound(errResp),
			sdkerror.WithErrorNetworkNotFound(errResp),
			sdkerror.WithErrorVolumeExceedQuota(errResp),
			sdkerror.WithErrorVolumeSizeExceedGlobalQuota(errResp),
			sdkerror.WithErrorSecgroupNotFound(errResp),
			sdkerror.WithErrorServerExceedFloatingIpQuota(errResp),
			sdkerror.WithErrorServerImageNotSupported(errResp),
			sdkerror.WithErrorServerFlavorNotSupported(errResp),
			sdkerror.WithErrorProjectConflict(errResp),
			sdkerror.WithErrorServerCreateBillingPaymentMethodNotAllowed(errResp)).
			WithParameters(opts.ToMap()).
			WithKVparameters("projectId", s.getProjectId()).
			WithErrorCategories(sdkerror.ErrCatVServer)
	}

	return resp.ToEntityServer(), nil
}

func (s *ComputeServiceV2) GetServerById(opts IGetServerByIdRequest) (*entity.Server, sdkerror.Error) {
	url := getServerByIdUrl(s.VServerClient, opts)
	resp := new(GetServerByIdResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithHeader("User-Agent", opts.ParseUserAgent()).
		WithOkCodes(200).
		WithJsonResponse(resp).
		WithJsonError(errResp)

	if _, sdkErr := s.VServerClient.Get(url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.WithErrorServerNotFound(errResp)).
			WithParameters(opts.ToMap()).
			WithKVparameters("projectId", s.getProjectId())
	}

	return resp.ToEntityServer(), nil
}

func (s *ComputeServiceV2) DeleteServerById(opts IDeleteServerByIdRequest) sdkerror.Error {
	url := deleteServerByIdUrl(s.VServerClient, opts)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithHeader("User-Agent", opts.ParseUserAgent()).
		WithOkCodes(202).
		WithJsonBody(opts.ToRequestBody()).
		WithJsonError(errResp)

	if _, sdkErr := s.VServerClient.Delete(url, req); sdkErr != nil {
		return sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.WithErrorServerNotFound(errResp),
			sdkerror.WithErrorServerDeleteDeletingServer(errResp),
			sdkerror.WithErrorServerUpdatingSecgroups(errResp),
			sdkerror.WithErrorServerDeleteBillingServer(errResp),
			sdkerror.WithErrorServerDeleteCreatingServer(errResp),
			sdkerror.WithErrorVolumeInProcess(errResp)).
			WithKVparameters("projectId", s.getProjectId(),
				"serverId", opts.GetServerId())
	}

	return nil
}

func (s *ComputeServiceV2) UpdateServerSecgroupsByServerId(opts IUpdateServerSecgroupsByServerIdRequest) (*entity.Server, sdkerror.Error) {
	url := updateServerSecgroupsByServerIdUrl(s.VServerClient, opts)
	resp := new(UpdateServerSecgroupsByServerIdResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithHeader("User-Agent", opts.ParseUserAgent()).
		WithOkCodes(202).
		WithJsonBody(opts.ToRequestBody()).
		WithJsonResponse(resp).
		WithJsonError(errResp)

	if _, sdkErr := s.VServerClient.Put(url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.WithErrorServerNotFound(errResp),
			sdkerror.WithErrorServerExpired(errResp),
			sdkerror.WithErrorServerUpdatingSecgroups(errResp),
			sdkerror.WithErrorSecgroupNotFound(errResp)).
			WithKVparameters("projectId", s.getProjectId(),
				"serverId", opts.GetServerId(),
				"secgroupIds", opts.GetListSecgroupsIds())
	}

	return resp.ToEntityServer(), nil
}

func (s *ComputeServiceV2) AttachBlockVolume(opts IAttachBlockVolumeRequest) sdkerror.Error {
	url := attachBlockVolumeUrl(s.VServerClient, opts)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithOkCodes(202).
		WithJsonBody(map[string]interface{}{}).
		WithJsonError(errResp)

	if _, sdkErr := s.VServerClient.Put(url, req); sdkErr != nil {
		return sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.WithErrorVolumeNotFound(errResp),
			sdkerror.WithErrorServerNotFound(errResp),
			sdkerror.WithErrorVolumeAvailable(errResp),
			sdkerror.WithErrorVolumeInProcess(errResp),
			sdkerror.WithErrorVolumeAlreadyAttached(errResp),
			sdkerror.WithErrorServerAttachEncryptedVolume(errResp),
			sdkerror.WithErrorVolumeAlreadyAttachedThisServer(errResp),
			sdkerror.WithErrorServerAttachVolumeQuotaExceeded(errResp)).
			WithKVparameters("projectId", s.getProjectId(),
				"volumeId", opts.GetBlockVolumeId(),
				"serverId", opts.GetServerId())
	}

	return nil
}

func (s *ComputeServiceV2) DetachBlockVolume(opts IDetachBlockVolumeRequest) sdkerror.Error {
	url := detachBlockVolumeUrl(s.VServerClient, opts)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithOkCodes(202).
		WithJsonBody(map[string]interface{}{}).
		WithJsonError(errResp)

	if _, sdkErr := s.VServerClient.Put(url, req); sdkErr != nil {
		return sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.WithErrorVolumeNotFound(errResp),
			sdkerror.WithErrorVolumeInProcess(errResp),
			sdkerror.WithErrorServerNotFound(errResp),
			sdkerror.WithErrorVolumeIsMigrating(errResp),
			sdkerror.WithErrorVolumeAvailable(errResp)).
			WithKVparameters("projectId", s.getProjectId(),
				"volumeId", opts.GetBlockVolumeId(),
				"serverId", opts.GetServerId())
	}

	return nil
}

func (s *ComputeServiceV2) AttachFloatingIp(opts IAttachFloatingIpRequest) sdkerror.Error {
	url := attachFloatingIpUrl(s.VServerClient, opts)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithHeader("User-Agent", opts.ParseUserAgent()).
		WithOkCodes(204).
		WithJsonBody(opts.ToRequestBody()).
		WithJsonError(errResp)

	if _, sdkErr := s.VServerClient.Put(url, req); sdkErr != nil {
		return sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.WithErrorServerNotFound(errResp),
			sdkerror.WithErrorServerCanNotAttachFloatingIp(errResp),
			sdkerror.WithErrorInternalNetworkInterfaceNotFound(errResp)).
			WithParameters(opts.ToMap()).
			WithKVparameters("projectId", s.getProjectId())
	}

	return nil
}

func (s *ComputeServiceV2) DetachFloatingIp(opts IDetachFloatingIpRequest) sdkerror.Error {
	url := detachFloatingIpUrl(s.VServerClient, opts)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithHeader("User-Agent", opts.ParseUserAgent()).
		WithOkCodes(204).
		WithJsonBody(opts.ToRequestBody()).
		WithJsonError(errResp)

	if _, sdkErr := s.VServerClient.Put(url, req); sdkErr != nil {
		return sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.WithErrorWanIpAvailable(errResp),
			sdkerror.WithErrorServerNotFound(errResp),
			sdkerror.WithErrorWanIdNotFound(errResp),
			sdkerror.WithErrorInternalNetworkInterfaceNotFound(errResp)).
			WithParameters(opts.ToMap()).
			WithKVparameters("projectId", s.getProjectId())
	}

	return nil
}

func (s *ComputeServiceV2) ListServerGroupPolicies(opts IListServerGroupPoliciesRequest) (*entity.ListServerGroupPolicies, sdkerror.Error) {
	url := listServerGroupPoliciesUrl(s.VServerClient)
	resp := new(ListServerGroupPoliciesResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithOkCodes(200).
		WithJsonResponse(resp).
		WithJsonError(errResp)

	if opts != nil {
		req = req.WithHeader("User-Agent", opts.ParseUserAgent())
	}

	if _, sdkErr := s.VServerClient.Get(url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp).
			WithKVparameters("projectId", s.getProjectId())
	}

	return resp.ToEntityListServerGroupPolicies(), nil
}

func (s *ComputeServiceV2) DeleteServerGroupById(opts IDeleteServerGroupByIdRequest) sdkerror.Error {
	url := deleteServerGroupByIdUrl(s.VServerClient, opts)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithHeader("User-Agent", opts.ParseUserAgent()).
		WithOkCodes(204).
		WithJsonError(errResp)

	if _, sdkErr := s.VServerClient.Delete(url, req); sdkErr != nil {
		return sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.WithErrorServerGroupNotFound(errResp),
			sdkerror.WithErrorServerGroupInUse(errResp)).
			WithParameters(opts.ToMap()).
			WithKVparameters("projectId", s.getProjectId(),
				"serverGroupId", opts.GetServerGroupId())
	}

	return nil
}

func (s *ComputeServiceV2) ListServerGroups(opts IListServerGroupsRequest) (*entity.ListServerGroups, sdkerror.Error) {
	url := listServerGroupsUrl(s.VServerClient, opts)
	resp := new(ListServerGroupsResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithHeader("User-Agent", opts.ParseUserAgent()).
		WithOkCodes(200).
		WithJsonResponse(resp).
		WithJsonError(errResp)

	if _, sdkErr := s.VServerClient.Get(url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp).
			WithParameters(opts.ToMap()).
			WithKVparameters("projectId", s.getProjectId())
	}

	return resp.ToEntityListServerGroups(), nil
}

func (s *ComputeServiceV2) CreateServerGroup(opts ICreateServerGroupRequest) (*entity.ServerGroup, sdkerror.Error) {
	url := createServerGroupUrl(s.VServerClient, opts)
	resp := new(CreateServerGroupResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithHeader("User-Agent", opts.ParseUserAgent()).
		WithOkCodes(201).
		WithJsonBody(opts.ToRequestBody()).
		WithJsonResponse(resp).
		WithJsonError(errResp)

	if _, sdkErr := s.VServerClient.Post(url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.WithErrorServerGroupNameMustBeUnique(errResp)).
			WithParameters(opts.ToMap()).
			WithKVparameters("projectId", s.getProjectId())
	}

	return resp.ToEntityServerGroup(), nil
}
