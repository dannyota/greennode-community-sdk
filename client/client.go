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
		projectId  string
		zoneId     string
		userId     string
		httpClient svcclient.HttpClient
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

func (s *client) WithHttpClient(client svcclient.HttpClient) Client {
	s.httpClient = client
	return s
}

func (s *client) WithContext(ctx context.Context) Client {
	s.context = ctx
	return s
}

func (s *client) WithAuthOption(authOpts svcclient.AuthOpts, authConfig SdkConfigure) Client {
	if s.httpClient == nil {
		s.httpClient = svcclient.NewHttpClient(s.context)
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
		s.httpClient = svcclient.NewHttpClient(s.context)
	}

	s.httpClient.WithRetryCount(retry)
	return s
}

func (s *client) WithKvDefaultHeaders(args ...string) Client {
	if s.httpClient == nil {
		s.httpClient = svcclient.NewHttpClient(s.context)
	}

	s.httpClient.WithKvDefaultHeaders(args...)
	return s
}

func (s *client) WithSleep(sleep time.Duration) Client {
	if s.httpClient == nil {
		s.httpClient = svcclient.NewHttpClient(s.context)
	}

	s.httpClient.WithSleep(sleep)
	return s
}

func (s *client) WithProjectId(projectId string) Client {
	s.projectId = projectId
	if s.httpClient == nil {
		return s
	}

	// So it needs to reconfigure the gateway project id
	if s.vserverGateway != nil {
		s.vserverGateway = gateway.NewVServerGateway(s.vserverGateway.GetEndpoint(), s.projectId, s.httpClient)
	}

	if s.vlbGateway != nil {
		s.vlbGateway = gateway.NewVLBGateway(
			s.vlbGateway.GetEndpoint(),
			s.vserverGateway.GetEndpoint(),
			s.projectId,
			s.httpClient,
		)
	}

	if s.vnetworkGateway != nil {
		s.vnetworkGateway = gateway.NewVNetworkGateway(
			s.vnetworkGateway.GetEndpoint(),
			s.zoneId,
			s.projectId,
			s.userId,
			s.httpClient,
		)
	}

	if s.vdnsGateway != nil {
		s.vdnsGateway = gateway.NewVDnsGateway(s.vdnsGateway.GetEndpoint(), s.projectId, s.httpClient)
	}

	return s
}

func (s *client) WithUserId(userId string) Client {
	s.userId = userId
	if s.vnetworkGateway != nil {
		s.vnetworkGateway = gateway.NewVNetworkGateway(
			s.vnetworkGateway.GetEndpoint(),
			s.zoneId,
			s.projectId,
			s.userId,
			s.httpClient,
		)
	}

	return s
}

func (s *client) Configure(sdkCfg SdkConfigure) Client {
	s.projectId = sdkCfg.GetProjectId()
	s.userId = sdkCfg.GetUserId()
	if s.httpClient == nil {
		s.httpClient = svcclient.NewHttpClient(s.context)
	}

	if s.iamGateway == nil && sdkCfg.GetIamEndpoint() != "" {
		s.iamGateway = gateway.NewIamGateway(sdkCfg.GetIamEndpoint(), s.projectId, s.httpClient)
	}

	if s.vserverGateway == nil && sdkCfg.GetVServerEndpoint() != "" {
		s.vserverGateway = gateway.NewVServerGateway(
			sdkCfg.GetVServerEndpoint(),
			s.projectId,
			s.httpClient,
		)
	}

	if s.vlbGateway == nil && sdkCfg.GetVLBEndpoint() != "" && sdkCfg.GetVServerEndpoint() != "" {
		s.vlbGateway = gateway.NewVLBGateway(
			sdkCfg.GetVLBEndpoint(),
			sdkCfg.GetVServerEndpoint(),
			s.projectId,
			s.httpClient,
		)
	}

	if s.vnetworkGateway == nil && sdkCfg.GetVNetworkEndpoint() != "" {
		s.vnetworkGateway = gateway.NewVNetworkGateway(
			sdkCfg.GetVNetworkEndpoint(),
			sdkCfg.GetZoneId(),
			s.projectId,
			s.userId,
			s.httpClient,
		)
	}

	if s.glbGateway == nil && sdkCfg.GetGLBEndpoint() != "" {
		s.glbGateway = gateway.NewGLBGateway(sdkCfg.GetGLBEndpoint(), s.httpClient)
	}

	if s.vdnsGateway == nil && sdkCfg.GetVDnsEndpoint() != "" {
		s.vdnsGateway = gateway.NewVDnsGateway(sdkCfg.GetVDnsEndpoint(), s.projectId, s.httpClient)
	}

	s.httpClient.WithReauthFunc(svcclient.IamOauth2, s.usingIamOauth2AsAuthOption(sdkCfg))
	s.userAgent = sdkCfg.GetUserAgent()

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
			identityv2.NewGetAccessTokenRequest(authConfig.GetClientId(), authConfig.GetClientSecret()))
		if err != nil {
			return nil, err
		}

		return token.ToSdkAuthentication(), nil
	}

	return authFunc
}

func (s *client) GetUserAgent() string {
	return s.userAgent
}
