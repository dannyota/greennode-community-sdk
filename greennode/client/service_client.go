package client

import (
	"strings"
	"time"

	"github.com/imroc/req/v3"
	sdkerror "github.com/dannyota/greennode-community-sdk/v2/greennode/sdkerror"
)

type serviceClient struct {
	name        string
	endpoint    string
	projectID   string
	zoneID      string
	userID      string
	moreHeaders map[string]string
	client      HTTPClient
}

func NewServiceClient() ServiceClient {
	return &serviceClient{}
}

func (sc *serviceClient) WithEndpoint(endpoint string) ServiceClient {
	sc.endpoint = normalizeURL(endpoint)
	return sc
}

func (sc *serviceClient) WithName(name string) ServiceClient {
	sc.name = name
	return sc
}

func (sc *serviceClient) WithZoneID(zoneID string) ServiceClient {
	sc.zoneID = zoneID
	return sc
}

func (sc *serviceClient) WithUserID(userID string) ServiceClient {
	sc.userID = userID
	return sc
}

func (sc *serviceClient) WithProjectID(projectID string) ServiceClient {
	sc.projectID = projectID
	return sc
}

func (sc *serviceClient) WithMoreHeaders(moreHeaders map[string]string) ServiceClient {
	sc.moreHeaders = moreHeaders
	return sc
}

func (sc *serviceClient) WithKVheader(key string, value string) ServiceClient {
	sc.moreHeaders[key] = value
	return sc
}

func (sc *serviceClient) WithClient(client HTTPClient) ServiceClient {
	sc.client = client
	return sc
}

func (sc *serviceClient) ServiceURL(parts ...string) string {
	return sc.endpoint + strings.Join(parts, "/")
}

func (sc *serviceClient) Post(url string, req Request) (*req.Response, sdkerror.Error) {
	return sc.client.DoRequest(url, req.WithRequestMethod(MethodPost))
}

func (sc *serviceClient) Get(url string, req Request) (*req.Response, sdkerror.Error) {
	return sc.client.DoRequest(url, req.WithRequestMethod(MethodGet))
}

func (sc *serviceClient) Delete(url string, req Request) (*req.Response, sdkerror.Error) {
	return sc.client.DoRequest(url, req.WithRequestMethod(MethodDelete))
}

func (sc *serviceClient) Put(url string, req Request) (*req.Response, sdkerror.Error) {
	return sc.client.DoRequest(url, req.WithRequestMethod(MethodPut))
}

func (sc *serviceClient) Patch(url string, req Request) (*req.Response, sdkerror.Error) {
	return sc.client.DoRequest(url, req.WithRequestMethod(MethodPatch))
}

func (sc *serviceClient) GetProjectID() string {
	return sc.projectID
}

func (sc *serviceClient) GetZoneID() string {
	return sc.zoneID
}

func (sc *serviceClient) GetUserID() string {
	return sc.userID
}

type sdkAuthentication struct {
	accessToken string
	expiresAt   int64
}

func NewSdkAuthentication() SdkAuthentication {
	return &sdkAuthentication{}
}

func (a *sdkAuthentication) WithAccessToken(accessToken string) SdkAuthentication {
	a.accessToken = accessToken
	return a
}

func (a *sdkAuthentication) WithExpiresAt(expiresAt int64) SdkAuthentication {
	a.expiresAt = expiresAt
	return a
}

func (a *sdkAuthentication) NeedReauth() bool {
	if a.accessToken == "" {
		return true
	}

	ea := time.Unix(0, a.expiresAt)
	return time.Until(ea) < 5*time.Minute
}

func (a *sdkAuthentication) UpdateAuth(auth SdkAuthentication) {
	a.accessToken = auth.AccessToken()
	a.expiresAt = auth.ExpiresAt()
}

func (a *sdkAuthentication) AccessToken() string {
	return a.accessToken
}

func (a *sdkAuthentication) ExpiresAt() int64 {
	return a.expiresAt
}

func normalizeURL(u string) string {
	if !strings.HasSuffix(u, "/") {
		return u + "/"
	}
	return u
}
