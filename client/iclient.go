package client

import (
	"context"
	"time"

	svcclient "github.com/dannyota/greennode-community-sdk/v2/greennode/client"
	"github.com/dannyota/greennode-community-sdk/v2/greennode/gateway"
)

type IClient interface {
	// List of builder methods
	WithHttpClient(phttpClient svcclient.IHttpClient) IClient
	WithContext(pctx context.Context) IClient
	WithAuthOption(pauthOption svcclient.AuthOpts, pauthConfig ISdkConfigure) IClient
	WithKvDefaultHeaders(pargs ...string) IClient
	WithRetryCount(pretry int) IClient
	WithSleep(psleep time.Duration) IClient
	WithProjectId(pprojectId string) IClient

	// List of functional methods
	Configure(psdkCfg ISdkConfigure) IClient
	GetUserAgent() string

	// List of gateways
	IamGateway() gateway.IIamGateway
	VServerGateway() gateway.IVServerGateway
	VLBGateway() gateway.IVLBGateway
	VBackUpGateway() gateway.IVBackUpGateway
	VNetworkGateway() gateway.IVNetworkGateway
	GLBGateway() gateway.IGLBGateway
	VDnsGateway() gateway.IVDnsGateway
}
