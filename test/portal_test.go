package test

import (
	"context"
	"errors"
	"testing"

	"github.com/dannyota/greennode-community-sdk/v2/greennode/sdkerror"
	portalv1 "github.com/dannyota/greennode-community-sdk/v2/greennode/services/portal/v1"
	portalv2 "github.com/dannyota/greennode-community-sdk/v2/greennode/services/portal/v2"
)

func TestGetPortalInfoFailed(t *testing.T) {
	backendProjectID := getValueOfEnv("BACKEND_PROJECT_ID")
	vngcloud := invalidSdkConfig()
	opt := portalv1.NewGetPortalInfoRequest(backendProjectID)
	portal, err := vngcloud.VServerGateway().V1().PortalService().GetPortalInfo(context.Background(), opt)

	if err == nil {
		t.Errorf("Expect error but got nil")
	}

	if portal != nil {
		t.Errorf("Expect portal to be nil but got %+v", portal)
	}

	var sdkErr *sdkerror.SdkError
	if errors.As(err, &sdkErr) {
		if !sdkErr.IsError(sdkerror.EcAuthenticationFailed) {
			t.Errorf("Expect error code to be %s but got %s", sdkerror.EcAuthenticationFailed, sdkErr.ErrorCode())
		}
	} else {
		t.Errorf("Expected SdkError")
	}

	t.Log("RESULT:", err)
	t.Log("PASS")
}

func TestGetPortalInfoSuccess(t *testing.T) {
	backendProjectID := getValueOfEnv("BACKEND_PROJECT_ID")
	vngcloud := validSdkConfig()
	opt := portalv1.NewGetPortalInfoRequest(backendProjectID)
	portal, err := vngcloud.VServerGateway().V1().PortalService().GetPortalInfo(context.Background(), opt)

	if err != nil {
		t.Errorf("Expect error to be nil but got %+v", err)
	}

	if portal == nil {
		t.Errorf("Expect portal not to be nil but got nil")
	}

	t.Log("RESULT:", portal)
	t.Log("PASS")
}

func TestGetPortalInfoSuccess2(t *testing.T) {
	backendProjectID := getValueOfEnv("USER_11412_OS_PROJECT_ID")
	vngcloud := validUser11412SdkConfig()
	opt := portalv1.NewGetPortalInfoRequest(backendProjectID)
	portal, err := vngcloud.VServerGateway().V1().PortalService().GetPortalInfo(context.Background(), opt)

	if err != nil {
		t.Errorf("Expect error to be nil but got %+v", err)
	}

	if portal == nil {
		t.Errorf("Expect portal not to be nil but got nil")
	}

	t.Log("RESULT:", portal)
	t.Log("PASS")
}

func TestGetPortalInfoFailure(t *testing.T) {
	backendProjectID := getValueOfEnv("FAKE_BACKEND_PROJECT_ID")
	vngcloud := validSdkConfig()
	opt := portalv1.NewGetPortalInfoRequest(backendProjectID)
	portal, err := vngcloud.VServerGateway().V1().PortalService().GetPortalInfo(context.Background(), opt)

	if err == nil {
		t.Errorf("Expect error to be nil but got %+v", err)
	}

	if portal != nil {
		t.Errorf("Expect portal not to be nil but got nil")
	}

	t.Log("RESULT:", err)
	t.Log("PASS")
}

func TestListAllQuotaSuccess(t *testing.T) {
	vngcloud := validSdkConfig()
	quotas, err := vngcloud.VServerGateway().V2().PortalService().ListAllQuotaUsed(context.Background())

	if err != nil {
		t.Errorf("Expect error to be nil but got %+v", err)
	}

	if quotas == nil {
		t.Errorf("Expect quotas not to be nil but got nil")
	}

	t.Log("RESULT:", quotas)
	t.Log("PASS")
}

func TestGetQuotaByNameFailure(t *testing.T) {
	vngcloud := validSdkConfig()
	opt := portalv2.NewGetQuotaByNameRequest("fake-quota-name")
	quota, err := vngcloud.VServerGateway().V2().PortalService().GetQuotaByName(context.Background(), opt)

	if err == nil {
		t.Errorf("Expect error but got nil")
	}

	if quota != nil {
		t.Errorf("Expect quota to be nil but got %+v", quota)
	}

	t.Log("RESULT:", err)
	t.Log("PASS")
}

func TestGetQuotaByNamePass(t *testing.T) {
	vngcloud := validSdkConfig()
	opt := portalv2.NewGetQuotaByNameRequest(portalv2.QtVolumeAttachLimit)
	quota, err := vngcloud.VServerGateway().V2().PortalService().GetQuotaByName(context.Background(), opt)

	if err != nil {
		t.Errorf("Expect error to be nil but got %+v", err)
	}

	if quota == nil {
		t.Errorf("Expect quota not to be nil but got nil")
	}

	t.Log("RESULT:", quota)
	t.Log("PASS")
}

func TestListProjects(t *testing.T) {
	vngcloud := validSdkConfig()
	projects, err := vngcloud.VServerGateway().V1().PortalService().ListProjects(context.Background(), portalv1.NewListProjectsRequest())
	if err != nil {
		t.Log("Error: ", err)
	}

	t.Log("Result: ", projects.At(0))
}

func TestListPortalUser11412(t *testing.T) {
	vngcloud := validUser11412()
	projects, err := vngcloud.VServerGateway().V1().PortalService().ListProjects(context.Background(), portalv1.NewListProjectsRequest())
	if err != nil {
		t.Log("Error: ", err)
	}

	t.Log("Result: ", projects.At(0))
}

func TestListZones(t *testing.T) {
	vngcloud := validSdkConfig()
	zones, err := vngcloud.VServerGateway().V1().PortalService().ListZones(context.Background())
	if err != nil {
		t.Log("Error: ", err)
	}

	t.Log("Result: ", zones.Items)
}
