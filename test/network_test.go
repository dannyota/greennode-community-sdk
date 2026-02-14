package test

import (
	"testing"

	networkv2 "github.com/dannyota/greennode-community-sdk/v2/greennode/services/network/v2"
)

func TestGetNetworkByIdFailure(t *testing.T) {
	vngcloud := validSdkConfig()
	opt := networkv2.NewGetNetworkByIdRequest("net-4f35f173-e0fe-4202-9c2b-5121b558bcd2")
	network, err := vngcloud.VServerGateway().V2().NetworkService().GetNetworkById(opt)

	if err == nil {
		t.Errorf("Expect error not to be nil but got nil")
	}

	if network != nil {
		t.Errorf("Expect portal to be nil but got %+v", network)
	}

	t.Log("RESULT:", err)
	t.Log("PASS")
}

func TestGetNetworkByIdSuccess(t *testing.T) {
	vngcloud := validSdkConfig()
	opt := networkv2.NewGetNetworkByIdRequest("net-4f35f173-e0fe-4202-9c2b-5121b558bcd3")
	network, err := vngcloud.VServerGateway().V2().NetworkService().GetNetworkById(opt)

	if err != nil {
		t.Fatalf("Expect error to be nil but got %+v", err)
	}

	if network == nil {
		t.Fatalf("Expect portal not to be nil but got nil")
	}

	t.Log("RESULT:", network)
	t.Log("PASS")
}
