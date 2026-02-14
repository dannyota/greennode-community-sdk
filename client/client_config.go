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
		WithUserId(puserId string) SdkConfigure
		WithZoneId(pzoneId string) SdkConfigure
		WithClientId(pclientId string) SdkConfigure
		WithClientSecret(pclientSecret string) SdkConfigure
		WithUserAgent(puserAgent string) SdkConfigure
		WithProjectId(pprojectId string) SdkConfigure
		WithIamEndpoint(piamEndpoint string) SdkConfigure
		WithVServerEndpoint(pvserverEndpoint string) SdkConfigure
		WithVLBEndpoint(pvlbEndpoint string) SdkConfigure
		WithVNetworkEndpoint(pvnetworkEndpoint string) SdkConfigure
		WithVDnsEndpoint(pvdnsEndpoint string) SdkConfigure
		WithGLBEndpoint(pvlbEndpoint string) SdkConfigure
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

func (s *sdkConfigure) WithUserAgent(puserAgent string) SdkConfigure {
	s.userAgent = puserAgent
	return s
}

func (s *sdkConfigure) WithClientId(pclientId string) SdkConfigure {
	s.clientId = pclientId
	return s
}

func (s *sdkConfigure) WithClientSecret(pclientSecret string) SdkConfigure {
	s.clientSecret = pclientSecret
	return s
}

func (s *sdkConfigure) WithUserId(puserId string) SdkConfigure {
	s.userId = puserId
	return s
}

func (s *sdkConfigure) WithZoneId(pzoneId string) SdkConfigure {
	s.zoneId = pzoneId
	return s
}

func (s *sdkConfigure) WithProjectId(pprojectId string) SdkConfigure {
	s.projectId = pprojectId
	return s
}

func (s *sdkConfigure) WithIamEndpoint(piamEndpoint string) SdkConfigure {
	s.iamEndpoint = normalizeURL(piamEndpoint)
	return s
}

func (s *sdkConfigure) WithVServerEndpoint(pvserverEndpoint string) SdkConfigure {
	s.vserverEndpoint = normalizeURL(pvserverEndpoint)
	return s
}

func (s *sdkConfigure) WithVLBEndpoint(pvlbEndpoint string) SdkConfigure {
	s.vlbEndpoint = normalizeURL(pvlbEndpoint)
	return s
}

func (s *sdkConfigure) WithVNetworkEndpoint(pvnetworkEndpoint string) SdkConfigure {
	s.vnetworkEndpoint = normalizeURL(pvnetworkEndpoint)
	return s
}

func (s *sdkConfigure) WithVDnsEndpoint(pvdnsEndpoint string) SdkConfigure {
	s.vdnsEndpoint = normalizeURL(pvdnsEndpoint)
	return s
}

func (s *sdkConfigure) WithGLBEndpoint(pvlbEndpoint string) SdkConfigure {
	s.glbEndpoint = normalizeURL(pvlbEndpoint)
	return s
}

func normalizeURL(u string) string {
	if !strings.HasSuffix(u, "/") {
		return u + "/"
	}
	return u
}
