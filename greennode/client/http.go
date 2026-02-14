package client

import (
	"context"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/imroc/req/v3"

	sdkerror "github.com/dannyota/greennode-community-sdk/v2/greennode/sdkerror"
)

const (
	IAMOauth2 AuthOpts = "IamOauth2"
)

type HTTPClient interface {
	WithRetryCount(retryCount int) HTTPClient
	WithTimeout(timeout time.Duration) HTTPClient
	WithSleep(sleep time.Duration) HTTPClient
	WithKvDefaultHeaders(args ...string) HTTPClient
	WithReauthFunc(authOpt AuthOpts, reauthFunc func() (SdkAuthentication, sdkerror.Error)) HTTPClient

	DoRequest(url string, req Request) (*req.Response, sdkerror.Error)
}

type (
	httpClient struct {
		context    context.Context
		retryCount int
		client     *req.Client

		reauthFunc   func() (SdkAuthentication, sdkerror.Error)
		reauthOption AuthOpts

		accessToken    SdkAuthentication
		defaultHeaders map[string]string

		mut       *sync.RWMutex
		reauthmut *reauthlock
	}

	reauthlock struct {
		sync.RWMutex
		ongoing *reauthFuture
	}

	reauthFuture struct {
		done chan struct{}
		err  sdkerror.Error
	}

	AuthOpts string
)

func NewHTTPClient(ctx context.Context) HTTPClient {
	return &httpClient{
		context:    ctx,
		retryCount: 0,
		client: req.NewClient().
			SetCommonRetryCount(3).
			SetCommonRetryFixedInterval(10).
			SetTimeout(120 * time.Second),
		mut:       new(sync.RWMutex),
		reauthmut: new(reauthlock),
	}
}

func (hc *httpClient) WithRetryCount(retryCount int) HTTPClient {
	hc.client.SetCommonRetryCount(retryCount)
	return hc
}

func (hc *httpClient) WithTimeout(timeout time.Duration) HTTPClient {
	hc.client.SetTimeout(timeout)
	return hc
}

func (hc *httpClient) WithSleep(sleep time.Duration) HTTPClient {
	hc.client.SetCommonRetryFixedInterval(sleep)
	return hc
}

func (hc *httpClient) WithKvDefaultHeaders(args ...string) HTTPClient {
	if hc.defaultHeaders == nil {
		hc.defaultHeaders = make(map[string]string)
	}

	if len(args)%2 != 0 {
		args = append(args, "")
	}

	for i := 0; i < len(args); i += 2 {
		hc.defaultHeaders[args[i]] = args[i+1]
	}

	return hc
}

func (hc *httpClient) WithReauthFunc(authOpt AuthOpts, reauthFunc func() (SdkAuthentication, sdkerror.Error)) HTTPClient {
	hc.reauthFunc = reauthFunc
	hc.reauthOption = authOpt
	return hc
}

func (hc *httpClient) DoRequest(url string, preq Request) (*req.Response, sdkerror.Error) {
	req := hc.prepareRequest(preq)

	resp, sdkErr := hc.executeRequest(url, req, preq)
	if sdkErr != nil {
		return resp, sdkErr
	}

	return hc.handleResponse(url, resp, preq)
}

func (hc *httpClient) prepareRequest(preq Request) *req.Request {
	req := hc.client.R().SetContext(hc.context).SetHeaders(hc.getDefaultHeaders()).SetHeaders(preq.MoreHeaders())

	if opt := preq.RequestBody(); opt != nil {
		req.SetBodyJsonMarshal(opt)
	}

	if opt := preq.JSONResponse(); opt != nil {
		req.SetSuccessResult(opt)
	}

	if opt := preq.JSONError(); opt != nil {
		req.SetErrorResult(opt)
	}

	return req
}

func (hc *httpClient) executeRequest(url string, req *req.Request, preq Request) (*req.Response, sdkerror.Error) {
	if hc.needReauth(preq) {
		return hc.handleReauthBeforeRequest(url, preq)
	}

	resp, err := hc.executeHTTPMethod(url, req, preq)

	if err != nil && resp == nil {
		return resp, sdkerror.ErrorHandler(err)
	}

	return resp, nil
}

func (hc *httpClient) executeHTTPMethod(url string, req *req.Request, preq Request) (*req.Response, error) {
	switch strings.ToUpper(preq.RequestMethod()) {
	case "POST":
		return req.Post(url)
	case "GET":
		return req.Get(url)
	case "DELETE":
		return req.Delete(url)
	case "PUT":
		return req.Put(url)
	case "PATCH":
		return req.Patch(url)
	default:
		return nil, nil
	}
}

func (hc *httpClient) handleReauthBeforeRequest(url string, req Request) (*req.Response, sdkerror.Error) {
	if !req.SkipAuthentication() && hc.reauthFunc != nil {
		if sdkErr := hc.reauthenticate(); sdkErr != nil {
			return nil, sdkErr
		}
		return hc.DoRequest(url, req)
	}
	return nil, nil
}

func (hc *httpClient) handleResponse(url string, resp *req.Response, req Request) (*req.Response, sdkerror.Error) {
	if resp == nil || resp.Response == nil {
		return nil, sdkerror.ErrorHandler(nil, sdkerror.WithErrorUnexpected(resp))
	}

	if sdkErr := hc.handleStatusCode(url, resp, req); sdkErr != nil {
		return nil, sdkErr
	}

	if req.ContainsOkCode(resp.StatusCode) {
		return resp, nil
	}

	return resp, sdkerror.ErrorHandler(resp.Err)
}

func (hc *httpClient) handleStatusCode(url string, resp *req.Response, req Request) sdkerror.Error {
	switch resp.StatusCode {
	case http.StatusUnauthorized:
		return hc.handleUnauthorized(url, resp, req)
	case http.StatusTooManyRequests:
		return sdkerror.SdkErrorHandler(
			defaultErrorResponse(resp.Err, url, req, resp), nil,
			sdkerror.WithErrorPermissionDenied())
	case http.StatusInternalServerError:
		return sdkerror.SdkErrorHandler(
			defaultErrorResponse(resp.Err, url, req, resp), nil,
			sdkerror.WithErrorInternalServerError())
	case http.StatusServiceUnavailable:
		return sdkerror.SdkErrorHandler(
			defaultErrorResponse(resp.Err, url, req, resp), nil,
			sdkerror.WithErrorServiceMaintenance())
	case http.StatusForbidden:
		return sdkerror.SdkErrorHandler(
			defaultErrorResponse(resp.Err, url, req, resp), nil,
			sdkerror.WithErrorPermissionDenied())
	}
	return nil
}

func (hc *httpClient) handleUnauthorized(url string, resp *req.Response, req Request) sdkerror.Error {
	if !req.SkipAuthentication() && hc.reauthFunc != nil {
		if sdkErr := hc.reauthenticate(); sdkErr != nil {
			return sdkErr
		}
		// Note: This will cause recursion - returning to trigger DoRequest again
		_, err := hc.DoRequest(url, req)
		return err
	}
	return defaultErrorResponse(resp.Err, url, req, resp)
}

func (hc *httpClient) needReauth(req Request) bool {
	if req.SkipAuthentication() {
		return false
	}

	if hc.accessToken == nil {
		return true
	}

	return hc.accessToken.NeedReauth()
}

func (hc *httpClient) reauthenticate() sdkerror.Error {
	if hc.reauthFunc == nil {
		return sdkerror.ErrorHandler(nil, sdkerror.WithErrorReauthFuncNotSet())
	}

	hc.reauthmut.Lock()
	ongoing := hc.reauthmut.ongoing
	if ongoing == nil {
		hc.reauthmut.ongoing = newReauthFuture()
	}
	hc.reauthmut.Unlock()

	if ongoing != nil {
		return ongoing.get()
	}

	auth, sdkerr := hc.reauthFunc()
	hc.reauthmut.Lock()
	hc.reauthmut.ongoing.set(sdkerr)
	hc.reauthmut.ongoing = nil
	hc.reauthmut.Unlock()

	hc.setAccessToken(auth)

	return sdkerr
}

func (hc *httpClient) setAccessToken(newToken SdkAuthentication) HTTPClient {
	hc.mut.Lock()
	defer hc.mut.Unlock()
	if newToken != nil {
		hc.accessToken = newToken
		hc.WithKvDefaultHeaders("Authorization", "Bearer "+hc.accessToken.AccessToken())
	}

	return hc
}

func (hc *httpClient) getDefaultHeaders() map[string]string {
	hc.mut.RLock()
	defer hc.mut.RUnlock()
	if hc.defaultHeaders == nil {
		hc.defaultHeaders = make(map[string]string)
	}

	return hc.defaultHeaders
}

func newReauthFuture() *reauthFuture {
	return &reauthFuture{
		done: make(chan struct{}),
		err:  nil,
	}
}

func (f *reauthFuture) get() sdkerror.Error {
	<-f.done
	return f.err
}

func (f *reauthFuture) set(err sdkerror.Error) {
	f.err = err
	close(f.done)
}

func defaultErrorResponse(err error, url string, req Request, resp *req.Response) sdkerror.Error {
	headers := req.MoreHeaders()

	// Remove sensitive information
	if headers != nil {
		delete(headers, "Authorization")
	}

	return sdkerror.ErrorHandler(err).WithKVparameters(
		"statusCode", resp.StatusCode,
		"url", url,
		"method", req.RequestMethod(),
		"requestHeaders", headers,
		"responseHeaders", resp.Header,
	)
}
