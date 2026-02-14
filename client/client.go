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
	_ IClient = new(client)
)

type (
	client struct {
		context    context.Context
		projectId  string
		zoneId     string
		userId     string
		httpClient svcclient.IHttpClient
		userAgent  string
		authOpt    svcclient.AuthOpts

		iamGateway      gateway.IIamGateway
		vserverGateway  gateway.IVServerGateway
		vlbGateway      gateway.IVLBGateway
		vbackupGateway  gateway.IVBackUpGateway
		vnetworkGateway gateway.IVNetworkGateway
		glbGateway      gateway.IGLBGateway
		vdnsGateway     gateway.IVDnsGateway
	}
)

func NewClient(pctx context.Context) IClient {
	c := new(client)
	c.context = pctx

	return c
}

func NewSdkConfigure() ISdkConfigure {
	return &sdkConfigure{}
}

func (s *client) WithHttpClient(pclient svcclient.IHttpClient) IClient {
	s.httpClient = pclient
	return s
}

func (s *client) WithContext(pctx context.Context) IClient {
	s.context = pctx
	return s
}

func (s *client) WithAuthOption(pauthOpts svcclient.AuthOpts, pauthConfig ISdkConfigure) IClient {
	if s.httpClient == nil {
		s.httpClient = svcclient.NewHttpClient(s.context)
	}

	s.authOpt = pauthOpts // Assign the auth option to the client

	switch pauthOpts {
	case svcclient.IamOauth2:
		s.httpClient.WithReauthFunc(svcclient.IamOauth2, s.usingIamOauth2AsAuthOption(pauthConfig)).
			WithKvDefaultHeaders("Content-Type", "application/json")
	default:
		s.httpClient.WithReauthFunc(svcclient.IamOauth2, s.usingIamOauth2AsAuthOption(pauthConfig)).
			WithKvDefaultHeaders("Content-Type", "application/json")
	}

	return s
}

func (s *client) WithRetryCount(pretry int) IClient {
	if s.httpClient == nil {
		s.httpClient = svcclient.NewHttpClient(s.context)
	}

	s.httpClient.WithRetryCount(pretry)
	return s
}

func (s *client) WithKvDefaultHeaders(pargs ...string) IClient {
	if s.httpClient == nil {
		s.httpClient = svcclient.NewHttpClient(s.context)
	}

	s.httpClient.WithKvDefaultHeaders(pargs...)
	return s
}

func (s *client) WithSleep(psleep time.Duration) IClient {
	if s.httpClient == nil {
		s.httpClient = svcclient.NewHttpClient(s.context)
	}

	s.httpClient.WithSleep(psleep)
	return s
}

func (s *client) WithProjectId(pprojectId string) IClient {
	s.projectId = pprojectId
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

func (s *client) WithUserId(puserId string) IClient {
	s.userId = puserId
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

func (s *client) Configure(psdkCfg ISdkConfigure) IClient {
	s.projectId = psdkCfg.GetProjectId()
	s.userId = psdkCfg.GetUserId()
	if s.httpClient == nil {
		s.httpClient = svcclient.NewHttpClient(s.context)
	}

	if s.iamGateway == nil && psdkCfg.GetIamEndpoint() != "" {
		s.iamGateway = gateway.NewIamGateway(psdkCfg.GetIamEndpoint(), s.projectId, s.httpClient)
	}

	if s.vserverGateway == nil && psdkCfg.GetVServerEndpoint() != "" {
		s.vserverGateway = gateway.NewVServerGateway(
			psdkCfg.GetVServerEndpoint(),
			s.projectId,
			s.httpClient,
		)
	}

	if s.vlbGateway == nil && psdkCfg.GetVLBEndpoint() != "" && psdkCfg.GetVServerEndpoint() != "" {
		s.vlbGateway = gateway.NewVLBGateway(
			psdkCfg.GetVLBEndpoint(),
			psdkCfg.GetVServerEndpoint(),
			s.projectId,
			s.httpClient,
		)
	}

	if s.vnetworkGateway == nil && psdkCfg.GetVNetworkEndpoint() != "" {
		s.vnetworkGateway = gateway.NewVNetworkGateway(
			psdkCfg.GetVNetworkEndpoint(),
			psdkCfg.GetZoneId(),
			s.projectId,
			s.userId,
			s.httpClient,
		)
	}

	if s.glbGateway == nil && psdkCfg.GetGLBEndpoint() != "" {
		s.glbGateway = gateway.NewGLBGateway(psdkCfg.GetGLBEndpoint(), s.httpClient)
	}

	if s.vdnsGateway == nil && psdkCfg.GetVDnsEndpoint() != "" {
		s.vdnsGateway = gateway.NewVDnsGateway(psdkCfg.GetVDnsEndpoint(), s.projectId, s.httpClient)
	}

	s.httpClient.WithReauthFunc(svcclient.IamOauth2, s.usingIamOauth2AsAuthOption(psdkCfg))
	s.userAgent = psdkCfg.GetUserAgent()

	return s
}

func (s *client) IamGateway() gateway.IIamGateway {
	return s.iamGateway
}

func (s *client) VServerGateway() gateway.IVServerGateway {
	return s.vserverGateway
}

func (s *client) VLBGateway() gateway.IVLBGateway {
	return s.vlbGateway
}

func (s *client) VBackUpGateway() gateway.IVBackUpGateway {
	return s.vbackupGateway
}

func (s *client) VNetworkGateway() gateway.IVNetworkGateway {
	return s.vnetworkGateway
}

func (s *client) GLBGateway() gateway.IGLBGateway {
	return s.glbGateway
}

func (s *client) VDnsGateway() gateway.IVDnsGateway {
	return s.vdnsGateway
}

func (s *client) usingIamOauth2AsAuthOption(pauthConfig ISdkConfigure) func() (svcclient.ISdkAuthentication, sdkerror.IError) {
	authFunc := func() (svcclient.ISdkAuthentication, sdkerror.IError) {
		token, err := s.iamGateway.V2().IdentityService().GetAccessToken(
			identityv2.NewGetAccessTokenRequest(pauthConfig.GetClientId(), pauthConfig.GetClientSecret()))
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
