package client

import (
	"context"
	"strings"

	"github.com/imroc/req/v3"
)

type ServiceClient struct {
	endpoint  string
	projectID string
	zoneID    string
	client    *HTTPClient
}

func NewServiceClient() *ServiceClient {
	return &ServiceClient{}
}

func (sc *ServiceClient) WithEndpoint(endpoint string) *ServiceClient {
	sc.endpoint = NormalizeURL(endpoint)
	return sc
}

func (sc *ServiceClient) WithZoneID(zoneID string) *ServiceClient {
	sc.zoneID = zoneID
	return sc
}

func (sc *ServiceClient) WithProjectID(projectID string) *ServiceClient {
	sc.projectID = projectID
	return sc
}

func (sc *ServiceClient) WithClient(client *HTTPClient) *ServiceClient {
	sc.client = client
	return sc
}

func (sc *ServiceClient) ServiceURL(parts ...string) string {
	return sc.endpoint + strings.Join(parts, "/")
}

func (sc *ServiceClient) Post(ctx context.Context, url string, req *Request) (*req.Response, error) {
	return sc.client.DoRequest(ctx, url, req.WithRequestMethod(MethodPost))
}

func (sc *ServiceClient) Get(ctx context.Context, url string, req *Request) (*req.Response, error) {
	return sc.client.DoRequest(ctx, url, req.WithRequestMethod(MethodGet))
}

func (sc *ServiceClient) Delete(ctx context.Context, url string, req *Request) (*req.Response, error) {
	return sc.client.DoRequest(ctx, url, req.WithRequestMethod(MethodDelete))
}

func (sc *ServiceClient) Put(ctx context.Context, url string, req *Request) (*req.Response, error) {
	return sc.client.DoRequest(ctx, url, req.WithRequestMethod(MethodPut))
}

func (sc *ServiceClient) Patch(ctx context.Context, url string, req *Request) (*req.Response, error) {
	return sc.client.DoRequest(ctx, url, req.WithRequestMethod(MethodPatch))
}

func (sc *ServiceClient) ProjectID() string {
	return sc.projectID
}

func (sc *ServiceClient) ZoneID() string {
	return sc.zoneID
}

func NormalizeURL(u string) string {
	if !strings.HasSuffix(u, "/") {
		return u + "/"
	}
	return u
}
