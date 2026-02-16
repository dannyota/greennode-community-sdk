package v2

import (
	"testing"
)

func TestAccessToken_Fields(t *testing.T) {
	at := AccessToken{Token: "tok", ExpiresAt: 123456789}
	if at.Token != "tok" {
		t.Fatalf("Token: got %q", at.Token)
	}
	if at.ExpiresAt != 123456789 {
		t.Fatalf("ExpiresAt: got %d", at.ExpiresAt)
	}
}
