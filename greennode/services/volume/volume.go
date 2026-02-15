package volume

import (
	"github.com/dannyota/greennode-community-sdk/v2/greennode/client"
	"github.com/dannyota/greennode-community-sdk/v2/greennode/entity"
	sdkerror "github.com/dannyota/greennode-community-sdk/v2/greennode/sdkerror"
	volumev1 "github.com/dannyota/greennode-community-sdk/v2/greennode/services/volume/v1"
	volumev2 "github.com/dannyota/greennode-community-sdk/v2/greennode/services/volume/v2"
)

type VolumeServiceV2 interface {
	CreateBlockVolume(opts *volumev2.CreateBlockVolumeRequest) (*entity.Volume, sdkerror.Error)
	DeleteBlockVolumeByID(opts *volumev2.DeleteBlockVolumeByIDRequest) sdkerror.Error
	ListBlockVolumes(opts *volumev2.ListBlockVolumesRequest) (*entity.ListVolumes, sdkerror.Error)
	GetBlockVolumeByID(opts *volumev2.GetBlockVolumeByIDRequest) (*entity.Volume, sdkerror.Error)
	ResizeBlockVolumeByID(opts *volumev2.ResizeBlockVolumeByIDRequest) (*entity.Volume, sdkerror.Error)
	GetUnderBlockVolumeID(opts *volumev2.GetUnderBlockVolumeIDRequest) (*entity.Volume, sdkerror.Error)
	MigrateBlockVolumeByID(opts *volumev2.MigrateBlockVolumeByIDRequest) sdkerror.Error

	// Snapshot
	ListSnapshotsByBlockVolumeID(opts *volumev2.ListSnapshotsByBlockVolumeIDRequest) (*entity.ListSnapshots, sdkerror.Error)
	CreateSnapshotByBlockVolumeID(opts *volumev2.CreateSnapshotByBlockVolumeIDRequest) (*entity.Snapshot, sdkerror.Error)
	DeleteSnapshotByID(opts *volumev2.DeleteSnapshotByIDRequest) sdkerror.Error
}

type VolumeServiceV1 interface {
	// Volume Type Api group
	GetVolumeTypeByID(opts *volumev1.GetVolumeTypeByIDRequest) (*entity.VolumeType, sdkerror.Error)
	GetListVolumeTypes(opts *volumev1.GetListVolumeTypeRequest) (*entity.ListVolumeType, sdkerror.Error)
	GetVolumeTypeZones(opts *volumev1.GetVolumeTypeZonesRequest) (*entity.ListVolumeTypeZones, sdkerror.Error)
	GetDefaultVolumeType() (*entity.VolumeType, sdkerror.Error)
}

func NewVolumeServiceV2(svcClient client.ServiceClient) VolumeServiceV2 {
	return &volumev2.VolumeServiceV2{
		VServerClient: svcClient,
	}
}

func NewVolumeServiceV1(svcClient client.ServiceClient) VolumeServiceV1 {
	return &volumev1.VolumeServiceV1{
		VServerClient: svcClient,
	}
}
