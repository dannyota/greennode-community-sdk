package v2

import (
	"fmt"
	"net/url"
	"strconv"
)

type ListPeeringsRequest struct {
	Name string
	Page int
	Size int
}

func (r *ListPeeringsRequest) ToListQuery() (string, error) {
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

func (r *ListPeeringsRequest) getDefaultQuery() string {
	return fmt.Sprintf("name=&page=%d&size=%d", defaultPageListPeerings, defaultSizeListPeerings)
}

func NewListPeeringsRequest(page, size int) *ListPeeringsRequest {
	return &ListPeeringsRequest{
		Page: page,
		Size: size,
	}
}
