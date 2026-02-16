//go:build integration

package test

import (
	"context"
	"errors"
	"testing"

	"github.com/dannyota/greennode-community-sdk/v2/greennode"
	"github.com/dannyota/greennode-community-sdk/v2/greennode/sdkerror"
	identityv2 "github.com/dannyota/greennode-community-sdk/v2/greennode/services/identity/v2"
)

func TestAuthenFailed(t *testing.T) {
	clientID := "cc136360-709c-4248-9358-e8e96c74480a"
	clientSecret := "invalid-secret"

	vngcloud, _ := greennode.NewClient(context.Background(), greennode.Config{
		ClientID:        clientID,
		ClientSecret:    clientSecret,
		IAMEndpoint:     "https://iamapis.vngcloud.vn/accounts-api",
		VServerEndpoint: "https://hcm-3.api.vngcloud.vn/vserver/vserver-gateway",
		RetryCount:      1,
		SleepDuration:   10,
	})
	opt := identityv2.NewGetAccessTokenRequest(clientID, clientSecret)
	token, err := vngcloud.Identity.GetAccessToken(context.Background(), opt)

	if err == nil {
		t.Error("Error MUST not be nil")
	}

	if token != nil {
		t.Error("Token MUST be nil")
	}

	var sdkErr *sdkerror.SdkError
	if errors.As(err, &sdkErr) {
		if !sdkErr.IsError(sdkerror.EcAuthenticationFailed) {
			t.Error("Error MUST be VngCloudIamAuthenticationFailed")
		}
	} else {
		t.Error("Expected SdkError")
	}

	t.Log("RESULT:", err)
	t.Log("PASS")
}

func TestAuthenPass(t *testing.T) {
	clientID, clientSecret := getEnv()
	vngcloud := validSdkConfig()
	opt := identityv2.NewGetAccessTokenRequest(clientID, clientSecret)
	token, err := vngcloud.Identity.GetAccessToken(context.Background(), opt)

	if err != nil {
		t.Fatalf("Expected no error but got: %v", err)
	}
	if token == nil {
		t.Fatal("Expected token to be non-nil")
	}

	t.Log("RESULT:", token)
}

func TestSuperAdminAuthenPass(t *testing.T) {
	clientID, clientSecret := getValueOfEnv("VNGCLOUD_SUPER_CLIENT_ID"), getValueOfEnv("VNGCLOUD_SUPER_CLIENT_SECRET")
	vngcloud := validSdkConfig()
	opt := identityv2.NewGetAccessTokenRequest(clientID, clientSecret)
	token, err := vngcloud.Identity.GetAccessToken(context.Background(), opt)

	if err != nil || token == nil {
		t.Fatal("This testcase MUST pass")
	}

	t.Log("RESULT:", token.Token)
}

func TestSecondaryUserAuthenPass(t *testing.T) {
	clientID, clientSecret := getValueOfEnv("SECONDARY_CLIENT_ID"), getValueOfEnv("SECONDARY_CLIENT_SECRET")
	vngcloud := validSuperSdkConfig2()
	opt := identityv2.NewGetAccessTokenRequest(clientID, clientSecret)
	token, err := vngcloud.Identity.GetAccessToken(context.Background(), opt)

	if err != nil || token == nil {
		t.Fatal("This testcase MUST pass")
	}

	t.Log("RESULT:", token)
}
