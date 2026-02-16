package client

import (
	"testing"
	"time"
)

func TestToken_ZeroValue(t *testing.T) {
	tok := &Token{}
	if tok.AccessToken != "" {
		t.Fatal("expected empty access token")
	}
	if tok.ExpiresAt != 0 {
		t.Fatal("expected zero expiresAt")
	}
}

func TestNeedsReauth_EmptyToken(t *testing.T) {
	tok := &Token{}
	if !tok.NeedsReauth() {
		t.Fatal("empty token should need reauth")
	}
}

func TestNeedsReauth_ExpiredToken(t *testing.T) {
	tok := &Token{
		AccessToken: "tok",
		ExpiresAt:   time.Now().Add(-1 * time.Hour).UnixNano(),
	}
	if !tok.NeedsReauth() {
		t.Fatal("expired token should need reauth")
	}
}

func TestNeedsReauth_ExpiringWithin5Min(t *testing.T) {
	tok := &Token{
		AccessToken: "tok",
		ExpiresAt:   time.Now().Add(3 * time.Minute).UnixNano(),
	}
	if !tok.NeedsReauth() {
		t.Fatal("token expiring within 5 min should need reauth")
	}
}

func TestNeedsReauth_ValidToken(t *testing.T) {
	tok := &Token{
		AccessToken: "tok",
		ExpiresAt:   time.Now().Add(1 * time.Hour).UnixNano(),
	}
	if tok.NeedsReauth() {
		t.Fatal("valid token should not need reauth")
	}
}
