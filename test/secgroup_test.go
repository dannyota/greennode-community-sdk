package test

import (
	"testing"

	networkv2 "github.com/dannyota/greennode-community-sdk/v2/greennode/services/network/v2"
)

func TestGetSecgroupByIDSuccess(t *testing.T) {
	secgroupID := "secg-d803abe8-2cf9-4b46-b3e7-94d8ab3c94ca"
	vngcloud := validSdkConfig()
	opt := networkv2.NewGetSecgroupByIDRequest(secgroupID)
	secgroup, err := vngcloud.VServerGateway().V2().NetworkService().GetSecgroupByID(opt)

	if err != nil {
		t.Errorf("Expect error to be nil but got %+v", err)
	}

	if secgroup == nil {
		t.Errorf("Expect portal not to be nil but got nil")
	}

	t.Log("RESULT:", secgroup)
	t.Log("PASS")
}

func TestGetSecgroupByIDFailure(t *testing.T) {
	secgroupID := "secg-90d617b4-b893-407b-a9a8-3bd80c177920"
	vngcloud := validSdkConfig()
	opt := networkv2.NewGetSecgroupByIDRequest(secgroupID)
	secgroup, err := vngcloud.VServerGateway().V2().NetworkService().GetSecgroupByID(opt)

	if err == nil {
		t.Errorf("Expect error not to be nil but got nil")
	}

	if secgroup != nil {
		t.Errorf("Expect portal to be nil but got %+v", secgroup)
	}

	t.Log("RESULT:", err)
	t.Log("PASS")
}

func TestCreateSecgroupSameNameFailure(t *testing.T) {
	secgroupName := "cuongdm3-temporal"
	vngcloud := validSdkConfig()
	opt := networkv2.NewCreateSecgroupRequest(secgroupName, "this is a test")
	secgroup, err := vngcloud.VServerGateway().V2().NetworkService().CreateSecgroup(opt)

	if err == nil {
		t.Errorf("Expect error not to be nil but got nil")
	}

	if secgroup != nil {
		t.Errorf("Expect portal to be nil but got %+v", secgroup)
	}

	t.Log("RESULT:", err)
	t.Log("PASS")
}

func TestCreateSecgroupSuccess(t *testing.T) {
	secgroupName := "cuongdm3-temporal-1"
	vngcloud := validSdkConfig()
	opt := networkv2.NewCreateSecgroupRequest(secgroupName, "this is a test")
	secgroup, err := vngcloud.VServerGateway().V2().NetworkService().CreateSecgroup(opt)

	if err != nil {
		t.Errorf("Expect error to be nil but got %+v", err)
	}

	if secgroup == nil {
		t.Errorf("Expect portal not to be nil but got nil")
	}

	t.Log("RESULT:", secgroup)
	t.Log("PASS")
}

func TestDeleteSecgroupByIDFailure(t *testing.T) {
	secgroupID := "secg-90d617b4-b893-407b-a9a8-3bd80c177920"
	vngcloud := validSdkConfig()
	opt := networkv2.NewDeleteSecgroupByIDRequest(secgroupID)
	err := vngcloud.VServerGateway().V2().NetworkService().DeleteSecgroupByID(opt)

	if err == nil {
		t.Errorf("Expect error not to be nil but got nil")
	}

	t.Log("RESULT:", err)
	t.Log("PASS")
}

func TestDeleteSecgroupByIDSuccess(t *testing.T) {
	secgroupID := "secg-3787f73d-d62b-49ca-96cd-226b7dc8ead4"
	vngcloud := validSdkConfig()
	opt := networkv2.NewDeleteSecgroupByIDRequest(secgroupID)
	err := vngcloud.VServerGateway().V2().NetworkService().DeleteSecgroupByID(opt)

	if err != nil {
		t.Errorf("Expect error to be nil but got %+v", err)
	}

	t.Log("RESULT:", err)
	t.Log("PASS")
}

func TestListAllServerBySecgroupIDSuccess(t *testing.T) {
	secgroupID := "secg-1395e86c-9631-4c13-xxxx-e41be5bdaab3"
	vngcloud := validSdkConfig()
	opt := networkv2.NewListAllServersBySecgroupIDRequest(secgroupID)
	serbvers, err := vngcloud.VServerGateway().V2().NetworkService().ListAllServersBySecgroupID(opt)

	if err != nil {
		t.Errorf("Expect error to be nil but got %+v", err)
	}

	if serbvers == nil {
		t.Errorf("Expect portal not to be nil but got nil")
	}

	t.Log("RESULT:", serbvers)
	t.Log("PASS")
}

func TestListSecgroupSuccess(t *testing.T) {
	vngcloud := validSdkConfig()
	opt := networkv2.NewListSecgroupRequest()
	secgroups, err := vngcloud.VServerGateway().V2().NetworkService().ListSecgroup(opt.AddUserAgent("test"))

	if err != nil {
		t.Errorf("Expect error to be nil but got %+v", err)
	}

	if secgroups == nil {
		t.Errorf("Expect portal not to be nil but got nil")
	}

	t.Log("RESULT:", secgroups)
	if secgroups == nil || len(secgroups.Items) == 0 {
		t.Log("No secgroup found")
	}
	for _, secgroup := range secgroups.Items {
		t.Logf("Secgroup: %+v", secgroup)
	}
	t.Log("PASS")
}
