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
	IAMOauth2     AuthOpts = "IamOauth2"
	IAMUserOauth2 AuthOpts = "IamUserOauth2"
)

type (
	HTTPClient struct {
		retryCount int
		client     *req.Client

		reauthFunc   func(ctx context.Context) (*SdkAuthentication, error)
		reauthOption AuthOpts

		accessToken    *SdkAuthentication
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
		err  error
	}

	AuthOpts string
)

func NewHTTPClient() *HTTPClient {
	return &HTTPClient{
		retryCount: 0,
		client: req.NewClient().
			SetCommonRetryCount(3).
			SetCommonRetryFixedInterval(10).
			SetTimeout(120 * time.Second),
		mut:       new(sync.RWMutex),
		reauthmut: new(reauthlock),
	}
}

func (hc *HTTPClient) WithRetryCount(retryCount int) *HTTPClient {
	hc.client.SetCommonRetryCount(retryCount)
	return hc
}

func (hc *HTTPClient) WithTimeout(timeout time.Duration) *HTTPClient {
	hc.client.SetTimeout(timeout)
	return hc
}

func (hc *HTTPClient) WithSleep(sleep time.Duration) *HTTPClient {
	hc.client.SetCommonRetryFixedInterval(sleep)
	return hc
}

func (hc *HTTPClient) WithKvDefaultHeaders(args ...string) *HTTPClient {
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

func (hc *HTTPClient) WithReauthFunc(authOpt AuthOpts, reauthFunc func(ctx context.Context) (*SdkAuthentication, error)) *HTTPClient {
	hc.reauthFunc = reauthFunc
	hc.reauthOption = authOpt
	return hc
}

func (hc *HTTPClient) DoRequest(ctx context.Context, url string, preq *Request) (*req.Response, error) {
	req := hc.prepareRequest(ctx, preq)

	resp, sdkErr := hc.executeRequest(ctx, url, req, preq)
	if sdkErr != nil {
		return resp, sdkErr
	}

	return hc.handleResponse(ctx, url, resp, preq)
}

func (hc *HTTPClient) prepareRequest(ctx context.Context, preq *Request) *req.Request {
	req := hc.client.R().SetContext(ctx).SetHeaders(hc.getDefaultHeaders()).SetHeaders(preq.MoreHeaders())

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

func (hc *HTTPClient) executeRequest(ctx context.Context, url string, req *req.Request, preq *Request) (*req.Response, error) {
	if hc.needReauth(preq) {
		return hc.handleReauthBeforeRequest(ctx, url, preq)
	}

	resp, err := hc.executeHTTPMethod(url, req, preq)

	if err != nil && resp == nil {
		return resp, sdkerror.ErrorHandler(err)
	}

	return resp, nil
}

func (hc *HTTPClient) executeHTTPMethod(url string, req *req.Request, preq *Request) (*req.Response, error) {
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

func (hc *HTTPClient) handleReauthBeforeRequest(ctx context.Context, url string, req *Request) (*req.Response, error) {
	if !req.SkipAuthentication() && hc.reauthFunc != nil {
		if sdkErr := hc.reauthenticate(ctx); sdkErr != nil {
			return nil, sdkErr
		}
		return hc.DoRequest(ctx, url, req)
	}
	return nil, nil
}

func (hc *HTTPClient) handleResponse(ctx context.Context, url string, resp *req.Response, req *Request) (*req.Response, error) {
	if resp == nil || resp.Response == nil {
		return nil, sdkerror.NewUnexpectedError(resp)
	}

	if sdkErr := hc.handleStatusCode(ctx, url, resp, req); sdkErr != nil {
		return nil, sdkErr
	}

	if req.ContainsOkCode(resp.StatusCode) {
		return resp, nil
	}

	return resp, sdkerror.ErrorHandler(resp.Err)
}

func (hc *HTTPClient) handleStatusCode(ctx context.Context, url string, resp *req.Response, preq *Request) error {
	switch resp.StatusCode {
	case http.StatusUnauthorized:
		return hc.handleUnauthorized(ctx, url, resp, preq)
	case http.StatusTooManyRequests:
		return defaultErrorResponse(resp.Err, url, preq, resp).
			WithErrorCode(sdkerror.EcPermissionDenied).
			WithMessage("Permission Denied")
	case http.StatusInternalServerError:
		return defaultErrorResponse(resp.Err, url, preq, resp).
			WithErrorCode(sdkerror.EcInternalServerError).
			WithMessage("Internal Server Error")
	case http.StatusServiceUnavailable:
		return defaultErrorResponse(resp.Err, url, preq, resp).
			WithErrorCode(sdkerror.EcServiceMaintenance).
			WithMessage("Service Maintenance")
	case http.StatusForbidden:
		return defaultErrorResponse(resp.Err, url, preq, resp).
			WithErrorCode(sdkerror.EcPermissionDenied).
			WithMessage("Permission Denied")
	}
	return nil
}

func (hc *HTTPClient) handleUnauthorized(ctx context.Context, url string, resp *req.Response, req *Request) error {
	if !req.SkipAuthentication() && hc.reauthFunc != nil {
		if sdkErr := hc.reauthenticate(ctx); sdkErr != nil {
			return sdkErr
		}
		// Note: This will cause recursion - returning to trigger DoRequest again
		_, err := hc.DoRequest(ctx, url, req)
		return err
	}
	return defaultErrorResponse(resp.Err, url, req, resp)
}

func (hc *HTTPClient) needReauth(req *Request) bool {
	if req.SkipAuthentication() {
		return false
	}

	if hc.accessToken == nil {
		return true
	}

	return hc.accessToken.NeedReauth()
}

func (hc *HTTPClient) reauthenticate(ctx context.Context) error {
	if hc.reauthFunc == nil {
		return sdkerror.NewReauthFuncNotSet()
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

	auth, sdkerr := hc.reauthFunc(ctx)
	hc.reauthmut.Lock()
	hc.reauthmut.ongoing.set(sdkerr)
	hc.reauthmut.ongoing = nil
	hc.reauthmut.Unlock()

	hc.setAccessToken(auth)

	return sdkerr
}

func (hc *HTTPClient) setAccessToken(newToken *SdkAuthentication) *HTTPClient {
	hc.mut.Lock()
	defer hc.mut.Unlock()
	if newToken != nil {
		hc.accessToken = newToken
		hc.WithKvDefaultHeaders("Authorization", "Bearer "+hc.accessToken.AccessToken())
	}

	return hc
}

func (hc *HTTPClient) getDefaultHeaders() map[string]string {
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

func (f *reauthFuture) get() error {
	<-f.done
	return f.err
}

func (f *reauthFuture) set(err error) {
	f.err = err
	close(f.done)
}

func defaultErrorResponse(err error, url string, req *Request, resp *req.Response) *sdkerror.SdkError {
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
