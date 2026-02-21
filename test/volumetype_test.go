//go:build integration

package test

import (
	"context"
	"testing"

	volumev1 "danny.vn/greennode/greennode/services/volume/v1"
)

func TestGetVolumeTypeFailure(t *testing.T) {
	vngcloud := validSdkConfig()
	opt := volumev1.NewGetVolumeTypeByIDRequest("fake-id")
	volume, sdkerr := vngcloud.VolumeV1.GetVolumeTypeByID(context.Background(), opt)
	if sdkerr == nil {
		t.Fatalf("Expect error but got nil")
	}

	if volume != nil {
		t.Fatalf("Expect nil but got %v", volume)
	}

	t.Log("Result: ", sdkerr)
	t.Log("PASS")
}

func TestGetVolumeTypeSuccess(t *testing.T) {
	vngcloud := validSdkConfig()
	opt := volumev1.NewGetVolumeTypeByIDRequest("vtype-2fc64a6c-38e3-4f08-93a5-18018cb3ab23")
	volume, sdkerr := vngcloud.VolumeV1.GetVolumeTypeByID(context.Background(), opt)

	if sdkerr != nil {
		t.Fatalf("Expect nil but got %v", sdkerr)
	}

	if volume == nil {
		t.Fatalf("Expect not nil but got nil")
	}

	t.Logf("Result: %+v", volume)
	t.Log("PASS")
}

func TestGetDefaultVolumeType(t *testing.T) {
	vngcloud := validSdkConfig()
	volType, sdkerr := vngcloud.VolumeV1.GetDefaultVolumeType(context.Background())
	if sdkerr != nil {
		t.Fatalf("Expect nil but got %v", sdkerr)
	}
	if volType == nil {
		t.Fatalf("Expect not nil but got nil")
	}

	t.Log("Result: ", volType)
}

func TestGetVolumeTypeZones(t *testing.T) {
	vngcloud := validSdkConfig()
	opt := volumev1.NewGetVolumeTypeZonesRequest("HCM03-1A")

	volType, sdkerr := vngcloud.VolumeV1.GetVolumeTypeZones(context.Background(), opt)
	if sdkerr != nil {
		t.Fatalf("Expect nil but got %v", sdkerr)
	}
	if volType == nil {
		t.Fatalf("Expect not nil but got nil")
	}

	t.Log("Result: ", volType)
}

func TestGetVolumeTypes(t *testing.T) {
	vngcloud := validSdkConfig()
	opt := volumev1.NewListVolumeTypeRequest("0745BE12-9433-4DD4-90A1-384631504EBE")

	volType, sdkerr := vngcloud.VolumeV1.GetListVolumeTypes(context.Background(), opt)
	if sdkerr != nil {
		t.Fatalf("Expect nil but got %v", sdkerr)
	}
	if volType == nil {
		t.Fatalf("Expect not nil but got nil")
	}

	t.Log("Result: ", volType)
}
