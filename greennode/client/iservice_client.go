package client

import (
	"github.com/imroc/req/v3"

	sdkerror "github.com/dannyota/greennode-community-sdk/v2/greennode/sdkerror"
)

type ServiceClient interface {
	WithEndpoint(endpoint string) ServiceClient
	WithName(name string) ServiceClient
	WithProjectId(projectId string) ServiceClient
	WithZoneId(zoneId string) ServiceClient
	WithUserId(userId string) ServiceClient
	WithMoreHeaders(moreHeaders map[string]string) ServiceClient
	WithKVheader(key string, value string) ServiceClient
	WithClient(client HttpClient) ServiceClient
	ServiceURL(parts ...string) string
	GetProjectId() string
	GetZoneId() string
	GetUserId() string

	Post(url string, req Request) (*req.Response, sdkerror.Error)
	Get(url string, req Request) (*req.Response, sdkerror.Error)
	Delete(url string, req Request) (*req.Response, sdkerror.Error)
	Put(url string, req Request) (*req.Response, sdkerror.Error)
	Patch(url string, req Request) (*req.Response, sdkerror.Error)
}

type SdkAuthentication interface {
	WithAccessToken(accessToken string) SdkAuthentication
	WithExpiresAt(expiresAt int64) SdkAuthentication
	UpdateAuth(auth SdkAuthentication)
	NeedReauth() bool
	GetAccessToken() string
	GetExpiresAt() int64
}
