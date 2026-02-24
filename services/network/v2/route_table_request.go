package v2

import (
	"fmt"
	"net/url"
	"strconv"
)

type ListRouteTablesRequest struct {
	Name string
	Page int
	Size int
}

func (r *ListRouteTablesRequest) ToListQuery() (string, error) {
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

func (r *ListRouteTablesRequest) getDefaultQuery() string {
	return fmt.Sprintf("name=&page=%d&size=%d", defaultPageListRouteTables, defaultSizeListRouteTables)
}

func NewListRouteTablesRequest(page, size int) *ListRouteTablesRequest {
	return &ListRouteTablesRequest{
		Page: page,
		Size: size,
	}
}
