package client

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"strings"
	"sync"
	"time"

	sdkerror "github.com/dannyota/greennode-community-sdk/v2/greennode/sdkerror"
)

const (
	IAMOauth2     AuthOpts = "IamOauth2"
	IAMUserOauth2 AuthOpts = "IamUserOauth2"
)

type (
	HTTPClient struct {
		retryCount    int
		retryInterval time.Duration
		client        *http.Client

		reauthFunc   func(ctx context.Context) (*Token, error)
		reauthOption AuthOpts

		token          *Token
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
		retryCount:    3,
		retryInterval: 10 * time.Second,
		client:        &http.Client{Timeout: 120 * time.Second},
		mut:           new(sync.RWMutex),
		reauthmut:     new(reauthlock),
	}
}

func (hc *HTTPClient) WithRetryCount(retryCount int) *HTTPClient {
	hc.retryCount = retryCount
	return hc
}

func (hc *HTTPClient) WithTimeout(timeout time.Duration) *HTTPClient {
	hc.client.Timeout = timeout
	return hc
}

func (hc *HTTPClient) WithSleep(sleep time.Duration) *HTTPClient {
	hc.retryInterval = sleep
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

func (hc *HTTPClient) WithReauthFunc(authOpt AuthOpts, reauthFunc func(ctx context.Context) (*Token, error)) *HTTPClient {
	hc.reauthFunc = reauthFunc
	hc.reauthOption = authOpt
	return hc
}

func (hc *HTTPClient) DoRequest(ctx context.Context, url string, preq *Request) (*http.Response, error) {
	resp, sdkErr := hc.executeRequest(ctx, url, preq)
	if sdkErr != nil {
		return resp, sdkErr
	}

	return hc.handleResponse(ctx, url, resp, preq)
}

func (hc *HTTPClient) prepareRequest(ctx context.Context, url string, preq *Request) (*http.Request, error) {
	method := strings.ToUpper(string(preq.method))
	if method == "" {
		method = "GET"
	}

	var body io.Reader
	if preq.jsonBody != nil {
		jsonData, err := json.Marshal(preq.jsonBody)
		if err != nil {
			return nil, sdkerror.ErrorHandler(err)
		}
		body = bytes.NewReader(jsonData)
	}

	req, err := http.NewRequestWithContext(ctx, method, url, body)
	if err != nil {
		return nil, sdkerror.ErrorHandler(err)
	}

	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	for k, v := range hc.getDefaultHeaders() {
		req.Header.Set(k, v)
	}
	for k, v := range preq.moreHeaders {
		req.Header.Set(k, v)
	}

	return req, nil
}

func (hc *HTTPClient) executeRequest(ctx context.Context, url string, preq *Request) (*http.Response, error) {
	if hc.needReauth(preq) {
		return hc.handleReauthBeforeRequest(ctx, url, preq)
	}

	req, err := hc.prepareRequest(ctx, url, preq)
	if err != nil {
		return nil, err
	}

	resp, err := hc.doWithRetry(req)
	if err != nil {
		return nil, sdkerror.ErrorHandler(err)
	}

	return resp, nil
}

func (hc *HTTPClient) doWithRetry(req *http.Request) (*http.Response, error) {
	var (
		resp *http.Response
		err  error
		body []byte
	)

	// Buffer the request body so we can retry
	if req.Body != nil {
		body, err = io.ReadAll(req.Body)
		req.Body.Close()
		if err != nil {
			return nil, err
		}
	}

	for attempt := 0; attempt <= hc.retryCount; attempt++ {
		if body != nil {
			req.Body = io.NopCloser(bytes.NewReader(body))
		}

		resp, err = hc.client.Do(req)
		if err == nil {
			break
		}

		if attempt < hc.retryCount {
			time.Sleep(hc.retryInterval)
		}
	}

	return resp, err
}

func (hc *HTTPClient) handleReauthBeforeRequest(ctx context.Context, url string, req *Request) (*http.Response, error) {
	if !req.skipAuth && hc.reauthFunc != nil {
		if sdkErr := hc.reauthenticate(ctx); sdkErr != nil {
			return nil, sdkErr
		}
		return hc.DoRequest(ctx, url, req)
	}
	return nil, nil
}

func (hc *HTTPClient) handleResponse(ctx context.Context, url string, resp *http.Response, preq *Request) (*http.Response, error) {
	if resp == nil {
		return nil, sdkerror.NewUnexpectedError(nil)
	}

	defer resp.Body.Close()
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return resp, sdkerror.ErrorHandler(err)
	}

	if sdkErr := hc.handleStatusCode(ctx, url, resp, preq); sdkErr != nil {
		return nil, sdkErr
	}

	if preq.containsOKCode(resp.StatusCode) {
		if preq.jsonResponse != nil && len(bodyBytes) > 0 {
			if err := json.Unmarshal(bodyBytes, preq.jsonResponse); err != nil {
				return resp, sdkerror.ErrorHandler(err)
			}
		}
		return resp, nil
	}

	if preq.jsonError != nil && len(bodyBytes) > 0 {
		json.Unmarshal(bodyBytes, preq.jsonError) //nolint:errcheck // best-effort error body parse
	}

	return resp, nil
}

func (hc *HTTPClient) handleStatusCode(ctx context.Context, url string, resp *http.Response, preq *Request) error {
	switch resp.StatusCode {
	case http.StatusUnauthorized:
		return hc.handleUnauthorized(ctx, url, resp, preq)
	case http.StatusTooManyRequests:
		return defaultErrorResponse(nil, url, preq, resp).
			WithErrorCode(sdkerror.EcPermissionDenied).
			WithMessage("Permission Denied")
	case http.StatusInternalServerError:
		return defaultErrorResponse(nil, url, preq, resp).
			WithErrorCode(sdkerror.EcInternalServerError).
			WithMessage("Internal Server Error")
	case http.StatusServiceUnavailable:
		return defaultErrorResponse(nil, url, preq, resp).
			WithErrorCode(sdkerror.EcServiceMaintenance).
			WithMessage("Service Maintenance")
	case http.StatusForbidden:
		return defaultErrorResponse(nil, url, preq, resp).
			WithErrorCode(sdkerror.EcPermissionDenied).
			WithMessage("Permission Denied")
	}
	return nil
}

func (hc *HTTPClient) handleUnauthorized(ctx context.Context, url string, resp *http.Response, req *Request) error {
	if !req.skipAuth && hc.reauthFunc != nil {
		if sdkErr := hc.reauthenticate(ctx); sdkErr != nil {
			return sdkErr
		}
		_, err := hc.DoRequest(ctx, url, req)
		return err
	}
	return defaultErrorResponse(nil, url, req, resp)
}

func (hc *HTTPClient) needReauth(req *Request) bool {
	if req.skipAuth {
		return false
	}

	if hc.token == nil {
		return true
	}

	return hc.token.NeedsReauth()
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

	hc.setToken(auth)

	return sdkerr
}

func (hc *HTTPClient) setToken(newToken *Token) *HTTPClient {
	hc.mut.Lock()
	defer hc.mut.Unlock()
	if newToken != nil {
		hc.token = newToken
		hc.WithKvDefaultHeaders("Authorization", "Bearer "+hc.token.AccessToken)
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

func defaultErrorResponse(err error, url string, req *Request, resp *http.Response) *sdkerror.SdkError {
	headers := req.moreHeaders

	// Remove sensitive information
	if headers != nil {
		delete(headers, "Authorization")
	}

	return sdkerror.ErrorHandler(err).WithKVparameters(
		"statusCode", resp.StatusCode,
		"url", url,
		"method", string(req.method),
		"requestHeaders", headers,
		"responseHeaders", resp.Header,
	)
}
