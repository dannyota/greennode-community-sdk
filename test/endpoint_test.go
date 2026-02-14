package test

import (
	"net/url"
	"testing"

	networkv1 "github.com/dannyota/greennode-community-sdk/v2/greennode/services/network/v1"
)

func TestGetEndpointSuccess(t *testing.T) {
	vngcloud := validSdkConfig()
	opt := networkv1.NewGetEndpointByIDRequest("enp-7575cb25-0033-4c26-9145-53cd90d7778c")

	lb, sdkerr := vngcloud.VNetworkGateway().V1().NetworkService().GetEndpointByID(opt)
	if sdkerr != nil {
		t.Errorf("Expect nil but got %+v", sdkerr.ErrorCode())
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
		"cuongdm3-test",
		"f3d11a4c-f071-4009-88a6-4a21346c8708",
		"net-5ac170fc-834a-4621-b512-481e09b82fc8",
		"sub-0c508dd6-5af6-4f0e-a860-35346b530cf1",
	).WithDescription(
		"This is the service endpoint for vStorage APIs, established by the VKS product. " +
			"Please refrain from DELETING it manually.",
	)

	lb, sdkerr := vngcloud.VNetworkGateway().V1().NetworkService().CreateEndpoint(opt)
	if sdkerr != nil {
		t.Fatalf("Expect nil but got %+v", sdkerr.Err())
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
		"tytv2-test",
		"c36bb265-f569-4748-a03a-fca52c7588ea",
		"net-dc14bb60-d500-40b5-945f-218540990187",
		"sub-3f7a1d9b-1d68-44d0-a14f-4cc6bf18a7c4",
	).WithDescription(
		"This is the service endpoint for vStorage APIs, established by the VKS product. "+
			"Please refrain from DELETING it manually.",
	).AddNetworking("HCM-1B", "sub-3f7a1d9b-1d68-44d0-a14f-4cc6bf18a7c4").
		AddNetworking("HCM-1A", "sub-85ba01f6-02ec-4dfc-8884-ee0036c68a5b").
		WithScaling(1, 3).
		WithEnableDnsName(true).WithBuyMorePoc(false).WithPoc(false).WithEnableAutoRenew(true)

	lb, sdkerr := vngcloud.VNetworkGateway().V1().NetworkService().CreateEndpoint(opt)
	if sdkerr != nil {
		t.Fatalf("Expect nil but got %+v", sdkerr.Err())
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

	sdkerr := vngcloud.VNetworkGateway().V1().NetworkService().DeleteEndpointByID(opt)
	if sdkerr != nil {
		t.Fatalf("Expect nil but got %+v", sdkerr.ErrorCode())
	}

	t.Log("PASS")
}

func TestListEndpoints(t *testing.T) {
	vngcloud := validSdkConfig()
	opt := networkv1.NewListEndpointsRequest(1, 100).WithUuid("enp-9349271b-af44-4e39-8829-615d945fa6c2")

	lb, sdkerr := vngcloud.VNetworkGateway().V1().NetworkService().ListEndpoints(opt)
	if sdkerr != nil {
		t.Fatalf("Expect nil but got %+v", sdkerr.ErrorCode())
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

func TestListEndpointTags(t *testing.T) {
	vngcloud := validSuperSdkHcm03bConfig()
	opt := networkv1.NewListTagsByEndpointIDRequest(
		"95174",
		"pro-b2fff1cf-6d72-4643-a8e7-5907bc9e439c",
		"enp-3fe5d1e9-679e-4eb8-ad35-d9ce53243259",
	)

	lb, sdkerr := vngcloud.VNetworkGateway().InternalV1().NetworkService().ListTagsByEndpointID(opt)
	if sdkerr != nil {
		t.Logf("Expect nil but got %+v", sdkerr.ErrorCode())
	}

	if lb == nil {
		t.Logf("Expect not nil but got nil")
	}

	t.Log("Result: ", lb)
	t.Log("PASS")
}

func TestCreateEndpointTags(t *testing.T) {
	vngcloud := validSuperSdkConfig()
	opt := networkv1.NewCreateTagsWithEndpointIDRequest(
		"60108",
		"pro-88265bae-d2ef-424b-b8a7-9eeb08aec1f7",
		"enp-7e8e4476-feeb-414c-ac03-3501aae607d0",
	).AddTag("cuongdm3", "test")

	sdkerr := vngcloud.VNetworkGateway().InternalV1().NetworkService().CreateTagsWithEndpointID(opt)
	if sdkerr != nil {
		t.Logf("Expect nil but got %+v", sdkerr.ErrorCode())
	}

	t.Log("Result: ", sdkerr)
	t.Log("PASS")
}

func TestDeleteTagByEndpointID(t *testing.T) {
	vngcloud := validSuperSdkConfig()
	opt := networkv1.NewDeleteTagOfEndpointRequest(
		"60108",
		"pro-88265bae-d2ef-424b-b8a7-9eeb08aec1f7",
		"tag-6ceb41e1-47e9-43f0-94dd-521a1af870ee",
	)

	sdkerr := vngcloud.VNetworkGateway().InternalV1().NetworkService().DeleteTagOfEndpoint(opt)
	if sdkerr != nil {
		t.Logf("Expect nil but got %+v", sdkerr.ErrorCode())
	}

	t.Log("Result: ", sdkerr)
	t.Log("PASS")
}

func TestUpdateEndpointTag(t *testing.T) {
	vngcloud := validSuperSdkConfig()
	opt := networkv1.NewUpdateTagValueOfEndpointRequest(
		"60108",
		"pro-88265bae-d2ef-424b-b8a7-9eeb08aec1f7",
		"tag-c6d6e343-ed13-4bf1-bf2e-e63a1a5e0eab",
		"cuonghahahah",
	)

	sdkerr := vngcloud.VNetworkGateway().InternalV1().NetworkService().UpdateTagValueOfEndpoint(opt)
	if sdkerr != nil {
		t.Logf("Expect nil but got %+v", sdkerr.ErrorCode())
	}

	t.Log("Result: ", sdkerr)
	t.Log("PASS")
}
