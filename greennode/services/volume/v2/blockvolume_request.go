package v2

import (
	"fmt"
	"net/url"
	"strconv"

	"github.com/dannyota/greennode-community-sdk/v2/greennode/services/common"
)

func NewCreateBlockVolumeRequest(volumeName, volumeType string, size int64) ICreateBlockVolumeRequest {
	opt := new(CreateBlockVolumeRequest)
	opt.VolumeTypeId = volumeType
	opt.CreatedFrom = CreateFromNew
	opt.Name = volumeName
	opt.Size = size

	return opt
}

func NewDeleteBlockVolumeByIdRequest(volumeId string) IDeleteBlockVolumeByIdRequest {
	opt := new(DeleteBlockVolumeByIdRequest)
	opt.BlockVolumeId = volumeId
	return opt
}

func NewListBlockVolumesRequest(page, size int) IListBlockVolumesRequest {
	opt := new(ListBlockVolumesRequest)
	opt.Page = page
	opt.Size = size
	return opt
}

func NewGetBlockVolumeByIdRequest(volumeId string) IGetBlockVolumeByIdRequest {
	opt := new(GetBlockVolumeByIdRequest)
	opt.BlockVolumeId = volumeId
	return opt
}

func NewResizeBlockVolumeByIdRequest(volumeId, volumeType string, size int) IResizeBlockVolumeByIdRequest {
	opt := new(ResizeBlockVolumeByIdRequest)
	opt.BlockVolumeId = volumeId
	opt.NewSize = size
	opt.VolumeTypeID = volumeType
	return opt
}

func NewGetUnderVolumeIdRequest(volumeId string) IGetUnderBlockVolumeIdRequest {
	opt := new(GetUnderBlockVolumeIdRequest)
	opt.BlockVolumeId = volumeId
	return opt
}

func NewMigrateBlockVolumeByIdRequest(volumeId, volumeType string) IMigrateBlockVolumeByIdRequest {
	opt := new(MigrateBlockVolumeByIdRequest)
	opt.BlockVolumeId = volumeId
	opt.VolumeTypeId = volumeType
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
	BackupVolumePointId    string                  `json:"backupVolumePointId,omitempty"`
	CreatedFrom            CreateVolumeFrom        `json:"createdFrom,omitempty"`
	EncryptionType         EncryptType             `json:"encryptionType,omitempty"`
	MultiAttach            bool                    `json:"multiAttach,omitempty"`
	Name                   string                  `json:"name"`
	Size                   int64                   `json:"size"`
	VolumeTypeId           string                  `json:"volumeTypeId"`
	Tags                   []VolumeTag             `json:"tags,omitempty"`
	IsPoc                  bool                    `json:"isPoc,omitempty"`
	IsEnableAutoRenew      bool                    `json:"isEnableAutoRenew,omitempty"`
	ConfigureVolumeRestore *ConfigureVolumeRestore `json:"configVolumeRestore,omitempty"`
	Zone                   string                  `json:"zoneId,omitempty"`
	PoolName               string                  `json:"poolName,omitempty"`
}

type DeleteBlockVolumeByIdRequest struct {
	common.BlockVolumeCommon
}

type ResizeBlockVolumeByIdRequest struct {
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

type GetBlockVolumeByIdRequest struct {
	common.BlockVolumeCommon
}

type GetUnderBlockVolumeIdRequest struct {
	common.BlockVolumeCommon
}

type MigrateBlockVolumeByIdRequest struct {
	Action         MigrateAction `json:"action"`
	ConfirmMigrate bool
	Tags           []common.Tag `json:"tags"`
	VolumeTypeId   string       `json:"volumeTypeId"`
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
		SnapshotVolumePointId string `json:"snapshotVolumePointId"`
		VolumeTypeId          string `json:"volumeTypeId"`
	}
)

func (s *CreateBlockVolumeRequest) ToRequestBody() interface{} {
	return s
}

func (s *CreateBlockVolumeRequest) WithZone(zone string) ICreateBlockVolumeRequest {
	s.Zone = zone
	return s
}

func (s *CreateBlockVolumeRequest) WithPoolName(poolName string) ICreateBlockVolumeRequest {
	s.PoolName = poolName
	return s
}

func (s *CreateBlockVolumeRequest) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"backupVolumePointId": s.BackupVolumePointId,
		"createdFrom":         s.CreatedFrom,
		"encryptionType":      s.EncryptionType,
		"multiAttach":         s.MultiAttach,
		"name":                s.Name,
		"size":                s.Size,
		"volumeTypeId":        s.VolumeTypeId,
		"tags":                s.Tags,
		"isPoc":               s.IsPoc,
		"isEnableAutoRenew":   s.IsEnableAutoRenew,
		"configVolumeRestore": s.ConfigureVolumeRestore,
	}
}

func (s *CreateBlockVolumeRequest) GetListParameters() []interface{} {
	return []interface{}{
		"backupVolumePointId", s.BackupVolumePointId,
		"createdFrom", s.CreatedFrom,
		"encryptionType", s.EncryptionType,
		"multiAttach", s.MultiAttach,
		"name", s.Name,
		"size", s.Size,
		"volumeTypeId", s.VolumeTypeId,
		"tags", s.Tags,
		"isPoc", s.IsPoc,
		"isEnableAutoRenew", s.IsEnableAutoRenew,
		"configVolumeRestore", s.ConfigureVolumeRestore,
	}
}

func (s *CreateBlockVolumeRequest) GetVolumeName() string {
	return s.Name
}

func (s *CreateBlockVolumeRequest) GetVolumeType() string {
	return s.VolumeTypeId
}

func (s *CreateBlockVolumeRequest) GetZone() string {
	return s.Zone
}

func (s *CreateBlockVolumeRequest) GetPoolName() string {
	return s.PoolName
}

func (s *CreateBlockVolumeRequest) GetSize() int64 {
	return s.Size
}

func (s *CreateBlockVolumeRequest) WithPoc(isPoc bool) ICreateBlockVolumeRequest {
	s.IsPoc = isPoc
	return s
}

func (s *CreateBlockVolumeRequest) WithAutoRenew(val bool) ICreateBlockVolumeRequest {
	s.IsEnableAutoRenew = val
	return s
}

func (s *CreateBlockVolumeRequest) WithMultiAttach(multiAttach bool) ICreateBlockVolumeRequest {
	s.MultiAttach = multiAttach
	return s
}

func (s *CreateBlockVolumeRequest) WithSize(size int64) ICreateBlockVolumeRequest {
	s.Size = size
	return s
}

func (s *CreateBlockVolumeRequest) WithEncryptionType(et EncryptType) ICreateBlockVolumeRequest {
	s.EncryptionType = et
	return s
}

func (s *CreateBlockVolumeRequest) WithVolumeType(volumeTypeId string) ICreateBlockVolumeRequest {
	s.VolumeTypeId = volumeTypeId
	return s
}

func (s *CreateBlockVolumeRequest) WithVolumeRestoreFromSnapshot(snapshotID, volumeTypeID string) ICreateBlockVolumeRequest {
	s.CreatedFrom = CreateFromSnapshot
	s.ConfigureVolumeRestore = &ConfigureVolumeRestore{
		SnapshotVolumePointId: snapshotID,
		VolumeTypeId:          volumeTypeID,
	}

	return s
}

func (s *CreateBlockVolumeRequest) WithTags(tags ...string) ICreateBlockVolumeRequest {
	if s.Tags == nil {
		s.Tags = make([]VolumeTag, 0)
	}

	if len(tags)%2 != 0 {
		tags = append(tags, "none")
	}

	for i := 0; i < len(tags); i += 2 {
		s.Tags = append(s.Tags, VolumeTag{Key: tags[i], Value: tags[i+1]})
	}

	return s
}

func (s *ListBlockVolumesRequest) ToQuery() (string, error) {
	v := url.Values{}
	v.Set("name", s.Name)
	if s.Page > 0 {
		v.Set("page", strconv.Itoa(s.Page))
	}
	if s.Size > 0 {
		v.Set("size", strconv.Itoa(s.Size))
	}
	return v.Encode(), nil
}

func (s *ListBlockVolumesRequest) GetDefaultQuery() string {
	return fmt.Sprintf("page=%d&size=%d&name=", defaultPageListBlockVolumesRequest, defaultSizeListBlockVolumesRequest)
}

func (s *ListBlockVolumesRequest) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"name": s.Name,
		"page": s.Page,
		"size": s.Size,
	}
}

func (s *ListBlockVolumesRequest) WithName(name string) IListBlockVolumesRequest {
	s.Name = name
	return s
}

func (s *ResizeBlockVolumeByIdRequest) ToRequestBody() interface{} {
	return s
}

func (s *ResizeBlockVolumeByIdRequest) GetSize() int {
	return s.NewSize
}

func (s *ResizeBlockVolumeByIdRequest) GetVolumeTypeId() string {
	return s.VolumeTypeID
}

func (s *MigrateBlockVolumeByIdRequest) ToRequestBody() interface{} {
	return s
}

func (s *MigrateBlockVolumeByIdRequest) WithTags(tags ...string) IMigrateBlockVolumeByIdRequest {
	if s.Tags == nil {
		s.Tags = make([]common.Tag, 0)
	}

	if len(tags)%2 != 0 {
		tags = append(tags, "none")
	}

	for i := 0; i < len(tags); i += 2 {
		s.Tags = append(s.Tags, common.Tag{Key: tags[i], Value: tags[i+1]})
	}

	return s
}

func (s *MigrateBlockVolumeByIdRequest) WithAction(action MigrateAction) IMigrateBlockVolumeByIdRequest {
	switch action {
	case InitMigrateAction, ProcessMigrateAction, ConfirmMigrateAction:
		s.Action = action
	default:
		s.Action = InitMigrateAction
	}

	return s
}

func (s *MigrateBlockVolumeByIdRequest) WithConfirm(confirm bool) IMigrateBlockVolumeByIdRequest {
	s.ConfirmMigrate = confirm
	return s
}

func (s *MigrateBlockVolumeByIdRequest) IsConfirm() bool {
	return s.ConfirmMigrate
}
