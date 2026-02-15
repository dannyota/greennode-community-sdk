package client

import (
	"context"
	"time"

	svcclient "github.com/dannyota/greennode-community-sdk/v2/greennode/client"
	"github.com/dannyota/greennode-community-sdk/v2/greennode/gateway"
	identityv2 "github.com/dannyota/greennode-community-sdk/v2/greennode/services/identity/v2"
)

type Client struct {
	context    context.Context
	projectID  string
	zoneID     string
	userID     string
	httpClient svcclient.HTTPClient
	userAgent  string
	authOpt    svcclient.AuthOpts

	iamGateway      *gateway.IAMGateway
	vserverGateway  *gateway.VServerGateway
	vlbGateway      *gateway.VLBGateway
	vnetworkGateway *gateway.VNetworkGateway
	glbGateway      *gateway.GLBGateway
	vdnsGateway     *gateway.VDnsGateway
}

func NewClient(ctx context.Context) *Client {
	c := new(Client)
	c.context = ctx

	return c
}

func NewSdkConfigure() *SdkConfigure {
	return &SdkConfigure{}
}

func (c *Client) WithHTTPClient(client svcclient.HTTPClient) *Client {
	c.httpClient = client
	return c
}

func (c *Client) WithContext(ctx context.Context) *Client {
	c.context = ctx
	return c
}

func (c *Client) WithAuthOption(authOpts svcclient.AuthOpts, authConfig *SdkConfigure) *Client {
	if c.httpClient == nil {
		c.httpClient = svcclient.NewHTTPClient(c.context)
	}

	c.authOpt = authOpts // Assign the auth option to the client

	switch authOpts {
	case svcclient.IAMOauth2:
		c.httpClient.WithReauthFunc(svcclient.IAMOauth2, c.usingIAMOauth2AsAuthOption(authConfig)).
			WithKvDefaultHeaders("Content-Type", "application/json")
	default:
		c.httpClient.WithReauthFunc(svcclient.IAMOauth2, c.usingIAMOauth2AsAuthOption(authConfig)).
			WithKvDefaultHeaders("Content-Type", "application/json")
	}

	return c
}

func (c *Client) WithRetryCount(retry int) *Client {
	if c.httpClient == nil {
		c.httpClient = svcclient.NewHTTPClient(c.context)
	}

	c.httpClient.WithRetryCount(retry)
	return c
}

func (c *Client) WithKvDefaultHeaders(args ...string) *Client {
	if c.httpClient == nil {
		c.httpClient = svcclient.NewHTTPClient(c.context)
	}

	c.httpClient.WithKvDefaultHeaders(args...)
	return c
}

func (c *Client) WithSleep(sleep time.Duration) *Client {
	if c.httpClient == nil {
		c.httpClient = svcclient.NewHTTPClient(c.context)
	}

	c.httpClient.WithSleep(sleep)
	return c
}

func (c *Client) WithProjectID(projectID string) *Client {
	c.projectID = projectID
	if c.httpClient == nil {
		return c
	}

	// So it needs to reconfigure the gateway project id
	if c.vserverGateway != nil {
		c.vserverGateway = gateway.NewVServerGateway(c.vserverGateway.GetEndpoint(), c.projectID, c.httpClient)
	}

	if c.vlbGateway != nil {
		c.vlbGateway = gateway.NewVLBGateway(
			c.vlbGateway.GetEndpoint(),
			c.vserverGateway.GetEndpoint(),
			c.projectID,
			c.httpClient,
		)
	}

	if c.vnetworkGateway != nil {
		c.vnetworkGateway = gateway.NewVNetworkGateway(
			c.vnetworkGateway.GetEndpoint(),
			c.zoneID,
			c.projectID,
			c.userID,
			c.httpClient,
		)
	}

	if c.vdnsGateway != nil {
		c.vdnsGateway = gateway.NewVDnsGateway(c.vdnsGateway.GetEndpoint(), c.projectID, c.httpClient)
	}

	return c
}

func (c *Client) WithUserID(userID string) *Client {
	c.userID = userID
	if c.vnetworkGateway != nil {
		c.vnetworkGateway = gateway.NewVNetworkGateway(
			c.vnetworkGateway.GetEndpoint(),
			c.zoneID,
			c.projectID,
			c.userID,
			c.httpClient,
		)
	}

	return c
}

func (c *Client) Configure(sdkCfg *SdkConfigure) *Client {
	c.projectID = sdkCfg.GetProjectID()
	c.userID = sdkCfg.GetUserID()
	if c.httpClient == nil {
		c.httpClient = svcclient.NewHTTPClient(c.context)
	}

	if c.iamGateway == nil && sdkCfg.IAMEndpoint() != "" {
		c.iamGateway = gateway.NewIAMGateway(sdkCfg.IAMEndpoint(), c.projectID, c.httpClient)
	}

	if c.vserverGateway == nil && sdkCfg.VServerEndpoint() != "" {
		c.vserverGateway = gateway.NewVServerGateway(
			sdkCfg.VServerEndpoint(),
			c.projectID,
			c.httpClient,
		)
	}

	if c.vlbGateway == nil && sdkCfg.VLBEndpoint() != "" && sdkCfg.VServerEndpoint() != "" {
		c.vlbGateway = gateway.NewVLBGateway(
			sdkCfg.VLBEndpoint(),
			sdkCfg.VServerEndpoint(),
			c.projectID,
			c.httpClient,
		)
	}

	if c.vnetworkGateway == nil && sdkCfg.VNetworkEndpoint() != "" {
		c.vnetworkGateway = gateway.NewVNetworkGateway(
			sdkCfg.VNetworkEndpoint(),
			sdkCfg.GetZoneID(),
			c.projectID,
			c.userID,
			c.httpClient,
		)
	}

	if c.glbGateway == nil && sdkCfg.GLBEndpoint() != "" {
		c.glbGateway = gateway.NewGLBGateway(sdkCfg.GLBEndpoint(), c.httpClient)
	}

	if c.vdnsGateway == nil && sdkCfg.VDnsEndpoint() != "" {
		c.vdnsGateway = gateway.NewVDnsGateway(sdkCfg.VDnsEndpoint(), c.projectID, c.httpClient)
	}

	c.httpClient.WithReauthFunc(svcclient.IAMOauth2, c.usingIAMOauth2AsAuthOption(sdkCfg))
	c.userAgent = sdkCfg.UserAgent()
	if c.userAgent != "" {
		c.httpClient.WithKvDefaultHeaders("User-Agent", c.userAgent)
	}

	return c
}

func (c *Client) IAMGateway() *gateway.IAMGateway {
	return c.iamGateway
}

func (c *Client) VServerGateway() *gateway.VServerGateway {
	return c.vserverGateway
}

func (c *Client) VLBGateway() *gateway.VLBGateway {
	return c.vlbGateway
}

func (c *Client) VNetworkGateway() *gateway.VNetworkGateway {
	return c.vnetworkGateway
}

func (c *Client) GLBGateway() *gateway.GLBGateway {
	return c.glbGateway
}

func (c *Client) VDnsGateway() *gateway.VDnsGateway {
	return c.vdnsGateway
}

func (c *Client) usingIAMOauth2AsAuthOption(authConfig *SdkConfigure) func() (svcclient.SdkAuthentication, error) {
	authFunc := func() (svcclient.SdkAuthentication, error) {
		token, err := c.iamGateway.V2().IdentityService().GetAccessToken(
			identityv2.NewGetAccessTokenRequest(authConfig.GetClientID(), authConfig.GetClientSecret()))
		if err != nil {
			return nil, err
		}

		return token.ToSdkAuthentication(), nil
	}

	return authFunc
}

func (c *Client) UserAgent() string {
	return c.userAgent
}
