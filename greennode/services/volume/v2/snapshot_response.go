package v2

import "github.com/dannyota/greennode-community-sdk/v2/greennode/entity"

type ListSnapshotsByBlockVolumeIDResponse struct {
	Items      []Snapshot `json:"items"`
	Page       int        `json:"page"`
	PageSize   int        `json:"pageSize"`
	TotalPages int        `json:"totalPages"`
	TotalItems int        `json:"totalItems"`
}

type CreateSnapshotByBlockVolumeIDResponse struct {
	Snapshot
}

type (
	Snapshot struct {
		BackendID           string         `json:"backendId"`
		BackendSnapshotID   *string        `json:"backendSnapshotId"`
		BackendSnapshotName *string        `json:"backendSnapshotName"`
		CreatedAt           string         `json:"createdAt"`
		DeletedAt           *string        `json:"deletedAt"`
		ID                  string         `json:"id"`
		ParentID            *string        `json:"parentId"`
		ParentType          string         `json:"parentType"`
		ProjectID           string         `json:"projectId"`
		Size                int64          `json:"size"`
		SnapshotConfig      Config         `json:"snapshotConfig"`
		SnapshotTime        int64          `json:"snapshotTime"`
		SnapshotVolumeID    string         `json:"snapshotVolumeId"`
		Status              string         `json:"status"`
		UpdatedAt           string         `json:"updatedAt"`
		UserID              int64          `json:"userId"`
		VolumeID            string         `json:"volumeId"`
		VolumeSnapshot      VolumeSnapshot `json:"volumeSnapshot"`
		Name                string         `json:"name"`
		Description         string         `json:"description"`
	}

	Config struct {
		IsPermanently bool  `json:"isPermanently"`
		RetainedDays  int64 `json:"retainedDays"`
	}

	VolumeSnapshot struct {
		BackendPool      string  `json:"backendPool"`
		BackendPrefix    string  `json:"backendPrefix"`
		BackendStatus    string  `json:"backendStatus"`
		BackendUuid      string  `json:"backendUuid"`
		BootIndex        int     `json:"bootIndex"`
		Bootable         bool    `json:"bootable"`
		CreatedAt        string  `json:"createdAt"`
		ID               string  `json:"id"`
		MultiAttach      bool    `json:"multiAttach"`
		Name             string  `json:"name"`
		Product          string  `json:"product"`
		ProjectID        string  `json:"projectId"`
		Size             int64   `json:"size"`
		Status           string  `json:"status"`
		Type             string  `json:"type"`
		UpdatedAt        string  `json:"updatedAt"`
		Uuid             string  `json:"uuid"`
		VolumeID         int64   `json:"volumeId"`
		VolumeTypeID     string  `json:"volumeTypeId"`
		VolumeTypeZoneID string  `json:"volumeTypeZoneId"`
		EncryptionType   *string `json:"encryptionType"`
	}
)

func (s *ListSnapshotsByBlockVolumeIDResponse) ToEntityListSnapshots() *entity.ListSnapshots {
	sl := new(entity.ListSnapshots)

	for _, item := range s.Items {
		sl.Items = append(sl.Items, item.toEntitySnapshot())
	}

	sl.TotalPages = s.TotalPages
	sl.Page = s.Page
	sl.PageSize = s.PageSize
	sl.TotalItems = s.TotalItems

	return sl
}

func (s *Snapshot) toEntitySnapshot() *entity.Snapshot {
	return &entity.Snapshot{
		ID:        s.ID,
		Name:      s.Name,
		CreatedAt: s.CreatedAt,
		VolumeID:  s.VolumeID,
		Status:    s.Status,
		Size:      s.Size,
	}
}

func (s *CreateSnapshotByBlockVolumeIDResponse) ToEntitySnapshot() *entity.Snapshot {
	return s.toEntitySnapshot()
}
