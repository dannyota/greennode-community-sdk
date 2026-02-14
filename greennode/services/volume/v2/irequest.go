package v2

type ICreateBlockVolumeRequest interface {
	ToRequestBody() any
	ToMap() map[string]any
	ListParameters() []any
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
	ToMap() map[string]any
}

type IGetBlockVolumeByIDRequest interface {
	GetBlockVolumeID() string
}

type IResizeBlockVolumeByIDRequest interface {
	ToRequestBody() any
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
	ToRequestBody() any
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
	ToRequestBody() any
	WithTags(tags ...string) IMigrateBlockVolumeByIDRequest
	WithAction(action MigrateAction) IMigrateBlockVolumeByIDRequest
	WithConfirm(confirm bool) IMigrateBlockVolumeByIDRequest
	IsConfirm() bool
}
