package client

import "strings"

type SdkConfigure struct {
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

func (c *SdkConfigure) GetClientID() string {
	return c.clientID
}

func (c *SdkConfigure) GetClientSecret() string {
	return c.clientSecret
}

func (c *SdkConfigure) GetProjectID() string {
	return c.projectID
}

func (c *SdkConfigure) GetUserID() string {
	return c.userID
}

func (c *SdkConfigure) GetZoneID() string {
	return c.zoneID
}

func (c *SdkConfigure) IAMEndpoint() string {
	return c.iamEndpoint
}

func (c *SdkConfigure) VServerEndpoint() string {
	return c.vserverEndpoint
}

func (c *SdkConfigure) VLBEndpoint() string {
	return c.vlbEndpoint
}

func (c *SdkConfigure) GLBEndpoint() string {
	return c.glbEndpoint
}

func (c *SdkConfigure) VNetworkEndpoint() string {
	return c.vnetworkEndpoint
}

func (c *SdkConfigure) VDnsEndpoint() string {
	return c.vdnsEndpoint
}

func (c *SdkConfigure) UserAgent() string {
	return c.userAgent
}

func (c *SdkConfigure) WithUserAgent(userAgent string) *SdkConfigure {
	c.userAgent = userAgent
	return c
}

func (c *SdkConfigure) WithClientID(clientID string) *SdkConfigure {
	c.clientID = clientID
	return c
}

func (c *SdkConfigure) WithClientSecret(clientSecret string) *SdkConfigure {
	c.clientSecret = clientSecret
	return c
}

func (c *SdkConfigure) WithUserID(userID string) *SdkConfigure {
	c.userID = userID
	return c
}

func (c *SdkConfigure) WithZoneID(zoneID string) *SdkConfigure {
	c.zoneID = zoneID
	return c
}

func (c *SdkConfigure) WithProjectID(projectID string) *SdkConfigure {
	c.projectID = projectID
	return c
}

func (c *SdkConfigure) WithIAMEndpoint(iamEndpoint string) *SdkConfigure {
	c.iamEndpoint = normalizeURL(iamEndpoint)
	return c
}

func (c *SdkConfigure) WithVServerEndpoint(vserverEndpoint string) *SdkConfigure {
	c.vserverEndpoint = normalizeURL(vserverEndpoint)
	return c
}

func (c *SdkConfigure) WithVLBEndpoint(vlbEndpoint string) *SdkConfigure {
	c.vlbEndpoint = normalizeURL(vlbEndpoint)
	return c
}

func (c *SdkConfigure) WithVNetworkEndpoint(vnetworkEndpoint string) *SdkConfigure {
	c.vnetworkEndpoint = normalizeURL(vnetworkEndpoint)
	return c
}

func (c *SdkConfigure) WithVDnsEndpoint(vdnsEndpoint string) *SdkConfigure {
	c.vdnsEndpoint = normalizeURL(vdnsEndpoint)
	return c
}

func (c *SdkConfigure) WithGLBEndpoint(vlbEndpoint string) *SdkConfigure {
	c.glbEndpoint = normalizeURL(vlbEndpoint)
	return c
}

func normalizeURL(u string) string {
	if !strings.HasSuffix(u, "/") {
		return u + "/"
	}
	return u
}
