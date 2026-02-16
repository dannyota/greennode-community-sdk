package v2

import (
	"encoding/json"
	"testing"
)

const sampleBlockVolumeJSON = `{
	"data": {
		"uuid": "vol-uuid-123",
		"projectId": "proj-1",
		"name": "my-volume",
		"size": 100,
		"status": "AVAILABLE",
		"volumeTypeId": "ssd-type",
		"volumeTypeZoneName": "zone-1",
		"iops": "3000",
		"serverId": "srv-1",
		"createdAt": "2024-01-01T00:00:00Z",
		"updatedAt": null,
		"bootable": false,
		"encryptionType": null,
		"bootIndex": 0,
		"multiAttach": false,
		"serverIdList": ["srv-1", "srv-2"],
		"location": null,
		"product": "standard",
		"persistentVolume": true,
		"migrateState": "none",
		"zone": {"uuid": "zone-uuid-1"}
	}
}`

func TestGetBlockVolumeByIDResponse_Unmarshal(t *testing.T) {
	var resp GetBlockVolumeByIDResponse
	if err := json.Unmarshal([]byte(sampleBlockVolumeJSON), &resp); err != nil {
		t.Fatalf("unmarshal error: %v", err)
	}
	if resp.Data.UUID != "vol-uuid-123" {
		t.Fatalf("UUID: got %q", resp.Data.UUID)
	}
	if resp.Data.Size != 100 {
		t.Fatalf("Size: got %d", resp.Data.Size)
	}
}

func TestGetBlockVolumeByIDResponse_ToEntityVolume(t *testing.T) {
	var resp GetBlockVolumeByIDResponse
	if err := json.Unmarshal([]byte(sampleBlockVolumeJSON), &resp); err != nil {
		t.Fatalf("unmarshal error: %v", err)
	}
	vol := resp.ToEntityVolume()

	if vol.ID != "vol-uuid-123" {
		t.Fatalf("ID: got %q", vol.ID)
	}
	if vol.Name != "my-volume" {
		t.Fatalf("Name: got %q", vol.Name)
	}
	if vol.Size != 100 {
		t.Fatalf("Size: got %d", vol.Size)
	}
	if vol.Status != "AVAILABLE" {
		t.Fatalf("Status: got %q", vol.Status)
	}
	if vol.VmID != "srv-1" {
		t.Fatalf("VmID: got %q", vol.VmID)
	}
	if len(vol.AttachedMachine) != 2 {
		t.Fatalf("AttachedMachine: got %d items", len(vol.AttachedMachine))
	}
	if vol.VolumeTypeID != "ssd-type" {
		t.Fatalf("VolumeTypeID: got %q", vol.VolumeTypeID)
	}
	if vol.MigrateState != "none" {
		t.Fatalf("MigrateState: got %q", vol.MigrateState)
	}
	if vol.ZoneID != "zone-uuid-1" {
		t.Fatalf("ZoneID: got %q", vol.ZoneID)
	}
}

func TestCreateBlockVolumeResponse_ToEntityVolume(t *testing.T) {
	var resp CreateBlockVolumeResponse
	if err := json.Unmarshal([]byte(sampleBlockVolumeJSON), &resp); err != nil {
		t.Fatalf("unmarshal error: %v", err)
	}
	vol := resp.ToEntityVolume()
	if vol.ID != "vol-uuid-123" {
		t.Fatalf("ID: got %q", vol.ID)
	}
}

func TestListBlockVolumesResponse_ToEntityListVolumes(t *testing.T) {
	jsonData := `{
		"page": 1,
		"pageSize": 10,
		"totalPage": 1,
		"totalItem": 2,
		"listData": [
			{"uuid": "v1", "name": "vol-1", "size": 50, "status": "AVAILABLE", "zone": {"uuid": "z1"}},
			{"uuid": "v2", "name": "vol-2", "size": 100, "status": "IN-USE", "zone": {"uuid": "z2"}}
		]
	}`
	var resp ListBlockVolumesResponse
	if err := json.Unmarshal([]byte(jsonData), &resp); err != nil {
		t.Fatalf("unmarshal error: %v", err)
	}
	lst := resp.ToEntityListVolumes()
	if lst.Len() != 2 {
		t.Fatalf("expected 2 items, got %d", lst.Len())
	}
	if lst.Items[0].ID != "v1" {
		t.Fatalf("first item ID: got %q", lst.Items[0].ID)
	}
	if lst.Items[1].Name != "vol-2" {
		t.Fatalf("second item Name: got %q", lst.Items[1].Name)
	}
}

func TestResizeBlockVolumeByIDResponse_ToEntityVolume(t *testing.T) {
	var resp ResizeBlockVolumeByIDResponse
	if err := json.Unmarshal([]byte(sampleBlockVolumeJSON), &resp); err != nil {
		t.Fatalf("unmarshal error: %v", err)
	}
	vol := resp.ToEntityVolume()
	if vol.ID != "vol-uuid-123" {
		t.Fatalf("ID: got %q", vol.ID)
	}
}

func TestGetUnderBlockVolumeIDResponse_ToEntityVolume(t *testing.T) {
	jsonData := `{"uuid": "under-123"}`
	var resp GetUnderBlockVolumeIDResponse
	if err := json.Unmarshal([]byte(jsonData), &resp); err != nil {
		t.Fatalf("unmarshal error: %v", err)
	}
	vol := resp.ToEntityVolume()
	if vol.UnderID != "under-123" {
		t.Fatalf("UnderID: got %q", vol.UnderID)
	}
}
