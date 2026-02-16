package test

import (
	"bufio"
	"context"
	"errors"
	"os"
	"strings"
	"testing"

	"github.com/dannyota/greennode-community-sdk/v2/greennode"
	"github.com/dannyota/greennode-community-sdk/v2/greennode/sdkerror"
	identityv2 "github.com/dannyota/greennode-community-sdk/v2/greennode/services/identity/v2"
)

func readEnvFile(path string) (map[string]string, error) {
	data, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer data.Close()

	result := make(map[string]string)
	scanner := bufio.NewScanner(data)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		key, value, ok := strings.Cut(line, "=")
		if !ok {
			continue
		}
		result[strings.TrimSpace(key)] = strings.TrimSpace(value)
	}
	return result, scanner.Err()
}

func getEnv() (string, string) {
	envFile, _ := readEnvFile("./env.yaml")
	clientID := envFile["VNGCLOUD_CLIENT_ID"]
	clientSecret := envFile["VNGCLOUD_CLIENT_SECRET"]

	return clientID, clientSecret
}

func getValueOfEnv(key string) string {
	envFile, _ := readEnvFile("./env.yaml")
	value := envFile[key]
	return value
}

// newClientFromEnvKeys creates a client using the given env key names for client ID/secret
// and an optional set of config overrides.
func newClientFromEnvKeys(
	clientIDKey, clientSecretKey string,
	overrides ...func(*greennode.Config),
) *greennode.Client {
	cfg := greennode.Config{
		ClientID:     getValueOfEnv(clientIDKey),
		ClientSecret: getValueOfEnv(clientSecretKey),
		IAMEndpoint:  "https://iamapis.vngcloud.vn/accounts-api",
		RetryCount:   1,
		SleepDuration: 10,
	}
	for _, o := range overrides {
		o(&cfg)
	}
	c, _ := greennode.NewClient(context.Background(), cfg)
	return c
}

func validSdkConfig() *greennode.Client {
	clientID, clientSecret := getEnv()
	c, _ := greennode.NewClient(context.Background(), greennode.Config{
		ClientID:         clientID,
		ClientSecret:     clientSecret,
		UserID:           getValueOfEnv("VNGCLOUD_USER_ID"),
		ZoneID:           getValueOfEnv("VNGCLOUD_ZONE_ID"),
		ProjectID:        getValueOfEnv("VNGCLOUD_PROJECT_ID"),
		IAMEndpoint:      "https://iamapis.vngcloud.vn/accounts-api",
		VServerEndpoint:  "https://hcm-3.api.vngcloud.vn/vserver/vserver-gateway",
		VLBEndpoint:      "https://hcm-3.api.vngcloud.vn/vserver/vlb-gateway",
		VNetworkEndpoint: "https://hcm-3.console.vngcloud.vn/vserver/vnetwork-gateway/vnetwork",
		GLBEndpoint:      "https://glb.console.vngcloud.vn/glb-controller/",
		DNSEndpoint:      "https://vdns.api.vngcloud.vn/",
		RetryCount:       1,
		SleepDuration:    10,
	})
	return c
}

func validUserSdkConfig() *greennode.Client {
	c, _ := greennode.NewClient(context.Background(), greennode.Config{
		ClientID:         getValueOfEnv("USER_CLIENT_ID"),
		ClientSecret:     getValueOfEnv("USER_CLIENT_SECRET"),
		UserID:           getValueOfEnv("VNGCLOUD_USER_ID"),
		ZoneID:           getValueOfEnv("VNGCLOUD_ZONE_ID"),
		ProjectID:        getValueOfEnv("USER_PROJECT"),
		IAMEndpoint:      "https://iamapis.vngcloud.vn/accounts-api",
		VServerEndpoint:  "https://hcm-3.api.vngcloud.vn/vserver/vserver-gateway",
		VLBEndpoint:      "https://hcm-3.api.vngcloud.vn/vserver/vlb-gateway",
		VNetworkEndpoint: "https://hcm-3.console.vngcloud.vn/vserver/vnetwork-gateway/vnetwork",
		RetryCount:       1,
		SleepDuration:    10,
	})
	return c
}

func validAltUserSdkConfig() *greennode.Client {
	c, _ := greennode.NewClient(context.Background(), greennode.Config{
		ClientID:         getValueOfEnv("ALT_USER_CLIENT_ID"),
		ClientSecret:     getValueOfEnv("ALT_USER_CLIENT_SECRET"),
		ProjectID:        getValueOfEnv("ALT_USER_PROJECT_ID"),
		IAMEndpoint:      "https://iamapis.vngcloud.vn/accounts-api",
		VServerEndpoint:  "https://hcm-3.api.vngcloud.vn/vserver/vserver-gateway",
		VLBEndpoint:      "https://hcm-3.api.vngcloud.vn/vserver/vlb-gateway",
		VNetworkEndpoint: "https://hcm-3.console.vngcloud.vn/vserver/vnetwork-gateway/vnetwork",
		RetryCount:       1,
		SleepDuration:    10,
	})
	return c
}

func validSdkConfigHanRegion() *greennode.Client {
	clientID, clientSecret := getEnv()
	c, _ := greennode.NewClient(context.Background(), greennode.Config{
		ClientID:         clientID,
		ClientSecret:     clientSecret,
		UserID:           getValueOfEnv("VNGCLOUD_USER_ID"),
		ZoneID:           getValueOfEnv("VNGCLOUD_ZONE_ID"),
		ProjectID:        getValueOfEnv("HAN01_PROJECT_ID"),
		IAMEndpoint:      "https://iamapis.vngcloud.vn/accounts-api",
		VServerEndpoint:  "https://han-1.api.vngcloud.vn/vserver/vserver-gateway",
		VLBEndpoint:      "https://han-1.api.vngcloud.vn/vserver/vlb-gateway",
		VNetworkEndpoint: "https://hcm-3.console.vngcloud.vn/vserver/vnetwork-gateway/vnetwork",
		UserAgent:        "vngcloud-go-sdk",
		RetryCount:       1,
		SleepDuration:    10,
	})
	return c
}

func validHcm3bSdkConfig() *greennode.Client {
	c, _ := greennode.NewClient(context.Background(), greennode.Config{
		ClientID:         getValueOfEnv("HCM3B_CLIENT_ID"),
		ClientSecret:     getValueOfEnv("HCM3B_CLIENT_SECRET"),
		ZoneID:           getValueOfEnv("VNGCLOUD_ZONE_ID"),
		ProjectID:        getValueOfEnv("HCM3B_PROJECT_ID"),
		IAMEndpoint:      "https://iamapis.vngcloud.vn/accounts-api",
		VServerEndpoint:  "https://hcm-3.api.vngcloud.vn/vserver/vserver-gateway",
		VLBEndpoint:      "https://hcm-3.api.vngcloud.vn/vserver/vlb-gateway",
		VNetworkEndpoint: "https://hcm-3.console.vngcloud.vn/vserver/vnetwork-gateway/vnetwork",
		RetryCount:       1,
		SleepDuration:    10,
	})
	return c
}

func validHcm3bSuperSdkConfig() *greennode.Client {
	c, _ := greennode.NewClient(context.Background(), greennode.Config{
		ClientID:         getValueOfEnv("HCM3B_SUPER_CLIENT_ID"),
		ClientSecret:     getValueOfEnv("HCM3B_SUPER_CLIENT_SECRET"),
		ZoneID:           getValueOfEnv("VNGCLOUD_ZONE_ID"),
		ProjectID:        getValueOfEnv("HCM3B_PROJECT_ID"),
		IAMEndpoint:      "https://iamapis.vngcloud.vn/accounts-api",
		VServerEndpoint:  "https://hcm-3.api.vngcloud.vn/vserver/vserver-gateway",
		VLBEndpoint:      "https://hcm-3.api.vngcloud.vn/vserver/vlb-gateway",
		VNetworkEndpoint: "https://hcm-3.console.vngcloud.vn/vserver/vnetwork-gateway/vnetwork",
		RetryCount:       1,
		SleepDuration:    10,
	})
	return c
}

func validSecondaryUserSdkConfig() *greennode.Client {
	c, _ := greennode.NewClient(context.Background(), greennode.Config{
		ClientID:         getValueOfEnv("SECONDARY_CLIENT_ID"),
		ClientSecret:     getValueOfEnv("SECONDARY_CLIENT_SECRET"),
		ZoneID:           getValueOfEnv("VNGCLOUD_ZONE_ID"),
		ProjectID:        getValueOfEnv("SECONDARY_PROJECT_ID"),
		IAMEndpoint:      "https://iamapis.vngcloud.vn/accounts-api",
		VServerEndpoint:  "https://hcm-3.api.vngcloud.vn/vserver/vserver-gateway",
		VLBEndpoint:      "https://hcm-3.api.vngcloud.vn/vserver/vlb-gateway",
		VNetworkEndpoint: "https://hcm-3.console.vngcloud.vn/vserver/vnetwork-gateway/vnetwork",
		RetryCount:       1,
		SleepDuration:    10,
	})
	return c
}

func customerSdkConfig() *greennode.Client {
	c, _ := greennode.NewClient(context.Background(), greennode.Config{
		IAMEndpoint:     "https://iamapis.vngcloud.vn/accounts-api",
		VServerEndpoint: "https://hcm-3.api.vngcloud.vn/vserver/vserver-gateway",
		VLBEndpoint:     "https://hcm-3.api.vngcloud.vn/vserver/vlb-gateway",
		RetryCount:      1,
		SleepDuration:   10,
	})
	return c
}

func validTestProjectSdkConfig() *greennode.Client {
	c, _ := greennode.NewClient(context.Background(), greennode.Config{
		ClientID:        getValueOfEnv("TEST_PROJECT_CLIENT_ID"),
		ClientSecret:    getValueOfEnv("TEST_PROJECT_CLIENT_SECRET"),
		ProjectID:       getValueOfEnv("TEST_PROJECT_ID"),
		IAMEndpoint:     "https://iamapis.vngcloud.vn/accounts-api",
		VServerEndpoint: "https://hcm-3.api.vngcloud.vn/vserver/vserver-gateway",
		VLBEndpoint:     "https://hcm-3.api.vngcloud.vn/vserver/vlb-gateway",
		RetryCount:      1,
		SleepDuration:   10,
	})
	return c
}

func validSuperSdkConfig() *greennode.Client {
	c, _ := greennode.NewClient(context.Background(), greennode.Config{
		ClientID:         getValueOfEnv("VNGCLOUD_SUPER_CLIENT_ID"),
		ClientSecret:     getValueOfEnv("VNGCLOUD_SUPER_CLIENT_SECRET"),
		ZoneID:           getValueOfEnv("VNGCLOUD_ZONE_ID"),
		ProjectID:        getValueOfEnv("VNGCLOUD_PROJECT_ID"),
		IAMEndpoint:      "https://iamapis.vngcloud.vn/accounts-api",
		VServerEndpoint:  "https://hcm-3.api.vngcloud.vn/vserver/vserver-gateway",
		VLBEndpoint:      "https://hcm-3.api.vngcloud.vn/vserver/vlb-gateway",
		VNetworkEndpoint: "https://vnetwork-hcm03-api.vngcloud.vn/vnetwork-gateway",
		RetryCount:       1,
		SleepDuration:    10,
	})
	return c
}

func validSuperSdkHcm03bConfig() *greennode.Client {
	c, _ := greennode.NewClient(context.Background(), greennode.Config{
		ClientID:         getValueOfEnv("VNGCLOUD_PROD_HCM03B_CLIENT_ID"),
		ClientSecret:     getValueOfEnv("VNGCLOUD_PROD_HCM03B_CLIENT_SECRET"),
		ZoneID:           getValueOfEnv("VNGCLOUD_ZONE_ID"),
		ProjectID:        getValueOfEnv("VNGCLOUD_PROD_HCM03B_PROJECT_ID"),
		IAMEndpoint:      "https://iamapis.vngcloud.vn/accounts-api",
		VServerEndpoint:  "https://hcm-3.api.vngcloud.vn/vserver/vserver-gateway",
		VLBEndpoint:      "https://hcm-3.api.vngcloud.vn/vserver/vlb-gateway",
		VNetworkEndpoint: "https://vnetwork-hcm03-api.vngcloud.vn/vnetwork-gateway",
		RetryCount:       1,
		SleepDuration:    10,
	})
	return c
}

func validSuperWithTargetProjectSdkConfig() *greennode.Client {
	c, _ := greennode.NewClient(context.Background(), greennode.Config{
		ClientID:        getValueOfEnv("VNGCLOUD_SUPER_CLIENT_ID"),
		ClientSecret:    getValueOfEnv("VNGCLOUD_SUPER_CLIENT_SECRET"),
		ProjectID:       getValueOfEnv("TARGET_USER_PROJECT_ID"),
		IAMEndpoint:     "https://iamapis.vngcloud.vn/accounts-api",
		VServerEndpoint: "https://hcm-3.api.vngcloud.vn/vserver/vserver-gateway",
		VLBEndpoint:     "https://hcm-3.api.vngcloud.vn/vserver/vlb-gateway",
		RetryCount:      1,
		SleepDuration:   10,
	})
	return c
}

func validTargetUserSdkConfig() *greennode.Client {
	c, _ := greennode.NewClient(context.Background(), greennode.Config{
		ClientID:         getValueOfEnv("TARGET_USER_CLIENT_ID"),
		ClientSecret:     getValueOfEnv("TARGET_USER_CLIENT_SECRET"),
		ProjectID:        getValueOfEnv("TARGET_USER_PROJECT_ID"),
		IAMEndpoint:      "https://iamapis.vngcloud.vn/accounts-api",
		VServerEndpoint:  "https://hcm-3.api.vngcloud.vn/vserver/vserver-gateway",
		VLBEndpoint:      "https://hcm-3.api.vngcloud.vn/vserver/vlb-gateway",
		VNetworkEndpoint: "https://vnetwork-hcm03-api.vngcloud.vn/vnetwork-gateway",
		RetryCount:       1,
		SleepDuration:    10,
	})
	return c
}

func validSuperSdkConfig2() *greennode.Client {
	c, _ := greennode.NewClient(context.Background(), greennode.Config{
		ClientID:        getValueOfEnv("VNGCLOUD_SUPER_CLIENT_ID"),
		ClientSecret:    getValueOfEnv("VNGCLOUD_SUPER_CLIENT_SECRET"),
		ProjectID:       getValueOfEnv("SECONDARY_PROJECT_ID"),
		IAMEndpoint:     "https://iamapis.vngcloud.vn/accounts-api",
		VServerEndpoint: "https://hcm-3.api.vngcloud.vn/vserver/vserver-gateway",
		VLBEndpoint:     "https://hcm-3.api.vngcloud.vn/vserver/vlb-gateway",
		RetryCount:      1,
		SleepDuration:   10,
	})
	return c
}

func validServiceAccountSdkConfig() *greennode.Client {
	c, _ := greennode.NewClient(context.Background(), greennode.Config{
		ClientID:        getValueOfEnv("SERVICE_ACCOUNT_CLIENT_ID"),
		ClientSecret:    getValueOfEnv("SERVICE_ACCOUNT_CLIENT_SECRET"),
		ProjectID:       getValueOfEnv("SERVICE_ACCOUNT_PROJECT_ID"),
		IAMEndpoint:     "https://iamapis.vngcloud.vn/accounts-api",
		VServerEndpoint: "https://hcm-3.api.vngcloud.vn/vserver/vserver-gateway",
		RetryCount:      1,
		SleepDuration:   10,
	})
	return c
}

func invalidSdkConfig() *greennode.Client {
	c, _ := greennode.NewClient(context.Background(), greennode.Config{
		ClientID:        "invalid-id",
		ClientSecret:    "invalid-secret",
		IAMEndpoint:     "https://iamapis.vngcloud.vn/accounts-api",
		VServerEndpoint: "https://hcm-3.api.vngcloud.vn/vserver/vserver-gateway",
		RetryCount:      1,
		SleepDuration:   10,
	})
	return c
}

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
