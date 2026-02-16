package v2

import (
	"context"

	"github.com/dannyota/greennode-community-sdk/v2/greennode/client"
	sdkerror "github.com/dannyota/greennode-community-sdk/v2/greennode/sdkerror"
	"github.com/dannyota/greennode-community-sdk/v2/greennode/services/common"
)

type ComputeServiceV2 struct {
	Client *client.ServiceClient
}


const (
	defaultOffsetListServerGroups = 0
	defaultLimitListServerGroups  = 10
)

func (s *ComputeServiceV2) CreateServer(ctx context.Context, opts *CreateServerRequest) (*Server, error) {
	url := createServerURL(s.Client)
	resp := new(CreateServerResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithOKCodes(202).
		WithJSONBody(opts).
		WithJSONResponse(resp).
		WithJSONError(errResp)

	if _, sdkErr := s.Client.Post(ctx, url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.EcPurchaseIssue,
			sdkerror.EcVServerSubnetNotFound,
			sdkerror.EcVServerImageNotFound,
			sdkerror.EcVServerServerExceedQuota,
			sdkerror.EcVServerServerExceedCpuQuota,
			sdkerror.EcVServerServerFlavorSystemExceedQuota,
			sdkerror.EcVServerVolumeTypeNotFound,
			sdkerror.EcVServerNetworkNotFound,
			sdkerror.EcVServerVolumeExceedQuota,
			sdkerror.EcVServerVolumeSizeExceedGlobalQuota,
			sdkerror.EcVServerSecgroupNotFound,
			sdkerror.EcVServerServerExceedFloatingIpQuota,
			sdkerror.EcVServerServerImageNotSupported,
			sdkerror.EcVServerFlavorNotSupported,
			sdkerror.EcProjectConflict,
			sdkerror.EcVServerCreateBillingPaymentMethodNotAllowed).
			WithParameters(common.StructToMap(opts)).
			WithKVparameters("projectId", s.Client.ProjectID).
			AppendCategories(sdkerror.ErrCatVServer)
	}

	return resp.ToEntityServer(), nil
}

func (s *ComputeServiceV2) GetServerByID(ctx context.Context, opts *GetServerByIDRequest) (*Server, error) {
	url := getServerByIDURL(s.Client, opts)
	resp := new(GetServerByIDResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithOKCodes(200).
		WithJSONResponse(resp).
		WithJSONError(errResp)

	if _, sdkErr := s.Client.Get(ctx, url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.EcVServerServerNotFound).
			WithParameters(common.StructToMap(opts)).
			WithKVparameters("projectId", s.Client.ProjectID)
	}

	return resp.ToEntityServer(), nil
}

func (s *ComputeServiceV2) DeleteServerByID(ctx context.Context, opts *DeleteServerByIDRequest) error {
	url := deleteServerByIDURL(s.Client, opts)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithOKCodes(202).
		WithJSONBody(opts).
		WithJSONError(errResp)

	if _, sdkErr := s.Client.Delete(ctx, url, req); sdkErr != nil {
		return sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.EcVServerServerNotFound,
			sdkerror.EcVServerServerDeleteDeletingServer,
			sdkerror.EcVServerServerUpdatingSecgroups,
			sdkerror.EcVServerServerDeleteBillingServer,
			sdkerror.EcVServerServerDeleteCreatingServer,
			sdkerror.EcVServerVolumeInProcess).
			WithKVparameters("projectId", s.Client.ProjectID,
				"serverId", opts.ServerID)
	}

	return nil
}

func (s *ComputeServiceV2) UpdateServerSecgroupsByServerID(ctx context.Context, opts *UpdateServerSecgroupsByServerIDRequest) (*Server, error) {
	url := updateServerSecgroupsByServerIDURL(s.Client, opts)
	resp := new(UpdateServerSecgroupsByServerIDResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithOKCodes(202).
		WithJSONBody(opts).
		WithJSONResponse(resp).
		WithJSONError(errResp)

	if _, sdkErr := s.Client.Put(ctx, url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.EcVServerServerNotFound,
			sdkerror.EcVServerServerExpired,
			sdkerror.EcVServerServerUpdatingSecgroups,
			sdkerror.EcVServerSecgroupNotFound).
			WithKVparameters("projectId", s.Client.ProjectID,
				"serverId", opts.ServerID,
				"secgroupIds", opts.Secgroups)
	}

	return resp.ToEntityServer(), nil
}

func (s *ComputeServiceV2) AttachBlockVolume(ctx context.Context, opts *AttachBlockVolumeRequest) error {
	url := attachBlockVolumeURL(s.Client, opts)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithOKCodes(202).
		WithJSONBody(map[string]any{}).
		WithJSONError(errResp)

	if _, sdkErr := s.Client.Put(ctx, url, req); sdkErr != nil {
		return sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.EcVServerVolumeNotFound,
			sdkerror.EcVServerServerNotFound,
			sdkerror.EcVServerVolumeAvailable,
			sdkerror.EcVServerVolumeInProcess,
			sdkerror.EcVServerVolumeAlreadyAttached,
			sdkerror.EcVServerServerAttachEncryptedVolume,
			sdkerror.EcVServerVolumeAlreadyAttachedThisServer,
			sdkerror.EcVServerServerVolumeAttachQuotaExceeded).
			WithKVparameters("projectId", s.Client.ProjectID,
				"volumeId", opts.BlockVolumeID,
				"serverId", opts.ServerID)
	}

	return nil
}

func (s *ComputeServiceV2) DetachBlockVolume(ctx context.Context, opts *DetachBlockVolumeRequest) error {
	url := detachBlockVolumeURL(s.Client, opts)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithOKCodes(202).
		WithJSONBody(map[string]any{}).
		WithJSONError(errResp)

	if _, sdkErr := s.Client.Put(ctx, url, req); sdkErr != nil {
		return sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.EcVServerVolumeNotFound,
			sdkerror.EcVServerVolumeInProcess,
			sdkerror.EcVServerServerNotFound,
			sdkerror.EcVServerVolumeIsMigrating,
			sdkerror.EcVServerVolumeAvailable).
			WithKVparameters("projectId", s.Client.ProjectID,
				"volumeId", opts.BlockVolumeID,
				"serverId", opts.ServerID)
	}

	return nil
}

func (s *ComputeServiceV2) AttachFloatingIp(ctx context.Context, opts *AttachFloatingIpRequest) error {
	url := attachFloatingIpURL(s.Client, opts)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithOKCodes(204).
		WithJSONBody(opts).
		WithJSONError(errResp)

	if _, sdkErr := s.Client.Put(ctx, url, req); sdkErr != nil {
		return sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.EcVServerServerNotFound,
			sdkerror.EcVServerServerCanNotAttachFloatingIp,
			sdkerror.EcVServerInternalNetworkInterfaceNotFound).
			WithParameters(common.StructToMap(opts)).
			WithKVparameters("projectId", s.Client.ProjectID)
	}

	return nil
}

func (s *ComputeServiceV2) DetachFloatingIp(ctx context.Context, opts *DetachFloatingIpRequest) error {
	url := detachFloatingIpURL(s.Client, opts)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithOKCodes(204).
		WithJSONBody(opts).
		WithJSONError(errResp)

	if _, sdkErr := s.Client.Put(ctx, url, req); sdkErr != nil {
		return sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.EcVServerWanIpAvailable,
			sdkerror.EcVServerServerNotFound,
			sdkerror.EcVServerWanIDNotFound,
			sdkerror.EcVServerInternalNetworkInterfaceNotFound).
			WithParameters(common.StructToMap(opts)).
			WithKVparameters("projectId", s.Client.ProjectID)
	}

	return nil
}

func (s *ComputeServiceV2) ListServerGroupPolicies(ctx context.Context, opts *ListServerGroupPoliciesRequest) (*ListServerGroupPolicies, error) {
	url := listServerGroupPoliciesURL(s.Client)
	resp := new(ListServerGroupPoliciesResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithOKCodes(200).
		WithJSONResponse(resp).
		WithJSONError(errResp)

	if _, sdkErr := s.Client.Get(ctx, url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp).
			WithKVparameters("projectId", s.Client.ProjectID)
	}

	return resp.ToEntityListServerGroupPolicies(), nil
}

func (s *ComputeServiceV2) DeleteServerGroupByID(ctx context.Context, opts *DeleteServerGroupByIDRequest) error {
	url := deleteServerGroupByIDURL(s.Client, opts)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithOKCodes(204).
		WithJSONError(errResp)

	if _, sdkErr := s.Client.Delete(ctx, url, req); sdkErr != nil {
		return sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.EcVServerServerGroupNotFound,
			sdkerror.EcVServerServerGroupInUse).
			WithParameters(common.StructToMap(opts)).
			WithKVparameters("projectId", s.Client.ProjectID,
				"serverGroupId", opts.ServerGroupID)
	}

	return nil
}

func (s *ComputeServiceV2) ListServerGroups(ctx context.Context, opts *ListServerGroupsRequest) (*ListServerGroups, error) {
	url := listServerGroupsURL(s.Client, opts)
	resp := new(ListServerGroupsResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithOKCodes(200).
		WithJSONResponse(resp).
		WithJSONError(errResp)

	if _, sdkErr := s.Client.Get(ctx, url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp).
			WithParameters(common.StructToMap(opts)).
			WithKVparameters("projectId", s.Client.ProjectID)
	}

	return resp.ToEntityListServerGroups(), nil
}

func (s *ComputeServiceV2) CreateServerGroup(ctx context.Context, opts *CreateServerGroupRequest) (*ServerGroup, error) {
	url := createServerGroupURL(s.Client, opts)
	resp := new(CreateServerGroupResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithOKCodes(201).
		WithJSONBody(opts).
		WithJSONResponse(resp).
		WithJSONError(errResp)

	if _, sdkErr := s.Client.Post(ctx, url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.EcVServerServerGroupNameMustBeUnique).
			WithParameters(common.StructToMap(opts)).
			WithKVparameters("projectId", s.Client.ProjectID)
	}

	return resp.ToEntityServerGroup(), nil
}
