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

func (s *serviceClient) WithEndpoint(endpoint string) ServiceClient {
	s.endpoint = normalizeURL(endpoint)
	return s
}

func (s *serviceClient) WithName(name string) ServiceClient {
	s.name = name
	return s
}

func (s *serviceClient) WithZoneID(zoneID string) ServiceClient {
	s.zoneID = zoneID
	return s
}

func (s *serviceClient) WithUserID(userID string) ServiceClient {
	s.userID = userID
	return s
}

func (s *serviceClient) WithProjectID(projectID string) ServiceClient {
	s.projectID = projectID
	return s
}

func (s *serviceClient) WithMoreHeaders(moreHeaders map[string]string) ServiceClient {
	s.moreHeaders = moreHeaders
	return s
}

func (s *serviceClient) WithKVheader(key string, value string) ServiceClient {
	s.moreHeaders[key] = value
	return s
}

func (s *serviceClient) WithClient(client HTTPClient) ServiceClient {
	s.client = client
	return s
}

func (s *serviceClient) ServiceURL(parts ...string) string {
	return s.endpoint + strings.Join(parts, "/")
}

func (s *serviceClient) Post(url string, req Request) (*req.Response, sdkerror.Error) {
	return s.client.DoRequest(url, req.WithRequestMethod(MethodPost))
}

func (s *serviceClient) Get(url string, req Request) (*req.Response, sdkerror.Error) {
	return s.client.DoRequest(url, req.WithRequestMethod(MethodGet))
}

func (s *serviceClient) Delete(url string, req Request) (*req.Response, sdkerror.Error) {
	return s.client.DoRequest(url, req.WithRequestMethod(MethodDelete))
}

func (s *serviceClient) Put(url string, req Request) (*req.Response, sdkerror.Error) {
	return s.client.DoRequest(url, req.WithRequestMethod(MethodPut))
}

func (s *serviceClient) Patch(url string, req Request) (*req.Response, sdkerror.Error) {
	return s.client.DoRequest(url, req.WithRequestMethod(MethodPatch))
}

func (s *serviceClient) GetProjectID() string {
	return s.projectID
}

func (s *serviceClient) GetZoneID() string {
	return s.zoneID
}

func (s *serviceClient) GetUserID() string {
	return s.userID
}

type sdkAuthentication struct {
	accessToken string
	expiresAt   int64
}

func NewSdkAuthentication() SdkAuthentication {
	return &sdkAuthentication{}
}

func (s *sdkAuthentication) WithAccessToken(accessToken string) SdkAuthentication {
	s.accessToken = accessToken
	return s
}

func (s *sdkAuthentication) WithExpiresAt(expiresAt int64) SdkAuthentication {
	s.expiresAt = expiresAt
	return s
}

func (s *sdkAuthentication) NeedReauth() bool {
	if s.accessToken == "" {
		return true
	}

	ea := time.Unix(0, s.expiresAt)
	return time.Until(ea) < 5*time.Minute
}

func (s *sdkAuthentication) UpdateAuth(auth SdkAuthentication) {
	s.accessToken = auth.AccessToken()
	s.expiresAt = auth.ExpiresAt()
}

func (s *sdkAuthentication) AccessToken() string {
	return s.accessToken
}

func (s *sdkAuthentication) ExpiresAt() int64 {
	return s.expiresAt
}

func normalizeURL(u string) string {
	if !strings.HasSuffix(u, "/") {
		return u + "/"
	}
	return u
}
