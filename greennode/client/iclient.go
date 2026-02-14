package client

import (
	"time"

	"github.com/imroc/req/v3"

	sdkerror "github.com/dannyota/greennode-community-sdk/v2/greennode/sdkerror"
)

type HttpClient interface {
	WithRetryCount(retryCount int) HttpClient
	WithTimeout(timeout time.Duration) HttpClient
	WithSleep(sleep time.Duration) HttpClient
	WithKvDefaultHeaders(args ...string) HttpClient
	WithReauthFunc(authOpt AuthOpts, reauthFunc func() (SdkAuthentication, sdkerror.Error)) HttpClient

	DoRequest(url string, req Request) (*req.Response, sdkerror.Error)
}

type Request interface {
	WithOkCodes(okCodes ...int) Request
	WithJsonBody(jsonBody interface{}) Request
	WithJsonResponse(jsonResponse interface{}) Request
	WithJsonError(jsonError interface{}) Request
	WithRequestMethod(method requestMethod) Request
	WithSkipAuth(skipAuth bool) Request
	WithHeader(key, value string) Request
	WithMapHeaders(headers map[string]string) Request
	WithUserId(userId string) Request

	GetRequestBody() interface{}
	GetRequestMethod() string
	GetMoreHeaders() map[string]string
	GetJsonResponse() interface{}
	GetJsonError() interface{}

	SetJsonResponse(jsonResponse interface{})
	SetJsonError(jsonError interface{})

	ContainsOkCode(code ...int) bool
	SkipAuthentication() bool
}
