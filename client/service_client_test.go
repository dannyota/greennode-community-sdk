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
			sc := &ServiceClient{Endpoint: NormalizeURL(tt.endpoint)}
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
		got := NormalizeURL(tt.input)
		if got != tt.want {
			t.Fatalf("NormalizeURL(%q) = %q, want %q", tt.input, got, tt.want)
		}
	}
}

func TestServiceClient_Fields(t *testing.T) {
	sc := &ServiceClient{
		Endpoint:  NormalizeURL("https://api.example.com"),
		ProjectID: "proj-1",
		ZoneID:    "zone-a",
	}

	if sc.ProjectID != "proj-1" {
		t.Fatalf("projectID: got %q", sc.ProjectID)
	}
	if sc.ZoneID != "zone-a" {
		t.Fatalf("zoneID: got %q", sc.ZoneID)
	}
}
