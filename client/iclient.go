package client

import (
	"context"
	"time"

	svcclient "github.com/dannyota/greennode-community-sdk/v2/greennode/client"
	"github.com/dannyota/greennode-community-sdk/v2/greennode/gateway"
)

type Client interface {
	// List of builder methods
	WithHttpClient(phttpClient svcclient.HttpClient) Client
	WithContext(pctx context.Context) Client
	WithAuthOption(pauthOption svcclient.AuthOpts, pauthConfig SdkConfigure) Client
	WithKvDefaultHeaders(pargs ...string) Client
	WithRetryCount(pretry int) Client
	WithSleep(psleep time.Duration) Client
	WithProjectId(pprojectId string) Client

	// List of functional methods
	Configure(psdkCfg SdkConfigure) Client
	GetUserAgent() string

	// List of gateways
	IamGateway() gateway.IamGateway
	VServerGateway() gateway.VServerGateway
	VLBGateway() gateway.VLBGateway
	VBackUpGateway() gateway.VBackUpGateway
	VNetworkGateway() gateway.VNetworkGateway
	GLBGateway() gateway.GLBGateway
	VDnsGateway() gateway.VDnsGateway
}
