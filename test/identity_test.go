package test

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/dannyota/greennode-community-sdk/v2/client"
	sdkerror "github.com/dannyota/greennode-community-sdk/v2/greennode/sdkerror"
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

func getEnvCuongDm4() (string, string) {
	envFile, _ := readEnvFile("./env.yaml")
	clientID := envFile["CUONGDM4_CLIENT_ID"]
	clientSecret := envFile["CUONGDM4_CLIENT_SECRET"]

	return clientID, clientSecret
}

func getEnvDevOps() (string, string) {
	envFile, _ := readEnvFile("./env.yaml")
	clientID := envFile["CLIENT_ID_DEVOPS"]
	clientSecret := envFile["CLIENT_SECRET_DEVOPS"]

	return clientID, clientSecret
}

func getValueOfEnv(key string) string {
	envFile, _ := readEnvFile("./env.yaml")
	value := envFile[key]
	return value
}

func validSdkConfig() client.Client {
	clientID, clientSecret := getEnv()
	sdkConfig := client.NewSdkConfigure().
		WithClientID(clientID).
		WithClientSecret(clientSecret).
		WithUserID(getValueOfEnv("VNGCLOUD_USER_ID")).
		WithZoneID(getValueOfEnv("VNGCLOUD_ZONE_ID")).
		WithProjectID(getValueOfEnv("VNGCLOUD_PROJECT_ID")).
		WithIamEndpoint("https://iamapis.vngcloud.vn/accounts-api").
		WithVServerEndpoint("https://hcm-3.api.vngcloud.vn/vserver/vserver-gateway").
		WithVLBEndpoint("https://hcm-3.api.vngcloud.vn/vserver/vlb-gateway").
		WithVNetworkEndpoint("https://vnetwork-hcm03.vngcloud.vn/vnetwork-gateway/vnetwork").
		WithVNetworkEndpoint("https://hcm-3.console.vngcloud.vn/vserver/vnetwork-gateway/vnetwork").
		WithGLBEndpoint("https://glb.console.vngcloud.vn/glb-controller/").
		WithVDnsEndpoint("https://vdns.api.vngcloud.vn/")

	return client.NewClient(context.TODO()).WithRetryCount(1).WithSleep(10).Configure(sdkConfig)
}

func validUserSdkConfig() client.Client {
	clientID, clientSecret := getValueOfEnv("USER_CLIENT_ID"), getValueOfEnv("USER_CLIENT_SECRET")
	sdkConfig := client.NewSdkConfigure().
		WithClientID(clientID).
		WithClientSecret(clientSecret).
		WithUserID(getValueOfEnv("VNGCLOUD_USER_ID")).
		WithZoneID(getValueOfEnv("VNGCLOUD_ZONE_ID")).
		WithProjectID(getValueOfEnv("USER_PROJECT")).
		WithIamEndpoint("https://iamapis.vngcloud.vn/accounts-api").
		WithVServerEndpoint("https://hcm-3.api.vngcloud.vn/vserver/vserver-gateway").
		WithVLBEndpoint("https://hcm-3.api.vngcloud.vn/vserver/vlb-gateway").
		WithVNetworkEndpoint("https://vnetwork-hcm03.vngcloud.vn/vnetwork-gateway/vnetwork").
		WithVNetworkEndpoint("https://hcm-3.console.vngcloud.vn/vserver/vnetwork-gateway/vnetwork")

	return client.NewClient(context.TODO()).WithRetryCount(1).WithSleep(10).Configure(sdkConfig)
}

func validUserSdkConfigForCuongDm4() client.Client {
	clientID, clientSecret := getEnvCuongDm4()
	sdkConfig := client.NewSdkConfigure().
		WithClientID(clientID).
		WithClientSecret(clientSecret).
		WithProjectID("pro-462803f3-6858-466f-bf05-df2b33faa360").
		WithIamEndpoint("https://iamapis.vngcloud.vn/accounts-api").
		WithVServerEndpoint("https://hcm-3.api.vngcloud.vn/vserver/vserver-gateway").
		WithVLBEndpoint("https://hcm-3.api.vngcloud.vn/vserver/vlb-gateway").
		WithVNetworkEndpoint("https://vnetwork-hcm03.vngcloud.vn/vnetwork-gateway/vnetwork").
		WithVNetworkEndpoint("https://hcm-3.console.vngcloud.vn/vserver/vnetwork-gateway/vnetwork")

	return client.NewClient(context.TODO()).WithRetryCount(1).WithSleep(10).Configure(sdkConfig)
}

func validSdkConfigHanRegion() client.Client {
	clientID, clientSecret := getEnv()
	sdkConfig := client.NewSdkConfigure().
		WithClientID(clientID).
		WithClientSecret(clientSecret).
		WithUserID(getValueOfEnv("VNGCLOUD_USER_ID")).
		WithZoneID(getValueOfEnv("VNGCLOUD_ZONE_ID")).
		WithProjectID(getValueOfEnv("HAN01_PROJECT_ID")).
		WithIamEndpoint("https://iamapis.vngcloud.vn/accounts-api").
		WithVServerEndpoint("https://han-1.api.vngcloud.vn/vserver/vserver-gateway").
		WithVLBEndpoint("https://han-1.api.vngcloud.vn/vserver/vlb-gateway").
		WithVNetworkEndpoint("https://vnetwork-hcm03.vngcloud.vn/vnetwork-gateway/vnetwork").
		WithVNetworkEndpoint("https://hcm-3.console.vngcloud.vn/vserver/vnetwork-gateway/vnetwork").
		WithUserAgent("vngcloud-go-sdk")

	return client.NewClient(context.TODO()).WithRetryCount(1).WithSleep(10).Configure(sdkConfig)
}

func validHcm3bSdkConfig() client.Client {
	sdkConfig := client.NewSdkConfigure().
		WithClientID(getValueOfEnv("HCM3B_CLIENT_ID")).
		WithClientSecret(getValueOfEnv("HCM3B_CLIENT_SECRET")).
		WithZoneID(getValueOfEnv("VNGCLOUD_ZONE_ID")).
		WithProjectID(getValueOfEnv("HCM3B_PROJECT_ID")).
		WithIamEndpoint("https://iamapis.vngcloud.vn/accounts-api").
		WithVServerEndpoint("https://hcm-3.api.vngcloud.vn/vserver/vserver-gateway").
		WithVLBEndpoint("https://hcm-3.api.vngcloud.vn/vserver/vlb-gateway").
		WithVNetworkEndpoint("https://vnetwork-hcm03.vngcloud.vn/vnetwork-gateway/vnetwork").
		WithVNetworkEndpoint("https://hcm-3.console.vngcloud.vn/vserver/vnetwork-gateway/vnetwork")

	return client.NewClient(context.TODO()).WithRetryCount(1).WithSleep(10).Configure(sdkConfig)
}

func validHcm3bSuperSdkConfig() client.Client {
	sdkConfig := client.NewSdkConfigure().
		WithClientID(getValueOfEnv("HCM3B_SUPER_CLIENT_ID")).
		WithClientSecret(getValueOfEnv("HCM3B_SUPER_CLIENT_SECRET")).
		WithZoneID(getValueOfEnv("VNGCLOUD_ZONE_ID")).
		WithProjectID(getValueOfEnv("HCM3B_PROJECT_ID")).
		WithIamEndpoint("https://iamapis.vngcloud.vn/accounts-api").
		WithVServerEndpoint("https://hcm-3.api.vngcloud.vn/vserver/vserver-gateway").
		WithVLBEndpoint("https://hcm-3.api.vngcloud.vn/vserver/vlb-gateway").
		WithVNetworkEndpoint("https://vnetwork-hcm03.vngcloud.vn/vnetwork-gateway/vnetwork").
		WithVNetworkEndpoint("https://hcm-3.console.vngcloud.vn/vserver/vnetwork-gateway/vnetwork")

	return client.NewClient(context.TODO()).WithRetryCount(1).WithSleep(10).Configure(sdkConfig)
}

func validVinhNt8SdkConfig() client.Client {
	sdkConfig := client.NewSdkConfigure().
		WithClientID(getValueOfEnv("VINHCLIENT_ID")).
		WithClientSecret(getValueOfEnv("VINHCLIENT_SECRET")).
		WithZoneID(getValueOfEnv("VNGCLOUD_ZONE_ID")).
		WithProjectID(getValueOfEnv("VINH_PROJECT_ID")).
		WithIamEndpoint("https://iamapis.vngcloud.vn/accounts-api").
		WithVServerEndpoint("https://hcm-3.api.vngcloud.vn/vserver/vserver-gateway").
		WithVLBEndpoint("https://hcm-3.api.vngcloud.vn/vserver/vlb-gateway").
		WithVNetworkEndpoint("https://vnetwork-hcm03.vngcloud.vn/vnetwork-gateway/vnetwork").
		WithVNetworkEndpoint("https://hcm-3.console.vngcloud.vn/vserver/vnetwork-gateway/vnetwork")

	return client.NewClient(context.TODO()).WithRetryCount(1).WithSleep(10).Configure(sdkConfig)
}

func customerSdkConfig() client.Client {
	sdkConfig := client.NewSdkConfigure().
		WithIamEndpoint("https://iamapis.vngcloud.vn/accounts-api").
		WithVServerEndpoint("https://hcm-3.api.vngcloud.vn/vserver/vserver-gateway").
		WithVLBEndpoint("https://hcm-3.api.vngcloud.vn/vserver/vlb-gateway")

	return client.NewClient(context.TODO()).WithRetryCount(1).WithSleep(10).Configure(sdkConfig)
}

func validSdkHannibalConfig() client.Client {
	clientID, clientSecret := getValueOfEnv("HANNIBAL_CLIENT_ID"), getValueOfEnv("HANNIBAL_CLIENT_SECRET")
	sdkConfig := client.NewSdkConfigure().
		WithClientID(clientID).
		WithClientSecret(clientSecret).
		WithProjectID(getValueOfEnv("HANNIBAL_PROJECT_ID")).
		WithIamEndpoint("https://iamapis.vngcloud.vn/accounts-api").
		WithVServerEndpoint("https://hcm-3.api.vngcloud.vn/vserver/vserver-gateway").
		WithVLBEndpoint("https://hcm-3.api.vngcloud.vn/vserver/vlb-gateway")

	return client.NewClient(context.TODO()).WithRetryCount(1).WithSleep(10).Configure(sdkConfig)
}

func validSuperSdkConfig() client.Client {
	clientID, clientSecret := getValueOfEnv("VNGCLOUD_SUPER_CLIENT_ID"), getValueOfEnv("VNGCLOUD_SUPER_CLIENT_SECRET")
	sdkConfig := client.NewSdkConfigure().
		WithClientID(clientID).
		WithClientSecret(clientSecret).
		WithZoneID(getValueOfEnv("VNGCLOUD_ZONE_ID")).
		WithProjectID(getValueOfEnv("VNGCLOUD_PROJECT_ID")).
		WithIamEndpoint("https://iamapis.vngcloud.vn/accounts-api").
		WithVServerEndpoint("https://hcm-3.api.vngcloud.vn/vserver/vserver-gateway").
		WithVLBEndpoint("https://hcm-3.api.vngcloud.vn/vserver/vlb-gateway").
		WithVNetworkEndpoint("https://vnetwork-hcm03-api.vngcloud.vn/vnetwork-gateway")

	return client.NewClient(context.TODO()).WithRetryCount(1).WithSleep(10).Configure(sdkConfig)
}

func validSuperSdkHcm03bConfig() client.Client {
	clientID, clientSecret := getValueOfEnv("VNGCLOUD_PROD_HCM03B_CLIENT_ID"), getValueOfEnv("VNGCLOUD_PROD_HCM03B_CLIENT_SECRET")
	sdkConfig := client.NewSdkConfigure().
		WithClientID(clientID).
		WithClientSecret(clientSecret).
		WithZoneID(getValueOfEnv("VNGCLOUD_ZONE_ID")).
		WithProjectID(getValueOfEnv("VNGCLOUD_PROD_HCM03B_PROJECT_ID")).
		WithIamEndpoint("https://iamapis.vngcloud.vn/accounts-api").
		WithVServerEndpoint("https://hcm-3.api.vngcloud.vn/vserver/vserver-gateway").
		WithVLBEndpoint("https://hcm-3.api.vngcloud.vn/vserver/vlb-gateway").
		WithVNetworkEndpoint("https://vnetwork-hcm03-api.vngcloud.vn/vnetwork-gateway")

	return client.NewClient(context.TODO()).WithRetryCount(1).WithSleep(10).Configure(sdkConfig)
}

func validUser11412SdkConfig() client.Client {
	clientID, clientSecret := getValueOfEnv("VNGCLOUD_SUPER_CLIENT_ID"), getValueOfEnv("VNGCLOUD_SUPER_CLIENT_SECRET")
	sdkConfig := client.NewSdkConfigure().
		WithClientID(clientID).
		WithClientSecret(clientSecret).
		WithProjectID(getValueOfEnv("USER_11412_PROJECT_ID")).
		WithIamEndpoint("https://iamapis.vngcloud.vn/accounts-api").
		WithVServerEndpoint("https://hcm-3.api.vngcloud.vn/vserver/vserver-gateway").
		WithVLBEndpoint("https://hcm-3.api.vngcloud.vn/vserver/vlb-gateway")

	return client.NewClient(context.TODO()).WithRetryCount(1).WithSleep(10).Configure(sdkConfig)
}

func validUser11412() client.Client {
	clientID, clientSecret := getValueOfEnv("USER_11412_CLIENT_ID"), getValueOfEnv("USER_11412_CLIENT_SECRET")
	sdkConfig := client.NewSdkConfigure().
		WithClientID(clientID).
		WithClientSecret(clientSecret).
		WithProjectID(getValueOfEnv("USER_11412_PROJECT_ID")).
		WithIamEndpoint("https://iamapis.vngcloud.vn/accounts-api").
		WithVServerEndpoint("https://hcm-3.api.vngcloud.vn/vserver/vserver-gateway").
		WithVLBEndpoint("https://hcm-3.api.vngcloud.vn/vserver/vlb-gateway").
		WithVNetworkEndpoint("https://vnetwork-hcm03-api.vngcloud.vn/vnetwork-gateway")

	return client.NewClient(context.TODO()).WithRetryCount(1).WithSleep(10).Configure(sdkConfig)
}

func validSuperSdkConfig2() client.Client {
	clientID, clientSecret := getValueOfEnv("VNGCLOUD_SUPER_CLIENT_ID"), getValueOfEnv("VNGCLOUD_SUPER_CLIENT_SECRET")
	sdkConfig := client.NewSdkConfigure().
		WithClientID(clientID).
		WithClientSecret(clientSecret).
		WithProjectID(getValueOfEnv("VINH_PROJECT_ID")).
		WithIamEndpoint("https://iamapis.vngcloud.vn/accounts-api").
		WithVServerEndpoint("https://hcm-3.api.vngcloud.vn/vserver/vserver-gateway").
		WithVLBEndpoint("https://hcm-3.api.vngcloud.vn/vserver/vlb-gateway")

	return client.NewClient(context.TODO()).WithRetryCount(1).WithSleep(10).Configure(sdkConfig)
}

func validSdkConfigDevops() client.Client {
	clientID, clientSecret := getEnvDevOps()
	sdkConfig := client.NewSdkConfigure().
		WithClientID(clientID).
		WithClientSecret(clientSecret).
		WithProjectID(getValueOfEnv("PROJECT_ID_DEVOPS")).
		WithIamEndpoint("https://iamapis.vngcloud.vn/accounts-api").
		WithVServerEndpoint("https://hcm-3.api.vngcloud.vn/vserver/vserver-gateway")

	return client.NewClient(context.TODO()).WithRetryCount(1).WithSleep(10).Configure(sdkConfig)
}

func invalidSdkConfig() client.Client {
	clientID := "invalid-id"
	clientSecret := "invalid-secret"
	sdkConfig := client.NewSdkConfigure().
		WithClientID(clientID).
		WithClientSecret(clientSecret).
		WithIamEndpoint("https://iamapis.vngcloud.vn/accounts-api").
		WithVServerEndpoint("https://hcm-3.api.vngcloud.vn/vserver/vserver-gateway")

	return client.NewClient(context.TODO()).WithRetryCount(1).WithSleep(10).Configure(sdkConfig)
}

func TestAuthenFailed(t *testing.T) {
	clientID := "cc136360-709c-4248-9358-e8e96c74480a"
	clientSecret := "invalid-secret"

	sdkConfig := client.NewSdkConfigure().
		WithClientID(clientID).
		WithClientSecret(clientSecret).
		WithIamEndpoint("https://iamapis.vngcloud.vn/accounts-api").
		WithVServerEndpoint("https://hcm-3.api.vngcloud.vn/vserver/vserver-gateway")

	vngcloud := client.NewClient(context.TODO()).WithRetryCount(1).WithSleep(10).Configure(sdkConfig)
	opt := identityv2.NewGetAccessTokenRequest(clientID, clientSecret)
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
	clientID, clientSecret := getEnv()
	vngcloud := validSdkConfig()
	opt := identityv2.NewGetAccessTokenRequest(clientID, clientSecret)
	token, err := vngcloud.IamGateway().V2().IdentityService().GetAccessToken(opt)

	if err != nil || token == nil {
		t.Error("This testcase MUST pass")
	}

	fmt.Println("RESULT 2:", token.Token)
	t.Log("RESULT:", token)
	t.Log("PASS")
}

func TestASuperuthenPass(t *testing.T) {
	clientID, clientSecret := getValueOfEnv("VNGCLOUD_SUPER_CLIENT_ID"), getValueOfEnv("VNGCLOUD_SUPER_CLIENT_SECRET")
	vngcloud := validSdkConfig()
	opt := identityv2.NewGetAccessTokenRequest(clientID, clientSecret)
	token, err := vngcloud.IamGateway().V2().IdentityService().GetAccessToken(opt)

	if err != nil || token == nil {
		t.Error("This testcase MUST pass")
	}

	fmt.Println("RESULT 2:", token.Token)
	t.Log("RESULT:", token.Token)
	t.Log("PASS")
}

func TestVinhAuthenPass(t *testing.T) {
	clientID, clientSecret := getValueOfEnv("VINHCLIENT_ID"), getValueOfEnv("VINHCLIENT_SECRET")
	vngcloud := validSuperSdkConfig2()
	opt := identityv2.NewGetAccessTokenRequest(clientID, clientSecret)
	token, err := vngcloud.IamGateway().V2().IdentityService().GetAccessToken(opt)

	if err != nil || token == nil {
		t.Error("This testcase MUST pass")
	}

	t.Log("RESULT:", token)
	t.Log("PASS")
}
