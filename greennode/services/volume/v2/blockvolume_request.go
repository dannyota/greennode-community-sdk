package v2

import (
	"fmt"
	"net/url"

	"strconv"

	"github.com/dannyota/greennode-community-sdk/v2/greennode/services/common"
)

func NewCreateBlockVolumeRequest(volumeName, volumeType string, size int64) *CreateBlockVolumeRequest {
	return &CreateBlockVolumeRequest{
		VolumeTypeID: volumeType,
		CreatedFrom:  CreateFromNew,
		Name:         volumeName,
		Size:         size,
	}
}

func NewDeleteBlockVolumeByIDRequest(volumeID string) *DeleteBlockVolumeByIDRequest {
	return &DeleteBlockVolumeByIDRequest{
		BlockVolumeID: volumeID,
	}
}

func NewListBlockVolumesRequest(page, size int) *ListBlockVolumesRequest {
	return &ListBlockVolumesRequest{
		Page: page,
		Size: size,
	}
}

func NewGetBlockVolumeByIDRequest(volumeID string) *GetBlockVolumeByIDRequest {
	return &GetBlockVolumeByIDRequest{
		BlockVolumeID: volumeID,
	}
}

func NewResizeBlockVolumeByIDRequest(volumeID, volumeType string, size int) *ResizeBlockVolumeByIDRequest {
	return &ResizeBlockVolumeByIDRequest{
		NewSize:       size,
		VolumeTypeID:  volumeType,
		BlockVolumeID: volumeID,
	}
}

func NewGetUnderVolumeIDRequest(volumeID string) *GetUnderBlockVolumeIDRequest {
	return &GetUnderBlockVolumeIDRequest{
		BlockVolumeID: volumeID,
	}
}

func NewMigrateBlockVolumeByIDRequest(volumeID, volumeType string) *MigrateBlockVolumeByIDRequest {
	return &MigrateBlockVolumeByIDRequest{
		Action:        InitMigrateAction,
		VolumeTypeID:  volumeType,
		BlockVolumeID: volumeID,
	}
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
	BlockVolumeID string
}

type ResizeBlockVolumeByIDRequest struct {
	NewSize       int    `json:"newSize"`         // NewSize is the new size of the volume, in GB
	VolumeTypeID  string `json:"newVolumeTypeId"` // VolumeTypeID is the type of the volume
	BlockVolumeID string
}

type ListBlockVolumesRequest struct {
	Name string
	Page int
	Size int
}

type AttachBlockVolumeRequest struct {
	BlockVolumeID string
	ServerID      string
}

type GetBlockVolumeByIDRequest struct {
	BlockVolumeID string
}

type GetUnderBlockVolumeIDRequest struct {
	BlockVolumeID string
}

type MigrateBlockVolumeByIDRequest struct {
	Action         MigrateAction `json:"action"`
	ConfirmMigrate bool
	Tags           []common.Tag `json:"tags"`
	VolumeTypeID   string       `json:"volumeTypeId"`
	Auto           bool
	BlockVolumeID  string
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

func (r *CreateBlockVolumeRequest) WithZone(zone string) *CreateBlockVolumeRequest {
	r.Zone = zone
	return r
}

func (r *CreateBlockVolumeRequest) WithPoolName(poolName string) *CreateBlockVolumeRequest {
	r.PoolName = poolName
	return r
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

func (r *CreateBlockVolumeRequest) WithPoc(isPoc bool) *CreateBlockVolumeRequest {
	r.IsPoc = isPoc
	return r
}

func (r *CreateBlockVolumeRequest) WithAutoRenew(val bool) *CreateBlockVolumeRequest {
	r.IsEnableAutoRenew = val
	return r
}

func (r *CreateBlockVolumeRequest) WithMultiAttach(multiAttach bool) *CreateBlockVolumeRequest {
	r.MultiAttach = multiAttach
	return r
}

func (r *CreateBlockVolumeRequest) WithSize(size int64) *CreateBlockVolumeRequest {
	r.Size = size
	return r
}

func (r *CreateBlockVolumeRequest) WithEncryptionType(et EncryptType) *CreateBlockVolumeRequest {
	r.EncryptionType = et
	return r
}

func (r *CreateBlockVolumeRequest) WithVolumeType(volumeTypeID string) *CreateBlockVolumeRequest {
	r.VolumeTypeID = volumeTypeID
	return r
}

func (r *CreateBlockVolumeRequest) WithVolumeRestoreFromSnapshot(snapshotID, volumeTypeID string) *CreateBlockVolumeRequest {
	r.CreatedFrom = CreateFromSnapshot
	r.ConfigureVolumeRestore = &ConfigureVolumeRestore{
		SnapshotVolumePointID: snapshotID,
		VolumeTypeID:          volumeTypeID,
	}

	return r
}

func (r *CreateBlockVolumeRequest) WithTags(tags ...string) *CreateBlockVolumeRequest {
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

func (r *ListBlockVolumesRequest) WithName(name string) *ListBlockVolumesRequest {
	r.Name = name
	return r
}

func (r *ResizeBlockVolumeByIDRequest) GetSize() int {
	return r.NewSize
}

func (r *ResizeBlockVolumeByIDRequest) GetVolumeTypeID() string {
	return r.VolumeTypeID
}

func (r *MigrateBlockVolumeByIDRequest) WithTags(tags ...string) *MigrateBlockVolumeByIDRequest {
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

func (r *MigrateBlockVolumeByIDRequest) WithAction(action MigrateAction) *MigrateBlockVolumeByIDRequest {
	switch action {
	case InitMigrateAction, ProcessMigrateAction, ConfirmMigrateAction:
		r.Action = action
	default:
		r.Action = InitMigrateAction
	}

	return r
}

func (r *MigrateBlockVolumeByIDRequest) WithConfirm(confirm bool) *MigrateBlockVolumeByIDRequest {
	r.ConfirmMigrate = confirm
	return r
}

func (r *MigrateBlockVolumeByIDRequest) IsConfirm() bool {
	return r.ConfirmMigrate
}
