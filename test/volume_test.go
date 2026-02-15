package test

import (
	"context"
	"testing"

	v2 "github.com/dannyota/greennode-community-sdk/v2/greennode/services/volume/v2"
)

func TestCreateVolumeFailure(t *testing.T) {
	vngcloud := validSdkConfig()
	opt := v2.NewCreateBlockVolumeRequest(
		"testsadsadsada",
		"vtype-7a7a8610-34f5-11ee-be56-0242ac120002",
		10,
	)
	volume, sdkerr := vngcloud.VServerGateway().V2().VolumeService().CreateBlockVolume(context.Background(), opt)
	if sdkerr == nil {
		t.Fatalf("Expect error but got nil")
	}

	if volume != nil {
		t.Fatalf("Expect nil but got %v", volume)
	}

	t.Log("Result: ", sdkerr)
	t.Log("PASS")
}

func TestCreateVolumeSuccess(t *testing.T) {
	vngcloud := validSdkConfig()
	opt := v2.NewCreateBlockVolumeRequest(
		"test-volume-tags",
		"vtype-7a7a8610-34f5-11ee-be56-0242ac120002",
		20,
	).WithTags("test-key", "test-value", "owner", "sdk-test")
	volume, sdkerr := vngcloud.VServerGateway().V2().VolumeService().CreateBlockVolume(context.Background(), opt)
	if sdkerr != nil {
		t.Fatalf("Expect nil but got %v", sdkerr)
	}

	if volume == nil {
		t.Fatalf("Expect not nil but got nil")
	}

	t.Log("Result: ", volume)
	t.Log("PASS")
}

func TestDeleteVolumeByIDFailure(t *testing.T) {
	vngcloud := validSdkConfig()
	opt := v2.NewDeleteBlockVolumeByIDRequest("this-is-fake")
	sdkerr := vngcloud.VServerGateway().V2().VolumeService().DeleteBlockVolumeByID(context.Background(), opt)
	if sdkerr == nil {
		t.Fatalf("Expect error but got nil")
	}

	t.Log("Result: ", sdkerr)
	t.Log("PASS")
}

func TestDeleteVolumeByIDSuccess(t *testing.T) {
	vngcloud := validSdkConfig()
	opt := v2.NewDeleteBlockVolumeByIDRequest("vol-51f71146-9c20-4615-a73e-a43a39bf03ea")
	sdkerr := vngcloud.VServerGateway().V2().VolumeService().DeleteBlockVolumeByID(context.Background(), opt)
	if sdkerr != nil {
		t.Fatalf("Expect nil but got %v", sdkerr)
	}

	t.Log("Result: ", sdkerr)
	t.Log("PASS")
}

func TestListBlockVolumeSuccess(t *testing.T) {
	vngcloud := validSdkConfig()
	opt := v2.NewListBlockVolumesRequest(1, 10)
	volumes, sdkerr := vngcloud.VServerGateway().V2().VolumeService().ListBlockVolumes(context.Background(), opt)
	if sdkerr != nil {
		t.Fatalf("Expect nil but got %v", sdkerr)
	}

	if volumes == nil {
		t.Fatalf("Expect not nil but got nil")
	}

	t.Log("Result: ", volumes)
	t.Log("PASS")
}

func TestListBlockVolumeWithNameSuccess(t *testing.T) {
	vngcloud := validSdkConfig()
	opt := v2.NewListBlockVolumesRequest(1, 10).WithName("pvc-24182151-aa4a-4a55-9572-f551c3d003aa")
	volumes, sdkerr := vngcloud.VServerGateway().V2().VolumeService().ListBlockVolumes(context.Background(), opt)
	if sdkerr != nil {
		t.Fatalf("Expect nil but got %v", sdkerr)
	}

	if volumes == nil {
		t.Fatalf("Expect not nil but got nil")
	}

	if volumes.Len() != 1 {
		t.Fatalf("Expect 1 but got %d", volumes.Len())
	}

	t.Log("Result: ", volumes)
	t.Log("PASS")
}

func TestListBlockVolumeWithFailure(t *testing.T) {
	vngcloud := validSdkConfig()
	opt := v2.NewListBlockVolumesRequest(1, -10)
	volumes, sdkerr := vngcloud.VServerGateway().V2().VolumeService().ListBlockVolumes(context.Background(), opt)
	if sdkerr == nil {
		t.Fatalf("Expect error but got nil")
	}

	if volumes != nil {
		t.Fatalf("Expect nil but got %v", volumes)
	}

	t.Log("Result: ", sdkerr)
	t.Log("PASS")
}

func TestGetBlockVolumeByIDFailure(t *testing.T) {
	vngcloud := validSdkConfig()
	opt := v2.NewGetBlockVolumeByIDRequest("vol-17dc6df0-43d3-4ad2-be88-69ddaef2f146")
	volume, sdkerr := vngcloud.VServerGateway().V2().VolumeService().GetBlockVolumeByID(context.Background(), opt)
	if sdkerr == nil {
		t.Fatalf("Expect error but got nil")
	}

	if volume != nil {
		t.Fatalf("Expect nil but got %v", volume)
	}

	t.Log("Result: ", sdkerr)
	t.Log("PASS")
}

func TestGetBlockVolumeByIDSuccess(t *testing.T) {
	vngcloud := validSdkConfig()
	opt := v2.NewGetBlockVolumeByIDRequest("vol-aa784f76-a13d-4f92-b807-d2df3180e030")
	volume, sdkerr := vngcloud.VServerGateway().V2().VolumeService().GetBlockVolumeByID(context.Background(), opt)
	if sdkerr != nil {
		t.Fatalf("Expect nil but got %v", sdkerr)
	}

	if volume == nil {
		t.Fatalf("Expect not nil but got nil")
	}

	t.Log("Result: ", volume)
	t.Log("PASS")
}

func TestResizeBlockVolumeFailure(t *testing.T) {
	vngcloud := validSdkConfig()
	opt := v2.NewResizeBlockVolumeByIDRequest(
		"vol-ae3fffe5-bd46-475f-bee3-3d5eff4a4b45",
		"vtype-9f811804-3574-466e-831c-f23d56ca6700",
		40)
	volume, sdkerr := vngcloud.VServerGateway().V2().VolumeService().ResizeBlockVolumeByID(context.Background(), opt)
	if sdkerr == nil {
		t.Fatalf("Expect error but got nil")
	}

	if volume != nil {
		t.Fatalf("Expect nil but got %v", volume)
	}

	t.Log("Result: ", sdkerr)
	t.Log("PASS")
}

func TestGetUnderVolumeIDFailure(t *testing.T) {
	vngcloud := validSdkConfig()
	opt := v2.NewGetUnderVolumeIDRequest("vol-ae3fffe5-bd46-475f-besadd")
	volume, sdkerr := vngcloud.VServerGateway().V2().VolumeService().GetUnderBlockVolumeID(context.Background(), opt)
	if sdkerr == nil {
		t.Fatalf("Expect error but got nil")
	}

	if volume != nil {
		t.Fatalf("Expect nil but got %v", volume)
	}

	t.Log("Result: ", sdkerr)
	t.Log("PASS")
}

func TestGetUnderVolumeIDSuccess(t *testing.T) {
	vngcloud := validSdkConfig()
	opt := v2.NewGetUnderVolumeIDRequest("vol-137f3dfc-9198-4d94-983f-6802e3c39e4f")
	volume, sdkerr := vngcloud.VServerGateway().V2().VolumeService().GetUnderBlockVolumeID(context.Background(), opt)
	if sdkerr != nil {
		t.Fatalf("Expect nil but got %v", sdkerr)
	}

	if volume == nil {
		t.Fatalf("Expect not nil but got nil")
	}

	t.Log("Result: ", volume)
	t.Log("PASS")
}

func TestMigrateBlockVolume(t *testing.T) {
	vtNvme := "vtype-7a7a8610-34f5-11ee-be56-0242ac120002"
	vtSsd := "vtype-7a7a8610-34f5-11ee-be56-0242ac120002"
	t.Log(vtSsd, vtNvme)

	vngcloud := validSdkConfig()
	opt := v2.NewMigrateBlockVolumeByIDRequest(
		"vol-78c26ee6-20b2-45a3-9f0a-e7728349c300",
		vtSsd).WithConfirm(true).WithAction(v2.ProcessMigrateAction)
	sdkerr := vngcloud.VServerGateway().V2().VolumeService().MigrateBlockVolumeByID(context.Background(), opt)

	t.Log("Error: ", sdkerr)
	t.Log("PASS")
}
