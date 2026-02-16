//go:build integration

package test

import (
	"context"
	"testing"

	v2 "github.com/dannyota/greennode-community-sdk/v2/greennode/services/loadbalancer/v2"
)

func TestListCertificates(t *testing.T) {
	vngcloud := validSdkConfig()
	opt := v2.NewListCertificatesRequest()
	certs, sdkerr := vngcloud.LoadBalancer.ListCertificates(context.Background(), opt)
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
	cert, sdkerr := vngcloud.LoadBalancer.GetCertificateByID(context.Background(), opt)
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
		"test-certificate",
		FakeCertificate,
		v2.ImportOptsTypeOptTLS,
	)
	privateKey := FakePrivateKey
	opt.PrivateKey = &privateKey

	cert, err := vngcloud.LoadBalancer.CreateCertificate(context.Background(), opt)
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
	err := vngcloud.LoadBalancer.DeleteCertificateByID(context.Background(), opt)
	if err != nil {
		t.Fatalf("Expect nil but got %+v", err)
	}

	t.Log("PASS")
}
