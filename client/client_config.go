package client

import "strings"

type (
	SdkConfigure interface {
		GetClientID() string
		GetClientSecret() string
		GetProjectID() string
		GetUserID() string
		GetZoneID() string
		IamEndpoint() string
		VServerEndpoint() string
		VLBEndpoint() string
		GLBEndpoint() string
		VNetworkEndpoint() string
		VDnsEndpoint() string
		UserAgent() string
		WithUserID(userID string) SdkConfigure
		WithZoneID(zoneID string) SdkConfigure
		WithClientID(clientID string) SdkConfigure
		WithClientSecret(clientSecret string) SdkConfigure
		WithUserAgent(userAgent string) SdkConfigure
		WithProjectID(projectID string) SdkConfigure
		WithIamEndpoint(iamEndpoint string) SdkConfigure
		WithVServerEndpoint(vserverEndpoint string) SdkConfigure
		WithVLBEndpoint(vlbEndpoint string) SdkConfigure
		WithVNetworkEndpoint(vnetworkEndpoint string) SdkConfigure
		WithVDnsEndpoint(vdnsEndpoint string) SdkConfigure
		WithGLBEndpoint(vlbEndpoint string) SdkConfigure
	}
)

type sdkConfigure struct {
	clientID         string
	clientSecret     string
	projectID        string
	zoneID           string
	userID           string
	iamEndpoint      string
	vserverEndpoint  string
	vlbEndpoint      string
	glbEndpoint      string
	vnetworkEndpoint string
	vdnsEndpoint     string
	userAgent        string
}

func (s *sdkConfigure) GetClientID() string {
	return s.clientID
}

func (s *sdkConfigure) GetClientSecret() string {
	return s.clientSecret
}

func (s *sdkConfigure) GetProjectID() string {
	return s.projectID
}

func (s *sdkConfigure) GetUserID() string {
	return s.userID
}

func (s *sdkConfigure) GetZoneID() string {
	return s.zoneID
}

func (s *sdkConfigure) IamEndpoint() string {
	return s.iamEndpoint
}

func (s *sdkConfigure) VServerEndpoint() string {
	return s.vserverEndpoint
}

func (s *sdkConfigure) VLBEndpoint() string {
	return s.vlbEndpoint
}

func (s *sdkConfigure) GLBEndpoint() string {
	return s.glbEndpoint
}

func (s *sdkConfigure) VNetworkEndpoint() string {
	return s.vnetworkEndpoint
}

func (s *sdkConfigure) VDnsEndpoint() string {
	return s.vdnsEndpoint
}

func (s *sdkConfigure) UserAgent() string {
	return s.userAgent
}

func (s *sdkConfigure) WithUserAgent(userAgent string) SdkConfigure {
	s.userAgent = userAgent
	return s
}

func (s *sdkConfigure) WithClientID(clientID string) SdkConfigure {
	s.clientID = clientID
	return s
}

func (s *sdkConfigure) WithClientSecret(clientSecret string) SdkConfigure {
	s.clientSecret = clientSecret
	return s
}

func (s *sdkConfigure) WithUserID(userID string) SdkConfigure {
	s.userID = userID
	return s
}

func (s *sdkConfigure) WithZoneID(zoneID string) SdkConfigure {
	s.zoneID = zoneID
	return s
}

func (s *sdkConfigure) WithProjectID(projectID string) SdkConfigure {
	s.projectID = projectID
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
