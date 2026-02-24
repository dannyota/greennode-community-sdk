//go:build integration

package test

import (
	"context"
	"net/url"
	"testing"

	networkv1 "danny.vn/greennode/services/network/v1"
)

func TestGetEndpointSuccess(t *testing.T) {
	vngcloud := validSdkConfig()
	opt := networkv1.NewGetEndpointByIDRequest("enp-7575cb25-0033-4c26-9145-53cd90d7778c")

	lb, sdkerr := vngcloud.NetworkV1.GetEndpointByID(context.Background(), opt)
	if sdkerr != nil {
		t.Errorf("Expect nil but got %+v", sdkerr)
	}

	if lb == nil {
		t.Errorf("Expect not nil but got nil")
	}

	t.Log("Result: ", lb)
	t.Log("PASS")
}

func TestCreateEndpoint(t *testing.T) {
	vngcloud := validSdkConfig()
	opt := networkv1.NewCreateEndpointRequest(
		"test-endpoint",
		"f3d11a4c-f071-4009-88a6-4a21346c8708",
		"net-5ac170fc-834a-4621-b512-481e09b82fc8",
		"sub-0c508dd6-5af6-4f0e-a860-35346b530cf1",
	)
	opt.ResourceInfo.Description = "This is the service endpoint for vStorage APIs, established by the VKS product. " +
		"Please refrain from DELETING it manually."

	lb, sdkerr := vngcloud.NetworkV1.CreateEndpoint(context.Background(), opt)
	if sdkerr != nil {
		t.Fatalf("Expect nil but got %+v", sdkerr)
	}

	if lb == nil {
		t.Fatalf("Expect not nil but got nil")
	}

	t.Log("Result: ", lb)
	t.Log("PASS")
}

func TestCreateEndpointInternal(t *testing.T) {
	vngcloud := validSuperSdkConfig()
	opt := networkv1.NewCreateEndpointRequest(
		"test-endpoint-internal",
		"c36bb265-f569-4748-a03a-fca52c7588ea",
		"net-dc14bb60-d500-40b5-945f-218540990187",
		"sub-3f7a1d9b-1d68-44d0-a14f-4cc6bf18a7c4",
	)
	opt.ResourceInfo.Description = "This is the service endpoint for vStorage APIs, established by the VKS product. " +
		"Please refrain from DELETING it manually."
	opt.ResourceInfo.Networking = append(opt.ResourceInfo.Networking,
		networkv1.EndpointNetworking{Zone: "HCM-1B", SubnetUuid: "sub-3f7a1d9b-1d68-44d0-a14f-4cc6bf18a7c4"},
		networkv1.EndpointNetworking{Zone: "HCM-1A", SubnetUuid: "sub-85ba01f6-02ec-4dfc-8884-ee0036c68a5b"},
	)
	opt.ResourceInfo.Scaling = networkv1.EndpointScaling{MinSize: 1, MaxSize: 3}
	opt.ResourceInfo.EnableDnsName = true
	opt.ResourceInfo.IsBuyMorePoc = false
	opt.ResourceInfo.IsPoc = false
	opt.ResourceInfo.IsEnableAutoRenew = true

	lb, sdkerr := vngcloud.NetworkV1.CreateEndpoint(context.Background(), opt)
	if sdkerr != nil {
		t.Fatalf("Expect nil but got %+v", sdkerr)
	}

	if lb == nil {
		t.Fatalf("Expect not nil but got nil")
	}

	t.Log("Result: ", lb)
	t.Log("PASS")
}

func TestDeleteEndpoint(t *testing.T) {
	vngcloud := validSdkConfig()
	opt := networkv1.NewDeleteEndpointByIDRequest(
		"enp-56d7359f-4b9a-4f01-a210-54523c6d0c88",
		"net-5ac170fc-834a-4621-b512-481e09b82fc8",
		"b9ba2b16-389e-48b7-9e75-4c991239da27",
	)

	sdkerr := vngcloud.NetworkV1.DeleteEndpointByID(context.Background(), opt)
	if sdkerr != nil {
		t.Fatalf("Expect nil but got %+v", sdkerr)
	}

	t.Log("PASS")
}

func TestListEndpoints(t *testing.T) {
	vngcloud := validSdkConfig()
	opt := networkv1.NewListEndpointsRequest(1, 100)
	opt.Uuid = "enp-9349271b-af44-4e39-8829-615d945fa6c2"

	lb, sdkerr := vngcloud.NetworkV1.ListEndpoints(context.Background(), opt)
	if sdkerr != nil {
		t.Fatalf("Expect nil but got %+v", sdkerr)
	}

	if lb == nil {
		t.Fatalf("Expect not nil but got nil")
	}

	t.Log("Result: ", lb.At(0))
	t.Log("PASS")
}

func TestEndpoint(t *testing.T) {
	raw := `{"page":1,"size":10,"search":[{"field":"vpcId","value":"net-5ac170fc-834a-4621-b512-481e09b82fc8"}]}`
	encode := url.QueryEscape(raw)

	t.Log("Encode: ", encode)
}

