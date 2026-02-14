package v2

import (
	"fmt"
	"net/url"

	"strconv"

	"github.com/dannyota/greennode-community-sdk/v2/greennode/services/common"
)

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

func NewCreateBlockVolumeRequest(volumeName, volumeType string, size int64) *CreateBlockVolumeRequest {
	opt := new(CreateBlockVolumeRequest)
	opt.VolumeTypeID = volumeType
	opt.CreatedFrom = CreateFromNew
	opt.Name = volumeName
	opt.Size = size

	return opt
}

func NewDeleteBlockVolumeByIDRequest(volumeID string) *DeleteBlockVolumeByIDRequest {
	opt := new(DeleteBlockVolumeByIDRequest)
	opt.BlockVolumeID = volumeID
	return opt
}

func NewListBlockVolumesRequest(page, size int) *ListBlockVolumesRequest {
	opt := new(ListBlockVolumesRequest)
	opt.Page = page
	opt.Size = size
	return opt
}

func NewGetBlockVolumeByIDRequest(volumeID string) *GetBlockVolumeByIDRequest {
	opt := new(GetBlockVolumeByIDRequest)
	opt.BlockVolumeID = volumeID
	return opt
}

func NewResizeBlockVolumeByIDRequest(volumeID, volumeType string, size int) *ResizeBlockVolumeByIDRequest {
	opt := new(ResizeBlockVolumeByIDRequest)
	opt.BlockVolumeID = volumeID
	opt.NewSize = size
	opt.VolumeTypeID = volumeType
	return opt
}

func NewGetUnderVolumeIDRequest(volumeID string) *GetUnderBlockVolumeIDRequest {
	opt := new(GetUnderBlockVolumeIDRequest)
	opt.BlockVolumeID = volumeID
	return opt
}

func NewMigrateBlockVolumeByIDRequest(volumeID, volumeType string) *MigrateBlockVolumeByIDRequest {
	opt := new(MigrateBlockVolumeByIDRequest)
	opt.BlockVolumeID = volumeID
	opt.VolumeTypeID = volumeType
	opt.Action = InitMigrateAction
	return opt
}

const (
	CreateFromNew      = CreateVolumeFrom("NEW")
	CreateFromSnapshot = CreateVolumeFrom("SNAPSHOT")

	AesXtsPlain64_128 = EncryptType("aes-xts-plain64_128")
	AesXtsPlain64_256 = EncryptType("aes-xts-plain64_256")

	InitMigrateAction    = MigrateAction("INIT-MIGRATE")
	ProcessMigrateAction = MigrateAction("MIGRATE")
	ConfirmMigrateAction = MigrateAction("CONFIRM-MIGRATE")
)

type CreateBlockVolumeRequest struct {
	BackupVolumePointID    string                  `json:"backupVolumePointId,omitempty"`
	CreatedFrom            CreateVolumeFrom        `json:"createdFrom,omitempty"`
	EncryptionType         EncryptType             `json:"encryptionType,omitempty"`
	MultiAttach            bool                    `json:"multiAttach,omitempty"`
	Name                   string                  `json:"name"`
	Size                   int64                   `json:"size"`
	VolumeTypeID           string                  `json:"volumeTypeId"`
	Tags                   []VolumeTag             `json:"tags,omitempty"`
	IsPoc                  bool                    `json:"isPoc,omitempty"`
	IsEnableAutoRenew      bool                    `json:"isEnableAutoRenew,omitempty"`
	ConfigureVolumeRestore *ConfigureVolumeRestore `json:"configVolumeRestore,omitempty"`
	Zone                   string                  `json:"zoneId,omitempty"`
	PoolName               string                  `json:"poolName,omitempty"`
}

type DeleteBlockVolumeByIDRequest struct {
	common.BlockVolumeCommon
}

type ResizeBlockVolumeByIDRequest struct {
	NewSize      int    `json:"newSize"`         // NewSize is the new size of the volume, in GB
	VolumeTypeID string `json:"newVolumeTypeId"` // VolumeTypeID is the type of the volume
	common.BlockVolumeCommon
}

type ListBlockVolumesRequest struct {
	Name string
	Page int
	Size int
}

type AttachBlockVolumeRequest struct {
	common.BlockVolumeCommon
	common.ServerCommon
}

type GetBlockVolumeByIDRequest struct {
	common.BlockVolumeCommon
}

type GetUnderBlockVolumeIDRequest struct {
	common.BlockVolumeCommon
}

type MigrateBlockVolumeByIDRequest struct {
	Action         MigrateAction `json:"action"`
	ConfirmMigrate bool
	Tags           []common.Tag `json:"tags"`
	VolumeTypeID   string       `json:"volumeTypeId"`
	Auto           bool

	common.BlockVolumeCommon
}

type (
	MigrateAction    string
	CreateVolumeFrom string
	EncryptType      string

	VolumeTag struct {
		Key   string `json:"key"`
		Value string `json:"value"`
	}

	ConfigureVolumeRestore struct {
		SnapshotVolumePointID string `json:"snapshotVolumePointId"`
		VolumeTypeID          string `json:"volumeTypeId"`
	}
)

func (r *CreateBlockVolumeRequest) ToRequestBody() any {
	return r
}

func (r *CreateBlockVolumeRequest) WithZone(zone string) ICreateBlockVolumeRequest {
	r.Zone = zone
	return r
}

func (r *CreateBlockVolumeRequest) WithPoolName(poolName string) ICreateBlockVolumeRequest {
	r.PoolName = poolName
	return r
}

func (r *CreateBlockVolumeRequest) ToMap() map[string]any {
	return map[string]any{
		"backupVolumePointId": r.BackupVolumePointID,
		"createdFrom":         r.CreatedFrom,
		"encryptionType":      r.EncryptionType,
		"multiAttach":         r.MultiAttach,
		"name":                r.Name,
		"size":                r.Size,
		"volumeTypeId":        r.VolumeTypeID,
		"tags":                r.Tags,
		"isPoc":               r.IsPoc,
		"isEnableAutoRenew":   r.IsEnableAutoRenew,
		"configVolumeRestore": r.ConfigureVolumeRestore,
	}
}

func (r *CreateBlockVolumeRequest) ListParameters() []any {
	return []any{
		"backupVolumePointId", r.BackupVolumePointID,
		"createdFrom", r.CreatedFrom,
		"encryptionType", r.EncryptionType,
		"multiAttach", r.MultiAttach,
		"name", r.Name,
		"size", r.Size,
		"volumeTypeId", r.VolumeTypeID,
		"tags", r.Tags,
		"isPoc", r.IsPoc,
		"isEnableAutoRenew", r.IsEnableAutoRenew,
		"configVolumeRestore", r.ConfigureVolumeRestore,
	}
}

func (r *CreateBlockVolumeRequest) GetVolumeName() string {
	return r.Name
}

func (r *CreateBlockVolumeRequest) GetVolumeType() string {
	return r.VolumeTypeID
}

func (r *CreateBlockVolumeRequest) GetZone() string {
	return r.Zone
}

func (r *CreateBlockVolumeRequest) GetPoolName() string {
	return r.PoolName
}

func (r *CreateBlockVolumeRequest) GetSize() int64 {
	return r.Size
}

func (r *CreateBlockVolumeRequest) WithPoc(isPoc bool) ICreateBlockVolumeRequest {
	r.IsPoc = isPoc
	return r
}

func (r *CreateBlockVolumeRequest) WithAutoRenew(val bool) ICreateBlockVolumeRequest {
	r.IsEnableAutoRenew = val
	return r
}

func (r *CreateBlockVolumeRequest) WithMultiAttach(multiAttach bool) ICreateBlockVolumeRequest {
	r.MultiAttach = multiAttach
	return r
}

func (r *CreateBlockVolumeRequest) WithSize(size int64) ICreateBlockVolumeRequest {
	r.Size = size
	return r
}

func (r *CreateBlockVolumeRequest) WithEncryptionType(et EncryptType) ICreateBlockVolumeRequest {
	r.EncryptionType = et
	return r
}

func (r *CreateBlockVolumeRequest) WithVolumeType(volumeTypeID string) ICreateBlockVolumeRequest {
	r.VolumeTypeID = volumeTypeID
	return r
}

func (r *CreateBlockVolumeRequest) WithVolumeRestoreFromSnapshot(snapshotID, volumeTypeID string) ICreateBlockVolumeRequest {
	r.CreatedFrom = CreateFromSnapshot
	r.ConfigureVolumeRestore = &ConfigureVolumeRestore{
		SnapshotVolumePointID: snapshotID,
		VolumeTypeID:          volumeTypeID,
	}

	return r
}

func (r *CreateBlockVolumeRequest) WithTags(tags ...string) ICreateBlockVolumeRequest {
	if r.Tags == nil {
		r.Tags = make([]VolumeTag, 0)
	}

	if len(tags)%2 != 0 {
		tags = append(tags, "none")
	}

	for i := 0; i < len(tags); i += 2 {
		r.Tags = append(r.Tags, VolumeTag{Key: tags[i], Value: tags[i+1]})
	}

	return r
}

func (r *ListBlockVolumesRequest) ToQuery() (string, error) {
	v := url.Values{}
	v.Set("name", r.Name)
	if r.Page > 0 {
		v.Set("page", strconv.Itoa(r.Page))
	}
	if r.Size > 0 {
		v.Set("size", strconv.Itoa(r.Size))
	}
	return v.Encode(), nil
}

func (r *ListBlockVolumesRequest) GetDefaultQuery() string {
	return fmt.Sprintf("page=%d&size=%d&name=", defaultPageListBlockVolumesRequest, defaultSizeListBlockVolumesRequest)
}

func (r *ListBlockVolumesRequest) ToMap() map[string]any {
	return map[string]any{
		"name": r.Name,
		"page": r.Page,
		"size": r.Size,
	}
}

func (r *ListBlockVolumesRequest) WithName(name string) IListBlockVolumesRequest {
	r.Name = name
	return r
}

func (r *ResizeBlockVolumeByIDRequest) ToRequestBody() any {
	return r
}

func (r *ResizeBlockVolumeByIDRequest) GetSize() int {
	return r.NewSize
}

func (r *ResizeBlockVolumeByIDRequest) GetVolumeTypeID() string {
	return r.VolumeTypeID
}

func (r *MigrateBlockVolumeByIDRequest) ToRequestBody() any {
	return r
}

func (r *MigrateBlockVolumeByIDRequest) WithTags(tags ...string) IMigrateBlockVolumeByIDRequest {
	if r.Tags == nil {
		r.Tags = make([]common.Tag, 0)
	}

	if len(tags)%2 != 0 {
		tags = append(tags, "none")
	}

	for i := 0; i < len(tags); i += 2 {
		r.Tags = append(r.Tags, common.Tag{Key: tags[i], Value: tags[i+1]})
	}

	return r
}

func (r *MigrateBlockVolumeByIDRequest) WithAction(action MigrateAction) IMigrateBlockVolumeByIDRequest {
	switch action {
	case InitMigrateAction, ProcessMigrateAction, ConfirmMigrateAction:
		r.Action = action
	default:
		r.Action = InitMigrateAction
	}

	return r
}

func (r *MigrateBlockVolumeByIDRequest) WithConfirm(confirm bool) IMigrateBlockVolumeByIDRequest {
	r.ConfirmMigrate = confirm
	return r
}

func (r *MigrateBlockVolumeByIDRequest) IsConfirm() bool {
	return r.ConfirmMigrate
}
