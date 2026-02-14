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
	IamOauth2 AuthOpts = "IamOauth2"
)

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

func NewHttpClient(ctx context.Context) HttpClient {
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

func (s *httpClient) WithRetryCount(retryCount int) HttpClient {
	s.client.SetCommonRetryCount(retryCount)
	return s
}

func (s *httpClient) WithTimeout(timeout time.Duration) HttpClient {
	s.client.SetTimeout(timeout)
	return s
}

func (s *httpClient) WithSleep(sleep time.Duration) HttpClient {
	s.client.SetCommonRetryFixedInterval(sleep)
	return s
}

func (s *httpClient) WithKvDefaultHeaders(args ...string) HttpClient {
	if s.defaultHeaders == nil {
		s.defaultHeaders = make(map[string]string)
	}

	if len(args)%2 != 0 {
		args = append(args, "")
	}

	for i := 0; i < len(args); i += 2 {
		s.defaultHeaders[args[i]] = args[i+1]
	}

	return s
}

func (s *httpClient) WithReauthFunc(authOpt AuthOpts, reauthFunc func() (SdkAuthentication, sdkerror.Error)) HttpClient {
	s.reauthFunc = reauthFunc
	s.reauthOption = authOpt
	return s
}

func (s *httpClient) DoRequest(url string, preq Request) (*req.Response, sdkerror.Error) {
	req := s.prepareRequest(preq)

	resp, sdkErr := s.executeRequest(url, req, preq)
	if sdkErr != nil {
		return resp, sdkErr
	}

	return s.handleResponse(url, resp, preq)
}

func (s *httpClient) prepareRequest(preq Request) *req.Request {
	req := s.client.R().SetContext(s.context).SetHeaders(s.getDefaultHeaders()).SetHeaders(preq.GetMoreHeaders())

	if opt := preq.GetRequestBody(); opt != nil {
		req.SetBodyJsonMarshal(opt)
	}

	if opt := preq.GetJsonResponse(); opt != nil {
		req.SetSuccessResult(opt)
	}

	if opt := preq.GetJsonError(); opt != nil {
		req.SetErrorResult(opt)
	}

	return req
}

func (s *httpClient) executeRequest(url string, req *req.Request, preq Request) (*req.Response, sdkerror.Error) {
	if s.needReauth(preq) {
		return s.handleReauthBeforeRequest(url, preq)
	}

	resp, err := s.executeHttpMethod(url, req, preq)

	if err != nil && resp == nil {
		return resp, sdkerror.ErrorHandler(err)
	}

	return resp, nil
}

func (s *httpClient) executeHttpMethod(url string, req *req.Request, preq Request) (*req.Response, error) {
	switch strings.ToUpper(preq.GetRequestMethod()) {
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

func (s *httpClient) handleReauthBeforeRequest(url string, req Request) (*req.Response, sdkerror.Error) {
	if !req.SkipAuthentication() && s.reauthFunc != nil {
		if sdkErr := s.reauthenticate(); sdkErr != nil {
			return nil, sdkErr
		}
		return s.DoRequest(url, req)
	}
	return nil, nil
}

func (s *httpClient) handleResponse(url string, resp *req.Response, req Request) (*req.Response, sdkerror.Error) {
	if resp == nil || resp.Response == nil {
		return nil, sdkerror.ErrorHandler(nil, sdkerror.WithErrorUnexpected(resp))
	}

	if sdkErr := s.handleStatusCode(url, resp, req); sdkErr != nil {
		return nil, sdkErr
	}

	if req.ContainsOkCode(resp.StatusCode) {
		return resp, nil
	}

	return resp, sdkerror.ErrorHandler(resp.Err)
}

func (s *httpClient) handleStatusCode(url string, resp *req.Response, req Request) sdkerror.Error {
	switch resp.StatusCode {
	case http.StatusUnauthorized:
		return s.handleUnauthorized(url, resp, req)
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

func (s *httpClient) handleUnauthorized(url string, resp *req.Response, req Request) sdkerror.Error {
	if !req.SkipAuthentication() && s.reauthFunc != nil {
		if sdkErr := s.reauthenticate(); sdkErr != nil {
			return sdkErr
		}
		// Note: This will cause recursion - returning to trigger DoRequest again
		_, err := s.DoRequest(url, req)
		return err
	}
	return defaultErrorResponse(resp.Err, url, req, resp)
}

func (s *httpClient) needReauth(req Request) bool {
	if req.SkipAuthentication() {
		return false
	}

	if s.accessToken == nil {
		return true
	}

	return s.accessToken.NeedReauth()
}

func (s *httpClient) reauthenticate() sdkerror.Error {
	if s.reauthFunc == nil {
		return sdkerror.ErrorHandler(nil, sdkerror.WithErrorReauthFuncNotSet())
	}

	s.reauthmut.Lock()
	ongoing := s.reauthmut.ongoing
	if ongoing == nil {
		s.reauthmut.ongoing = newReauthFuture()
	}
	s.reauthmut.Unlock()

	if ongoing != nil {
		return ongoing.get()
	}

	auth, sdkerr := s.reauthFunc()
	s.reauthmut.Lock()
	s.reauthmut.ongoing.set(sdkerr)
	s.reauthmut.ongoing = nil
	s.reauthmut.Unlock()

	s.setAccessToken(auth)

	return sdkerr
}

func (s *httpClient) setAccessToken(newToken SdkAuthentication) HttpClient {
	s.mut.Lock()
	defer s.mut.Unlock()
	if newToken != nil {
		s.accessToken = newToken
		s.WithKvDefaultHeaders("Authorization", "Bearer "+s.accessToken.GetAccessToken())
	}

	return s
}

func (s *httpClient) getDefaultHeaders() map[string]string {
	s.mut.RLock()
	defer s.mut.RUnlock()
	if s.defaultHeaders == nil {
		s.defaultHeaders = make(map[string]string)
	}

	return s.defaultHeaders
}

func newReauthFuture() *reauthFuture {
	return &reauthFuture{
		done: make(chan struct{}),
		err:  nil,
	}
}

func (s *reauthFuture) get() sdkerror.Error {
	<-s.done
	return s.err
}

func (s *reauthFuture) set(err sdkerror.Error) {
	s.err = err
	close(s.done)
}

func defaultErrorResponse(err error, url string, req Request, resp *req.Response) sdkerror.Error {
	headers := req.GetMoreHeaders()

	// Remove sensitive information
	if headers != nil {
		delete(headers, "Authorization")
	}

	return sdkerror.ErrorHandler(err).WithKVparameters(
		"statusCode", resp.StatusCode,
		"url", url,
		"method", req.GetRequestMethod(),
		"requestHeaders", headers,
		"responseHeaders", resp.Header,
	)
}
