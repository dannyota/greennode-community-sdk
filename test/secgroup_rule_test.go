//go:build integration

package test

import (
	"context"
	"testing"

	networkv2 "danny.vn/greennode/services/network/v2"
)

func TestCreateSecgroupRuleSuccess(t *testing.T) {
	vngcloud := validSdkConfig()
	opt := networkv2.NewCreateSecgroupRuleRequest(
		networkv2.SecgroupRuleDirectionIngress,
		networkv2.SecgroupRuleEtherTypeIPv4,
		networkv2.SecgroupRuleProtocolTCP,
		6444, 6450,
		"0.0.0.0",
		"secg-f5a2de4d-f8a2-4269-bcba-7c50f09ee838",
		"test2",
	)
	secgroupRule, err := vngcloud.Network.CreateSecgroupRule(context.Background(), opt)

	if err != nil {
		t.Fatalf("Expect error to be nil but got %+v", err)
	}

	if secgroupRule == nil {
		t.Fatalf("Expect portal not to be nil but got nil")
	}

	t.Log("RESULT:", secgroupRule)
	t.Log("PASS")
}

func TestCreateSecgroupRuleFailure(t *testing.T) {
	vngcloud := validSdkConfig()
	opt := networkv2.NewCreateSecgroupRuleRequest(
		networkv2.SecgroupRuleDirectionIngress,
		networkv2.SecgroupRuleEtherTypeIPv4,
		networkv2.SecgroupRuleProtocolTCP,
		6443, 6443,
		"0.0.0.0",
		"secg-f5a2de4d-f8a2-4269-bcba-7c50f09ee840",
		"test",
	)
	secgroupRule, err := vngcloud.Network.CreateSecgroupRule(context.Background(), opt)

	if err == nil {
		t.Errorf("Expect error not to be nil but got nil")
	}

	if secgroupRule != nil {
		t.Errorf("Expect portal to be nil but got %+v", secgroupRule)
	}

	t.Log("RESULT:", err)
	t.Log("PASS")
}

func TestDeleteSecgroupRuleFailure(t *testing.T) {
	vngcloud := validSdkConfig()
	opt := networkv2.NewDeleteSecgroupRuleByIDRequest("secr-8c44dc7f-3916-4952-8a01-884ad9360f00")
	err := vngcloud.Network.DeleteSecgroupRuleByID(context.Background(), opt)

	if err == nil {
		t.Errorf("Expect error not to be nil but got nil")
	}

	t.Log("RESULT:", err)
	t.Log("PASS")
}

func TestDeleteSecgroupRuleSuccess(t *testing.T) {
	vngcloud := validSdkConfig()
	opt := networkv2.NewDeleteSecgroupRuleByIDRequest("secr-d8c83235-93ae-4de1-85c2-365a32494121")
	err := vngcloud.Network.DeleteSecgroupRuleByID(context.Background(), opt)

	if err != nil {
		t.Fatalf("Expect error to be nil but got %+v", err)
	}

	t.Log("RESULT:", err)
	t.Log("PASS")
}

func TestListSecgroupRulesBySecgroupIDFailure(t *testing.T) {
	vngcloud := validSdkConfig()
	opt := networkv2.NewListSecgroupRulesBySecgroupIDRequest("secg-f5a2de4d-f8a2-4269-bcba-7c50f09ee840")
	rules, err := vngcloud.Network.ListSecgroupRulesBySecgroupID(context.Background(), opt)

	if err != nil {
		t.Fatalf("Expect error to be nil but got %+v", err)
	}

	if rules == nil {
		t.Fatalf("Expect rules not to be nil but got nil")
	}

	t.Log("RESULT:", rules)
	t.Log("PASS")
}

func TestListSecgroupRulesBySecgroupIDSuccess(t *testing.T) {
	vngcloud := validSdkConfig()
	opt := networkv2.NewListSecgroupRulesBySecgroupIDRequest("secg-f5a2de4d-f8a2-4269-bcba-7c50f09ee838")
	rules, err := vngcloud.Network.ListSecgroupRulesBySecgroupID(context.Background(), opt)

	if err != nil {
		t.Fatalf("Expect error to be nil but got %+v", err)
	}

	if rules == nil {
		t.Fatalf("Expect rules not to be nil but got nil")
	}

	t.Log("RESULT:", rules)
	t.Log("PASS")
}
