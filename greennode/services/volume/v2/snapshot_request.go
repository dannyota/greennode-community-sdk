package v2

import (
	"fmt"
	"net/url"
	"strconv"

	"github.com/dannyota/greennode-community-sdk/v2/greennode/services/common"
)

func NewListSnapshotsByBlockVolumeIDRequest(page, size int, blockVolumeID string) IListSnapshotsByBlockVolumeIDRequest {
	opt := new(ListSnapshotsByBlockVolumeIDRequest)
	opt.BlockVolumeID = blockVolumeID
	opt.Page = page
	opt.Size = size

	return opt
}

func NewCreateSnapshotByBlockVolumeIDRequest(name, blockVolumeID string) ICreateSnapshotByBlockVolumeIDRequest {
	opt := new(CreateSnapshotByBlockVolumeIDRequest)
	opt.Name = name
	opt.BlockVolumeID = blockVolumeID

	return opt
}

func NewDeleteSnapshotByIDRequest(snapshotID string) IDeleteSnapshotByIDRequest {
	opt := new(DeleteSnapshotByIDRequest)
	opt.BlockVolumeID = "undefined"
	opt.SnapshotID = snapshotID

	return opt
}

type ListSnapshotsByBlockVolumeIDRequest struct {
	Page int
	Size int

	common.BlockVolumeCommon
}

type CreateSnapshotByBlockVolumeIDRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Permanently bool   `json:"isPermanently"`
	RetainedDay uint64 `json:"retainedDay"`

	common.BlockVolumeCommon
}

type DeleteSnapshotByIDRequest struct {
	common.BlockVolumeCommon
	common.SnapshotCommon
}

func (s *ListSnapshotsByBlockVolumeIDRequest) GetDefaultQuery() string {
	return fmt.Sprintf("page=%d&size=%d", defaultPageListSnapshotsByBlockVolumeIDRequest, defaultSizeListSnapshotsByBlockVolumeIDRequest)
}

func (s *ListSnapshotsByBlockVolumeIDRequest) ToQuery() (string, error) {
	v := url.Values{}
	if s.Page > 0 {
		v.Set("page", strconv.Itoa(s.Page))
	}
	if s.Size > 0 {
		v.Set("size", strconv.Itoa(s.Size))
	}
	return v.Encode(), nil
}

func (s *CreateSnapshotByBlockVolumeIDRequest) ToRequestBody() interface{} {
	return s
}

func (s *CreateSnapshotByBlockVolumeIDRequest) WithDescription(desc string) ICreateSnapshotByBlockVolumeIDRequest {
	s.Description = desc
	return s
}

func (s *CreateSnapshotByBlockVolumeIDRequest) WithPermanently(val bool) ICreateSnapshotByBlockVolumeIDRequest {
	s.Permanently = val
	return s
}

func (s *CreateSnapshotByBlockVolumeIDRequest) WithRetainedDay(val uint64) ICreateSnapshotByBlockVolumeIDRequest {
	s.RetainedDay = val
	return s
}
