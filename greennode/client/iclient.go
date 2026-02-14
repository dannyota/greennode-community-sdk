package client

import (
	"time"

	"github.com/imroc/req/v3"

	sdkerror "github.com/dannyota/greennode-community-sdk/v2/greennode/sdkerror"
)

type HttpClient interface {
	WithRetryCount(pretryCount int) HttpClient
	WithTimeout(ptimeout time.Duration) HttpClient
	WithSleep(psleep time.Duration) HttpClient
	WithKvDefaultHeaders(pargs ...string) HttpClient
	WithReauthFunc(pauthOpt AuthOpts, preauthFunc func() (SdkAuthentication, sdkerror.Error)) HttpClient

	DoRequest(purl string, preq Request) (*req.Response, sdkerror.Error)
}

type Request interface {
	WithOkCodes(pokCodes ...int) Request
	WithJsonBody(pjsonBody interface{}) Request
	WithJsonResponse(pjsonResponse interface{}) Request
	WithJsonError(pjsonError interface{}) Request
	WithRequestMethod(pmethod requestMethod) Request
	WithSkipAuth(pskipAuth bool) Request
	WithHeader(pkey, pvalue string) Request
	WithMapHeaders(pheaders map[string]string) Request
	WithUserId(puserId string) Request

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
