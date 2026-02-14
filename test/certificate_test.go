package test

import (
	"testing"

	v2 "github.com/dannyota/greennode-community-sdk/v2/greennode/services/loadbalancer/v2"
)

func TestListCertificates(t *testing.T) {
	vngcloud := validSdkConfig()
	opt := v2.NewListCertificatesRequest()
	certs, sdkerr := vngcloud.VLBGateway().V2().LoadBalancerService().ListCertificates(opt)
	if sdkerr != nil {
		t.Fatalf("Expect nil but got %+v", sdkerr)
	}

	if certs == nil {
		t.Fatalf("Expect not nil but got nil")
	}

	t.Log("Result: ", certs)
	for _, pkg := range certs.Certificates {
		t.Logf("Package: %+v", pkg)
	}
	t.Log("PASS")
}

func TestGetCertificateByID(t *testing.T) {
	vngcloud := validSdkConfig()
	opt := v2.NewGetCertificateByIDRequest("secret-84cb7a5e-b949-4f1b-a2e8-d2752e6e1181")
	cert, sdkerr := vngcloud.VLBGateway().V2().LoadBalancerService().GetCertificateByID(opt)
	if sdkerr != nil {
		t.Fatalf("Expect nil but got %+v", sdkerr)
	}

	if cert == nil {
		t.Fatalf("Expect not nil but got nil")
	}

	t.Logf("Result: %+v", cert)
	t.Log("PASS")
}

func TestCreateCertificate(t *testing.T) {
	vngcloud := validSdkConfig()

	opt := v2.NewCreateCertificateRequest(
		"annd2-haha",
		FakeCertificate,
		v2.ImportOptsTypeOptTLS,
	).WithPrivateKey(FakePrivateKey)

	cert, err := vngcloud.VLBGateway().V2().LoadBalancerService().CreateCertificate(opt)
	if err != nil {
		t.Fatalf("Expect nil but got %+v", err)
	}

	if cert == nil {
		t.Fatalf("Expect not nil but got nil")
	}

	t.Logf("Result: %+v", cert)
	t.Log("PASS")
}

func TestDeleteCertificateByID(t *testing.T) {
	vngcloud := validSdkConfig()
	opt := v2.NewDeleteCertificateByIDRequest(FakeSecretID)
	err := vngcloud.VLBGateway().V2().LoadBalancerService().DeleteCertificateByID(opt)
	if err != nil {
		t.Fatalf("Expect nil but got %+v", err)
	}

	t.Log("PASS")
}
