package v2

import (
	"testing"
)

func TestSecgroup_Fields(t *testing.T) {
	sg := Secgroup{ID: "sg-1", Name: "test", Description: "desc", Status: "ACTIVE", CreatedAt: "2024-01-01", IsSystem: true}
	if sg.ID != "sg-1" || sg.Name != "test" || sg.Description != "desc" || sg.Status != "ACTIVE" {
		t.Fatalf("unexpected: %+v", sg)
	}
	if sg.CreatedAt != "2024-01-01" {
		t.Fatalf("CreatedAt: got %q", sg.CreatedAt)
	}
	if !sg.IsSystem {
		t.Fatal("IsSystem: expected true")
	}
}

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
