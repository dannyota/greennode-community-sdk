package client

import (
	"context"
	"time"

	svcclient "github.com/dannyota/greennode-community-sdk/v2/greennode/client"
	"github.com/dannyota/greennode-community-sdk/v2/greennode/gateway"
	sdkerror "github.com/dannyota/greennode-community-sdk/v2/greennode/sdkerror"
	identityv2 "github.com/dannyota/greennode-community-sdk/v2/greennode/services/identity/v2"
)

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

		iamGateway      gateway.IamGateway
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

func (s *client) WithHTTPClient(client svcclient.HTTPClient) Client {
	s.httpClient = client
	return s
}

func (s *client) WithContext(ctx context.Context) Client {
	s.context = ctx
	return s
}

func (s *client) WithAuthOption(authOpts svcclient.AuthOpts, authConfig SdkConfigure) Client {
	if s.httpClient == nil {
		s.httpClient = svcclient.NewHTTPClient(s.context)
	}

	s.authOpt = authOpts // Assign the auth option to the client

	switch authOpts {
	case svcclient.IamOauth2:
		s.httpClient.WithReauthFunc(svcclient.IamOauth2, s.usingIamOauth2AsAuthOption(authConfig)).
			WithKvDefaultHeaders("Content-Type", "application/json")
	default:
		s.httpClient.WithReauthFunc(svcclient.IamOauth2, s.usingIamOauth2AsAuthOption(authConfig)).
			WithKvDefaultHeaders("Content-Type", "application/json")
	}

	return s
}

func (s *client) WithRetryCount(retry int) Client {
	if s.httpClient == nil {
		s.httpClient = svcclient.NewHTTPClient(s.context)
	}

	s.httpClient.WithRetryCount(retry)
	return s
}

func (s *client) WithKvDefaultHeaders(args ...string) Client {
	if s.httpClient == nil {
		s.httpClient = svcclient.NewHTTPClient(s.context)
	}

	s.httpClient.WithKvDefaultHeaders(args...)
	return s
}

func (s *client) WithSleep(sleep time.Duration) Client {
	if s.httpClient == nil {
		s.httpClient = svcclient.NewHTTPClient(s.context)
	}

	s.httpClient.WithSleep(sleep)
	return s
}

func (s *client) WithProjectID(projectID string) Client {
	s.projectID = projectID
	if s.httpClient == nil {
		return s
	}

	// So it needs to reconfigure the gateway project id
	if s.vserverGateway != nil {
		s.vserverGateway = gateway.NewVServerGateway(s.vserverGateway.GetEndpoint(), s.projectID, s.httpClient)
	}

	if s.vlbGateway != nil {
		s.vlbGateway = gateway.NewVLBGateway(
			s.vlbGateway.GetEndpoint(),
			s.vserverGateway.GetEndpoint(),
			s.projectID,
			s.httpClient,
		)
	}

	if s.vnetworkGateway != nil {
		s.vnetworkGateway = gateway.NewVNetworkGateway(
			s.vnetworkGateway.GetEndpoint(),
			s.zoneID,
			s.projectID,
			s.userID,
			s.httpClient,
		)
	}

	if s.vdnsGateway != nil {
		s.vdnsGateway = gateway.NewVDnsGateway(s.vdnsGateway.GetEndpoint(), s.projectID, s.httpClient)
	}

	return s
}

func (s *client) WithUserID(userID string) Client {
	s.userID = userID
	if s.vnetworkGateway != nil {
		s.vnetworkGateway = gateway.NewVNetworkGateway(
			s.vnetworkGateway.GetEndpoint(),
			s.zoneID,
			s.projectID,
			s.userID,
			s.httpClient,
		)
	}

	return s
}

func (s *client) Configure(sdkCfg SdkConfigure) Client {
	s.projectID = sdkCfg.GetProjectID()
	s.userID = sdkCfg.GetUserID()
	if s.httpClient == nil {
		s.httpClient = svcclient.NewHTTPClient(s.context)
	}

	if s.iamGateway == nil && sdkCfg.IamEndpoint() != "" {
		s.iamGateway = gateway.NewIamGateway(sdkCfg.IamEndpoint(), s.projectID, s.httpClient)
	}

	if s.vserverGateway == nil && sdkCfg.VServerEndpoint() != "" {
		s.vserverGateway = gateway.NewVServerGateway(
			sdkCfg.VServerEndpoint(),
			s.projectID,
			s.httpClient,
		)
	}

	if s.vlbGateway == nil && sdkCfg.VLBEndpoint() != "" && sdkCfg.VServerEndpoint() != "" {
		s.vlbGateway = gateway.NewVLBGateway(
			sdkCfg.VLBEndpoint(),
			sdkCfg.VServerEndpoint(),
			s.projectID,
			s.httpClient,
		)
	}

	if s.vnetworkGateway == nil && sdkCfg.VNetworkEndpoint() != "" {
		s.vnetworkGateway = gateway.NewVNetworkGateway(
			sdkCfg.VNetworkEndpoint(),
			sdkCfg.GetZoneID(),
			s.projectID,
			s.userID,
			s.httpClient,
		)
	}

	if s.glbGateway == nil && sdkCfg.GLBEndpoint() != "" {
		s.glbGateway = gateway.NewGLBGateway(sdkCfg.GLBEndpoint(), s.httpClient)
	}

	if s.vdnsGateway == nil && sdkCfg.VDnsEndpoint() != "" {
		s.vdnsGateway = gateway.NewVDnsGateway(sdkCfg.VDnsEndpoint(), s.projectID, s.httpClient)
	}

	s.httpClient.WithReauthFunc(svcclient.IamOauth2, s.usingIamOauth2AsAuthOption(sdkCfg))
	s.userAgent = sdkCfg.UserAgent()

	return s
}

func (s *client) IamGateway() gateway.IamGateway {
	return s.iamGateway
}

func (s *client) VServerGateway() gateway.VServerGateway {
	return s.vserverGateway
}

func (s *client) VLBGateway() gateway.VLBGateway {
	return s.vlbGateway
}

func (s *client) VBackUpGateway() gateway.VBackUpGateway {
	return s.vbackupGateway
}

func (s *client) VNetworkGateway() gateway.VNetworkGateway {
	return s.vnetworkGateway
}

func (s *client) GLBGateway() gateway.GLBGateway {
	return s.glbGateway
}

func (s *client) VDnsGateway() gateway.VDnsGateway {
	return s.vdnsGateway
}

func (s *client) usingIamOauth2AsAuthOption(authConfig SdkConfigure) func() (svcclient.SdkAuthentication, sdkerror.Error) {
	authFunc := func() (svcclient.SdkAuthentication, sdkerror.Error) {
		token, err := s.iamGateway.V2().IdentityService().GetAccessToken(
			identityv2.NewGetAccessTokenRequest(authConfig.GetClientID(), authConfig.GetClientSecret()))
		if err != nil {
			return nil, err
		}

		return token.ToSdkAuthentication(), nil
	}

	return authFunc
}

func (s *client) UserAgent() string {
	return s.userAgent
}
