package volume

import (
	"github.com/dannyota/greennode-community-sdk/v2/greennode/client"
	"github.com/dannyota/greennode-community-sdk/v2/greennode/entity"
	volumev1 "github.com/dannyota/greennode-community-sdk/v2/greennode/services/volume/v1"
	volumev2 "github.com/dannyota/greennode-community-sdk/v2/greennode/services/volume/v2"
)

type VolumeServiceV2 interface {
	CreateBlockVolume(opts *volumev2.CreateBlockVolumeRequest) (*entity.Volume, error)
	DeleteBlockVolumeByID(opts *volumev2.DeleteBlockVolumeByIDRequest) error
	ListBlockVolumes(opts *volumev2.ListBlockVolumesRequest) (*entity.ListVolumes, error)
	GetBlockVolumeByID(opts *volumev2.GetBlockVolumeByIDRequest) (*entity.Volume, error)
	ResizeBlockVolumeByID(opts *volumev2.ResizeBlockVolumeByIDRequest) (*entity.Volume, error)
	GetUnderBlockVolumeID(opts *volumev2.GetUnderBlockVolumeIDRequest) (*entity.Volume, error)
	MigrateBlockVolumeByID(opts *volumev2.MigrateBlockVolumeByIDRequest) error

	// Snapshot
	ListSnapshotsByBlockVolumeID(opts *volumev2.ListSnapshotsByBlockVolumeIDRequest) (*entity.ListSnapshots, error)
	CreateSnapshotByBlockVolumeID(opts *volumev2.CreateSnapshotByBlockVolumeIDRequest) (*entity.Snapshot, error)
	DeleteSnapshotByID(opts *volumev2.DeleteSnapshotByIDRequest) error
}

type VolumeServiceV1 interface {
	// Volume Type Api group
	GetVolumeTypeByID(opts *volumev1.GetVolumeTypeByIDRequest) (*entity.VolumeType, error)
	GetListVolumeTypes(opts *volumev1.GetListVolumeTypeRequest) (*entity.ListVolumeType, error)
	GetVolumeTypeZones(opts *volumev1.GetVolumeTypeZonesRequest) (*entity.ListVolumeTypeZones, error)
	GetDefaultVolumeType() (*entity.VolumeType, error)
}

func NewVolumeServiceV2(svcClient client.ServiceClient) *volumev2.VolumeServiceV2 {
	return &volumev2.VolumeServiceV2{
		VServerClient: svcClient,
	}
}

func NewVolumeServiceV1(svcClient client.ServiceClient) *volumev1.VolumeServiceV1 {
	return &volumev1.VolumeServiceV1{
		VServerClient: svcClient,
	}
}
