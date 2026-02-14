package client

import (
	"context"
	"time"

	svcclient "github.com/dannyota/greennode-community-sdk/v2/greennode/client"
	"github.com/dannyota/greennode-community-sdk/v2/greennode/gateway"
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
