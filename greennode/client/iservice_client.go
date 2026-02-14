package client

import (
	"github.com/imroc/req/v3"

	sdkerror "github.com/dannyota/greennode-community-sdk/v2/greennode/sdkerror"
)

type IServiceClient interface {
	WithEndpoint(pendpoint string) IServiceClient
	WithName(pname string) IServiceClient
	WithProjectId(pprojectId string) IServiceClient
	WithZoneId(pzoneId string) IServiceClient
	WithUserId(puserId string) IServiceClient
	WithMoreHeaders(pmoreHeaders map[string]string) IServiceClient
	WithKVheader(pkey string, pvalue string) IServiceClient
	WithClient(pclient IHttpClient) IServiceClient
	ServiceURL(pparts ...string) string
	GetProjectId() string
	GetZoneId() string
	GetUserId() string

	Post(purl string, preq IRequest) (*req.Response, sdkerror.IError)
	Get(purl string, preq IRequest) (*req.Response, sdkerror.IError)
	Delete(purl string, preq IRequest) (*req.Response, sdkerror.IError)
	Put(purl string, preq IRequest) (*req.Response, sdkerror.IError)
	Patch(purl string, preq IRequest) (*req.Response, sdkerror.IError)
}

type ISdkAuthentication interface {
	WithAccessToken(paccessToken string) ISdkAuthentication
	WithExpiresAt(pexpiresAt int64) ISdkAuthentication
	UpdateAuth(pauth ISdkAuthentication)
	NeedReauth() bool
	GetAccessToken() string
	GetExpiresAt() int64
}
