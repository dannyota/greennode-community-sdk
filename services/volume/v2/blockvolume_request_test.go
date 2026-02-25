package v2

import (
	"encoding/json"
	"strings"
	"testing"
)

func TestNewCreateBlockVolumeRequest(t *testing.T) {
	r := NewCreateBlockVolumeRequest("my-vol", "ssd-type", 50)
	if r.Name != "my-vol" {
		t.Fatalf("Name: got %q", r.Name)
	}
	if r.VolumeTypeID != "ssd-type" {
		t.Fatalf("VolumeTypeID: got %q", r.VolumeTypeID)
	}
	if r.Size != 50 {
		t.Fatalf("Size: got %d", r.Size)
	}
	if r.CreatedFrom != CreateFromNew {
		t.Fatalf("CreatedFrom: got %q, want %q", r.CreatedFrom, CreateFromNew)
	}
}

func TestCreateBlockVolumeRequest_DirectFieldAccess(t *testing.T) {
	r := NewCreateBlockVolumeRequest("vol", "type", 10)
	r.Zone = "zone-a"
	r.PoolName = "pool-1"
	r.IsPoc = true
	r.IsEnableAutoRenew = true
	r.MultiAttach = true
	r.Size = 100
	r.EncryptionType = AesXtsPlain64_128
	r.VolumeTypeID = "new-type"
	r.Tags = NewVolumeTags("env", "prod", "team", "infra")

	if r.Zone != "zone-a" {
		t.Fatalf("Zone: got %q", r.Zone)
	}
	if r.PoolName != "pool-1" {
		t.Fatalf("PoolName: got %q", r.PoolName)
	}
	if !r.IsPoc {
		t.Fatal("IsPoc should be true")
	}
	if !r.IsEnableAutoRenew {
		t.Fatal("IsEnableAutoRenew should be true")
	}
	if !r.MultiAttach {
		t.Fatal("MultiAttach should be true")
	}
	if r.Size != 100 {
		t.Fatalf("Size: got %d", r.Size)
	}
	if r.EncryptionType != AesXtsPlain64_128 {
		t.Fatalf("EncryptionType: got %q", r.EncryptionType)
	}
	if r.VolumeTypeID != "new-type" {
		t.Fatalf("VolumeType: got %q", r.VolumeTypeID)
	}
	if len(r.Tags) != 2 {
		t.Fatalf("Tags: got %d, want 2", len(r.Tags))
	}
	if r.Tags[0].Key != "env" || r.Tags[0].Value != "prod" {
		t.Fatalf("Tags[0]: got %+v", r.Tags[0])
	}
}

func TestNewVolumeTags_OddCount(t *testing.T) {
	tags := NewVolumeTags("key1")
	if len(tags) != 1 {
		t.Fatalf("expected 1 tag, got %d", len(tags))
	}
	if tags[0].Value != "none" {
		t.Fatalf("odd tag value: got %q, want %q", tags[0].Value, "none")
	}
}

func TestCreateBlockVolumeRequest_VolumeRestoreFromSnapshot(t *testing.T) {
	r := NewCreateBlockVolumeRequest("v", "t", 1)
	r.CreatedFrom = CreateFromSnapshot
	r.ConfigureVolumeRestore = &ConfigureVolumeRestore{
		SnapshotVolumePointID: "snap-123",
		VolumeTypeID:          "vtype-456",
	}
	if r.CreatedFrom != CreateFromSnapshot {
		t.Fatalf("CreatedFrom: got %q", r.CreatedFrom)
	}
	if r.ConfigureVolumeRestore == nil {
		t.Fatal("ConfigureVolumeRestore should be set")
	}
	if r.ConfigureVolumeRestore.SnapshotVolumePointID != "snap-123" {
		t.Fatalf("got %q", r.ConfigureVolumeRestore.SnapshotVolumePointID)
	}
}

func TestNewDeleteBlockVolumeByIDRequest(t *testing.T) {
	r := NewDeleteBlockVolumeByIDRequest("vol-123")
	if r.BlockVolumeID != "vol-123" {
		t.Fatalf("got %q", r.BlockVolumeID)
	}
}

func TestNewListBlockVolumesRequest(t *testing.T) {
	r := NewListBlockVolumesRequest(2, 50)
	if r.Page != 2 || r.Size != 50 {
		t.Fatalf("got page=%d, size=%d", r.Page, r.Size)
	}
}

func TestListBlockVolumesRequest_ToQuery(t *testing.T) {
	r := NewListBlockVolumesRequest(1, 25)
	r.Name = "test-vol"
	q, err := r.ToQuery()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !strings.Contains(q, "name=test-vol") {
		t.Fatalf("query missing name: %q", q)
	}
	if !strings.Contains(q, "page=1") {
		t.Fatalf("query missing page: %q", q)
	}
	if !strings.Contains(q, "size=25") {
		t.Fatalf("query missing size: %q", q)
	}
}

func TestListBlockVolumesRequest_ToQuery_ZeroPageSize(t *testing.T) {
	r := &ListBlockVolumesRequest{Name: "x"}
	q, _ := r.ToQuery()
	// page=0 and size=0 should be omitted
	if strings.Contains(q, "page=") {
		t.Fatalf("zero page should be omitted: %q", q)
	}
	if strings.Contains(q, "size=") {
		t.Fatalf("zero size should be omitted: %q", q)
	}
}

func TestListBlockVolumesRequest_getDefaultQuery(t *testing.T) {
	r := &ListBlockVolumesRequest{}
	q := r.getDefaultQuery()
	if !strings.Contains(q, "page=1") {
		t.Fatalf("default query missing page: %q", q)
	}
	if !strings.Contains(q, "size=50") {
		t.Fatalf("default query missing size: %q", q)
	}
}

func TestCreateBlockVolumeRequest_JSONMarshal(t *testing.T) {
	r := NewCreateBlockVolumeRequest("vol", "ssd", 100)
	b, err := json.Marshal(r)
	if err != nil {
		t.Fatalf("marshal error: %v", err)
	}
	var m map[string]any
	if err := json.Unmarshal(b, &m); err != nil {
		t.Fatalf("unmarshal error: %v", err)
	}
	if m["name"] != "vol" {
		t.Fatalf("name: got %v", m["name"])
	}
	if m["volumeTypeId"] != "ssd" {
		t.Fatalf("volumeTypeId: got %v", m["volumeTypeId"])
	}
	if m["size"] != float64(100) {
		t.Fatalf("size: got %v", m["size"])
	}
	if m["createdFrom"] != "NEW" {
		t.Fatalf("createdFrom: got %v", m["createdFrom"])
	}
}
