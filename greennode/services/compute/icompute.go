package compute

import (
	"github.com/dannyota/greennode-community-sdk/v2/greennode/entity"
	sdkerror "github.com/dannyota/greennode-community-sdk/v2/greennode/sdkerror"
	computev2 "github.com/dannyota/greennode-community-sdk/v2/greennode/services/compute/v2"
)

type ComputeServiceV2 interface {
	CreateServer(popts computev2.ICreateServerRequest) (*entity.Server, sdkerror.Error)
	GetServerById(popts computev2.IGetServerByIdRequest) (*entity.Server, sdkerror.Error)
	DeleteServerById(popts computev2.IDeleteServerByIdRequest) sdkerror.Error
	UpdateServerSecgroupsByServerId(popts computev2.IUpdateServerSecgroupsByServerIdRequest) (*entity.Server, sdkerror.Error)
	AttachBlockVolume(popts computev2.IAttachBlockVolumeRequest) sdkerror.Error
	DetachBlockVolume(popts computev2.IDetachBlockVolumeRequest) sdkerror.Error
	AttachFloatingIp(popts computev2.IAttachFloatingIpRequest) sdkerror.Error
	DetachFloatingIp(popts computev2.IDetachFloatingIpRequest) sdkerror.Error
	ListServerGroupPolicies(popts computev2.IListServerGroupPoliciesRequest) (*entity.ListServerGroupPolicies, sdkerror.Error)
	DeleteServerGroupById(popts computev2.IDeleteServerGroupByIdRequest) sdkerror.Error
	ListServerGroups(popts computev2.IListServerGroupsRequest) (*entity.ListServerGroups, sdkerror.Error)
	CreateServerGroup(popts computev2.ICreateServerGroupRequest) (*entity.ServerGroup, sdkerror.Error)
}
