package v2

import (
	"testing"
)

func TestNewGetAccessTokenRequest(t *testing.T) {
	r := NewGetAccessTokenRequest("client-id", "client-secret")
	if r.ClientID != "client-id" {
		t.Fatalf("ClientID: got %q", r.ClientID)
	}
	if r.ClientSecret != "client-secret" {
		t.Fatalf("ClientSecret: got %q", r.ClientSecret)
	}
	if r.GrantType != "client_credentials" {
		t.Fatalf("GrantType: got %q", r.GrantType)
	}
}

func TestGetAccessTokenRequest_DirectFieldAccess(t *testing.T) {
	r := NewGetAccessTokenRequest("a", "b")
	r.ClientID = "new-id"
	r.ClientSecret = "new-secret"
	if r.ClientID != "new-id" {
		t.Fatalf("got %q", r.ClientID)
	}
	if r.ClientSecret != "new-secret" {
		t.Fatalf("got %q", r.ClientSecret)
	}
}
