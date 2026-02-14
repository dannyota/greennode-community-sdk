package volume

import (
	"github.com/dannyota/greennode-community-sdk/v2/greennode/entity"
	sdkerror "github.com/dannyota/greennode-community-sdk/v2/greennode/sdk_error"
	volumev1 "github.com/dannyota/greennode-community-sdk/v2/greennode/services/volume/v1"
	volumev2 "github.com/dannyota/greennode-community-sdk/v2/greennode/services/volume/v2"
)

type IVolumeServiceV2 interface {
	CreateBlockVolume(popts volumev2.ICreateBlockVolumeRequest) (*entity.Volume, sdkerror.IError)
	DeleteBlockVolumeById(popts volumev2.IDeleteBlockVolumeByIdRequest) sdkerror.IError
	ListBlockVolumes(popts volumev2.IListBlockVolumesRequest) (*entity.ListVolumes, sdkerror.IError)
	GetBlockVolumeById(popts volumev2.IGetBlockVolumeByIdRequest) (*entity.Volume, sdkerror.IError)
	ResizeBlockVolumeById(popts volumev2.IResizeBlockVolumeByIdRequest) (*entity.Volume, sdkerror.IError)
	GetUnderBlockVolumeId(popts volumev2.IGetUnderBlockVolumeIdRequest) (*entity.Volume, sdkerror.IError)
	MigrateBlockVolumeById(popts volumev2.IMigrateBlockVolumeByIdRequest) sdkerror.IError

	// Snapshot
	ListSnapshotsByBlockVolumeId(popts volumev2.IListSnapshotsByBlockVolumeIdRequest) (*entity.ListSnapshots, sdkerror.IError)
	CreateSnapshotByBlockVolumeId(popts volumev2.ICreateSnapshotByBlockVolumeIdRequest) (*entity.Snapshot, sdkerror.IError)
	DeleteSnapshotById(popts volumev2.IDeleteSnapshotByIdRequest) sdkerror.IError
}

type IVolumeServiceV1 interface {
	// Volume Type Api group
	GetVolumeTypeById(popts volumev1.IGetVolumeTypeByIdRequest) (*entity.VolumeType, sdkerror.IError)
	GetListVolumeTypes(popts volumev1.IGetListVolumeTypeRequest) (*entity.ListVolumeType, sdkerror.IError)
	GetVolumeTypeZones(popts volumev1.IGetVolumeTypeZonesRequest) (*entity.ListVolumeTypeZones, sdkerror.IError)
	GetDefaultVolumeType() (*entity.VolumeType, sdkerror.IError)
}
