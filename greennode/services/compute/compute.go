package compute

import (
	"github.com/dannyota/greennode-community-sdk/v2/greennode/client"
	"github.com/dannyota/greennode-community-sdk/v2/greennode/entity"
	computev2 "github.com/dannyota/greennode-community-sdk/v2/greennode/services/compute/v2"
)

type ComputeServiceV2 interface {
	CreateServer(opts *computev2.CreateServerRequest) (*entity.Server, error)
	GetServerByID(opts *computev2.GetServerByIDRequest) (*entity.Server, error)
	DeleteServerByID(opts *computev2.DeleteServerByIDRequest) error
	UpdateServerSecgroupsByServerID(opts *computev2.UpdateServerSecgroupsByServerIDRequest) (*entity.Server, error)
	AttachBlockVolume(opts *computev2.AttachBlockVolumeRequest) error
	DetachBlockVolume(opts *computev2.DetachBlockVolumeRequest) error
	AttachFloatingIp(opts *computev2.AttachFloatingIpRequest) error
	DetachFloatingIp(opts *computev2.DetachFloatingIpRequest) error
	ListServerGroupPolicies(opts *computev2.ListServerGroupPoliciesRequest) (*entity.ListServerGroupPolicies, error)
	DeleteServerGroupByID(opts *computev2.DeleteServerGroupByIDRequest) error
	ListServerGroups(opts *computev2.ListServerGroupsRequest) (*entity.ListServerGroups, error)
	CreateServerGroup(opts *computev2.CreateServerGroupRequest) (*entity.ServerGroup, error)
}

func NewComputeServiceV2(svcClient client.ServiceClient) *computev2.ComputeServiceV2 {
	return &computev2.ComputeServiceV2{
		VServerClient: svcClient,
	}
}
