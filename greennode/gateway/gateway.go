package gateway

import (
	"github.com/dannyota/greennode-community-sdk/v2/greennode/client"
	"github.com/dannyota/greennode-community-sdk/v2/greennode/services/dns"
	"github.com/dannyota/greennode-community-sdk/v2/greennode/services/glb"
)

var _ IIamGateway = &iamGateway{}

type iamGateway struct {
	iamGatewayV2 IIamGatewayV2
}

var _ IVServerGateway = &vserverGateway{}

type vserverGateway struct {
	endpoint                 string // Hold the endpoint of the vServer service
	vserverGatewayV1         IVServerGatewayV1
	vserverGatewayV2         IVServerGatewayV2
	vserverGatewayInternalV1 IVServerGatewayInternalV1
}

var _ IVLBGateway = &vlbGateway{}

type vlbGateway struct {
	endpoint           string // Hold the endpoint of the vLB service
	vlbGatewayInternal IVLBGatewayInternal
	vlbGatewayV2       IVLBGatewayV2
}

var _ IVNetworkGateway = &vnetworkGateway{}

type vnetworkGateway struct {
	endpoint                  string
	vnetworkGatewayV1         IVNetworkGatewayV1
	vnetworkGatewayV2         IVNetworkGatewayV1
	vnetworkGatewayInternalV1 IVNetworkGatewayInternalV1
}

func NewIamGateway(pendpoint, projectId string, phc client.IHttpClient) IIamGateway {
	iamSvcV2 := client.NewServiceClient().
		WithEndpoint(pendpoint + "v2").
		WithClient(phc).
		WithProjectId(projectId)

	return &iamGateway{
		iamGatewayV2: NewIamGatewayV2(iamSvcV2),
	}
}

func NewVServerGateway(pendpoint, pprojectId string, phc client.IHttpClient) IVServerGateway {
	vserverSvcV1 := client.NewServiceClient().
		WithEndpoint(pendpoint + "v1").
		WithClient(phc).
		WithProjectId(pprojectId)

	vserverInternalSvcV1 := client.NewServiceClient().
		WithEndpoint(pendpoint + "internal").
		WithClient(phc).
		WithProjectId(pprojectId)

	vserverSvcV2 := client.NewServiceClient().
		WithEndpoint(pendpoint + "v2").
		WithClient(phc).
		WithProjectId(pprojectId)

	return &vserverGateway{
		endpoint:                 pendpoint,
		vserverGatewayV1:         NewVServerGatewayV1(vserverSvcV1),
		vserverGatewayV2:         NewVServerGatewayV2(vserverSvcV2),
		vserverGatewayInternalV1: NewVServerGatewayInternalV1(vserverInternalSvcV1),
	}
}

func NewVLBGateway(plbEndpoint, pserverEndpoint, pprojectId string, phc client.IHttpClient) IVLBGateway {
	vlbSvcV2 := client.NewServiceClient().
		WithEndpoint(plbEndpoint + "v2").
		WithClient(phc).
		WithProjectId(pprojectId)

	vlbSvcIn := client.NewServiceClient().
		WithEndpoint(plbEndpoint + "internal").
		WithClient(phc).
		WithProjectId(pprojectId)

	vserverSvcV2 := client.NewServiceClient().
		WithEndpoint(pserverEndpoint + "v2").
		WithClient(phc).
		WithProjectId(pprojectId)

	return &vlbGateway{
		endpoint:           plbEndpoint,
		vlbGatewayV2:       NewVLBGatewayV2(vlbSvcV2, vserverSvcV2),
		vlbGatewayInternal: NewVLBGatewayInternal(vlbSvcIn),
	}
}

func NewVNetworkGateway(pendpoint, pzoneId, projectId, puserId string, phc client.IHttpClient) IVNetworkGateway {
	vnetworkSvcV1 := client.NewServiceClient().
		WithEndpoint(pendpoint + "vnetwork/v1").
		WithClient(phc).
		WithZoneId(pzoneId).
		WithProjectId(projectId).
		WithUserId(puserId)

	vnetworkSvcV2 := client.NewServiceClient().
		WithEndpoint(pendpoint + "vnetwork/az/v1").
		WithClient(phc).
		WithZoneId(pzoneId).
		WithProjectId(projectId).
		WithUserId(puserId)

	vnetworkSvcInternalV1 := client.NewServiceClient().
		WithEndpoint(pendpoint + "internal/v1").
		WithClient(phc).
		WithZoneId(pzoneId).
		WithProjectId(projectId).
		WithUserId(puserId)

	return &vnetworkGateway{
		endpoint:                  pendpoint,
		vnetworkGatewayV1:         NewVNetworkGatewayV1(vnetworkSvcV1),
		vnetworkGatewayV2:         NewVNetworkGatewayV1(vnetworkSvcV2),
		vnetworkGatewayInternalV1: NewVNetworkGatewayInternalV1(vnetworkSvcInternalV1),
	}
}

func (s *iamGateway) V2() IIamGatewayV2 {
	return s.iamGatewayV2
}

func (s *vserverGateway) V1() IVServerGatewayV1 {
	return s.vserverGatewayV1
}

func (s *vserverGateway) V2() IVServerGatewayV2 {
	return s.vserverGatewayV2
}
func (s *vserverGateway) InternalV1() IVServerGatewayInternalV1 {
	return s.vserverGatewayInternalV1
}

func (s *vlbGateway) Internal() IVLBGatewayInternal {
	return s.vlbGatewayInternal
}

func (s *vlbGateway) V2() IVLBGatewayV2 {
	return s.vlbGatewayV2
}

func (s *vserverGateway) GetEndpoint() string {
	return s.endpoint
}

func (s *vlbGateway) GetEndpoint() string {
	return s.endpoint
}

func (s *vnetworkGateway) V1() IVNetworkGatewayV1 {
	return s.vnetworkGatewayV1
}

func (s *vnetworkGateway) V2() IVNetworkGatewayV1 {
	return s.vnetworkGatewayV2
}

func (s *vnetworkGateway) GetEndpoint() string {
	return s.endpoint
}

func (s *vnetworkGateway) InternalV1() IVNetworkGatewayInternalV1 {
	return s.vnetworkGatewayInternalV1
}

var _ IGLBGateway = &glbGateway{}

type glbGateway struct {
	endpoint     string
	glbGatewayV1 IGLBGatewayV1
}

func NewGLBGateway(pendpoint string, phc client.IHttpClient) IGLBGateway {
	svcClient := client.NewServiceClient().
		WithEndpoint(pendpoint + "v1").
		WithClient(phc)

	return &glbGateway{
		endpoint:     pendpoint,
		glbGatewayV1: NewGLBGatewayV1(svcClient),
	}
}

func (s *glbGateway) V1() IGLBGatewayV1 {
	return s.glbGatewayV1
}

var _ IGLBGatewayV1 = &glbGatewayV1{}

type glbGatewayV1 struct {
	glbService glb.IGLBServiceV1
}

func (s *glbGatewayV1) GLBService() glb.IGLBServiceV1 {
	return s.glbService
}

func NewGLBGatewayV1(psvcClient client.IServiceClient) IGLBGatewayV1 {
	return &glbGatewayV1{
		glbService: glb.NewGLBServiceV1(psvcClient),
	}
}

var _ IVDnsGateway = &vdnsGateway{}

type vdnsGateway struct {
	endpoint           string
	dnsService         dns.IVDnsServiceV1
	dnsServiceInternal dns.IVDnsServiceInternal
}

func NewVDnsGateway(pendpoint, pprojectId string, phc client.IHttpClient) IVDnsGateway {
	svcClient := client.NewServiceClient().
		WithEndpoint(pendpoint + "v1").
		WithClient(phc).
		WithProjectId(pprojectId)

	internalClient := client.NewServiceClient().
		WithEndpoint(pendpoint + "internal/v1").
		WithClient(phc).
		WithProjectId(pprojectId)

	return &vdnsGateway{
		endpoint:           pendpoint,
		dnsService:         dns.NewVDnsServiceV1(svcClient),
		dnsServiceInternal: dns.NewVDnsServiceInternal(internalClient),
	}
}

func (s *vdnsGateway) V1() IVDnsGatewayV1 {
	return &vdnsGatewayV1{
		dnsService: s.dnsService,
	}
}

func (s *vdnsGateway) Internal() IVDnsGatewayInternal {
	return &vdnsGatewayInternal{
		dnsService: s.dnsServiceInternal,
	}
}

func (s *vdnsGateway) GetEndpoint() string {
	return s.endpoint
}

var _ IVDnsGatewayV1 = &vdnsGatewayV1{}

type vdnsGatewayV1 struct {
	dnsService dns.IVDnsServiceV1
}

func (s *vdnsGatewayV1) DnsService() dns.IVDnsServiceV1 {
	return s.dnsService
}

var _ IVDnsGatewayInternal = &vdnsGatewayInternal{}

type vdnsGatewayInternal struct {
	dnsService dns.IVDnsServiceInternal
}

func (s *vdnsGatewayInternal) DnsService() dns.IVDnsServiceInternal {
	return s.dnsService
}
