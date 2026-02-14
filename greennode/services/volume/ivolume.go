package volume

import (
	"github.com/dannyota/greennode-community-sdk/v2/greennode/entity"
	sdkerror "github.com/dannyota/greennode-community-sdk/v2/greennode/sdkerror"
	volumev1 "github.com/dannyota/greennode-community-sdk/v2/greennode/services/volume/v1"
	volumev2 "github.com/dannyota/greennode-community-sdk/v2/greennode/services/volume/v2"
)

type VolumeServiceV2 interface {
	CreateBlockVolume(opts volumev2.ICreateBlockVolumeRequest) (*entity.Volume, sdkerror.Error)
	DeleteBlockVolumeByID(opts volumev2.IDeleteBlockVolumeByIDRequest) sdkerror.Error
	ListBlockVolumes(opts volumev2.IListBlockVolumesRequest) (*entity.ListVolumes, sdkerror.Error)
	GetBlockVolumeByID(opts volumev2.IGetBlockVolumeByIDRequest) (*entity.Volume, sdkerror.Error)
	ResizeBlockVolumeByID(opts volumev2.IResizeBlockVolumeByIDRequest) (*entity.Volume, sdkerror.Error)
	GetUnderBlockVolumeID(opts volumev2.IGetUnderBlockVolumeIDRequest) (*entity.Volume, sdkerror.Error)
	MigrateBlockVolumeByID(opts volumev2.IMigrateBlockVolumeByIDRequest) sdkerror.Error

	// Snapshot
	ListSnapshotsByBlockVolumeID(opts volumev2.IListSnapshotsByBlockVolumeIDRequest) (*entity.ListSnapshots, sdkerror.Error)
	CreateSnapshotByBlockVolumeID(opts volumev2.ICreateSnapshotByBlockVolumeIDRequest) (*entity.Snapshot, sdkerror.Error)
	DeleteSnapshotByID(opts volumev2.IDeleteSnapshotByIDRequest) sdkerror.Error
}

type VolumeServiceV1 interface {
	// Volume Type Api group
	GetVolumeTypeByID(opts volumev1.IGetVolumeTypeByIDRequest) (*entity.VolumeType, sdkerror.Error)
	GetListVolumeTypes(opts volumev1.IGetListVolumeTypeRequest) (*entity.ListVolumeType, sdkerror.Error)
	GetVolumeTypeZones(opts volumev1.IGetVolumeTypeZonesRequest) (*entity.ListVolumeTypeZones, sdkerror.Error)
	GetDefaultVolumeType() (*entity.VolumeType, sdkerror.Error)
}
