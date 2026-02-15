package gateway

import (
	"github.com/dannyota/greennode-community-sdk/v2/greennode/client"
	"github.com/dannyota/greennode-community-sdk/v2/greennode/services/dns"
	"github.com/dannyota/greennode-community-sdk/v2/greennode/services/glb"
)

type IAMGateway struct {
	iamGatewayV2 *IAMGatewayV2
}

type VServerGateway struct {
	endpoint                 string
	vserverGatewayV1         *VServerGatewayV1
	vserverGatewayV2         *VServerGatewayV2
	vserverGatewayInternalV1 *VServerGatewayInternalV1
}

type VLBGateway struct {
	endpoint           string
	vlbGatewayInternal *VLBGatewayInternal
	vlbGatewayV2       *VLBGatewayV2
}

type VNetworkGateway struct {
	endpoint                  string
	vnetworkGatewayV1         *VNetworkGatewayV1
	vnetworkGatewayV2         *VNetworkGatewayV2
	vnetworkGatewayInternalV1 *VNetworkGatewayInternalV1
}

func NewIAMGateway(endpoint, projectID string, hc client.HTTPClient) *IAMGateway {
	iamSvcV2 := client.NewServiceClient().
		WithEndpoint(endpoint + "v2").
		WithClient(hc).
		WithProjectID(projectID)

	return &IAMGateway{
		iamGatewayV2: NewIAMGatewayV2(iamSvcV2),
	}
}

func NewVServerGateway(endpoint, projectID string, hc client.HTTPClient) *VServerGateway {
	vserverSvcV1 := client.NewServiceClient().
		WithEndpoint(endpoint + "v1").
		WithClient(hc).
		WithProjectID(projectID)

	vserverInternalSvcV1 := client.NewServiceClient().
		WithEndpoint(endpoint + "internal").
		WithClient(hc).
		WithProjectID(projectID)

	vserverSvcV2 := client.NewServiceClient().
		WithEndpoint(endpoint + "v2").
		WithClient(hc).
		WithProjectID(projectID)

	return &VServerGateway{
		endpoint:                 endpoint,
		vserverGatewayV1:         NewVServerGatewayV1(vserverSvcV1),
		vserverGatewayV2:         NewVServerGatewayV2(vserverSvcV2),
		vserverGatewayInternalV1: NewVServerGatewayInternalV1(vserverInternalSvcV1),
	}
}

func NewVLBGateway(lbEndpoint, serverEndpoint, projectID string, hc client.HTTPClient) *VLBGateway {
	vlbSvcV2 := client.NewServiceClient().
		WithEndpoint(lbEndpoint + "v2").
		WithClient(hc).
		WithProjectID(projectID)

	vlbSvcIn := client.NewServiceClient().
		WithEndpoint(lbEndpoint + "internal").
		WithClient(hc).
		WithProjectID(projectID)

	vserverSvcV2 := client.NewServiceClient().
		WithEndpoint(serverEndpoint + "v2").
		WithClient(hc).
		WithProjectID(projectID)

	return &VLBGateway{
		endpoint:           lbEndpoint,
		vlbGatewayV2:       NewVLBGatewayV2(vlbSvcV2, vserverSvcV2),
		vlbGatewayInternal: NewVLBGatewayInternal(vlbSvcIn),
	}
}

func NewVNetworkGateway(endpoint, zoneID, projectID, userID string, hc client.HTTPClient) *VNetworkGateway {
	vnetworkSvcV1 := client.NewServiceClient().
		WithEndpoint(endpoint + "vnetwork/v1").
		WithClient(hc).
		WithZoneID(zoneID).
		WithProjectID(projectID).
		WithUserID(userID)

	vnetworkSvcV2 := client.NewServiceClient().
		WithEndpoint(endpoint + "vnetwork/az/v1").
		WithClient(hc).
		WithZoneID(zoneID).
		WithProjectID(projectID).
		WithUserID(userID)

	vnetworkSvcInternalV1 := client.NewServiceClient().
		WithEndpoint(endpoint + "internal/v1").
		WithClient(hc).
		WithZoneID(zoneID).
		WithProjectID(projectID).
		WithUserID(userID)

	return &VNetworkGateway{
		endpoint:                  endpoint,
		vnetworkGatewayV1:         NewVNetworkGatewayV1(vnetworkSvcV1),
		vnetworkGatewayV2:         NewVNetworkGatewayV2(vnetworkSvcV2),
		vnetworkGatewayInternalV1: NewVNetworkGatewayInternalV1(vnetworkSvcInternalV1),
	}
}

func (g *IAMGateway) V2() *IAMGatewayV2 {
	return g.iamGatewayV2
}

func (g *VServerGateway) V1() *VServerGatewayV1 {
	return g.vserverGatewayV1
}

func (g *VServerGateway) V2() *VServerGatewayV2 {
	return g.vserverGatewayV2
}

func (g *VServerGateway) InternalV1() *VServerGatewayInternalV1 {
	return g.vserverGatewayInternalV1
}

func (g *VLBGateway) Internal() *VLBGatewayInternal {
	return g.vlbGatewayInternal
}

func (g *VLBGateway) V2() *VLBGatewayV2 {
	return g.vlbGatewayV2
}

func (g *VServerGateway) GetEndpoint() string {
	return g.endpoint
}

func (g *VLBGateway) GetEndpoint() string {
	return g.endpoint
}

func (g *VNetworkGateway) V1() *VNetworkGatewayV1 {
	return g.vnetworkGatewayV1
}

func (g *VNetworkGateway) V2() *VNetworkGatewayV2 {
	return g.vnetworkGatewayV2
}

func (g *VNetworkGateway) GetEndpoint() string {
	return g.endpoint
}

func (g *VNetworkGateway) InternalV1() *VNetworkGatewayInternalV1 {
	return g.vnetworkGatewayInternalV1
}

type GLBGateway struct {
	endpoint     string
	glbGatewayV1 *GLBGatewayV1
}

func NewGLBGateway(endpoint string, hc client.HTTPClient) *GLBGateway {
	svcClient := client.NewServiceClient().
		WithEndpoint(endpoint + "v1").
		WithClient(hc)

	return &GLBGateway{
		endpoint:     endpoint,
		glbGatewayV1: NewGLBGatewayV1(svcClient),
	}
}

func (g *GLBGateway) V1() *GLBGatewayV1 {
	return g.glbGatewayV1
}

type GLBGatewayV1 struct {
	glbService glb.GLBServiceV1
}

func (g *GLBGatewayV1) GLBService() glb.GLBServiceV1 {
	return g.glbService
}

func NewGLBGatewayV1(svcClient client.ServiceClient) *GLBGatewayV1 {
	return &GLBGatewayV1{
		glbService: glb.NewGLBServiceV1(svcClient),
	}
}

type VDnsGateway struct {
	endpoint           string
	dnsService         dns.VDnsServiceV1
	dnsServiceInternal dns.VDnsServiceInternal
}

func NewVDnsGateway(endpoint, projectID string, hc client.HTTPClient) *VDnsGateway {
	svcClient := client.NewServiceClient().
		WithEndpoint(endpoint + "v1").
		WithClient(hc).
		WithProjectID(projectID)

	internalClient := client.NewServiceClient().
		WithEndpoint(endpoint + "internal/v1").
		WithClient(hc).
		WithProjectID(projectID)

	return &VDnsGateway{
		endpoint:           endpoint,
		dnsService:         dns.NewVDnsServiceV1(svcClient),
		dnsServiceInternal: dns.NewVDnsServiceInternal(internalClient),
	}
}

func (g *VDnsGateway) V1() *VDnsGatewayV1 {
	return &VDnsGatewayV1{
		dnsService: g.dnsService,
	}
}

func (g *VDnsGateway) Internal() *VDnsGatewayInternal {
	return &VDnsGatewayInternal{
		dnsService: g.dnsServiceInternal,
	}
}

func (g *VDnsGateway) GetEndpoint() string {
	return g.endpoint
}

type VDnsGatewayV1 struct {
	dnsService dns.VDnsServiceV1
}

func (g *VDnsGatewayV1) DnsService() dns.VDnsServiceV1 {
	return g.dnsService
}

type VDnsGatewayInternal struct {
	dnsService dns.VDnsServiceInternal
}

func (g *VDnsGatewayInternal) DnsService() dns.VDnsServiceInternal {
	return g.dnsService
}
