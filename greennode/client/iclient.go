package client

import (
	"time"

	"github.com/imroc/req/v3"

	sdkerror "github.com/dannyota/greennode-community-sdk/v2/greennode/sdk_error"
)

type IHttpClient interface {
	WithRetryCount(pretryCount int) IHttpClient
	WithTimeout(ptimeout time.Duration) IHttpClient
	WithSleep(psleep time.Duration) IHttpClient
	WithKvDefaultHeaders(pargs ...string) IHttpClient
	WithReauthFunc(pauthOpt AuthOpts, preauthFunc func() (ISdkAuthentication, sdkerror.IError)) IHttpClient

	DoRequest(purl string, preq IRequest) (*req.Response, sdkerror.IError)
}

type IRequest interface {
	WithOkCodes(pokCodes ...int) IRequest
	WithJsonBody(pjsonBody interface{}) IRequest
	WithJsonResponse(pjsonResponse interface{}) IRequest
	WithJsonError(pjsonError interface{}) IRequest
	WithRequestMethod(pmethod requestMethod) IRequest
	WithSkipAuth(pskipAuth bool) IRequest
	WithHeader(pkey, pvalue string) IRequest
	WithMapHeaders(pheaders map[string]string) IRequest
	WithUserId(puserId string) IRequest

	GetRequestBody() interface{}
	GetRequestMethod() string
	GetMoreHeaders() map[string]string
	GetJsonResponse() interface{}
	GetJsonError() interface{}

	SetJsonResponse(pjsonResponse interface{})
	SetJsonError(pjsonError interface{})

	ContainsOkCode(pcode ...int) bool
	SkipAuthentication() bool
}
