package greennode

import (
	"context"
	"net/http"
	"testing"

	"github.com/dannyota/greennode-community-sdk/greennode/option"
)

func TestNewClientWithHTTPClient(t *testing.T) {
	custom := &http.Client{}
	c, err := NewClient(context.Background(), Config{Region: "hcm-3"},
		option.WithHTTPClient(custom))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	_ = c
}

func TestNewClientBackwardCompat(t *testing.T) {
	_, err := NewClient(context.Background(), Config{Region: "hcm-3"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}
