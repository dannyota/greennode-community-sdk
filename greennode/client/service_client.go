package client

import (
	"context"
	"strings"
	"time"

	"github.com/imroc/req/v3"
)

type ServiceClient struct {
	name        string
	endpoint    string
	projectID   string
	zoneID      string
	userID      string
	moreHeaders map[string]string
	client      *HTTPClient
}

func NewServiceClient() *ServiceClient {
	return &ServiceClient{}
}

func (sc *ServiceClient) WithEndpoint(endpoint string) *ServiceClient {
	sc.endpoint = normalizeURL(endpoint)
	return sc
}

func (sc *ServiceClient) WithName(name string) *ServiceClient {
	sc.name = name
	return sc
}

func (sc *ServiceClient) WithZoneID(zoneID string) *ServiceClient {
	sc.zoneID = zoneID
	return sc
}

func (sc *ServiceClient) WithUserID(userID string) *ServiceClient {
	sc.userID = userID
	return sc
}

func (sc *ServiceClient) WithProjectID(projectID string) *ServiceClient {
	sc.projectID = projectID
	return sc
}

func (sc *ServiceClient) WithMoreHeaders(moreHeaders map[string]string) *ServiceClient {
	sc.moreHeaders = moreHeaders
	return sc
}

func (sc *ServiceClient) WithKVheader(key string, value string) *ServiceClient {
	sc.moreHeaders[key] = value
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

func (sc *ServiceClient) GetProjectID() string {
	return sc.projectID
}

func (sc *ServiceClient) GetZoneID() string {
	return sc.zoneID
}

func (sc *ServiceClient) GetUserID() string {
	return sc.userID
}

type SdkAuthentication struct {
	accessToken string
	expiresAt   int64
}

func NewSdkAuthentication() *SdkAuthentication {
	return &SdkAuthentication{}
}

func (a *SdkAuthentication) WithAccessToken(accessToken string) *SdkAuthentication {
	a.accessToken = accessToken
	return a
}

func (a *SdkAuthentication) WithExpiresAt(expiresAt int64) *SdkAuthentication {
	a.expiresAt = expiresAt
	return a
}

func (a *SdkAuthentication) NeedReauth() bool {
	if a.accessToken == "" {
		return true
	}

	ea := time.Unix(0, a.expiresAt)
	return time.Until(ea) < 5*time.Minute
}

func (a *SdkAuthentication) UpdateAuth(auth *SdkAuthentication) {
	a.accessToken = auth.AccessToken()
	a.expiresAt = auth.ExpiresAt()
}

func (a *SdkAuthentication) AccessToken() string {
	return a.accessToken
}

func (a *SdkAuthentication) ExpiresAt() int64 {
	return a.expiresAt
}

func normalizeURL(u string) string {
	if !strings.HasSuffix(u, "/") {
		return u + "/"
	}
	return u
}
