//go:build integration

package test

import (
	"context"
	"testing"

	networkv2 "danny.vn/greennode/greennode/services/network/v2"
)

func TestCreateVirtualAddressCrossProject(t *testing.T) {
	virtualAddressName := "test-virtual-address"
	projectID := "pro-5ce9da27-8ac9-40db-8743-d80f6cbf1491"
	subnetID := "sub-33ec4719-915f-4818-85b0-f17bdf7f899b"

	vngcloud := validSdkConfigHanRegion()
	opt := networkv2.NewCreateVirtualAddressCrossProjectRequest(
		virtualAddressName, projectID, subnetID)
	opt.Description = "Private DNS endpoint address for VPC created by vDNS. Please DO NOT delete this address."
	vaddr, err := vngcloud.Network.CreateVirtualAddressCrossProject(context.Background(), opt)

	if err != nil {
		t.Errorf("Expect error to be nil but got %+v", err)
	}

	if vaddr == nil {
		t.Errorf("Expect portal not to be nil but got nil")
	}

	t.Log("RESULT:", vaddr)
	t.Log("PASS")
}

func TestDeleteVirtualAddressByID(t *testing.T) {
	virtualAddressID := "vip-1a17ffb3-28e5-4a7a-a4e0-17af09de28aa"

	vngcloud := validSdkConfigHanRegion()
	opt := networkv2.NewDeleteVirtualAddressByIDRequest(virtualAddressID)
	err := vngcloud.Network.DeleteVirtualAddressByID(context.Background(), opt)

	if err != nil {
		t.Errorf("Expect error to be nil but got %+v", err)
	}

	t.Log("PASS")
}

func TestGetVirtualAddessByID(t *testing.T) {
	virtualAddressID := "vip-0d2402cf-49e8-43bf-abbe-b707597320e0"

	vngcloud := validSdkConfigHanRegion()
	opt := networkv2.NewGetVirtualAddressByIDRequest(virtualAddressID)
	vaddr, err := vngcloud.Network.GetVirtualAddressByID(context.Background(), opt)

	if err != nil {
		t.Errorf("Expect error to be nil but got %+v", err)
	}

	if vaddr == nil {
		t.Errorf("Expect portal not to be nil but got nil")
	}

	t.Log("RESULT:", vaddr)
	t.Log("PASS")
}

func TestListAddressPairsByVirtualAddressID(t *testing.T) {
	virtualAddressID := "vip-0d2402cf-49e8-43bf-abbe-b707597320e9"

	vngcloud := validSdkConfigHanRegion()
	opt := networkv2.NewListAddressPairsByVirtualAddressIDRequest(virtualAddressID)
	pairs, err := vngcloud.Network.ListAddressPairsByVirtualAddressID(context.Background(), opt)

	if err != nil {
		t.Errorf("Expect error to be nil but got %+v", err)
	}

	if pairs == nil {
		t.Errorf("Expect portal not to be nil but got nil")
	}

	t.Log("RESULT:", pairs.At(0))
	t.Log("PASS")
}
