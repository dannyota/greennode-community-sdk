package common

import (
	"testing"
)

// ---------------------------------------------------------------------------
// Ptr
// ---------------------------------------------------------------------------

func TestPtr(t *testing.T) {
	tests := []struct {
		name string
		fn   func(t *testing.T)
	}{
		{"int", func(t *testing.T) {
			p := Ptr(42)
			if *p != 42 {
				t.Fatalf("got %d, want 42", *p)
			}
		}},
		{"string", func(t *testing.T) {
			p := Ptr("hello")
			if *p != "hello" {
				t.Fatalf("got %q, want %q", *p, "hello")
			}
		}},
		{"bool", func(t *testing.T) {
			p := Ptr(true)
			if *p != true {
				t.Fatal("got false, want true")
			}
		}},
		{"struct", func(t *testing.T) {
			type S struct{ X int }
			p := Ptr(S{X: 7})
			if p.X != 7 {
				t.Fatalf("got %d, want 7", p.X)
			}
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, tt.fn)
	}
}

// ---------------------------------------------------------------------------
// StructToMap
// ---------------------------------------------------------------------------

func TestStructToMap(t *testing.T) {
	tests := []struct {
		name     string
		input    any
		wantKeys []string
		wantVals map[string]any
	}{
		{
			name: "struct with json tags",
			input: struct {
				Name string `json:"name"`
				Age  int    `json:"age"`
			}{"Alice", 30},
			wantKeys: []string{"name", "age"},
			wantVals: map[string]any{"name": "Alice", "age": float64(30)},
		},
		{
			name:     "empty struct",
			input:    struct{}{},
			wantKeys: nil,
			wantVals: map[string]any{},
		},
		{
			name: "omitempty skips zero values",
			input: struct {
				A string `json:"a,omitempty"`
				B string `json:"b"`
			}{A: "", B: ""},
			wantKeys: []string{"b"},
			wantVals: map[string]any{"b": ""},
		},
		{
			name:     "nil input",
			input:    nil,
			wantKeys: nil,
			wantVals: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := StructToMap(tt.input)

			if tt.wantVals == nil {
				if m != nil {
					t.Fatalf("expected nil map, got %v", m)
				}
				return
			}

			for _, k := range tt.wantKeys {
				if _, ok := m[k]; !ok {
					t.Errorf("missing key %q in map", k)
				}
			}

			for k, want := range tt.wantVals {
				got, ok := m[k]
				if !ok {
					t.Errorf("missing key %q", k)
					continue
				}
				if got != want {
					t.Errorf("key %q: got %v, want %v", k, got, want)
				}
			}
		})
	}
}

// ---------------------------------------------------------------------------
// Project
// ---------------------------------------------------------------------------

func TestProject(t *testing.T) {
	var p Project
	if p.GetProjectID() != "" {
		t.Fatal("zero-value project should return empty string")
	}

	p.ID = "proj-123"
	if got := p.GetProjectID(); got != "proj-123" {
		t.Fatalf("got %q, want %q", got, "proj-123")
	}
}

// ---------------------------------------------------------------------------
// Paging
// ---------------------------------------------------------------------------

func TestPaging(t *testing.T) {
	var p Paging
	if p.GetPage() != 0 || p.GetSize() != 0 {
		t.Fatal("zero-value paging should return 0")
	}

	p.Page = 3
	p.Size = 25
	if p.GetPage() != 3 {
		t.Fatalf("page: got %d, want 3", p.GetPage())
	}
	if p.GetSize() != 25 {
		t.Fatalf("size: got %d, want 25", p.GetSize())
	}
}

// ---------------------------------------------------------------------------
// PortalUser
// ---------------------------------------------------------------------------

func TestPortalUser(t *testing.T) {
	var pu PortalUser
	if pu.GetPortalUserID() != "" {
		t.Fatal("zero-value should return empty string")
	}

	pu.ID = "user-456"
	if got := pu.GetPortalUserID(); got != "user-456" {
		t.Fatalf("got %q, want %q", got, "user-456")
	}

	headers := pu.GetMapHeaders()
	if headers["portal-user-id"] != "user-456" {
		t.Fatalf("header: got %q, want %q", headers["portal-user-id"], "user-456")
	}
}

// ---------------------------------------------------------------------------
// ID-holder structs (table-driven)
// ---------------------------------------------------------------------------

func TestIDHolderStructs(t *testing.T) {
	tests := []struct {
		name   string
		set    func(id string)
		get    func() string
		field  string
		wantID string
	}{
		{"BlockVolumeCommon", nil, nil, "BlockVolumeID", "bv-1"},
		{"VolumeTypeCommon", nil, nil, "VolumeTypeID", "vt-1"},
		{"ServerCommon", nil, nil, "ServerID", "sv-1"},
		{"ServerGroupCommon", nil, nil, "ServerGroupID", "sg-1"},
		{"NetworkCommon", nil, nil, "NetworkID", "net-1"},
		{"InternalNetworkInterfaceCommon", nil, nil, "InternalNetworkInterfaceID", "ini-1"},
		{"WanCommon", nil, nil, "WanID", "wan-1"},
		{"SecgroupCommon", nil, nil, "SecgroupID", "scg-1"},
		{"LoadBalancerCommon", nil, nil, "LoadBalancerID", "lb-1"},
		{"ListenerCommon", nil, nil, "ListenerID", "lis-1"},
		{"PoolCommon", nil, nil, "PoolID", "pool-1"},
		{"PolicyCommon", nil, nil, "PolicyID", "pol-1"},
		{"PoolMemberCommon", nil, nil, "PoolMemberID", "pm-1"},
		{"SnapshotCommon", nil, nil, "SnapshotID", "snap-1"},
		{"SubnetCommon", nil, nil, "SubnetID", "sub-1"},
		{"EndpointCommon", nil, nil, "EndpointID", "ep-1"},
		{"VirtualAddressCommon", nil, nil, "VirtualAddressID", "va-1"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			switch tt.name {
			case "BlockVolumeCommon":
				s := BlockVolumeCommon{BlockVolumeID: tt.wantID}
				if s.GetBlockVolumeID() != tt.wantID {
					t.Fatalf("got %q, want %q", s.GetBlockVolumeID(), tt.wantID)
				}
			case "VolumeTypeCommon":
				s := VolumeTypeCommon{VolumeTypeID: tt.wantID}
				if s.GetVolumeTypeID() != tt.wantID {
					t.Fatalf("got %q, want %q", s.GetVolumeTypeID(), tt.wantID)
				}
			case "ServerCommon":
				s := ServerCommon{ServerID: tt.wantID}
				if s.GetServerID() != tt.wantID {
					t.Fatalf("got %q, want %q", s.GetServerID(), tt.wantID)
				}
			case "ServerGroupCommon":
				s := ServerGroupCommon{ServerGroupID: tt.wantID}
				if s.GetServerGroupID() != tt.wantID {
					t.Fatalf("got %q, want %q", s.GetServerGroupID(), tt.wantID)
				}
			case "NetworkCommon":
				s := NetworkCommon{NetworkID: tt.wantID}
				if s.GetNetworkID() != tt.wantID {
					t.Fatalf("got %q, want %q", s.GetNetworkID(), tt.wantID)
				}
			case "InternalNetworkInterfaceCommon":
				s := InternalNetworkInterfaceCommon{InternalNetworkInterfaceID: tt.wantID}
				if s.GetInternalNetworkInterfaceID() != tt.wantID {
					t.Fatalf("got %q, want %q", s.GetInternalNetworkInterfaceID(), tt.wantID)
				}
			case "WanCommon":
				s := WanCommon{WanID: tt.wantID}
				if s.GetWanID() != tt.wantID {
					t.Fatalf("got %q, want %q", s.GetWanID(), tt.wantID)
				}
			case "SecgroupCommon":
				s := SecgroupCommon{SecgroupID: tt.wantID}
				if s.GetSecgroupID() != tt.wantID {
					t.Fatalf("got %q, want %q", s.GetSecgroupID(), tt.wantID)
				}
			case "LoadBalancerCommon":
				s := LoadBalancerCommon{LoadBalancerID: tt.wantID}
				if s.GetLoadBalancerID() != tt.wantID {
					t.Fatalf("got %q, want %q", s.GetLoadBalancerID(), tt.wantID)
				}
			case "ListenerCommon":
				s := ListenerCommon{ListenerID: tt.wantID}
				if s.GetListenerID() != tt.wantID {
					t.Fatalf("got %q, want %q", s.GetListenerID(), tt.wantID)
				}
			case "PoolCommon":
				s := PoolCommon{PoolID: tt.wantID}
				if s.GetPoolID() != tt.wantID {
					t.Fatalf("got %q, want %q", s.GetPoolID(), tt.wantID)
				}
			case "PolicyCommon":
				s := PolicyCommon{PolicyID: tt.wantID}
				if s.GetPolicyID() != tt.wantID {
					t.Fatalf("got %q, want %q", s.GetPolicyID(), tt.wantID)
				}
			case "PoolMemberCommon":
				s := PoolMemberCommon{PoolMemberID: tt.wantID}
				if s.GetPoolMemberID() != tt.wantID {
					t.Fatalf("got %q, want %q", s.GetPoolMemberID(), tt.wantID)
				}
			case "SnapshotCommon":
				s := SnapshotCommon{SnapshotID: tt.wantID}
				if s.GetSnapshotID() != tt.wantID {
					t.Fatalf("got %q, want %q", s.GetSnapshotID(), tt.wantID)
				}
			case "SubnetCommon":
				s := SubnetCommon{SubnetID: tt.wantID}
				if s.GetSubnetID() != tt.wantID {
					t.Fatalf("got %q, want %q", s.GetSubnetID(), tt.wantID)
				}
			case "EndpointCommon":
				s := EndpointCommon{EndpointID: tt.wantID}
				if s.GetEndpointID() != tt.wantID {
					t.Fatalf("got %q, want %q", s.GetEndpointID(), tt.wantID)
				}
			case "VirtualAddressCommon":
				s := VirtualAddressCommon{VirtualAddressID: tt.wantID}
				if s.GetVirtualAddressID() != tt.wantID {
					t.Fatalf("got %q, want %q", s.GetVirtualAddressID(), tt.wantID)
				}
			}
		})
	}
}

// ---------------------------------------------------------------------------
// Tag
// ---------------------------------------------------------------------------

func TestTag(t *testing.T) {
	tag := Tag{Key: "env", Value: "prod", IsEdited: true}
	if tag.Key != "env" || tag.Value != "prod" || !tag.IsEdited {
		t.Fatalf("unexpected tag: %+v", tag)
	}
}

// ---------------------------------------------------------------------------
// Zone constants
// ---------------------------------------------------------------------------

func TestZoneConstants(t *testing.T) {
	zones := []Zone{HCM_03_1A_ZONE, HCM_03_1B_ZONE, HCM_03_1C_ZONE, HCM_03_BKK_01_ZONE, HAN_01_1A_ZONE}
	for _, z := range zones {
		if z == "" {
			t.Fatal("zone constant should not be empty")
		}
	}
}
