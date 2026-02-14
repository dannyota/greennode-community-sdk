package test

import (
	"testing"

	networkv2 "github.com/dannyota/greennode-community-sdk/v2/greennode/services/network/v2"
)

func TestGetSubnetByIDSuccess(t *testing.T) {
	vngcloud := validSdkConfig()
	opt := networkv2.NewGetSubnetByIDRequest(getValueOfEnv("NETWORK_ID"), getValueOfEnv("SUBNET_ID"))
	network, err := vngcloud.VServerGateway().V2().NetworkService().GetSubnetByID(opt)

	if err != nil {
		t.Fatalf("Expect error to be nil but got %+v", err)
	}

	if network == nil {
		t.Fatalf("Expect portal not to be nil but got nil")
	}

	t.Log("RESULT:", network)
	t.Log("PASS")
}

func TestUpdateSubnetByID(t *testing.T) {
	vngcloud := validSdkConfig()
	updateBody := networkv2.UpdateSubnetBody{
		Name: "subnet-1",
		CIDR: "10.30.0.0/24",
		SecondarySubnetRequests: []networkv2.SecondarySubnetUpdateBody{
			{Name: "subnet3", CIDR: "10.30.6.0/24"},
			{Name: "subnet2", CIDR: "10.30.7.0/24"},
		},
	}

	opt := networkv2.NewUpdateSubnetByIDRequest(getValueOfEnv("NETWORK_ID"), getValueOfEnv("SUBNET_ID"), &updateBody)
	network, err := vngcloud.VServerGateway().V2().NetworkService().UpdateSubnetByID(opt)

	if err != nil {
		t.Fatalf("Expect error to be nil but got %+v", err)
	}

	if network == nil {
		t.Fatalf("Expect portal not to be nil but got nil")
	}

	t.Log("RESULT:", network)
	t.Log("PASS")
}
