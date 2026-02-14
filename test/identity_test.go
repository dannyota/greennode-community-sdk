package test

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/dannyota/greennode-community-sdk/v2/client"
	sdkerror "github.com/dannyota/greennode-community-sdk/v2/greennode/sdk_error"
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
	clientId := envFile["VNGCLOUD_CLIENT_ID"]
	clientSecret := envFile["VNGCLOUD_CLIENT_SECRET"]

	return clientId, clientSecret
}

func getEnvCuongDm4() (string, string) {
	envFile, _ := readEnvFile("./env.yaml")
	clientId := envFile["CUONGDM4_CLIENT_ID"]
	clientSecret := envFile["CUONGDM4_CLIENT_SECRET"]

	return clientId, clientSecret
}

func getEnvDevOps() (string, string) {
	envFile, _ := readEnvFile("./env.yaml")
	clientId := envFile["CLIENT_ID_DEVOPS"]
	clientSecret := envFile["CLIENT_SECRET_DEVOPS"]

	return clientId, clientSecret
}

func getValueOfEnv(pkey string) string {
	envFile, _ := readEnvFile("./env.yaml")
	value := envFile[pkey]
	return value
}

func validSdkConfig() client.IClient {
	clientId, clientSecret := getEnv()
	sdkConfig := client.NewSdkConfigure().
		WithClientId(clientId).
		WithClientSecret(clientSecret).
		WithUserId(getValueOfEnv("VNGCLOUD_USER_ID")).
		WithZoneId(getValueOfEnv("VNGCLOUD_ZONE_ID")).
		WithProjectId(getValueOfEnv("VNGCLOUD_PROJECT_ID")).
		WithIamEndpoint("https://iamapis.vngcloud.vn/accounts-api").
		WithVServerEndpoint("https://hcm-3.api.vngcloud.vn/vserver/vserver-gateway").
		WithVLBEndpoint("https://hcm-3.api.vngcloud.vn/vserver/vlb-gateway").
		WithVNetworkEndpoint("https://vnetwork-hcm03.vngcloud.vn/vnetwork-gateway/vnetwork").
		WithVNetworkEndpoint("https://hcm-3.console.vngcloud.vn/vserver/vnetwork-gateway/vnetwork").
		WithGLBEndpoint("https://glb.console.vngcloud.vn/glb-controller/").
		WithVDnsEndpoint("https://vdns.api.vngcloud.vn/")

	return client.NewClient(context.TODO()).WithRetryCount(1).WithSleep(10).Configure(sdkConfig)
}

func validUserSdkConfig() client.IClient {
	clientId, clientSecret := getValueOfEnv("USER_CLIENT_ID"), getValueOfEnv("USER_CLIENT_SECRET")
	sdkConfig := client.NewSdkConfigure().
		WithClientId(clientId).
		WithClientSecret(clientSecret).
		WithUserId(getValueOfEnv("VNGCLOUD_USER_ID")).
		WithZoneId(getValueOfEnv("VNGCLOUD_ZONE_ID")).
		WithProjectId(getValueOfEnv("USER_PROJECT")).
		WithIamEndpoint("https://iamapis.vngcloud.vn/accounts-api").
		WithVServerEndpoint("https://hcm-3.api.vngcloud.vn/vserver/vserver-gateway").
		WithVLBEndpoint("https://hcm-3.api.vngcloud.vn/vserver/vlb-gateway").
		WithVNetworkEndpoint("https://vnetwork-hcm03.vngcloud.vn/vnetwork-gateway/vnetwork").
		WithVNetworkEndpoint("https://hcm-3.console.vngcloud.vn/vserver/vnetwork-gateway/vnetwork")

	return client.NewClient(context.TODO()).WithRetryCount(1).WithSleep(10).Configure(sdkConfig)
}

func validUserSdkConfigForCuongDm4() client.IClient {
	clientId, clientSecret := getEnvCuongDm4()
	sdkConfig := client.NewSdkConfigure().
		WithClientId(clientId).
		WithClientSecret(clientSecret).
		WithProjectId("pro-462803f3-6858-466f-bf05-df2b33faa360").
		WithIamEndpoint("https://iamapis.vngcloud.vn/accounts-api").
		WithVServerEndpoint("https://hcm-3.api.vngcloud.vn/vserver/vserver-gateway").
		WithVLBEndpoint("https://hcm-3.api.vngcloud.vn/vserver/vlb-gateway").
		WithVNetworkEndpoint("https://vnetwork-hcm03.vngcloud.vn/vnetwork-gateway/vnetwork").
		WithVNetworkEndpoint("https://hcm-3.console.vngcloud.vn/vserver/vnetwork-gateway/vnetwork")

	return client.NewClient(context.TODO()).WithRetryCount(1).WithSleep(10).Configure(sdkConfig)
}

func validSdkConfigHanRegion() client.IClient {
	clientId, clientSecret := getEnv()
	sdkConfig := client.NewSdkConfigure().
		WithClientId(clientId).
		WithClientSecret(clientSecret).
		WithUserId(getValueOfEnv("VNGCLOUD_USER_ID")).
		WithZoneId(getValueOfEnv("VNGCLOUD_ZONE_ID")).
		WithProjectId(getValueOfEnv("HAN01_PROJECT_ID")).
		WithIamEndpoint("https://iamapis.vngcloud.vn/accounts-api").
		WithVServerEndpoint("https://han-1.api.vngcloud.vn/vserver/vserver-gateway").
		WithVLBEndpoint("https://han-1.api.vngcloud.vn/vserver/vlb-gateway").
		WithVNetworkEndpoint("https://vnetwork-hcm03.vngcloud.vn/vnetwork-gateway/vnetwork").
		WithVNetworkEndpoint("https://hcm-3.console.vngcloud.vn/vserver/vnetwork-gateway/vnetwork").
		WithUserAgent("vngcloud-go-sdk")

	return client.NewClient(context.TODO()).WithRetryCount(1).WithSleep(10).Configure(sdkConfig)
}

func validHcm3bSdkConfig() client.IClient {
	sdkConfig := client.NewSdkConfigure().
		WithClientId(getValueOfEnv("HCM3B_CLIENT_ID")).
		WithClientSecret(getValueOfEnv("HCM3B_CLIENT_SECRET")).
		WithZoneId(getValueOfEnv("VNGCLOUD_ZONE_ID")).
		WithProjectId(getValueOfEnv("HCM3B_PROJECT_ID")).
		WithIamEndpoint("https://iamapis.vngcloud.vn/accounts-api").
		WithVServerEndpoint("https://hcm-3.api.vngcloud.vn/vserver/vserver-gateway").
		WithVLBEndpoint("https://hcm-3.api.vngcloud.vn/vserver/vlb-gateway").
		WithVNetworkEndpoint("https://vnetwork-hcm03.vngcloud.vn/vnetwork-gateway/vnetwork").
		WithVNetworkEndpoint("https://hcm-3.console.vngcloud.vn/vserver/vnetwork-gateway/vnetwork")

	return client.NewClient(context.TODO()).WithRetryCount(1).WithSleep(10).Configure(sdkConfig)
}

func validHcm3bSuperSdkConfig() client.IClient {
	sdkConfig := client.NewSdkConfigure().
		WithClientId(getValueOfEnv("HCM3B_SUPER_CLIENT_ID")).
		WithClientSecret(getValueOfEnv("HCM3B_SUPER_CLIENT_SECRET")).
		WithZoneId(getValueOfEnv("VNGCLOUD_ZONE_ID")).
		WithProjectId(getValueOfEnv("HCM3B_PROJECT_ID")).
		WithIamEndpoint("https://iamapis.vngcloud.vn/accounts-api").
		WithVServerEndpoint("https://hcm-3.api.vngcloud.vn/vserver/vserver-gateway").
		WithVLBEndpoint("https://hcm-3.api.vngcloud.vn/vserver/vlb-gateway").
		WithVNetworkEndpoint("https://vnetwork-hcm03.vngcloud.vn/vnetwork-gateway/vnetwork").
		WithVNetworkEndpoint("https://hcm-3.console.vngcloud.vn/vserver/vnetwork-gateway/vnetwork")

	return client.NewClient(context.TODO()).WithRetryCount(1).WithSleep(10).Configure(sdkConfig)
}

func validVinhNt8SdkConfig() client.IClient {
	sdkConfig := client.NewSdkConfigure().
		WithClientId(getValueOfEnv("VINHCLIENT_ID")).
		WithClientSecret(getValueOfEnv("VINHCLIENT_SECRET")).
		WithZoneId(getValueOfEnv("VNGCLOUD_ZONE_ID")).
		WithProjectId(getValueOfEnv("VINH_PROJECT_ID")).
		WithIamEndpoint("https://iamapis.vngcloud.vn/accounts-api").
		WithVServerEndpoint("https://hcm-3.api.vngcloud.vn/vserver/vserver-gateway").
		WithVLBEndpoint("https://hcm-3.api.vngcloud.vn/vserver/vlb-gateway").
		WithVNetworkEndpoint("https://vnetwork-hcm03.vngcloud.vn/vnetwork-gateway/vnetwork").
		WithVNetworkEndpoint("https://hcm-3.console.vngcloud.vn/vserver/vnetwork-gateway/vnetwork")

	return client.NewClient(context.TODO()).WithRetryCount(1).WithSleep(10).Configure(sdkConfig)
}

func customerSdkConfig() client.IClient {
	sdkConfig := client.NewSdkConfigure().
		WithIamEndpoint("https://iamapis.vngcloud.vn/accounts-api").
		WithVServerEndpoint("https://hcm-3.api.vngcloud.vn/vserver/vserver-gateway").
		WithVLBEndpoint("https://hcm-3.api.vngcloud.vn/vserver/vlb-gateway")

	return client.NewClient(context.TODO()).WithRetryCount(1).WithSleep(10).Configure(sdkConfig)
}

func validSdkHannibalConfig() client.IClient {
	clientId, clientSecret := getValueOfEnv("HANNIBAL_CLIENT_ID"), getValueOfEnv("HANNIBAL_CLIENT_SECRET")
	sdkConfig := client.NewSdkConfigure().
		WithClientId(clientId).
		WithClientSecret(clientSecret).
		WithProjectId(getValueOfEnv("HANNIBAL_PROJECT_ID")).
		WithIamEndpoint("https://iamapis.vngcloud.vn/accounts-api").
		WithVServerEndpoint("https://hcm-3.api.vngcloud.vn/vserver/vserver-gateway").
		WithVLBEndpoint("https://hcm-3.api.vngcloud.vn/vserver/vlb-gateway")

	return client.NewClient(context.TODO()).WithRetryCount(1).WithSleep(10).Configure(sdkConfig)
}

func validSuperSdkConfig() client.IClient {
	clientId, clientSecret := getValueOfEnv("VNGCLOUD_SUPER_CLIENT_ID"), getValueOfEnv("VNGCLOUD_SUPER_CLIENT_SECRET")
	sdkConfig := client.NewSdkConfigure().
		WithClientId(clientId).
		WithClientSecret(clientSecret).
		WithZoneId(getValueOfEnv("VNGCLOUD_ZONE_ID")).
		WithProjectId(getValueOfEnv("VNGCLOUD_PROJECT_ID")).
		WithIamEndpoint("https://iamapis.vngcloud.vn/accounts-api").
		WithVServerEndpoint("https://hcm-3.api.vngcloud.vn/vserver/vserver-gateway").
		WithVLBEndpoint("https://hcm-3.api.vngcloud.vn/vserver/vlb-gateway").
		WithVNetworkEndpoint("https://vnetwork-hcm03-api.vngcloud.vn/vnetwork-gateway")

	return client.NewClient(context.TODO()).WithRetryCount(1).WithSleep(10).Configure(sdkConfig)
}

func validSuperSdkHcm03bConfig() client.IClient {
	clientId, clientSecret := getValueOfEnv("VNGCLOUD_PROD_HCM03B_CLIENT_ID"), getValueOfEnv("VNGCLOUD_PROD_HCM03B_CLIENT_SECRET")
	sdkConfig := client.NewSdkConfigure().
		WithClientId(clientId).
		WithClientSecret(clientSecret).
		WithZoneId(getValueOfEnv("VNGCLOUD_ZONE_ID")).
		WithProjectId(getValueOfEnv("VNGCLOUD_PROD_HCM03B_PROJECT_ID")).
		WithIamEndpoint("https://iamapis.vngcloud.vn/accounts-api").
		WithVServerEndpoint("https://hcm-3.api.vngcloud.vn/vserver/vserver-gateway").
		WithVLBEndpoint("https://hcm-3.api.vngcloud.vn/vserver/vlb-gateway").
		WithVNetworkEndpoint("https://vnetwork-hcm03-api.vngcloud.vn/vnetwork-gateway")

	return client.NewClient(context.TODO()).WithRetryCount(1).WithSleep(10).Configure(sdkConfig)
}

func validUser11412SdkConfig() client.IClient {
	clientId, clientSecret := getValueOfEnv("VNGCLOUD_SUPER_CLIENT_ID"), getValueOfEnv("VNGCLOUD_SUPER_CLIENT_SECRET")
	sdkConfig := client.NewSdkConfigure().
		WithClientId(clientId).
		WithClientSecret(clientSecret).
		WithProjectId(getValueOfEnv("USER_11412_PROJECT_ID")).
		WithIamEndpoint("https://iamapis.vngcloud.vn/accounts-api").
		WithVServerEndpoint("https://hcm-3.api.vngcloud.vn/vserver/vserver-gateway").
		WithVLBEndpoint("https://hcm-3.api.vngcloud.vn/vserver/vlb-gateway")

	return client.NewClient(context.TODO()).WithRetryCount(1).WithSleep(10).Configure(sdkConfig)
}

func validUser11412() client.IClient {
	clientId, clientSecret := getValueOfEnv("USER_11412_CLIENT_ID"), getValueOfEnv("USER_11412_CLIENT_SECRET")
	sdkConfig := client.NewSdkConfigure().
		WithClientId(clientId).
		WithClientSecret(clientSecret).
		WithProjectId(getValueOfEnv("USER_11412_PROJECT_ID")).
		WithIamEndpoint("https://iamapis.vngcloud.vn/accounts-api").
		WithVServerEndpoint("https://hcm-3.api.vngcloud.vn/vserver/vserver-gateway").
		WithVLBEndpoint("https://hcm-3.api.vngcloud.vn/vserver/vlb-gateway").
		WithVNetworkEndpoint("https://vnetwork-hcm03-api.vngcloud.vn/vnetwork-gateway")

	return client.NewClient(context.TODO()).WithRetryCount(1).WithSleep(10).Configure(sdkConfig)
}

func validSuperSdkConfig2() client.IClient {
	clientId, clientSecret := getValueOfEnv("VNGCLOUD_SUPER_CLIENT_ID"), getValueOfEnv("VNGCLOUD_SUPER_CLIENT_SECRET")
	sdkConfig := client.NewSdkConfigure().
		WithClientId(clientId).
		WithClientSecret(clientSecret).
		WithProjectId(getValueOfEnv("VINH_PROJECT_ID")).
		WithIamEndpoint("https://iamapis.vngcloud.vn/accounts-api").
		WithVServerEndpoint("https://hcm-3.api.vngcloud.vn/vserver/vserver-gateway").
		WithVLBEndpoint("https://hcm-3.api.vngcloud.vn/vserver/vlb-gateway")

	return client.NewClient(context.TODO()).WithRetryCount(1).WithSleep(10).Configure(sdkConfig)
}

func validSdkConfigDevops() client.IClient {
	clientId, clientSecret := getEnvDevOps()
	sdkConfig := client.NewSdkConfigure().
		WithClientId(clientId).
		WithClientSecret(clientSecret).
		WithProjectId(getValueOfEnv("PROJECT_ID_DEVOPS")).
		WithIamEndpoint("https://iamapis.vngcloud.vn/accounts-api").
		WithVServerEndpoint("https://hcm-3.api.vngcloud.vn/vserver/vserver-gateway")

	return client.NewClient(context.TODO()).WithRetryCount(1).WithSleep(10).Configure(sdkConfig)
}

func invalidSdkConfig() client.IClient {
	clientId := "invalid-id"
	clientSecret := "invalid-secret"
	sdkConfig := client.NewSdkConfigure().
		WithClientId(clientId).
		WithClientSecret(clientSecret).
		WithIamEndpoint("https://iamapis.vngcloud.vn/accounts-api").
		WithVServerEndpoint("https://hcm-3.api.vngcloud.vn/vserver/vserver-gateway")

	return client.NewClient(context.TODO()).WithRetryCount(1).WithSleep(10).Configure(sdkConfig)
}

func TestAuthenFailed(t *testing.T) {
	clientId := "cc136360-709c-4248-9358-e8e96c74480a"
	clientSecret := "invalid-secret"

	sdkConfig := client.NewSdkConfigure().
		WithClientId(clientId).
		WithClientSecret(clientSecret).
		WithIamEndpoint("https://iamapis.vngcloud.vn/accounts-api").
		WithVServerEndpoint("https://hcm-3.api.vngcloud.vn/vserver/vserver-gateway")

	vngcloud := client.NewClient(context.TODO()).WithRetryCount(1).WithSleep(10).Configure(sdkConfig)
	opt := identityv2.NewGetAccessTokenRequest(clientId, clientSecret)
	token, err := vngcloud.IamGateway().V2().IdentityService().GetAccessToken(opt)

	if err == nil {
		t.Error("Error MUST not be nil")
	}

	if token != nil {
		t.Error("Token MUST be nil")
	}

	if !err.IsError(sdkerror.EcAuthenticationFailed) {
		t.Error("Error MUST be VngCloudIamAuthenticationFailed")
	}

	t.Log("RESULT:", err)
	t.Log("PASS")
}

func TestAuthenPass(t *testing.T) {
	clientId, clientSecret := getEnv()
	vngcloud := validSdkConfig()
	opt := identityv2.NewGetAccessTokenRequest(clientId, clientSecret)
	token, err := vngcloud.IamGateway().V2().IdentityService().GetAccessToken(opt)

	if err != nil || token == nil {
		t.Error("This testcase MUST pass")
	}

	fmt.Println("RESULT 2:", token.Token)
	t.Log("RESULT:", token)
	t.Log("PASS")
}

func TestASuperuthenPass(t *testing.T) {
	clientId, clientSecret := getValueOfEnv("VNGCLOUD_SUPER_CLIENT_ID"), getValueOfEnv("VNGCLOUD_SUPER_CLIENT_SECRET")
	vngcloud := validSdkConfig()
	opt := identityv2.NewGetAccessTokenRequest(clientId, clientSecret)
	token, err := vngcloud.IamGateway().V2().IdentityService().GetAccessToken(opt)

	if err != nil || token == nil {
		t.Error("This testcase MUST pass")
	}

	fmt.Println("RESULT 2:", token.Token)
	t.Log("RESULT:", token.Token)
	t.Log("PASS")
}

func TestVinhAuthenPass(t *testing.T) {
	clientId, clientSecret := getValueOfEnv("VINHCLIENT_ID"), getValueOfEnv("VINHCLIENT_SECRET")
	vngcloud := validSuperSdkConfig2()
	opt := identityv2.NewGetAccessTokenRequest(clientId, clientSecret)
	token, err := vngcloud.IamGateway().V2().IdentityService().GetAccessToken(opt)

	if err != nil || token == nil {
		t.Error("This testcase MUST pass")
	}

	t.Log("RESULT:", token)
	t.Log("PASS")
}
