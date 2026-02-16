//go:build integration

package test

import (
	"context"
	"testing"

	networkv2 "github.com/dannyota/greennode-community-sdk/v2/greennode/services/network/v2"
)

func TestGetAllAddressPairByVirtualSubnetID(t *testing.T) {
	vngcloud := validSdkConfig()
	opt := networkv2.NewGetAllAddressPairByVirtualSubnetIDRequest(getValueOfEnv("VIRTUAL_SUBNET_ID"))
	network, err := vngcloud.Network.GetAllAddressPairByVirtualSubnetID(context.Background(), opt)

	if err != nil {
		t.Fatalf("Expect error to be nil but got %+v", err)
	}

	if network == nil {
		t.Fatalf("Expect portal not to be nil but got nil")
	}

	t.Log("RESULT:", network)
	for _, addressPair := range network {
		t.Logf("AddressPair: %+v", addressPair)
	}
	t.Log("PASS")
}

func TestSetAddressPairInVirtualSubnet(t *testing.T) {
	vngcloud := validSdkConfig()
	opt := networkv2.NewSetAddressPairInVirtualSubnetRequest(
		getValueOfEnv("VIRTUAL_SUBNET_ID"),
		getValueOfEnv("NETWORK_INTERFACE_ID"),
		"10.30.1.28/30",
	)
	network, err := vngcloud.Network.SetAddressPairInVirtualSubnet(context.Background(), opt)

	if err != nil {
		t.Fatalf("Expect error to be nil but got %+v", err)
	}

	if network == nil {
		t.Fatalf("Expect portal not to be nil but got nil")
	}

	t.Log("RESULT:", network)
}

func TestDeleteAddressPair(t *testing.T) {
	vngcloud := validSdkConfig()
	opt := networkv2.NewDeleteAddressPairRequest(getValueOfEnv("ADDRESS_PAIR_ID"))
	err := vngcloud.Network.DeleteAddressPair(context.Background(), opt)

	if err != nil {
		t.Fatalf("Expect error to be nil but got %+v", err)
	}

	t.Log("PASS")
}

func TestCreateAddressPair(t *testing.T) {
	vngcloud := validSdkConfigHanRegion()

	virtualAddressID := "vip-0d2402cf-49e8-43bf-abbe-b707597320e9"
	internalNicID := "net-in-3b076753-6561-4e3e-8a66-e10dc79cab2d"

	opt := networkv2.NewCreateAddressPairRequest(virtualAddressID, internalNicID)
	mode := networkv2.AddressPairModeActiveActive
	opt.Mode = &mode

	ap, err := vngcloud.Network.CreateAddressPair(context.Background(), opt)
	if err != nil {
		t.Fatalf("Expect error to be nil but got %+v", err)
	}

	if ap == nil {
		t.Fatalf("Expect portal not to be nil but got nil")
	}

	t.Log("RESULT:", ap)
}
