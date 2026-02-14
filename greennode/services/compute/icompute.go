package compute

import (
	"github.com/dannyota/greennode-community-sdk/v2/greennode/entity"
	sdkerror "github.com/dannyota/greennode-community-sdk/v2/greennode/sdkerror"
	computev2 "github.com/dannyota/greennode-community-sdk/v2/greennode/services/compute/v2"
)

type ComputeServiceV2 interface {
	CreateServer(opts computev2.ICreateServerRequest) (*entity.Server, sdkerror.Error)
	GetServerById(opts computev2.IGetServerByIdRequest) (*entity.Server, sdkerror.Error)
	DeleteServerById(opts computev2.IDeleteServerByIdRequest) sdkerror.Error
	UpdateServerSecgroupsByServerId(opts computev2.IUpdateServerSecgroupsByServerIdRequest) (*entity.Server, sdkerror.Error)
	AttachBlockVolume(opts computev2.IAttachBlockVolumeRequest) sdkerror.Error
	DetachBlockVolume(opts computev2.IDetachBlockVolumeRequest) sdkerror.Error
	AttachFloatingIp(opts computev2.IAttachFloatingIpRequest) sdkerror.Error
	DetachFloatingIp(opts computev2.IDetachFloatingIpRequest) sdkerror.Error
	ListServerGroupPolicies(opts computev2.IListServerGroupPoliciesRequest) (*entity.ListServerGroupPolicies, sdkerror.Error)
	DeleteServerGroupById(opts computev2.IDeleteServerGroupByIdRequest) sdkerror.Error
	ListServerGroups(opts computev2.IListServerGroupsRequest) (*entity.ListServerGroups, sdkerror.Error)
	CreateServerGroup(opts computev2.ICreateServerGroupRequest) (*entity.ServerGroup, sdkerror.Error)
}
