package v2

import (
	"fmt"
	"net/url"
	"strconv"
)

type ListSecgroupRequest struct {
	Name string
	Page int
	Size int
}

func (r *ListSecgroupRequest) ToListQuery() (string, error) {
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

func (r *ListSecgroupRequest) getDefaultQuery() string {
	return fmt.Sprintf("name=&page=%d&size=%d", defaultPageListSecgroups, defaultSizeListSecgroups)
}

func NewListSecgroupRequest(page, size int) *ListSecgroupRequest {
	return &ListSecgroupRequest{
		Page: page,
		Size: size,
	}
}

type DeleteSecgroupByIDRequest struct { //__________________________________________________________________________________
	SecgroupID string
}

type CreateSecgroupRequest struct { //__________________________________________________________________________________
	Name        string `json:"name"`
	Description string `json:"description"`
}

type GetSecgroupByIDRequest struct { //_________________________________________________________________________________
	SecgroupID string
}

func NewCreateSecgroupRequest(name, description string) *CreateSecgroupRequest {
	return &CreateSecgroupRequest{
		Name:        name,
		Description: description,
	}
}

func NewDeleteSecgroupByIDRequest(secgroupID string) *DeleteSecgroupByIDRequest {
	return &DeleteSecgroupByIDRequest{
		SecgroupID: secgroupID,
	}
}

func NewGetSecgroupByIDRequest(secgroupID string) *GetSecgroupByIDRequest {
	return &GetSecgroupByIDRequest{
		SecgroupID: secgroupID,
	}
}
