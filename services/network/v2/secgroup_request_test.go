package v2

import (
	"testing"
)

func TestNewCreateSecgroupRequest(t *testing.T) {
	r := NewCreateSecgroupRequest("my-sg", "test description")
	if r.Name != "my-sg" {
		t.Fatalf("Name: got %q", r.Name)
	}
	if r.Description != "test description" {
		t.Fatalf("Description: got %q", r.Description)
	}
}

func TestNewDeleteSecgroupByIDRequest(t *testing.T) {
	r := NewDeleteSecgroupByIDRequest("sg-123")
	if r.SecgroupID != "sg-123" {
		t.Fatalf("got %q", r.SecgroupID)
	}
}

func TestNewGetSecgroupByIDRequest(t *testing.T) {
	r := NewGetSecgroupByIDRequest("sg-456")
	if r.SecgroupID != "sg-456" {
		t.Fatalf("got %q", r.SecgroupID)
	}
}

func TestNewListSecgroupRequest(t *testing.T) {
	r := NewListSecgroupRequest(1, 10)
	if r == nil {
		t.Fatal("expected non-nil")
	}
	if r.Page != 1 {
		t.Fatalf("Page: got %d", r.Page)
	}
	if r.Size != 10 {
		t.Fatalf("Size: got %d", r.Size)
	}
}

func TestListSecgroupRequest_ToListQuery(t *testing.T) {
	r := &ListSecgroupRequest{Name: "web", Page: 2, Size: 5}
	q, err := r.ToListQuery()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if q == "" {
		t.Fatal("expected non-empty query")
	}
}
