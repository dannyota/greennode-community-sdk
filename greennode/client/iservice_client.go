package client

import (
	"github.com/imroc/req/v3"

	sdkerror "github.com/dannyota/greennode-community-sdk/v2/greennode/sdkerror"
)

type ServiceClient interface {
	WithEndpoint(endpoint string) ServiceClient
	WithName(name string) ServiceClient
	WithProjectID(projectID string) ServiceClient
	WithZoneID(zoneID string) ServiceClient
	WithUserID(userID string) ServiceClient
	WithMoreHeaders(moreHeaders map[string]string) ServiceClient
	WithKVheader(key string, value string) ServiceClient
	WithClient(client HTTPClient) ServiceClient
	ServiceURL(parts ...string) string
	GetProjectID() string
	GetZoneID() string
	GetUserID() string

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
	AccessToken() string
	ExpiresAt() int64
}
