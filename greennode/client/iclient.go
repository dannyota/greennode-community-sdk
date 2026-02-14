package client

import (
	"time"

	"github.com/imroc/req/v3"

	sdkerror "github.com/dannyota/greennode-community-sdk/v2/greennode/sdkerror"
)

type HTTPClient interface {
	WithRetryCount(retryCount int) HTTPClient
	WithTimeout(timeout time.Duration) HTTPClient
	WithSleep(sleep time.Duration) HTTPClient
	WithKvDefaultHeaders(args ...string) HTTPClient
	WithReauthFunc(authOpt AuthOpts, reauthFunc func() (SdkAuthentication, sdkerror.Error)) HTTPClient

	DoRequest(url string, req Request) (*req.Response, sdkerror.Error)
}

type Request interface {
	WithOkCodes(okCodes ...int) Request
	WithJSONBody(jsonBody interface{}) Request
	WithJSONResponse(jsonResponse interface{}) Request
	WithJSONError(jsonError interface{}) Request
	WithRequestMethod(method requestMethod) Request
	WithSkipAuth(skipAuth bool) Request
	WithHeader(key, value string) Request
	WithMapHeaders(headers map[string]string) Request
	WithUserID(userID string) Request

	RequestBody() interface{}
	RequestMethod() string
	MoreHeaders() map[string]string
	JSONResponse() interface{}
	JSONError() interface{}

	SetJSONResponse(jsonResponse interface{})
	SetJSONError(jsonError interface{})

	ContainsOkCode(code ...int) bool
	SkipAuthentication() bool
}
