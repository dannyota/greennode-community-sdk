package v2

import (
	"testing"
)

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
