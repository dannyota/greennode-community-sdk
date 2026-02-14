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

func (r *CreateSnapshotByBlockVolumeIDRequest) ToRequestBody() interface{} {
	return r
}

func (r *CreateSnapshotByBlockVolumeIDRequest) WithDescription(desc string) ICreateSnapshotByBlockVolumeIDRequest {
	r.Description = desc
	return r
}

func (r *CreateSnapshotByBlockVolumeIDRequest) WithPermanently(val bool) ICreateSnapshotByBlockVolumeIDRequest {
	r.Permanently = val
	return r
}

func (r *CreateSnapshotByBlockVolumeIDRequest) WithRetainedDay(val uint64) ICreateSnapshotByBlockVolumeIDRequest {
	r.RetainedDay = val
	return r
}
