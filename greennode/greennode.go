package greennode

import (
	"context"
	"time"

	"github.com/dannyota/greennode-community-sdk/v2/greennode/auth"
	"github.com/dannyota/greennode-community-sdk/v2/greennode/client"
	computev2 "github.com/dannyota/greennode-community-sdk/v2/greennode/services/compute/v2"
	dnsv1 "github.com/dannyota/greennode-community-sdk/v2/greennode/services/dns/v1"
	glbv1 "github.com/dannyota/greennode-community-sdk/v2/greennode/services/glb/v1"
	identityv2 "github.com/dannyota/greennode-community-sdk/v2/greennode/services/identity/v2"
	lbintervpc "github.com/dannyota/greennode-community-sdk/v2/greennode/services/loadbalancer/intervpc"
	lbv2 "github.com/dannyota/greennode-community-sdk/v2/greennode/services/loadbalancer/v2"
	networkv1 "github.com/dannyota/greennode-community-sdk/v2/greennode/services/network/v1"
	networkv2 "github.com/dannyota/greennode-community-sdk/v2/greennode/services/network/v2"
	portalv1 "github.com/dannyota/greennode-community-sdk/v2/greennode/services/portal/v1"
	portalv2 "github.com/dannyota/greennode-community-sdk/v2/greennode/services/portal/v2"
	serverv1 "github.com/dannyota/greennode-community-sdk/v2/greennode/services/server/v1"
	volumev1 "github.com/dannyota/greennode-community-sdk/v2/greennode/services/volume/v1"
	volumev2 "github.com/dannyota/greennode-community-sdk/v2/greennode/services/volume/v2"
)

// Config holds all configuration needed to create an SDK client.
type Config struct {
	Region       string // e.g. "hcm-3", "han-1" — derives default endpoint URLs
	ClientID     string
	ClientSecret string
	ProjectID    string
	UserID       string
	ZoneID       string
	UserAgent    string
	IAMAuth      *auth.IAMUserAuth // optional — enables IAM user auth instead of client credentials

	RetryCount    int
	SleepDuration time.Duration

	IAMEndpoint      string
	VServerEndpoint  string
	VLBEndpoint      string
	VNetworkEndpoint string
	GLBEndpoint      string
	DNSEndpoint      string
}

// Client provides flat access to all service APIs.
type Client struct {
	// Primary services (latest API version)
	LoadBalancer *lbv2.LoadBalancerServiceV2
	Compute      *computev2.ComputeServiceV2
	Network      *networkv2.NetworkServiceV2
	Volume       *volumev2.VolumeServiceV2
	DNS          *dnsv1.VDnsServiceV1
	GLB          *glbv1.GLBServiceV1
	Portal       *portalv2.PortalServiceV2
	Identity     *identityv2.IdentityServiceV2

	// Legacy / internal API versions
	NetworkV1            *networkv1.NetworkServiceV1
	NetworkAZ            *networkv2.NetworkServiceV2
	NetworkInternal      *networkv1.NetworkServiceInternalV1
	VolumeV1             *volumev1.VolumeServiceV1
	PortalV1             *portalv1.PortalServiceV1
	LoadBalancerInternal *lbintervpc.LoadBalancerServiceInternal
	ServerInternal       *serverv1.ServerServiceInternalV1
	DNSInternal          *dnsv1.VDnsServiceInternal
}

// NewClient creates a fully-wired SDK client from the given configuration.
func NewClient(ctx context.Context, cfg Config) (*Client, error) {
	resolveEndpoints(&cfg)

	hc := client.NewHTTPClient()

	if cfg.RetryCount > 0 {
		hc.WithRetryCount(cfg.RetryCount)
	}
	if cfg.SleepDuration > 0 {
		hc.WithSleep(cfg.SleepDuration)
	}
	hc.WithKvDefaultHeaders("Content-Type", "application/json")
	if cfg.UserAgent != "" {
		hc.WithKvDefaultHeaders("User-Agent", cfg.UserAgent)
	}

	c := &Client{}

	// IAM / Identity
	if cfg.IAMEndpoint != "" {
		iamSvc := newServiceClient(cfg.IAMEndpoint+"v2", cfg.ProjectID, "", hc)
		c.Identity = &identityv2.IdentityServiceV2{Client: iamSvc}
	}

	// VServer services (compute, network, volume, portal)
	if cfg.VServerEndpoint != "" {
		ep := client.NormalizeURL(cfg.VServerEndpoint)

		svcV2 := newServiceClient(ep+"v2", cfg.ProjectID, "", hc)
		c.Compute = &computev2.ComputeServiceV2{Client: svcV2}
		c.Network = &networkv2.NetworkServiceV2{Client: svcV2}
		c.Volume = &volumev2.VolumeServiceV2{Client: svcV2}
		c.Portal = &portalv2.PortalServiceV2{Client: svcV2}

		svcV1 := newServiceClient(ep+"v1", cfg.ProjectID, "", hc)
		c.VolumeV1 = &volumev1.VolumeServiceV1{Client: svcV1}
		c.PortalV1 = &portalv1.PortalServiceV1{Client: svcV1}

		svcInternal := newServiceClient(ep+"internal", cfg.ProjectID, "", hc)
		c.ServerInternal = &serverv1.ServerServiceInternalV1{Client: svcInternal}
	}

	// VLB (load balancer)
	if cfg.VLBEndpoint != "" {
		vlbEp := client.NormalizeURL(cfg.VLBEndpoint)

		vlbSvcV2 := newServiceClient(vlbEp+"v2", cfg.ProjectID, "", hc)
		var vserverSvcV2 *client.ServiceClient
		if cfg.VServerEndpoint != "" {
			vserverSvcV2 = newServiceClient(client.NormalizeURL(cfg.VServerEndpoint)+"v2", cfg.ProjectID, "", hc)
		}
		c.LoadBalancer = &lbv2.LoadBalancerServiceV2{Client: vlbSvcV2, ServerClient: vserverSvcV2}

		vlbSvcInternal := newServiceClient(vlbEp+"internal", cfg.ProjectID, "", hc)
		c.LoadBalancerInternal = &lbintervpc.LoadBalancerServiceInternal{Client: vlbSvcInternal}
	}

	// VNetwork
	if cfg.VNetworkEndpoint != "" {
		vnEp := client.NormalizeURL(cfg.VNetworkEndpoint)

		vnetV1 := newServiceClient(vnEp+"vnetwork/v1", cfg.ProjectID, cfg.ZoneID, hc)
		c.NetworkV1 = &networkv1.NetworkServiceV1{Client: vnetV1}

		vnetAZ := newServiceClient(vnEp+"vnetwork/az/v1", cfg.ProjectID, cfg.ZoneID, hc)
		c.NetworkAZ = &networkv2.NetworkServiceV2{Client: vnetAZ}

		vnetInternal := newServiceClient(vnEp+"internal/v1", cfg.ProjectID, cfg.ZoneID, hc)
		c.NetworkInternal = &networkv1.NetworkServiceInternalV1{Client: vnetInternal}
	}

	// GLB (global load balancer)
	if cfg.GLBEndpoint != "" {
		glbSvc := newServiceClient(cfg.GLBEndpoint+"v1", "", "", hc)
		c.GLB = &glbv1.GLBServiceV1{Client: glbSvc}
	}

	// DNS
	if cfg.DNSEndpoint != "" {
		dnsEp := client.NormalizeURL(cfg.DNSEndpoint)

		dnsSvc := newServiceClient(dnsEp+"v1", cfg.ProjectID, "", hc)
		c.DNS = &dnsv1.VDnsServiceV1{Client: dnsSvc}

		dnsInternalSvc := newServiceClient(dnsEp+"internal/v1", cfg.ProjectID, "", hc)
		c.DNSInternal = &dnsv1.VDnsServiceInternal{Client: dnsInternalSvc}
	}

	// Set up auth
	if cfg.IAMAuth != nil {
		hc.WithReauthFunc(client.IAMUserOauth2, iamUserReauthFunc(cfg.IAMAuth))
	} else if c.Identity != nil && cfg.ClientID != "" && cfg.ClientSecret != "" {
		hc.WithReauthFunc(client.IAMOauth2, reauthFunc(c.Identity, cfg.ClientID, cfg.ClientSecret))
	}

	return c, nil
}

func reauthFunc(identity *identityv2.IdentityServiceV2, clientID, clientSecret string) func(ctx context.Context) (*client.Token, error) {
	return func(ctx context.Context) (*client.Token, error) {
		token, err := identity.GetAccessToken(ctx, identityv2.NewGetAccessTokenRequest(clientID, clientSecret))
		if err != nil {
			return nil, err
		}
		return &client.Token{
			AccessToken: token.Token,
			ExpiresAt:   token.ExpiresAt,
		}, nil
	}
}

func iamUserReauthFunc(iamAuth *auth.IAMUserAuth) func(ctx context.Context) (*client.Token, error) {
	return func(ctx context.Context) (*client.Token, error) {
		token, expiresAt, err := iamAuth.Authenticate(ctx)
		if err != nil {
			return nil, err
		}
		return &client.Token{
			AccessToken: token,
			ExpiresAt:   expiresAt,
		}, nil
	}
}

func newServiceClient(endpoint, projectID, zoneID string, hc *client.HTTPClient) *client.ServiceClient {
	sc := client.NewServiceClient().
		WithEndpoint(endpoint).
		WithClient(hc)
	if projectID != "" {
		sc.WithProjectID(projectID)
	}
	if zoneID != "" {
		sc.WithZoneID(zoneID)
	}
	return sc
}
