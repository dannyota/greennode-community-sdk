package entity

import (
	"testing"
)

// ---------------------------------------------------------------------------
// Volume
// ---------------------------------------------------------------------------

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
	if !(Volume{Status: ServerStatusError}).IsError() {
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

// ---------------------------------------------------------------------------
// Server
// ---------------------------------------------------------------------------

func TestServer_CanDelete(t *testing.T) {
	tests := []struct {
		name   string
		status string
		want   bool
	}{
		{"active", ServerStatusActive, true},
		{"error", ServerStatusError, true},
		{"stopped", ServerStatusStopped, true},
		{"creating", "CREATING", false},
		{"deleting", "DELETING", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sv := Server{Status: tt.status}
			if got := sv.CanDelete(); got != tt.want {
				t.Fatalf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestServer_IsRunning(t *testing.T) {
	if !(Server{Status: ServerStatusActive}).IsRunning() {
		t.Fatal("expected running")
	}
	if (Server{Status: ServerStatusStopped}).IsRunning() {
		t.Fatal("expected not running")
	}
}

func TestServer_CanAttachFloatingIp(t *testing.T) {
	tests := []struct {
		name string
		sv   Server
		want bool
	}{
		{"no interfaces", Server{}, false},
		{"has interface, no floating", Server{
			InternalInterfaces: []NetworkInterface{{Uuid: "nic-1"}},
		}, true},
		{"has interface with floating", Server{
			InternalInterfaces: []NetworkInterface{{Uuid: "nic-1", FloatingIp: "1.2.3.4"}},
		}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.sv.CanAttachFloatingIp(); got != tt.want {
				t.Fatalf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestServer_GetInternalInterfaceWanInfo(t *testing.T) {
	sv := Server{
		InternalInterfaces: []NetworkInterface{
			{Uuid: "nic-1"},
			{Uuid: "nic-2", FloatingIp: "10.0.0.1", FloatingIpID: "fip-1"},
		},
	}
	uuid, fipID, fip, found := sv.GetInternalInterfaceWanInfo()
	if !found {
		t.Fatal("expected found")
	}
	if uuid != "nic-2" || fipID != "fip-1" || fip != "10.0.0.1" {
		t.Fatalf("got %q, %q, %q", uuid, fipID, fip)
	}
}

func TestServer_GetInternalInterfaceWanInfo_NotFound(t *testing.T) {
	sv := Server{
		InternalInterfaces: []NetworkInterface{{Uuid: "nic-1"}},
	}
	_, _, _, found := sv.GetInternalInterfaceWanInfo()
	if found {
		t.Fatal("expected not found")
	}
}

func TestServer_GetInternalNetworkInterfaceIDs(t *testing.T) {
	sv := Server{
		InternalInterfaces: []NetworkInterface{
			{Uuid: "nic-1"},
			{Uuid: "nic-2"},
		},
	}
	ids := sv.GetInternalNetworkInterfaceIDs()
	if len(ids) != 2 || ids[0] != "nic-1" || ids[1] != "nic-2" {
		t.Fatalf("got %v", ids)
	}
}

func TestServer_InternalNetworkInterfacePossible(t *testing.T) {
	if (Server{}).InternalNetworkInterfacePossible() {
		t.Fatal("empty server should return false")
	}
	if !(Server{InternalInterfaces: []NetworkInterface{{}}}).InternalNetworkInterfacePossible() {
		t.Fatal("expected true with interfaces")
	}
}

// ---------------------------------------------------------------------------
// Secgroup
// ---------------------------------------------------------------------------

func TestSecgroup_Fields(t *testing.T) {
	sg := Secgroup{ID: "sg-1", Name: "test", Description: "desc", Status: "ACTIVE"}
	if sg.ID != "sg-1" || sg.Name != "test" || sg.Description != "desc" || sg.Status != "ACTIVE" {
		t.Fatalf("unexpected: %+v", sg)
	}
}

// ---------------------------------------------------------------------------
// SecgroupRule list
// ---------------------------------------------------------------------------

func TestListSecgroupRules_Len_Get(t *testing.T) {
	lst := ListSecgroupRules{Items: []*SecgroupRule{{ID: "r1"}, {ID: "r2"}}}
	if lst.Len() != 2 {
		t.Fatalf("Len: got %d", lst.Len())
	}
	if lst.Get(0).ID != "r1" {
		t.Fatalf("Get(0): got %q", lst.Get(0).ID)
	}
	if lst.Get(-1) != nil {
		t.Fatal("Get(-1): expected nil")
	}
	if lst.Get(5) != nil {
		t.Fatal("Get(5): expected nil")
	}
}

// ---------------------------------------------------------------------------
// AccessToken
// ---------------------------------------------------------------------------

func TestAccessToken_ToSdkAuthentication(t *testing.T) {
	at := AccessToken{Token: "tok", ExpiresAt: 123456789}
	auth := at.ToSdkAuthentication()
	if auth.AccessToken() != "tok" {
		t.Fatalf("AccessToken: got %q", auth.AccessToken())
	}
	if auth.ExpiresAt() != 123456789 {
		t.Fatalf("ExpiresAt: got %d", auth.ExpiresAt())
	}
}
