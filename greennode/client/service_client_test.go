package client

import (
	"testing"
)

func TestServiceURL(t *testing.T) {
	tests := []struct {
		name     string
		endpoint string
		parts    []string
		want     string
	}{
		{
			name:     "no trailing slash",
			endpoint: "https://api.example.com",
			parts:    []string{"v2", "volumes"},
			want:     "https://api.example.com/v2/volumes",
		},
		{
			name:     "with trailing slash",
			endpoint: "https://api.example.com/",
			parts:    []string{"v2", "volumes"},
			want:     "https://api.example.com/v2/volumes",
		},
		{
			name:     "no parts",
			endpoint: "https://api.example.com",
			parts:    nil,
			want:     "https://api.example.com/",
		},
		{
			name:     "single part",
			endpoint: "https://api.example.com",
			parts:    []string{"health"},
			want:     "https://api.example.com/health",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sc := NewServiceClient().WithEndpoint(tt.endpoint)
			got := sc.ServiceURL(tt.parts...)
			if got != tt.want {
				t.Fatalf("got %q, want %q", got, tt.want)
			}
		})
	}
}

func TestNormalizeURL(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{"https://api.example.com", "https://api.example.com/"},
		{"https://api.example.com/", "https://api.example.com/"},
	}
	for _, tt := range tests {
		got := normalizeURL(tt.input)
		if got != tt.want {
			t.Fatalf("normalizeURL(%q) = %q, want %q", tt.input, got, tt.want)
		}
	}
}

func TestServiceClientBuilders(t *testing.T) {
	sc := NewServiceClient().
		WithEndpoint("https://api.example.com").
		WithName("test-service").
		WithProjectID("proj-1").
		WithZoneID("zone-a").
		WithUserID("user-x")

	if sc.GetProjectID() != "proj-1" {
		t.Fatalf("projectID: got %q", sc.GetProjectID())
	}
	if sc.GetZoneID() != "zone-a" {
		t.Fatalf("zoneID: got %q", sc.GetZoneID())
	}
	if sc.GetUserID() != "user-x" {
		t.Fatalf("userID: got %q", sc.GetUserID())
	}
}

func TestWithMoreHeaders(t *testing.T) {
	sc := NewServiceClient().WithMoreHeaders(map[string]string{"X-Test": "1"})
	if sc.moreHeaders["X-Test"] != "1" {
		t.Fatalf("got %v", sc.moreHeaders)
	}
}

// BUG: WithKVheader panics if moreHeaders is nil (no nil-map guard).
// This test documents the bug. If you fix it, change the test.
func TestWithKVheader_PanicOnNilMap(t *testing.T) {
	sc := NewServiceClient()
	defer func() {
		if r := recover(); r == nil {
			t.Fatal("expected panic when moreHeaders is nil, but got none — bug may have been fixed")
		}
	}()
	sc.WithKVheader("key", "value")
}
