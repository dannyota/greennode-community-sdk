package gateway

import (
	"github.com/dannyota/greennode-community-sdk/v2/greennode/client"
	"github.com/dannyota/greennode-community-sdk/v2/greennode/services/dns"
	"github.com/dannyota/greennode-community-sdk/v2/greennode/services/glb"
)

var _ IamGateway = &iamGateway{}

type iamGateway struct {
	iamGatewayV2 IamGatewayV2
}

var _ VServerGateway = &vserverGateway{}

type vserverGateway struct {
	endpoint                 string // Hold the endpoint of the vServer service
	vserverGatewayV1         VServerGatewayV1
	vserverGatewayV2         VServerGatewayV2
	vserverGatewayInternalV1 VServerGatewayInternalV1
}

var _ VLBGateway = &vlbGateway{}

type vlbGateway struct {
	endpoint           string // Hold the endpoint of the vLB service
	vlbGatewayInternal VLBGatewayInternal
	vlbGatewayV2       VLBGatewayV2
}

var _ VNetworkGateway = &vnetworkGateway{}

type vnetworkGateway struct {
	endpoint                  string
	vnetworkGatewayV1         VNetworkGatewayV1
	vnetworkGatewayV2         VNetworkGatewayV1
	vnetworkGatewayInternalV1 VNetworkGatewayInternalV1
}

func NewIamGateway(pendpoint, projectId string, phc client.HttpClient) IamGateway {
	iamSvcV2 := client.NewServiceClient().
		WithEndpoint(pendpoint + "v2").
		WithClient(phc).
		WithProjectId(projectId)

	return &iamGateway{
		iamGatewayV2: NewIamGatewayV2(iamSvcV2),
	}
}

func NewVServerGateway(pendpoint, pprojectId string, phc client.HttpClient) VServerGateway {
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

func NewVLBGateway(plbEndpoint, pserverEndpoint, pprojectId string, phc client.HttpClient) VLBGateway {
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

func NewVNetworkGateway(pendpoint, pzoneId, projectId, puserId string, phc client.HttpClient) VNetworkGateway {
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

func (s *iamGateway) V2() IamGatewayV2 {
	return s.iamGatewayV2
}

func (s *vserverGateway) V1() VServerGatewayV1 {
	return s.vserverGatewayV1
}

func (s *vserverGateway) V2() VServerGatewayV2 {
	return s.vserverGatewayV2
}
func (s *vserverGateway) InternalV1() VServerGatewayInternalV1 {
	return s.vserverGatewayInternalV1
}

func (s *vlbGateway) Internal() VLBGatewayInternal {
	return s.vlbGatewayInternal
}

func (s *vlbGateway) V2() VLBGatewayV2 {
	return s.vlbGatewayV2
}

func (s *vserverGateway) GetEndpoint() string {
	return s.endpoint
}

func (s *vlbGateway) GetEndpoint() string {
	return s.endpoint
}

func (s *vnetworkGateway) V1() VNetworkGatewayV1 {
	return s.vnetworkGatewayV1
}

func (s *vnetworkGateway) V2() VNetworkGatewayV1 {
	return s.vnetworkGatewayV2
}

func (s *vnetworkGateway) GetEndpoint() string {
	return s.endpoint
}

func (s *vnetworkGateway) InternalV1() VNetworkGatewayInternalV1 {
	return s.vnetworkGatewayInternalV1
}

var _ GLBGateway = &glbGateway{}

type glbGateway struct {
	endpoint     string
	glbGatewayV1 GLBGatewayV1
}

func NewGLBGateway(pendpoint string, phc client.HttpClient) GLBGateway {
	svcClient := client.NewServiceClient().
		WithEndpoint(pendpoint + "v1").
		WithClient(phc)

	return &glbGateway{
		endpoint:     pendpoint,
		glbGatewayV1: NewGLBGatewayV1(svcClient),
	}
}

func (s *glbGateway) V1() GLBGatewayV1 {
	return s.glbGatewayV1
}

var _ GLBGatewayV1 = &glbGatewayV1{}

type glbGatewayV1 struct {
	glbService glb.GLBServiceV1
}

func (s *glbGatewayV1) GLBService() glb.GLBServiceV1 {
	return s.glbService
}

func NewGLBGatewayV1(psvcClient client.ServiceClient) GLBGatewayV1 {
	return &glbGatewayV1{
		glbService: glb.NewGLBServiceV1(psvcClient),
	}
}

var _ VDnsGateway = &vdnsGateway{}

type vdnsGateway struct {
	endpoint           string
	dnsService         dns.VDnsServiceV1
	dnsServiceInternal dns.VDnsServiceInternal
}

func NewVDnsGateway(pendpoint, pprojectId string, phc client.HttpClient) VDnsGateway {
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

func (s *vdnsGateway) V1() VDnsGatewayV1 {
	return &vdnsGatewayV1{
		dnsService: s.dnsService,
	}
}

func (s *vdnsGateway) Internal() VDnsGatewayInternal {
	return &vdnsGatewayInternal{
		dnsService: s.dnsServiceInternal,
	}
}

func (s *vdnsGateway) GetEndpoint() string {
	return s.endpoint
}

var _ VDnsGatewayV1 = &vdnsGatewayV1{}

type vdnsGatewayV1 struct {
	dnsService dns.VDnsServiceV1
}

func (s *vdnsGatewayV1) DnsService() dns.VDnsServiceV1 {
	return s.dnsService
}

var _ VDnsGatewayInternal = &vdnsGatewayInternal{}

type vdnsGatewayInternal struct {
	dnsService dns.VDnsServiceInternal
}

func (s *vdnsGatewayInternal) DnsService() dns.VDnsServiceInternal {
	return s.dnsService
}
