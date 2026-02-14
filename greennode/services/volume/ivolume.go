package volume

import (
	"github.com/dannyota/greennode-community-sdk/v2/greennode/entity"
	sdkerror "github.com/dannyota/greennode-community-sdk/v2/greennode/sdkerror"
	volumev1 "github.com/dannyota/greennode-community-sdk/v2/greennode/services/volume/v1"
	volumev2 "github.com/dannyota/greennode-community-sdk/v2/greennode/services/volume/v2"
)

type VolumeServiceV2 interface {
	CreateBlockVolume(opts volumev2.ICreateBlockVolumeRequest) (*entity.Volume, sdkerror.Error)
	DeleteBlockVolumeById(opts volumev2.IDeleteBlockVolumeByIdRequest) sdkerror.Error
	ListBlockVolumes(opts volumev2.IListBlockVolumesRequest) (*entity.ListVolumes, sdkerror.Error)
	GetBlockVolumeById(opts volumev2.IGetBlockVolumeByIdRequest) (*entity.Volume, sdkerror.Error)
	ResizeBlockVolumeById(opts volumev2.IResizeBlockVolumeByIdRequest) (*entity.Volume, sdkerror.Error)
	GetUnderBlockVolumeId(opts volumev2.IGetUnderBlockVolumeIdRequest) (*entity.Volume, sdkerror.Error)
	MigrateBlockVolumeById(opts volumev2.IMigrateBlockVolumeByIdRequest) sdkerror.Error

	// Snapshot
	ListSnapshotsByBlockVolumeId(opts volumev2.IListSnapshotsByBlockVolumeIdRequest) (*entity.ListSnapshots, sdkerror.Error)
	CreateSnapshotByBlockVolumeId(opts volumev2.ICreateSnapshotByBlockVolumeIdRequest) (*entity.Snapshot, sdkerror.Error)
	DeleteSnapshotById(opts volumev2.IDeleteSnapshotByIdRequest) sdkerror.Error
}

type VolumeServiceV1 interface {
	// Volume Type Api group
	GetVolumeTypeById(opts volumev1.IGetVolumeTypeByIdRequest) (*entity.VolumeType, sdkerror.Error)
	GetListVolumeTypes(opts volumev1.IGetListVolumeTypeRequest) (*entity.ListVolumeType, sdkerror.Error)
	GetVolumeTypeZones(opts volumev1.IGetVolumeTypeZonesRequest) (*entity.ListVolumeTypeZones, sdkerror.Error)
	GetDefaultVolumeType() (*entity.VolumeType, sdkerror.Error)
}
