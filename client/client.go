package client

import (
	"context"
	"time"

	svcclient "github.com/dannyota/greennode-community-sdk/v2/greennode/client"
	"github.com/dannyota/greennode-community-sdk/v2/greennode/gateway"
	sdkerror "github.com/dannyota/greennode-community-sdk/v2/greennode/sdkerror"
	identityv2 "github.com/dannyota/greennode-community-sdk/v2/greennode/services/identity/v2"
)

type Client interface {
	// List of builder methods
	WithHTTPClient(httpClient svcclient.HTTPClient) Client
	WithContext(ctx context.Context) Client
	WithAuthOption(authOption svcclient.AuthOpts, authConfig SdkConfigure) Client
	WithKvDefaultHeaders(args ...string) Client
	WithRetryCount(retry int) Client
	WithSleep(sleep time.Duration) Client
	WithProjectID(projectID string) Client

	// List of functional methods
	Configure(sdkCfg SdkConfigure) Client
	UserAgent() string

	// List of gateways
	IAMGateway() gateway.IAMGateway
	VServerGateway() gateway.VServerGateway
	VLBGateway() gateway.VLBGateway
	VBackUpGateway() gateway.VBackUpGateway
	VNetworkGateway() gateway.VNetworkGateway
	GLBGateway() gateway.GLBGateway
	VDnsGateway() gateway.VDnsGateway
}

var (
	_ Client = new(client)
)

type (
	client struct {
		context    context.Context
		projectID  string
		zoneID     string
		userID     string
		httpClient svcclient.HTTPClient
		userAgent  string
		authOpt    svcclient.AuthOpts

		iamGateway      gateway.IAMGateway
		vserverGateway  gateway.VServerGateway
		vlbGateway      gateway.VLBGateway
		vbackupGateway  gateway.VBackUpGateway
		vnetworkGateway gateway.VNetworkGateway
		glbGateway      gateway.GLBGateway
		vdnsGateway     gateway.VDnsGateway
	}
)

func NewClient(ctx context.Context) Client {
	c := new(client)
	c.context = ctx

	return c
}

func NewSdkConfigure() SdkConfigure {
	return &sdkConfigure{}
}

func (c *client) WithHTTPClient(client svcclient.HTTPClient) Client {
	c.httpClient = client
	return c
}

func (c *client) WithContext(ctx context.Context) Client {
	c.context = ctx
	return c
}

func (c *client) WithAuthOption(authOpts svcclient.AuthOpts, authConfig SdkConfigure) Client {
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

func (c *client) WithRetryCount(retry int) Client {
	if c.httpClient == nil {
		c.httpClient = svcclient.NewHTTPClient(c.context)
	}

	c.httpClient.WithRetryCount(retry)
	return c
}

func (c *client) WithKvDefaultHeaders(args ...string) Client {
	if c.httpClient == nil {
		c.httpClient = svcclient.NewHTTPClient(c.context)
	}

	c.httpClient.WithKvDefaultHeaders(args...)
	return c
}

func (c *client) WithSleep(sleep time.Duration) Client {
	if c.httpClient == nil {
		c.httpClient = svcclient.NewHTTPClient(c.context)
	}

	c.httpClient.WithSleep(sleep)
	return c
}

func (c *client) WithProjectID(projectID string) Client {
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

func (c *client) WithUserID(userID string) Client {
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

func (c *client) Configure(sdkCfg SdkConfigure) Client {
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

	return c
}

func (c *client) IAMGateway() gateway.IAMGateway {
	return c.iamGateway
}

func (c *client) VServerGateway() gateway.VServerGateway {
	return c.vserverGateway
}

func (c *client) VLBGateway() gateway.VLBGateway {
	return c.vlbGateway
}

func (c *client) VBackUpGateway() gateway.VBackUpGateway {
	return c.vbackupGateway
}

func (c *client) VNetworkGateway() gateway.VNetworkGateway {
	return c.vnetworkGateway
}

func (c *client) GLBGateway() gateway.GLBGateway {
	return c.glbGateway
}

func (c *client) VDnsGateway() gateway.VDnsGateway {
	return c.vdnsGateway
}

func (c *client) usingIAMOauth2AsAuthOption(authConfig SdkConfigure) func() (svcclient.SdkAuthentication, sdkerror.Error) {
	authFunc := func() (svcclient.SdkAuthentication, sdkerror.Error) {
		token, err := c.iamGateway.V2().IdentityService().GetAccessToken(
			identityv2.NewGetAccessTokenRequest(authConfig.GetClientID(), authConfig.GetClientSecret()))
		if err != nil {
			return nil, err
		}

		return token.ToSdkAuthentication(), nil
	}

	return authFunc
}

func (c *client) UserAgent() string {
	return c.userAgent
}
