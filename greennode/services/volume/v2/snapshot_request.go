package v2

import (
	lfmt "fmt"
	lurl "net/url"
	lstrconv "strconv"

	lscommon "github.com/dannyota/greennode-community-sdk/v2/greennode/services/common"
)

func NewListSnapshotsByBlockVolumeIdRequest(ppage, psize int, pblockVolumeId string) IListSnapshotsByBlockVolumeIdRequest {
	opt := new(ListSnapshotsByBlockVolumeIdRequest)
	opt.BlockVolumeId = pblockVolumeId
	opt.Page = ppage
	opt.Size = psize

	return opt
}

func NewCreateSnapshotByBlockVolumeIdRequest(pname, pblockVolumeId string) ICreateSnapshotByBlockVolumeIdRequest {
	opt := new(CreateSnapshotByBlockVolumeIdRequest)
	opt.Name = pname
	opt.BlockVolumeId = pblockVolumeId

	return opt
}

func NewDeleteSnapshotByIdRequest(psnapshotId string) IDeleteSnapshotByIdRequest {
	opt := new(DeleteSnapshotByIdRequest)
	opt.BlockVolumeId = "undefined"
	opt.SnapshotId = psnapshotId

	return opt
}

type ListSnapshotsByBlockVolumeIdRequest struct {
	Page int
	Size int

	lscommon.BlockVolumeCommon
}

type CreateSnapshotByBlockVolumeIdRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Permanently bool   `json:"isPermanently"`
	RetainedDay uint64 `json:"retainedDay"`

	lscommon.BlockVolumeCommon
}

type DeleteSnapshotByIdRequest struct {
	lscommon.BlockVolumeCommon
	lscommon.SnapshotCommon
}

func (s *ListSnapshotsByBlockVolumeIdRequest) GetDefaultQuery() string {
	return lfmt.Sprintf("page=%d&size=%d", defaultPageListSnapshotsByBlockVolumeIdRequest, defaultSizeListSnapshotsByBlockVolumeIdRequest)
}

func (s *ListSnapshotsByBlockVolumeIdRequest) ToQuery() (string, error) {
	v := lurl.Values{}
	if s.Page > 0 {
		v.Set("page", lstrconv.Itoa(s.Page))
	}
	if s.Size > 0 {
		v.Set("size", lstrconv.Itoa(s.Size))
	}
	return v.Encode(), nil
}

func (s *CreateSnapshotByBlockVolumeIdRequest) ToRequestBody() interface{} {
	return s
}

func (s *CreateSnapshotByBlockVolumeIdRequest) WithDescription(pdesc string) ICreateSnapshotByBlockVolumeIdRequest {
	s.Description = pdesc
	return s
}

func (s *CreateSnapshotByBlockVolumeIdRequest) WithPermanently(pval bool) ICreateSnapshotByBlockVolumeIdRequest {
	s.Permanently = pval
	return s
}

func (s *CreateSnapshotByBlockVolumeIdRequest) WithRetainedDay(pval uint64) ICreateSnapshotByBlockVolumeIdRequest {
	s.RetainedDay = pval
	return s
}
