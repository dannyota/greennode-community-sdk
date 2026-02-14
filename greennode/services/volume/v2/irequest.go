package v2

type ICreateBlockVolumeRequest interface {
	ToRequestBody() interface{}
	ToMap() map[string]interface{}
	ListParameters() []interface{}
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
	WithVolumeType(volumeTypeID string) ICreateBlockVolumeRequest
	WithZone(zone string) ICreateBlockVolumeRequest
	WithPoolName(poolName string) ICreateBlockVolumeRequest
	WithVolumeRestoreFromSnapshot(snapshotID, volumeTypeID string) ICreateBlockVolumeRequest
	WithTags(tags ...string) ICreateBlockVolumeRequest
}

type IDeleteBlockVolumeByIDRequest interface {
	GetBlockVolumeID() string
}

type IListBlockVolumesRequest interface {
	WithName(name string) IListBlockVolumesRequest
	ToQuery() (string, error)
	GetDefaultQuery() string
	ToMap() map[string]interface{}
}

type IGetBlockVolumeByIDRequest interface {
	GetBlockVolumeID() string
}

type IResizeBlockVolumeByIDRequest interface {
	ToRequestBody() interface{}
	GetBlockVolumeID() string
	GetSize() int
	GetVolumeTypeID() string
}

type IListSnapshotsByBlockVolumeIDRequest interface {
	GetBlockVolumeID() string
	ToQuery() (string, error)
	GetDefaultQuery() string
}

type ICreateSnapshotByBlockVolumeIDRequest interface {
	GetBlockVolumeID() string
	ToRequestBody() interface{}
	WithDescription(desc string) ICreateSnapshotByBlockVolumeIDRequest
	WithPermanently(val bool) ICreateSnapshotByBlockVolumeIDRequest
	WithRetainedDay(val uint64) ICreateSnapshotByBlockVolumeIDRequest
}

type IDeleteSnapshotByIDRequest interface {
	GetSnapshotID() string
	GetBlockVolumeID() string
}

type IGetUnderBlockVolumeIDRequest interface {
	GetBlockVolumeID() string
}

type IMigrateBlockVolumeByIDRequest interface {
	GetBlockVolumeID() string
	ToRequestBody() interface{}
	WithTags(tags ...string) IMigrateBlockVolumeByIDRequest
	WithAction(action MigrateAction) IMigrateBlockVolumeByIDRequest
	WithConfirm(confirm bool) IMigrateBlockVolumeByIDRequest
	IsConfirm() bool
}
