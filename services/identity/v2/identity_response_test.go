package v2

import (
	"encoding/json"
	"testing"
	"time"
)

func TestGetAccessTokenResponse_Unmarshal(t *testing.T) {
	jsonData := `{
		"access_token": "eyJhbGciOiJSUzI1NiJ9.test",
		"expires_in": 3600,
		"token_type": "Bearer",
		"refresh_expires_in": 0
	}`

	var resp GetAccessTokenResponse
	if err := json.Unmarshal([]byte(jsonData), &resp); err != nil {
		t.Fatalf("unmarshal error: %v", err)
	}
	if resp.AccessToken != "eyJhbGciOiJSUzI1NiJ9.test" {
		t.Fatalf("AccessToken: got %q", resp.AccessToken)
	}
	if resp.ExpiresIn != 3600 {
		t.Fatalf("ExpiresIn: got %d", resp.ExpiresIn)
	}
	if resp.TokenType != "Bearer" {
		t.Fatalf("TokenType: got %q", resp.TokenType)
	}
}

func TestGetAccessTokenResponse_ToEntityAccessToken(t *testing.T) {
	before := time.Now()
	resp := &GetAccessTokenResponse{
		AccessToken: "my-token",
		ExpiresIn:   3600,
	}
	entity := resp.ToEntityAccessToken()
	after := time.Now()

	if entity.Token != "my-token" {
		t.Fatalf("Token: got %q", entity.Token)
	}

	// ExpiresAt should be approximately now + 3600 seconds (in nanoseconds)
	expectedMin := before.Add(3600 * time.Second).UnixNano()
	expectedMax := after.Add(3600 * time.Second).UnixNano()
	if entity.ExpiresAt < expectedMin || entity.ExpiresAt > expectedMax {
		t.Fatalf("ExpiresAt out of range: %d not in [%d, %d]",
			entity.ExpiresAt, expectedMin, expectedMax)
	}
}

func TestGetAccessTokenResponse_ToEntityAccessToken_Fields(t *testing.T) {
	resp := &GetAccessTokenResponse{
		AccessToken: "token-abc",
		ExpiresIn:   7200,
	}
	entity := resp.ToEntityAccessToken()

	if entity.Token != "token-abc" {
		t.Fatalf("Token: got %q", entity.Token)
	}
	if entity.ExpiresAt <= 0 {
		t.Fatalf("ExpiresAt should be positive: got %d", entity.ExpiresAt)
	}
}
