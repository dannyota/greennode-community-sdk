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

// NewVolumeTags creates a []VolumeTag from variadic key-value string pairs.
// If an odd number of strings is provided, the last value defaults to "none".
func NewVolumeTags(kvPairs ...string) []VolumeTag {
	if len(kvPairs)%2 != 0 {
		kvPairs = append(kvPairs, "none")
	}
	tags := make([]VolumeTag, 0, len(kvPairs)/2)
	for i := 0; i < len(kvPairs); i += 2 {
		tags = append(tags, VolumeTag{Key: kvPairs[i], Value: kvPairs[i+1]})
	}
	return tags
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

