package test

import (
	"bufio"
	"context"
	"errors"
	"os"
	"strings"
	"testing"

	"github.com/dannyota/greennode-community-sdk/v2/client"
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

// newClientFromEnvKeys creates a client using the given env key names for client ID/secret,
// an optional project ID key, and a list of endpoint configurators.
func newClientFromEnvKeys(
	clientIDKey, clientSecretKey string,
	opts ...func(*client.SdkConfigure),
) *client.Client {
	cid, csec := getValueOfEnv(clientIDKey), getValueOfEnv(clientSecretKey)
	sdkCfg := client.NewSdkConfigure().
		WithClientID(cid).
		WithClientSecret(csec).
		WithIAMEndpoint("https://iamapis.vngcloud.vn/accounts-api")
	for _, o := range opts {
		o(sdkCfg)
	}
	return client.NewClient().WithRetryCount(1).WithSleep(10).Configure(sdkCfg)
}

func getValueOfEnv(key string) string {
	envFile, _ := readEnvFile("./env.yaml")
	value := envFile[key]
	return value
}

func validSdkConfig() *client.Client {
	clientID, clientSecret := getEnv()
	sdkConfig := client.NewSdkConfigure().
		WithClientID(clientID).
		WithClientSecret(clientSecret).
		WithUserID(getValueOfEnv("VNGCLOUD_USER_ID")).
		WithZoneID(getValueOfEnv("VNGCLOUD_ZONE_ID")).
		WithProjectID(getValueOfEnv("VNGCLOUD_PROJECT_ID")).
		WithIAMEndpoint("https://iamapis.vngcloud.vn/accounts-api").
		WithVServerEndpoint("https://hcm-3.api.vngcloud.vn/vserver/vserver-gateway").
		WithVLBEndpoint("https://hcm-3.api.vngcloud.vn/vserver/vlb-gateway").
		WithVNetworkEndpoint("https://vnetwork-hcm03.vngcloud.vn/vnetwork-gateway/vnetwork").
		WithVNetworkEndpoint("https://hcm-3.console.vngcloud.vn/vserver/vnetwork-gateway/vnetwork").
		WithGLBEndpoint("https://glb.console.vngcloud.vn/glb-controller/").
		WithVDnsEndpoint("https://vdns.api.vngcloud.vn/")

	return client.NewClient().WithRetryCount(1).WithSleep(10).Configure(sdkConfig)
}

func validUserSdkConfig() *client.Client {
	clientID, clientSecret := getValueOfEnv("USER_CLIENT_ID"), getValueOfEnv("USER_CLIENT_SECRET")
	sdkConfig := client.NewSdkConfigure().
		WithClientID(clientID).
		WithClientSecret(clientSecret).
		WithUserID(getValueOfEnv("VNGCLOUD_USER_ID")).
		WithZoneID(getValueOfEnv("VNGCLOUD_ZONE_ID")).
		WithProjectID(getValueOfEnv("USER_PROJECT")).
		WithIAMEndpoint("https://iamapis.vngcloud.vn/accounts-api").
		WithVServerEndpoint("https://hcm-3.api.vngcloud.vn/vserver/vserver-gateway").
		WithVLBEndpoint("https://hcm-3.api.vngcloud.vn/vserver/vlb-gateway").
		WithVNetworkEndpoint("https://vnetwork-hcm03.vngcloud.vn/vnetwork-gateway/vnetwork").
		WithVNetworkEndpoint("https://hcm-3.console.vngcloud.vn/vserver/vnetwork-gateway/vnetwork")

	return client.NewClient().WithRetryCount(1).WithSleep(10).Configure(sdkConfig)
}

func validAltUserSdkConfig() *client.Client {
	clientID, clientSecret := getValueOfEnv("ALT_USER_CLIENT_ID"), getValueOfEnv("ALT_USER_CLIENT_SECRET")
	sdkConfig := client.NewSdkConfigure().
		WithClientID(clientID).
		WithClientSecret(clientSecret).
		WithProjectID(getValueOfEnv("ALT_USER_PROJECT_ID")).
		WithIAMEndpoint("https://iamapis.vngcloud.vn/accounts-api").
		WithVServerEndpoint("https://hcm-3.api.vngcloud.vn/vserver/vserver-gateway").
		WithVLBEndpoint("https://hcm-3.api.vngcloud.vn/vserver/vlb-gateway").
		WithVNetworkEndpoint("https://vnetwork-hcm03.vngcloud.vn/vnetwork-gateway/vnetwork").
		WithVNetworkEndpoint("https://hcm-3.console.vngcloud.vn/vserver/vnetwork-gateway/vnetwork")

	return client.NewClient().WithRetryCount(1).WithSleep(10).Configure(sdkConfig)
}

func validSdkConfigHanRegion() *client.Client {
	clientID, clientSecret := getEnv()
	sdkConfig := client.NewSdkConfigure().
		WithClientID(clientID).
		WithClientSecret(clientSecret).
		WithUserID(getValueOfEnv("VNGCLOUD_USER_ID")).
		WithZoneID(getValueOfEnv("VNGCLOUD_ZONE_ID")).
		WithProjectID(getValueOfEnv("HAN01_PROJECT_ID")).
		WithIAMEndpoint("https://iamapis.vngcloud.vn/accounts-api").
		WithVServerEndpoint("https://han-1.api.vngcloud.vn/vserver/vserver-gateway").
		WithVLBEndpoint("https://han-1.api.vngcloud.vn/vserver/vlb-gateway").
		WithVNetworkEndpoint("https://vnetwork-hcm03.vngcloud.vn/vnetwork-gateway/vnetwork").
		WithVNetworkEndpoint("https://hcm-3.console.vngcloud.vn/vserver/vnetwork-gateway/vnetwork").
		WithUserAgent("vngcloud-go-sdk")

	return client.NewClient().WithRetryCount(1).WithSleep(10).Configure(sdkConfig)
}

func validHcm3bSdkConfig() *client.Client {
	sdkConfig := client.NewSdkConfigure().
		WithClientID(getValueOfEnv("HCM3B_CLIENT_ID")).
		WithClientSecret(getValueOfEnv("HCM3B_CLIENT_SECRET")).
		WithZoneID(getValueOfEnv("VNGCLOUD_ZONE_ID")).
		WithProjectID(getValueOfEnv("HCM3B_PROJECT_ID")).
		WithIAMEndpoint("https://iamapis.vngcloud.vn/accounts-api").
		WithVServerEndpoint("https://hcm-3.api.vngcloud.vn/vserver/vserver-gateway").
		WithVLBEndpoint("https://hcm-3.api.vngcloud.vn/vserver/vlb-gateway").
		WithVNetworkEndpoint("https://vnetwork-hcm03.vngcloud.vn/vnetwork-gateway/vnetwork").
		WithVNetworkEndpoint("https://hcm-3.console.vngcloud.vn/vserver/vnetwork-gateway/vnetwork")

	return client.NewClient().WithRetryCount(1).WithSleep(10).Configure(sdkConfig)
}

func validHcm3bSuperSdkConfig() *client.Client {
	sdkConfig := client.NewSdkConfigure().
		WithClientID(getValueOfEnv("HCM3B_SUPER_CLIENT_ID")).
		WithClientSecret(getValueOfEnv("HCM3B_SUPER_CLIENT_SECRET")).
		WithZoneID(getValueOfEnv("VNGCLOUD_ZONE_ID")).
		WithProjectID(getValueOfEnv("HCM3B_PROJECT_ID")).
		WithIAMEndpoint("https://iamapis.vngcloud.vn/accounts-api").
		WithVServerEndpoint("https://hcm-3.api.vngcloud.vn/vserver/vserver-gateway").
		WithVLBEndpoint("https://hcm-3.api.vngcloud.vn/vserver/vlb-gateway").
		WithVNetworkEndpoint("https://vnetwork-hcm03.vngcloud.vn/vnetwork-gateway/vnetwork").
		WithVNetworkEndpoint("https://hcm-3.console.vngcloud.vn/vserver/vnetwork-gateway/vnetwork")

	return client.NewClient().WithRetryCount(1).WithSleep(10).Configure(sdkConfig)
}

func validSecondaryUserSdkConfig() *client.Client {
	sdkConfig := client.NewSdkConfigure().
		WithClientID(getValueOfEnv("SECONDARY_CLIENT_ID")).
		WithClientSecret(getValueOfEnv("SECONDARY_CLIENT_SECRET")).
		WithZoneID(getValueOfEnv("VNGCLOUD_ZONE_ID")).
		WithProjectID(getValueOfEnv("SECONDARY_PROJECT_ID")).
		WithIAMEndpoint("https://iamapis.vngcloud.vn/accounts-api").
		WithVServerEndpoint("https://hcm-3.api.vngcloud.vn/vserver/vserver-gateway").
		WithVLBEndpoint("https://hcm-3.api.vngcloud.vn/vserver/vlb-gateway").
		WithVNetworkEndpoint("https://vnetwork-hcm03.vngcloud.vn/vnetwork-gateway/vnetwork").
		WithVNetworkEndpoint("https://hcm-3.console.vngcloud.vn/vserver/vnetwork-gateway/vnetwork")

	return client.NewClient().WithRetryCount(1).WithSleep(10).Configure(sdkConfig)
}

func customerSdkConfig() *client.Client {
	sdkConfig := client.NewSdkConfigure().
		WithIAMEndpoint("https://iamapis.vngcloud.vn/accounts-api").
		WithVServerEndpoint("https://hcm-3.api.vngcloud.vn/vserver/vserver-gateway").
		WithVLBEndpoint("https://hcm-3.api.vngcloud.vn/vserver/vlb-gateway")

	return client.NewClient().WithRetryCount(1).WithSleep(10).Configure(sdkConfig)
}

func validTestProjectSdkConfig() *client.Client {
	clientID, clientSecret := getValueOfEnv("TEST_PROJECT_CLIENT_ID"), getValueOfEnv("TEST_PROJECT_CLIENT_SECRET")
	sdkConfig := client.NewSdkConfigure().
		WithClientID(clientID).
		WithClientSecret(clientSecret).
		WithProjectID(getValueOfEnv("TEST_PROJECT_ID")).
		WithIAMEndpoint("https://iamapis.vngcloud.vn/accounts-api").
		WithVServerEndpoint("https://hcm-3.api.vngcloud.vn/vserver/vserver-gateway").
		WithVLBEndpoint("https://hcm-3.api.vngcloud.vn/vserver/vlb-gateway")

	return client.NewClient().WithRetryCount(1).WithSleep(10).Configure(sdkConfig)
}

func validSuperSdkConfig() *client.Client {
	clientID, clientSecret := getValueOfEnv("VNGCLOUD_SUPER_CLIENT_ID"), getValueOfEnv("VNGCLOUD_SUPER_CLIENT_SECRET")
	sdkConfig := client.NewSdkConfigure().
		WithClientID(clientID).
		WithClientSecret(clientSecret).
		WithZoneID(getValueOfEnv("VNGCLOUD_ZONE_ID")).
		WithProjectID(getValueOfEnv("VNGCLOUD_PROJECT_ID")).
		WithIAMEndpoint("https://iamapis.vngcloud.vn/accounts-api").
		WithVServerEndpoint("https://hcm-3.api.vngcloud.vn/vserver/vserver-gateway").
		WithVLBEndpoint("https://hcm-3.api.vngcloud.vn/vserver/vlb-gateway").
		WithVNetworkEndpoint("https://vnetwork-hcm03-api.vngcloud.vn/vnetwork-gateway")

	return client.NewClient().WithRetryCount(1).WithSleep(10).Configure(sdkConfig)
}

func validSuperSdkHcm03bConfig() *client.Client {
	clientID, clientSecret := getValueOfEnv("VNGCLOUD_PROD_HCM03B_CLIENT_ID"), getValueOfEnv("VNGCLOUD_PROD_HCM03B_CLIENT_SECRET")
	sdkConfig := client.NewSdkConfigure().
		WithClientID(clientID).
		WithClientSecret(clientSecret).
		WithZoneID(getValueOfEnv("VNGCLOUD_ZONE_ID")).
		WithProjectID(getValueOfEnv("VNGCLOUD_PROD_HCM03B_PROJECT_ID")).
		WithIAMEndpoint("https://iamapis.vngcloud.vn/accounts-api").
		WithVServerEndpoint("https://hcm-3.api.vngcloud.vn/vserver/vserver-gateway").
		WithVLBEndpoint("https://hcm-3.api.vngcloud.vn/vserver/vlb-gateway").
		WithVNetworkEndpoint("https://vnetwork-hcm03-api.vngcloud.vn/vnetwork-gateway")

	return client.NewClient().WithRetryCount(1).WithSleep(10).Configure(sdkConfig)
}

func validSuperWithTargetProjectSdkConfig() *client.Client {
	clientID, clientSecret := getValueOfEnv("VNGCLOUD_SUPER_CLIENT_ID"), getValueOfEnv("VNGCLOUD_SUPER_CLIENT_SECRET")
	sdkConfig := client.NewSdkConfigure().
		WithClientID(clientID).
		WithClientSecret(clientSecret).
		WithProjectID(getValueOfEnv("TARGET_USER_PROJECT_ID")).
		WithIAMEndpoint("https://iamapis.vngcloud.vn/accounts-api").
		WithVServerEndpoint("https://hcm-3.api.vngcloud.vn/vserver/vserver-gateway").
		WithVLBEndpoint("https://hcm-3.api.vngcloud.vn/vserver/vlb-gateway")

	return client.NewClient().WithRetryCount(1).WithSleep(10).Configure(sdkConfig)
}

func validTargetUserSdkConfig() *client.Client {
	clientID, clientSecret := getValueOfEnv("TARGET_USER_CLIENT_ID"), getValueOfEnv("TARGET_USER_CLIENT_SECRET")
	sdkConfig := client.NewSdkConfigure().
		WithClientID(clientID).
		WithClientSecret(clientSecret).
		WithProjectID(getValueOfEnv("TARGET_USER_PROJECT_ID")).
		WithIAMEndpoint("https://iamapis.vngcloud.vn/accounts-api").
		WithVServerEndpoint("https://hcm-3.api.vngcloud.vn/vserver/vserver-gateway").
		WithVLBEndpoint("https://hcm-3.api.vngcloud.vn/vserver/vlb-gateway").
		WithVNetworkEndpoint("https://vnetwork-hcm03-api.vngcloud.vn/vnetwork-gateway")

	return client.NewClient().WithRetryCount(1).WithSleep(10).Configure(sdkConfig)
}

func validSuperSdkConfig2() *client.Client {
	clientID, clientSecret := getValueOfEnv("VNGCLOUD_SUPER_CLIENT_ID"), getValueOfEnv("VNGCLOUD_SUPER_CLIENT_SECRET")
	sdkConfig := client.NewSdkConfigure().
		WithClientID(clientID).
		WithClientSecret(clientSecret).
		WithProjectID(getValueOfEnv("SECONDARY_PROJECT_ID")).
		WithIAMEndpoint("https://iamapis.vngcloud.vn/accounts-api").
		WithVServerEndpoint("https://hcm-3.api.vngcloud.vn/vserver/vserver-gateway").
		WithVLBEndpoint("https://hcm-3.api.vngcloud.vn/vserver/vlb-gateway")

	return client.NewClient().WithRetryCount(1).WithSleep(10).Configure(sdkConfig)
}

func validServiceAccountSdkConfig() *client.Client {
	clientID, clientSecret := getValueOfEnv("SERVICE_ACCOUNT_CLIENT_ID"), getValueOfEnv("SERVICE_ACCOUNT_CLIENT_SECRET")
	sdkConfig := client.NewSdkConfigure().
		WithClientID(clientID).
		WithClientSecret(clientSecret).
		WithProjectID(getValueOfEnv("SERVICE_ACCOUNT_PROJECT_ID")).
		WithIAMEndpoint("https://iamapis.vngcloud.vn/accounts-api").
		WithVServerEndpoint("https://hcm-3.api.vngcloud.vn/vserver/vserver-gateway")

	return client.NewClient().WithRetryCount(1).WithSleep(10).Configure(sdkConfig)
}

func invalidSdkConfig() *client.Client {
	clientID := "invalid-id"
	clientSecret := "invalid-secret"
	sdkConfig := client.NewSdkConfigure().
		WithClientID(clientID).
		WithClientSecret(clientSecret).
		WithIAMEndpoint("https://iamapis.vngcloud.vn/accounts-api").
		WithVServerEndpoint("https://hcm-3.api.vngcloud.vn/vserver/vserver-gateway")

	return client.NewClient().WithRetryCount(1).WithSleep(10).Configure(sdkConfig)
}

func TestAuthenFailed(t *testing.T) {
	clientID := "cc136360-709c-4248-9358-e8e96c74480a"
	clientSecret := "invalid-secret"

	sdkConfig := client.NewSdkConfigure().
		WithClientID(clientID).
		WithClientSecret(clientSecret).
		WithIAMEndpoint("https://iamapis.vngcloud.vn/accounts-api").
		WithVServerEndpoint("https://hcm-3.api.vngcloud.vn/vserver/vserver-gateway")

	vngcloud := client.NewClient().WithRetryCount(1).WithSleep(10).Configure(sdkConfig)
	opt := identityv2.NewGetAccessTokenRequest(clientID, clientSecret)
	token, err := vngcloud.IAMGateway().V2().IdentityService().GetAccessToken(context.Background(), opt)

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
	token, err := vngcloud.IAMGateway().V2().IdentityService().GetAccessToken(context.Background(), opt)

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
	token, err := vngcloud.IAMGateway().V2().IdentityService().GetAccessToken(context.Background(), opt)

	if err != nil || token == nil {
		t.Fatal("This testcase MUST pass")
	}

	t.Log("RESULT:", token.Token)
}

func TestSecondaryUserAuthenPass(t *testing.T) {
	clientID, clientSecret := getValueOfEnv("SECONDARY_CLIENT_ID"), getValueOfEnv("SECONDARY_CLIENT_SECRET")
	vngcloud := validSuperSdkConfig2()
	opt := identityv2.NewGetAccessTokenRequest(clientID, clientSecret)
	token, err := vngcloud.IAMGateway().V2().IdentityService().GetAccessToken(context.Background(), opt)

	if err != nil || token == nil {
		t.Fatal("This testcase MUST pass")
	}

	t.Log("RESULT:", token)
}
