package v2

type ICreateBlockVolumeRequest interface {
	ToRequestBody() interface{}
	ToMap() map[string]interface{}
	GetListParameters() []interface{}
	GetSize() int64
	GetVolumeType() string
	GetVolumeName() string
	GetZone() string
	GetPoolName() string
	WithPoc(isPoc bool) ICreateBlockVolumeRequest
	WithAutoRenew(val bool) ICreateBlockVolumeRequest
	WithMultiAttach(multiAttach bool) ICreateBlockVolumeRequest
	WithSize(size int64) ICreateBlockVolumeRequest
	WithEncryptionType(et EncryptType) ICreateBlockVolumeRequest
	WithVolumeType(volumeTypeId string) ICreateBlockVolumeRequest
	WithZone(zone string) ICreateBlockVolumeRequest
	WithPoolName(poolName string) ICreateBlockVolumeRequest
	WithVolumeRestoreFromSnapshot(snapshotID, volumeTypeID string) ICreateBlockVolumeRequest
	WithTags(tags ...string) ICreateBlockVolumeRequest
}

type IDeleteBlockVolumeByIdRequest interface {
	GetBlockVolumeId() string
}

type IListBlockVolumesRequest interface {
	WithName(name string) IListBlockVolumesRequest
	ToQuery() (string, error)
	GetDefaultQuery() string
	ToMap() map[string]interface{}
}

type IGetBlockVolumeByIdRequest interface {
	GetBlockVolumeId() string
}

type IResizeBlockVolumeByIdRequest interface {
	ToRequestBody() interface{}
	GetBlockVolumeId() string
	GetSize() int
	GetVolumeTypeId() string
}

type IListSnapshotsByBlockVolumeIdRequest interface {
	GetBlockVolumeId() string
	ToQuery() (string, error)
	GetDefaultQuery() string
}

type ICreateSnapshotByBlockVolumeIdRequest interface {
	GetBlockVolumeId() string
	ToRequestBody() interface{}
	WithDescription(desc string) ICreateSnapshotByBlockVolumeIdRequest
	WithPermanently(val bool) ICreateSnapshotByBlockVolumeIdRequest
	WithRetainedDay(val uint64) ICreateSnapshotByBlockVolumeIdRequest
}

type IDeleteSnapshotByIdRequest interface {
	GetSnapshotId() string
	GetBlockVolumeId() string
}

type IGetUnderBlockVolumeIdRequest interface {
	GetBlockVolumeId() string
}

type IMigrateBlockVolumeByIdRequest interface {
	GetBlockVolumeId() string
	ToRequestBody() interface{}
	WithTags(tags ...string) IMigrateBlockVolumeByIdRequest
	WithAction(action MigrateAction) IMigrateBlockVolumeByIdRequest
	WithConfirm(confirm bool) IMigrateBlockVolumeByIdRequest
	IsConfirm() bool
}
