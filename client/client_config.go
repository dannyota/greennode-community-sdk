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

func (c *sdkConfigure) GetClientID() string {
	return c.clientID
}

func (c *sdkConfigure) GetClientSecret() string {
	return c.clientSecret
}

func (c *sdkConfigure) GetProjectID() string {
	return c.projectID
}

func (c *sdkConfigure) GetUserID() string {
	return c.userID
}

func (c *sdkConfigure) GetZoneID() string {
	return c.zoneID
}

func (c *sdkConfigure) IamEndpoint() string {
	return c.iamEndpoint
}

func (c *sdkConfigure) VServerEndpoint() string {
	return c.vserverEndpoint
}

func (c *sdkConfigure) VLBEndpoint() string {
	return c.vlbEndpoint
}

func (c *sdkConfigure) GLBEndpoint() string {
	return c.glbEndpoint
}

func (c *sdkConfigure) VNetworkEndpoint() string {
	return c.vnetworkEndpoint
}

func (c *sdkConfigure) VDnsEndpoint() string {
	return c.vdnsEndpoint
}

func (c *sdkConfigure) UserAgent() string {
	return c.userAgent
}

func (c *sdkConfigure) WithUserAgent(userAgent string) SdkConfigure {
	c.userAgent = userAgent
	return c
}

func (c *sdkConfigure) WithClientID(clientID string) SdkConfigure {
	c.clientID = clientID
	return c
}

func (c *sdkConfigure) WithClientSecret(clientSecret string) SdkConfigure {
	c.clientSecret = clientSecret
	return c
}

func (c *sdkConfigure) WithUserID(userID string) SdkConfigure {
	c.userID = userID
	return c
}

func (c *sdkConfigure) WithZoneID(zoneID string) SdkConfigure {
	c.zoneID = zoneID
	return c
}

func (c *sdkConfigure) WithProjectID(projectID string) SdkConfigure {
	c.projectID = projectID
	return c
}

func (c *sdkConfigure) WithIamEndpoint(iamEndpoint string) SdkConfigure {
	c.iamEndpoint = normalizeURL(iamEndpoint)
	return c
}

func (c *sdkConfigure) WithVServerEndpoint(vserverEndpoint string) SdkConfigure {
	c.vserverEndpoint = normalizeURL(vserverEndpoint)
	return c
}

func (c *sdkConfigure) WithVLBEndpoint(vlbEndpoint string) SdkConfigure {
	c.vlbEndpoint = normalizeURL(vlbEndpoint)
	return c
}

func (c *sdkConfigure) WithVNetworkEndpoint(vnetworkEndpoint string) SdkConfigure {
	c.vnetworkEndpoint = normalizeURL(vnetworkEndpoint)
	return c
}

func (c *sdkConfigure) WithVDnsEndpoint(vdnsEndpoint string) SdkConfigure {
	c.vdnsEndpoint = normalizeURL(vdnsEndpoint)
	return c
}

func (c *sdkConfigure) WithGLBEndpoint(vlbEndpoint string) SdkConfigure {
	c.glbEndpoint = normalizeURL(vlbEndpoint)
	return c
}

func normalizeURL(u string) string {
	if !strings.HasSuffix(u, "/") {
		return u + "/"
	}
	return u
}
