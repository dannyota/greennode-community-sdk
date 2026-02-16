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
	if r.GetSecgroupName() != "my-sg" {
		t.Fatalf("GetSecgroupName: got %q", r.GetSecgroupName())
	}
}

func TestNewDeleteSecgroupByIDRequest(t *testing.T) {
	r := NewDeleteSecgroupByIDRequest("sg-123")
	if r.GetSecgroupID() != "sg-123" {
		t.Fatalf("got %q", r.GetSecgroupID())
	}
}

func TestNewGetSecgroupByIDRequest(t *testing.T) {
	r := NewGetSecgroupByIDRequest("sg-456")
	if r.GetSecgroupID() != "sg-456" {
		t.Fatalf("got %q", r.GetSecgroupID())
	}
}

func TestNewListSecgroupRequest(t *testing.T) {
	r := NewListSecgroupRequest()
	if r == nil {
		t.Fatal("expected non-nil")
	}
}
