package compute

import (
	"github.com/dannyota/greennode-community-sdk/v2/greennode/entity"
	sdkerror "github.com/dannyota/greennode-community-sdk/v2/greennode/sdk_error"
	computev2 "github.com/dannyota/greennode-community-sdk/v2/greennode/services/compute/v2"
)

type IComputeServiceV2 interface {
	CreateServer(popts computev2.ICreateServerRequest) (*entity.Server, sdkerror.IError)
	GetServerById(popts computev2.IGetServerByIdRequest) (*entity.Server, sdkerror.IError)
	DeleteServerById(popts computev2.IDeleteServerByIdRequest) sdkerror.IError
	UpdateServerSecgroupsByServerId(popts computev2.IUpdateServerSecgroupsByServerIdRequest) (*entity.Server, sdkerror.IError)
	AttachBlockVolume(popts computev2.IAttachBlockVolumeRequest) sdkerror.IError
	DetachBlockVolume(popts computev2.IDetachBlockVolumeRequest) sdkerror.IError
	AttachFloatingIp(popts computev2.IAttachFloatingIpRequest) sdkerror.IError
	DetachFloatingIp(popts computev2.IDetachFloatingIpRequest) sdkerror.IError
	ListServerGroupPolicies(popts computev2.IListServerGroupPoliciesRequest) (*entity.ListServerGroupPolicies, sdkerror.IError)
	DeleteServerGroupById(popts computev2.IDeleteServerGroupByIdRequest) sdkerror.IError
	ListServerGroups(popts computev2.IListServerGroupsRequest) (*entity.ListServerGroups, sdkerror.IError)
	CreateServerGroup(popts computev2.ICreateServerGroupRequest) (*entity.ServerGroup, sdkerror.IError)
}
