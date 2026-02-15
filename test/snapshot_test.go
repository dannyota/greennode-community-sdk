package test

import (
	"context"
	"testing"

	v2 "github.com/dannyota/greennode-community-sdk/v2/greennode/services/volume/v2"
)

func TestListSnapshotFailure(t *testing.T) {
	vngcloud := validSdkConfig()
	opt := v2.NewListSnapshotsByBlockVolumeIDRequest(1, 10, "fsffsfsdfdsfsdf")
	_, sdkerr := vngcloud.VServerGateway().V2().VolumeService().ListSnapshotsByBlockVolumeID(context.Background(), opt)

	t.Log("Result: ", sdkerr)
	t.Log("PASS")
}

func TestListSnapshotSuccess(t *testing.T) {
	vngcloud := validSdkConfig()
	opt := v2.NewListSnapshotsByBlockVolumeIDRequest(1, 10, "vol-d360fd83-948d-4efa-ab46-aab97328e275")
	snapshots, sdkerr := vngcloud.VServerGateway().V2().VolumeService().ListSnapshotsByBlockVolumeID(context.Background(), opt)
	if sdkerr != nil {
		t.Fatalf("Expect nil but got %v", sdkerr)
	}

	if snapshots == nil {
		t.Fatalf("Expect not nil but got nil")
	}

	t.Log("Result: ", snapshots)
	t.Log("PASS")
}

func TestCreateSnapshotFailure(t *testing.T) {
	vngcloud := validSdkConfig()
	opt := v2.NewCreateSnapshotByBlockVolumeIDRequest(
		"teasdadasdadst",
		"vol-d360fd83-948d-4efa-ab46-aab97328e275").WithPermanently(true)
	_, sdkerr := vngcloud.VServerGateway().V2().VolumeService().CreateSnapshotByBlockVolumeID(context.Background(), opt)

	t.Log("Result: ", sdkerr)
	t.Log("PASS")
}

func TestDeleteSnapshot(t *testing.T) {
	vngcloud := validSdkConfig()
	opt := v2.NewDeleteSnapshotByIDRequest("snap-vol-pt-03e5891b-xxxx-4eb9-b2e6-be599f4e2a4b")
	sdkerr := vngcloud.VServerGateway().V2().VolumeService().DeleteSnapshotByID(context.Background(), opt)
	if sdkerr != nil {
		t.Fatalf("Expect nil but got %v", sdkerr)
	}

	t.Log("PASS")
}
