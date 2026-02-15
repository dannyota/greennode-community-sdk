package compute

import (
	"github.com/dannyota/greennode-community-sdk/v2/greennode/client"
	"github.com/dannyota/greennode-community-sdk/v2/greennode/entity"
	sdkerror "github.com/dannyota/greennode-community-sdk/v2/greennode/sdkerror"
	computev2 "github.com/dannyota/greennode-community-sdk/v2/greennode/services/compute/v2"
)

type ComputeServiceV2 interface {
	CreateServer(opts *computev2.CreateServerRequest) (*entity.Server, sdkerror.Error)
	GetServerByID(opts *computev2.GetServerByIDRequest) (*entity.Server, sdkerror.Error)
	DeleteServerByID(opts *computev2.DeleteServerByIDRequest) sdkerror.Error
	UpdateServerSecgroupsByServerID(opts *computev2.UpdateServerSecgroupsByServerIDRequest) (*entity.Server, sdkerror.Error)
	AttachBlockVolume(opts *computev2.AttachBlockVolumeRequest) sdkerror.Error
	DetachBlockVolume(opts *computev2.DetachBlockVolumeRequest) sdkerror.Error
	AttachFloatingIp(opts *computev2.AttachFloatingIpRequest) sdkerror.Error
	DetachFloatingIp(opts *computev2.DetachFloatingIpRequest) sdkerror.Error
	ListServerGroupPolicies(opts *computev2.ListServerGroupPoliciesRequest) (*entity.ListServerGroupPolicies, sdkerror.Error)
	DeleteServerGroupByID(opts *computev2.DeleteServerGroupByIDRequest) sdkerror.Error
	ListServerGroups(opts *computev2.ListServerGroupsRequest) (*entity.ListServerGroups, sdkerror.Error)
	CreateServerGroup(opts *computev2.CreateServerGroupRequest) (*entity.ServerGroup, sdkerror.Error)
}

func NewComputeServiceV2(svcClient client.ServiceClient) ComputeServiceV2 {
	return &computev2.ComputeServiceV2{
		VServerClient: svcClient,
	}
}
