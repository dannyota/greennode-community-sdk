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
	projectId   string
	zoneId      string
	userId      string
	moreHeaders map[string]string
	client      HttpClient
}

func NewServiceClient() ServiceClient {
	return &serviceClient{}
}

func (s *serviceClient) WithEndpoint(pendpoint string) ServiceClient {
	s.endpoint = normalizeURL(pendpoint)
	return s
}

func (s *serviceClient) WithName(pname string) ServiceClient {
	s.name = pname
	return s
}

func (s *serviceClient) WithZoneId(pzoneId string) ServiceClient {
	s.zoneId = pzoneId
	return s
}

func (s *serviceClient) WithUserId(puserId string) ServiceClient {
	s.userId = puserId
	return s
}

func (s *serviceClient) WithProjectId(pprojectId string) ServiceClient {
	s.projectId = pprojectId
	return s
}

func (s *serviceClient) WithMoreHeaders(pmoreHeaders map[string]string) ServiceClient {
	s.moreHeaders = pmoreHeaders
	return s
}

func (s *serviceClient) WithKVheader(pkey string, pvalue string) ServiceClient {
	s.moreHeaders[pkey] = pvalue
	return s
}

func (s *serviceClient) WithClient(pclient HttpClient) ServiceClient {
	s.client = pclient
	return s
}

func (s *serviceClient) ServiceURL(pparts ...string) string {
	return s.endpoint + strings.Join(pparts, "/")
}

func (s *serviceClient) Post(purl string, preq Request) (*req.Response, sdkerror.Error) {
	return s.client.DoRequest(purl, preq.WithRequestMethod(MethodPost))
}

func (s *serviceClient) Get(purl string, preq Request) (*req.Response, sdkerror.Error) {
	return s.client.DoRequest(purl, preq.WithRequestMethod(MethodGet))
}

func (s *serviceClient) Delete(purl string, preq Request) (*req.Response, sdkerror.Error) {
	return s.client.DoRequest(purl, preq.WithRequestMethod(MethodDelete))
}

func (s *serviceClient) Put(purl string, preq Request) (*req.Response, sdkerror.Error) {
	return s.client.DoRequest(purl, preq.WithRequestMethod(MethodPut))
}

func (s *serviceClient) Patch(purl string, preq Request) (*req.Response, sdkerror.Error) {
	return s.client.DoRequest(purl, preq.WithRequestMethod(MethodPatch))
}

func (s *serviceClient) GetProjectId() string {
	return s.projectId
}

func (s *serviceClient) GetZoneId() string {
	return s.zoneId
}

func (s *serviceClient) GetUserId() string {
	return s.userId
}

type sdkAuthentication struct {
	accessToken string
	expiresAt   int64
}

func NewSdkAuthentication() SdkAuthentication {
	return &sdkAuthentication{}
}

func (s *sdkAuthentication) WithAccessToken(paccessToken string) SdkAuthentication {
	s.accessToken = paccessToken
	return s
}

func (s *sdkAuthentication) WithExpiresAt(pexpiresAt int64) SdkAuthentication {
	s.expiresAt = pexpiresAt
	return s
}

func (s *sdkAuthentication) NeedReauth() bool {
	if s.accessToken == "" {
		return true
	}

	ea := time.Unix(0, s.expiresAt)
	return time.Until(ea) < 5*time.Minute
}

func (s *sdkAuthentication) UpdateAuth(pauth SdkAuthentication) {
	s.accessToken = pauth.GetAccessToken()
	s.expiresAt = pauth.GetExpiresAt()
}

func (s *sdkAuthentication) GetAccessToken() string {
	return s.accessToken
}

func (s *sdkAuthentication) GetExpiresAt() int64 {
	return s.expiresAt
}

func normalizeURL(u string) string {
	if !strings.HasSuffix(u, "/") {
		return u + "/"
	}
	return u
}
