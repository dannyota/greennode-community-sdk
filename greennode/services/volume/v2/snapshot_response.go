package v2

type ListSnapshotsByBlockVolumeIDResponse struct {
	Items      []snapshotResp `json:"items"`
	Page       int            `json:"page"`
	PageSize   int            `json:"pageSize"`
	TotalPages int            `json:"totalPages"`
	TotalItems int            `json:"totalItems"`
}

type CreateSnapshotByBlockVolumeIDResponse struct {
	snapshotResp
}

type (
	snapshotResp struct {
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

func (r *ListSnapshotsByBlockVolumeIDResponse) ToEntityListSnapshots() *ListSnapshots {
	sl := new(ListSnapshots)

	for _, item := range r.Items {
		sl.Items = append(sl.Items, item.toEntitySnapshot())
	}

	sl.TotalPages = r.TotalPages
	sl.Page = r.Page
	sl.PageSize = r.PageSize
	sl.TotalItems = r.TotalItems

	return sl
}

func (s *snapshotResp) toEntitySnapshot() *Snapshot {
	return &Snapshot{
		ID:        s.ID,
		Name:      s.Name,
		CreatedAt: s.CreatedAt,
		VolumeID:  s.VolumeID,
		Status:    s.Status,
		Size:      s.Size,
	}
}

func (r *CreateSnapshotByBlockVolumeIDResponse) ToEntitySnapshot() *Snapshot {
	return r.toEntitySnapshot()
}
