package v2

import (
	"github.com/dannyota/greennode-community-sdk/v2/greennode/client"
	"github.com/dannyota/greennode-community-sdk/v2/greennode/entity"
	sdkerror "github.com/dannyota/greennode-community-sdk/v2/greennode/sdkerror"
)

func (s *ComputeServiceV2) CreateServer(popts ICreateServerRequest) (*entity.Server, sdkerror.Error) {
	url := createServerUrl(s.VServerClient)
	resp := new(CreateServerResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithHeader("User-Agent", popts.ParseUserAgent()).
		WithOkCodes(202).
		WithJsonBody(popts.ToRequestBody()).
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
			WithParameters(popts.ToMap()).
			WithKVparameters("projectId", s.getProjectId()).
			WithErrorCategories(sdkerror.ErrCatVServer)
	}

	return resp.ToEntityServer(), nil
}

func (s *ComputeServiceV2) GetServerById(popts IGetServerByIdRequest) (*entity.Server, sdkerror.Error) {
	url := getServerByIdUrl(s.VServerClient, popts)
	resp := new(GetServerByIdResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithHeader("User-Agent", popts.ParseUserAgent()).
		WithOkCodes(200).
		WithJsonResponse(resp).
		WithJsonError(errResp)

	if _, sdkErr := s.VServerClient.Get(url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.WithErrorServerNotFound(errResp)).
			WithParameters(popts.ToMap()).
			WithKVparameters("projectId", s.getProjectId())
	}

	return resp.ToEntityServer(), nil
}

func (s *ComputeServiceV2) DeleteServerById(popts IDeleteServerByIdRequest) sdkerror.Error {
	url := deleteServerByIdUrl(s.VServerClient, popts)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithHeader("User-Agent", popts.ParseUserAgent()).
		WithOkCodes(202).
		WithJsonBody(popts.ToRequestBody()).
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
				"serverId", popts.GetServerId())
	}

	return nil
}

func (s *ComputeServiceV2) UpdateServerSecgroupsByServerId(popts IUpdateServerSecgroupsByServerIdRequest) (*entity.Server, sdkerror.Error) {
	url := updateServerSecgroupsByServerIdUrl(s.VServerClient, popts)
	resp := new(UpdateServerSecgroupsByServerIdResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithHeader("User-Agent", popts.ParseUserAgent()).
		WithOkCodes(202).
		WithJsonBody(popts.ToRequestBody()).
		WithJsonResponse(resp).
		WithJsonError(errResp)

	if _, sdkErr := s.VServerClient.Put(url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.WithErrorServerNotFound(errResp),
			sdkerror.WithErrorServerExpired(errResp),
			sdkerror.WithErrorServerUpdatingSecgroups(errResp),
			sdkerror.WithErrorSecgroupNotFound(errResp)).
			WithKVparameters("projectId", s.getProjectId(),
				"serverId", popts.GetServerId(),
				"secgroupIds", popts.GetListSecgroupsIds())
	}

	return resp.ToEntityServer(), nil
}

func (s *ComputeServiceV2) AttachBlockVolume(popts IAttachBlockVolumeRequest) sdkerror.Error {
	url := attachBlockVolumeUrl(s.VServerClient, popts)
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
				"volumeId", popts.GetBlockVolumeId(),
				"serverId", popts.GetServerId())
	}

	return nil
}

func (s *ComputeServiceV2) DetachBlockVolume(popts IDetachBlockVolumeRequest) sdkerror.Error {
	url := detachBlockVolumeUrl(s.VServerClient, popts)
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
				"volumeId", popts.GetBlockVolumeId(),
				"serverId", popts.GetServerId())
	}

	return nil
}

func (s *ComputeServiceV2) AttachFloatingIp(popts IAttachFloatingIpRequest) sdkerror.Error {
	url := attachFloatingIpUrl(s.VServerClient, popts)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithHeader("User-Agent", popts.ParseUserAgent()).
		WithOkCodes(204).
		WithJsonBody(popts.ToRequestBody()).
		WithJsonError(errResp)

	if _, sdkErr := s.VServerClient.Put(url, req); sdkErr != nil {
		return sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.WithErrorServerNotFound(errResp),
			sdkerror.WithErrorServerCanNotAttachFloatingIp(errResp),
			sdkerror.WithErrorInternalNetworkInterfaceNotFound(errResp)).
			WithParameters(popts.ToMap()).
			WithKVparameters("projectId", s.getProjectId())
	}

	return nil
}

func (s *ComputeServiceV2) DetachFloatingIp(popts IDetachFloatingIpRequest) sdkerror.Error {
	url := detachFloatingIpUrl(s.VServerClient, popts)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithHeader("User-Agent", popts.ParseUserAgent()).
		WithOkCodes(204).
		WithJsonBody(popts.ToRequestBody()).
		WithJsonError(errResp)

	if _, sdkErr := s.VServerClient.Put(url, req); sdkErr != nil {
		return sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.WithErrorWanIpAvailable(errResp),
			sdkerror.WithErrorServerNotFound(errResp),
			sdkerror.WithErrorWanIdNotFound(errResp),
			sdkerror.WithErrorInternalNetworkInterfaceNotFound(errResp)).
			WithParameters(popts.ToMap()).
			WithKVparameters("projectId", s.getProjectId())
	}

	return nil
}

func (s *ComputeServiceV2) ListServerGroupPolicies(popts IListServerGroupPoliciesRequest) (*entity.ListServerGroupPolicies, sdkerror.Error) {
	url := listServerGroupPoliciesUrl(s.VServerClient)
	resp := new(ListServerGroupPoliciesResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithOkCodes(200).
		WithJsonResponse(resp).
		WithJsonError(errResp)

	if popts != nil {
		req = req.WithHeader("User-Agent", popts.ParseUserAgent())
	}

	if _, sdkErr := s.VServerClient.Get(url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp).
			WithKVparameters("projectId", s.getProjectId())
	}

	return resp.ToEntityListServerGroupPolicies(), nil
}

func (s *ComputeServiceV2) DeleteServerGroupById(popts IDeleteServerGroupByIdRequest) sdkerror.Error {
	url := deleteServerGroupByIdUrl(s.VServerClient, popts)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithHeader("User-Agent", popts.ParseUserAgent()).
		WithOkCodes(204).
		WithJsonError(errResp)

	if _, sdkErr := s.VServerClient.Delete(url, req); sdkErr != nil {
		return sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.WithErrorServerGroupNotFound(errResp),
			sdkerror.WithErrorServerGroupInUse(errResp)).
			WithParameters(popts.ToMap()).
			WithKVparameters("projectId", s.getProjectId(),
				"serverGroupId", popts.GetServerGroupId())
	}

	return nil
}

func (s *ComputeServiceV2) ListServerGroups(popts IListServerGroupsRequest) (*entity.ListServerGroups, sdkerror.Error) {
	url := listServerGroupsUrl(s.VServerClient, popts)
	resp := new(ListServerGroupsResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithHeader("User-Agent", popts.ParseUserAgent()).
		WithOkCodes(200).
		WithJsonResponse(resp).
		WithJsonError(errResp)

	if _, sdkErr := s.VServerClient.Get(url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp).
			WithParameters(popts.ToMap()).
			WithKVparameters("projectId", s.getProjectId())
	}

	return resp.ToEntityListServerGroups(), nil
}

func (s *ComputeServiceV2) CreateServerGroup(popts ICreateServerGroupRequest) (*entity.ServerGroup, sdkerror.Error) {
	url := createServerGroupUrl(s.VServerClient, popts)
	resp := new(CreateServerGroupResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithHeader("User-Agent", popts.ParseUserAgent()).
		WithOkCodes(201).
		WithJsonBody(popts.ToRequestBody()).
		WithJsonResponse(resp).
		WithJsonError(errResp)

	if _, sdkErr := s.VServerClient.Post(url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.WithErrorServerGroupNameMustBeUnique(errResp)).
			WithParameters(popts.ToMap()).
			WithKVparameters("projectId", s.getProjectId())
	}

	return resp.ToEntityServerGroup(), nil
}
