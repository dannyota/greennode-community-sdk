package volume

import (
	"github.com/dannyota/greennode-community-sdk/v2/greennode/entity"
	sdkerror "github.com/dannyota/greennode-community-sdk/v2/greennode/sdkerror"
	volumev1 "github.com/dannyota/greennode-community-sdk/v2/greennode/services/volume/v1"
	volumev2 "github.com/dannyota/greennode-community-sdk/v2/greennode/services/volume/v2"
)

type VolumeServiceV2 interface {
	CreateBlockVolume(popts volumev2.ICreateBlockVolumeRequest) (*entity.Volume, sdkerror.Error)
	DeleteBlockVolumeById(popts volumev2.IDeleteBlockVolumeByIdRequest) sdkerror.Error
	ListBlockVolumes(popts volumev2.IListBlockVolumesRequest) (*entity.ListVolumes, sdkerror.Error)
	GetBlockVolumeById(popts volumev2.IGetBlockVolumeByIdRequest) (*entity.Volume, sdkerror.Error)
	ResizeBlockVolumeById(popts volumev2.IResizeBlockVolumeByIdRequest) (*entity.Volume, sdkerror.Error)
	GetUnderBlockVolumeId(popts volumev2.IGetUnderBlockVolumeIdRequest) (*entity.Volume, sdkerror.Error)
	MigrateBlockVolumeById(popts volumev2.IMigrateBlockVolumeByIdRequest) sdkerror.Error

	// Snapshot
	ListSnapshotsByBlockVolumeId(popts volumev2.IListSnapshotsByBlockVolumeIdRequest) (*entity.ListSnapshots, sdkerror.Error)
	CreateSnapshotByBlockVolumeId(popts volumev2.ICreateSnapshotByBlockVolumeIdRequest) (*entity.Snapshot, sdkerror.Error)
	DeleteSnapshotById(popts volumev2.IDeleteSnapshotByIdRequest) sdkerror.Error
}

type VolumeServiceV1 interface {
	// Volume Type Api group
	GetVolumeTypeById(popts volumev1.IGetVolumeTypeByIdRequest) (*entity.VolumeType, sdkerror.Error)
	GetListVolumeTypes(popts volumev1.IGetListVolumeTypeRequest) (*entity.ListVolumeType, sdkerror.Error)
	GetVolumeTypeZones(popts volumev1.IGetVolumeTypeZonesRequest) (*entity.ListVolumeTypeZones, sdkerror.Error)
	GetDefaultVolumeType() (*entity.VolumeType, sdkerror.Error)
}
