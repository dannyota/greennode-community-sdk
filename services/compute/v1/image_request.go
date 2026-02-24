package v1

import "net/url"

type ListOSImagesRequest struct {
	ZoneID string
}

func NewListOSImagesRequest(zoneID string) *ListOSImagesRequest {
	return &ListOSImagesRequest{ZoneID: zoneID}
}

func (r *ListOSImagesRequest) ToQuery() string {
	if r.ZoneID == "" {
		return ""
	}
	v := url.Values{}
	v.Set("zoneId", r.ZoneID)
	return "?" + v.Encode()
}
