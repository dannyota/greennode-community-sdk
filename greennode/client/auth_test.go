package client

import (
	"testing"
	"time"
)

func TestNewSdkAuthentication(t *testing.T) {
	auth := NewSdkAuthentication()
	if auth.AccessToken() != "" {
		t.Fatal("expected empty access token")
	}
	if auth.ExpiresAt() != 0 {
		t.Fatal("expected zero expiresAt")
	}
}

func TestNeedReauth_EmptyToken(t *testing.T) {
	auth := NewSdkAuthentication()
	if !auth.NeedReauth() {
		t.Fatal("empty token should need reauth")
	}
}

func TestNeedReauth_ExpiredToken(t *testing.T) {
	auth := NewSdkAuthentication().
		WithAccessToken("tok").
		WithExpiresAt(time.Now().Add(-1 * time.Hour).UnixNano())
	if !auth.NeedReauth() {
		t.Fatal("expired token should need reauth")
	}
}

func TestNeedReauth_ExpiringWithin5Min(t *testing.T) {
	auth := NewSdkAuthentication().
		WithAccessToken("tok").
		WithExpiresAt(time.Now().Add(3 * time.Minute).UnixNano())
	if !auth.NeedReauth() {
		t.Fatal("token expiring within 5 min should need reauth")
	}
}

func TestNeedReauth_ValidToken(t *testing.T) {
	auth := NewSdkAuthentication().
		WithAccessToken("tok").
		WithExpiresAt(time.Now().Add(1 * time.Hour).UnixNano())
	if auth.NeedReauth() {
		t.Fatal("valid token should not need reauth")
	}
}

func TestUpdateAuth(t *testing.T) {
	auth1 := NewSdkAuthentication().
		WithAccessToken("old").
		WithExpiresAt(100)

	auth2 := NewSdkAuthentication().
		WithAccessToken("new").
		WithExpiresAt(200)

	auth1.UpdateAuth(auth2)
	if auth1.AccessToken() != "new" {
		t.Fatalf("got %q, want %q", auth1.AccessToken(), "new")
	}
	if auth1.ExpiresAt() != 200 {
		t.Fatalf("got %d, want 200", auth1.ExpiresAt())
	}
}

func TestWithAccessToken(t *testing.T) {
	auth := NewSdkAuthentication().WithAccessToken("my-token")
	if auth.AccessToken() != "my-token" {
		t.Fatalf("got %q", auth.AccessToken())
	}
}

func TestWithExpiresAt(t *testing.T) {
	auth := NewSdkAuthentication().WithExpiresAt(12345)
	if auth.ExpiresAt() != 12345 {
		t.Fatalf("got %d", auth.ExpiresAt())
	}
}
