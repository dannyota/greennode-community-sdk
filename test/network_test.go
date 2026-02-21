//go:build integration

package test

import (
	"context"
	"testing"

	networkv2 "danny.vn/greennode/services/network/v2"
)

func TestGetNetworkByIDFailure(t *testing.T) {
	vngcloud := validSdkConfig()
	opt := networkv2.NewGetNetworkByIDRequest("net-4f35f173-e0fe-4202-9c2b-5121b558bcd2")
	network, err := vngcloud.Network.GetNetworkByID(context.Background(), opt)

	if err == nil {
		t.Errorf("Expect error not to be nil but got nil")
	}

	if network != nil {
		t.Errorf("Expect portal to be nil but got %+v", network)
	}

	t.Log("RESULT:", err)
	t.Log("PASS")
}

func TestGetNetworkByIDSuccess(t *testing.T) {
	vngcloud := validSdkConfig()
	opt := networkv2.NewGetNetworkByIDRequest("net-4f35f173-e0fe-4202-9c2b-5121b558bcd3")
	network, err := vngcloud.Network.GetNetworkByID(context.Background(), opt)

	if err != nil {
		t.Fatalf("Expect error to be nil but got %+v", err)
	}

	if network == nil {
		t.Fatalf("Expect portal not to be nil but got nil")
	}

	t.Log("RESULT:", network)
	t.Log("PASS")
}
