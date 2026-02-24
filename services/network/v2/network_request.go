package v2

import (
	"fmt"
	"net/url"
	"strconv"
)

func NewGetNetworkByIDRequest(networkID string) *GetNetworkByIDRequest {
	return &GetNetworkByIDRequest{
		NetworkID: networkID,
	}
}

type GetNetworkByIDRequest struct {
	NetworkID string
}

type ListNetworksRequest struct {
	Name string
	Page int
	Size int
}

func (r *ListNetworksRequest) ToListQuery() (string, error) {
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

func (r *ListNetworksRequest) getDefaultQuery() string {
	return fmt.Sprintf("name=&page=%d&size=%d", defaultPageListNetworks, defaultSizeListNetworks)
}

func NewListNetworksRequest(page, size int) *ListNetworksRequest {
	return &ListNetworksRequest{
		Page: page,
		Size: size,
	}
}
