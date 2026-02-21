package v2

import (
	"testing"
)

func TestVolume_CanDelete(t *testing.T) {
	tests := []struct {
		name string
		vol  Volume
		want bool
	}{
		{"available no attachments", Volume{Status: "AVAILABLE"}, true},
		{"error status", Volume{Status: "ERROR"}, true},
		{"in-use", Volume{Status: "IN-USE", VmID: "srv-1"}, false},
		{"available with VmID", Volume{Status: "AVAILABLE", VmID: "srv-1"}, false},
		{"available with attached machines", Volume{Status: "AVAILABLE", AttachedMachine: []string{"srv-1"}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.vol.CanDelete(); got != tt.want {
				t.Fatalf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVolume_IsAvailable(t *testing.T) {
	if !(Volume{Status: "AVAILABLE"}).IsAvailable() {
		t.Fatal("expected available")
	}
	if (Volume{Status: "IN-USE"}).IsAvailable() {
		t.Fatal("expected not available")
	}
}

func TestVolume_IsInUse(t *testing.T) {
	if !(Volume{Status: "IN-USE"}).IsInUse() {
		t.Fatal("expected in-use")
	}
	if (Volume{Status: "AVAILABLE"}).IsInUse() {
		t.Fatal("expected not in-use")
	}
}

func TestVolume_IsError(t *testing.T) {
	if !(Volume{Status: "ERROR"}).IsError() {
		t.Fatal("expected error")
	}
	if (Volume{Status: "AVAILABLE"}).IsError() {
		t.Fatal("expected not error")
	}
}

func TestVolume_AttachedTheInstance(t *testing.T) {
	vol := Volume{VmID: "srv-1", AttachedMachine: []string{"srv-2", "srv-3"}}
	if !vol.AttachedTheInstance("srv-1") {
		t.Fatal("expected match on VmID")
	}
	if !vol.AttachedTheInstance("srv-2") {
		t.Fatal("expected match on AttachedMachine")
	}
	if vol.AttachedTheInstance("srv-99") {
		t.Fatal("expected no match")
	}
}

func TestListVolumes_Len(t *testing.T) {
	lst := ListVolumes{Items: []*Volume{{ID: "a"}, {ID: "b"}}}
	if lst.Len() != 2 {
		t.Fatalf("got %d", lst.Len())
	}
}
