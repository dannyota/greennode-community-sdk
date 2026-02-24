package v2

import (
	"fmt"
	"net/url"
	"strconv"
)

type ListInterconnectsRequest struct {
	Name string
	Page int
	Size int
}

func (r *ListInterconnectsRequest) ToListQuery() (string, error) {
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

func (r *ListInterconnectsRequest) getDefaultQuery() string {
	return fmt.Sprintf("name=&page=%d&size=%d", defaultPageListInterconnects, defaultSizeListInterconnects)
}

func NewListInterconnectsRequest(page, size int) *ListInterconnectsRequest {
	return &ListInterconnectsRequest{
		Page: page,
		Size: size,
	}
}
