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
	return sc.HTTP.DoRequest(ctx, url, req.WithRequestMethod(MethodPost))
}

func (sc *ServiceClient) Get(ctx context.Context, url string, req *Request) (*http.Response, error) {
	return sc.HTTP.DoRequest(ctx, url, req.WithRequestMethod(MethodGet))
}

func (sc *ServiceClient) Delete(ctx context.Context, url string, req *Request) (*http.Response, error) {
	return sc.HTTP.DoRequest(ctx, url, req.WithRequestMethod(MethodDelete))
}

func (sc *ServiceClient) Put(ctx context.Context, url string, req *Request) (*http.Response, error) {
	return sc.HTTP.DoRequest(ctx, url, req.WithRequestMethod(MethodPut))
}

func (sc *ServiceClient) Patch(ctx context.Context, url string, req *Request) (*http.Response, error) {
	return sc.HTTP.DoRequest(ctx, url, req.WithRequestMethod(MethodPatch))
}

func NormalizeURL(u string) string {
	if !strings.HasSuffix(u, "/") {
		return u + "/"
	}
	return u
}
