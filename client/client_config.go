package client

import "strings"

type (
	SdkConfigure interface {
		GetClientId() string
		GetClientSecret() string
		GetProjectId() string
		GetUserId() string
		GetZoneId() string
		GetIamEndpoint() string
		GetVServerEndpoint() string
		GetVLBEndpoint() string
		GetGLBEndpoint() string
		GetVNetworkEndpoint() string
		GetVDnsEndpoint() string
		GetUserAgent() string
		WithUserId(userId string) SdkConfigure
		WithZoneId(zoneId string) SdkConfigure
		WithClientId(clientId string) SdkConfigure
		WithClientSecret(clientSecret string) SdkConfigure
		WithUserAgent(userAgent string) SdkConfigure
		WithProjectId(projectId string) SdkConfigure
		WithIamEndpoint(iamEndpoint string) SdkConfigure
		WithVServerEndpoint(vserverEndpoint string) SdkConfigure
		WithVLBEndpoint(vlbEndpoint string) SdkConfigure
		WithVNetworkEndpoint(vnetworkEndpoint string) SdkConfigure
		WithVDnsEndpoint(vdnsEndpoint string) SdkConfigure
		WithGLBEndpoint(vlbEndpoint string) SdkConfigure
	}
)

type sdkConfigure struct {
	clientId         string
	clientSecret     string
	projectId        string
	zoneId           string
	userId           string
	iamEndpoint      string
	vserverEndpoint  string
	vlbEndpoint      string
	glbEndpoint      string
	vnetworkEndpoint string
	vdnsEndpoint     string
	userAgent        string
}

func (s *sdkConfigure) GetClientId() string {
	return s.clientId
}

func (s *sdkConfigure) GetClientSecret() string {
	return s.clientSecret
}

func (s *sdkConfigure) GetProjectId() string {
	return s.projectId
}

func (s *sdkConfigure) GetUserId() string {
	return s.userId
}

func (s *sdkConfigure) GetZoneId() string {
	return s.zoneId
}

func (s *sdkConfigure) GetIamEndpoint() string {
	return s.iamEndpoint
}

func (s *sdkConfigure) GetVServerEndpoint() string {
	return s.vserverEndpoint
}

func (s *sdkConfigure) GetVLBEndpoint() string {
	return s.vlbEndpoint
}

func (s *sdkConfigure) GetGLBEndpoint() string {
	return s.glbEndpoint
}

func (s *sdkConfigure) GetVNetworkEndpoint() string {
	return s.vnetworkEndpoint
}

func (s *sdkConfigure) GetVDnsEndpoint() string {
	return s.vdnsEndpoint
}

func (s *sdkConfigure) GetUserAgent() string {
	return s.userAgent
}

func (s *sdkConfigure) WithUserAgent(userAgent string) SdkConfigure {
	s.userAgent = userAgent
	return s
}

func (s *sdkConfigure) WithClientId(clientId string) SdkConfigure {
	s.clientId = clientId
	return s
}

func (s *sdkConfigure) WithClientSecret(clientSecret string) SdkConfigure {
	s.clientSecret = clientSecret
	return s
}

func (s *sdkConfigure) WithUserId(userId string) SdkConfigure {
	s.userId = userId
	return s
}

func (s *sdkConfigure) WithZoneId(zoneId string) SdkConfigure {
	s.zoneId = zoneId
	return s
}

func (s *sdkConfigure) WithProjectId(projectId string) SdkConfigure {
	s.projectId = projectId
	return s
}

func (s *sdkConfigure) WithIamEndpoint(iamEndpoint string) SdkConfigure {
	s.iamEndpoint = normalizeURL(iamEndpoint)
	return s
}

func (s *sdkConfigure) WithVServerEndpoint(vserverEndpoint string) SdkConfigure {
	s.vserverEndpoint = normalizeURL(vserverEndpoint)
	return s
}

func (s *sdkConfigure) WithVLBEndpoint(vlbEndpoint string) SdkConfigure {
	s.vlbEndpoint = normalizeURL(vlbEndpoint)
	return s
}

func (s *sdkConfigure) WithVNetworkEndpoint(vnetworkEndpoint string) SdkConfigure {
	s.vnetworkEndpoint = normalizeURL(vnetworkEndpoint)
	return s
}

func (s *sdkConfigure) WithVDnsEndpoint(vdnsEndpoint string) SdkConfigure {
	s.vdnsEndpoint = normalizeURL(vdnsEndpoint)
	return s
}

func (s *sdkConfigure) WithGLBEndpoint(vlbEndpoint string) SdkConfigure {
	s.glbEndpoint = normalizeURL(vlbEndpoint)
	return s
}

func normalizeURL(u string) string {
	if !strings.HasSuffix(u, "/") {
		return u + "/"
	}
	return u
}
