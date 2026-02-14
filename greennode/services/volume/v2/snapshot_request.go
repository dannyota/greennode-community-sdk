package v2

import (
	"fmt"
	"net/url"
	"strconv"

	"github.com/dannyota/greennode-community-sdk/v2/greennode/services/common"
)

func NewListSnapshotsByBlockVolumeIdRequest(page, size int, blockVolumeId string) IListSnapshotsByBlockVolumeIdRequest {
	opt := new(ListSnapshotsByBlockVolumeIdRequest)
	opt.BlockVolumeId = blockVolumeId
	opt.Page = page
	opt.Size = size

	return opt
}

func NewCreateSnapshotByBlockVolumeIdRequest(name, blockVolumeId string) ICreateSnapshotByBlockVolumeIdRequest {
	opt := new(CreateSnapshotByBlockVolumeIdRequest)
	opt.Name = name
	opt.BlockVolumeId = blockVolumeId

	return opt
}

func NewDeleteSnapshotByIdRequest(snapshotId string) IDeleteSnapshotByIdRequest {
	opt := new(DeleteSnapshotByIdRequest)
	opt.BlockVolumeId = "undefined"
	opt.SnapshotId = snapshotId

	return opt
}

type ListSnapshotsByBlockVolumeIdRequest struct {
	Page int
	Size int

	common.BlockVolumeCommon
}

type CreateSnapshotByBlockVolumeIdRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Permanently bool   `json:"isPermanently"`
	RetainedDay uint64 `json:"retainedDay"`

	common.BlockVolumeCommon
}

type DeleteSnapshotByIdRequest struct {
	common.BlockVolumeCommon
	common.SnapshotCommon
}

func (s *ListSnapshotsByBlockVolumeIdRequest) GetDefaultQuery() string {
	return fmt.Sprintf("page=%d&size=%d", defaultPageListSnapshotsByBlockVolumeIdRequest, defaultSizeListSnapshotsByBlockVolumeIdRequest)
}

func (s *ListSnapshotsByBlockVolumeIdRequest) ToQuery() (string, error) {
	v := url.Values{}
	if s.Page > 0 {
		v.Set("page", strconv.Itoa(s.Page))
	}
	if s.Size > 0 {
		v.Set("size", strconv.Itoa(s.Size))
	}
	return v.Encode(), nil
}

func (s *CreateSnapshotByBlockVolumeIdRequest) ToRequestBody() interface{} {
	return s
}

func (s *CreateSnapshotByBlockVolumeIdRequest) WithDescription(desc string) ICreateSnapshotByBlockVolumeIdRequest {
	s.Description = desc
	return s
}

func (s *CreateSnapshotByBlockVolumeIdRequest) WithPermanently(val bool) ICreateSnapshotByBlockVolumeIdRequest {
	s.Permanently = val
	return s
}

func (s *CreateSnapshotByBlockVolumeIdRequest) WithRetainedDay(val uint64) ICreateSnapshotByBlockVolumeIdRequest {
	s.RetainedDay = val
	return s
}
