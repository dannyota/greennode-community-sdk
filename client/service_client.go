package client

import (
	"context"
	"net/http"
	"strings"
)

type ServiceClient struct {
	Endpoint  string
	ProjectID string
	ZoneID    string
	HTTP      *HTTPClient
}

func (sc *ServiceClient) ServiceURL(parts ...string) string {
	return sc.Endpoint + strings.Join(parts, "/")
}

func (sc *ServiceClient) Post(ctx context.Context, url string, req *Request) (*http.Response, error) {
	req.method = MethodPost
	return sc.HTTP.DoRequest(ctx, url, req)
}

func (sc *ServiceClient) Get(ctx context.Context, url string, req *Request) (*http.Response, error) {
	req.method = MethodGet
	return sc.HTTP.DoRequest(ctx, url, req)
}

func (sc *ServiceClient) Delete(ctx context.Context, url string, req *Request) (*http.Response, error) {
	req.method = MethodDelete
	return sc.HTTP.DoRequest(ctx, url, req)
}

func (sc *ServiceClient) Put(ctx context.Context, url string, req *Request) (*http.Response, error) {
	req.method = MethodPut
	return sc.HTTP.DoRequest(ctx, url, req)
}

func (sc *ServiceClient) Patch(ctx context.Context, url string, req *Request) (*http.Response, error) {
	req.method = MethodPatch
	return sc.HTTP.DoRequest(ctx, url, req)
}

func NormalizeURL(u string) string {
	if !strings.HasSuffix(u, "/") {
		return u + "/"
	}
	return u
}
