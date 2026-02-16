package v2

import (
	"fmt"
	"net/url"
	"strconv"
)

func NewListSnapshotsByBlockVolumeIDRequest(page, size int, blockVolumeID string) *ListSnapshotsByBlockVolumeIDRequest {
	return &ListSnapshotsByBlockVolumeIDRequest{
		Page:          page,
		Size:          size,
		BlockVolumeID: blockVolumeID,
	}
}

func NewCreateSnapshotByBlockVolumeIDRequest(name, blockVolumeID string) *CreateSnapshotByBlockVolumeIDRequest {
	return &CreateSnapshotByBlockVolumeIDRequest{
		Name:          name,
		BlockVolumeID: blockVolumeID,
	}
}

func NewDeleteSnapshotByIDRequest(snapshotID string) *DeleteSnapshotByIDRequest {
	return &DeleteSnapshotByIDRequest{
		BlockVolumeID: "undefined",
		SnapshotID:    snapshotID,
	}
}

type ListSnapshotsByBlockVolumeIDRequest struct {
	Page          int
	Size          int
	BlockVolumeID string
}

type CreateSnapshotByBlockVolumeIDRequest struct {
	Name          string `json:"name"`
	Description   string `json:"description"`
	Permanently   bool   `json:"isPermanently"`
	RetainedDay   uint64 `json:"retainedDay"`
	BlockVolumeID string
}

type DeleteSnapshotByIDRequest struct {
	BlockVolumeID string
	SnapshotID    string
}

func (r *ListSnapshotsByBlockVolumeIDRequest) GetDefaultQuery() string {
	return fmt.Sprintf("page=%d&size=%d", defaultPageListSnapshotsByBlockVolumeIDRequest, defaultSizeListSnapshotsByBlockVolumeIDRequest)
}

func (r *ListSnapshotsByBlockVolumeIDRequest) ToQuery() (string, error) {
	v := url.Values{}
	if r.Page > 0 {
		v.Set("page", strconv.Itoa(r.Page))
	}
	if r.Size > 0 {
		v.Set("size", strconv.Itoa(r.Size))
	}
	return v.Encode(), nil
}

