package client

import (
	"github.com/imroc/req/v3"

	sdkerror "github.com/dannyota/greennode-community-sdk/v2/greennode/sdkerror"
)

type ServiceClient interface {
	WithEndpoint(pendpoint string) ServiceClient
	WithName(pname string) ServiceClient
	WithProjectId(pprojectId string) ServiceClient
	WithZoneId(pzoneId string) ServiceClient
	WithUserId(puserId string) ServiceClient
	WithMoreHeaders(pmoreHeaders map[string]string) ServiceClient
	WithKVheader(pkey string, pvalue string) ServiceClient
	WithClient(pclient HttpClient) ServiceClient
	ServiceURL(pparts ...string) string
	GetProjectId() string
	GetZoneId() string
	GetUserId() string

	Post(purl string, preq Request) (*req.Response, sdkerror.Error)
	Get(purl string, preq Request) (*req.Response, sdkerror.Error)
	Delete(purl string, preq Request) (*req.Response, sdkerror.Error)
	Put(purl string, preq Request) (*req.Response, sdkerror.Error)
	Patch(purl string, preq Request) (*req.Response, sdkerror.Error)
}

type SdkAuthentication interface {
	WithAccessToken(paccessToken string) SdkAuthentication
	WithExpiresAt(pexpiresAt int64) SdkAuthentication
	UpdateAuth(pauth SdkAuthentication)
	NeedReauth() bool
	GetAccessToken() string
	GetExpiresAt() int64
}
