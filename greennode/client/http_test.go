package client

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"sync/atomic"
	"testing"
	"time"
)

func TestDoRequest_JSONResponseUnmarshal(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{"name": "alice"})
	}))
	defer ts.Close()

	type resp struct {
		Name string `json:"name"`
	}

	var got resp
	hc := NewHTTPClient().WithRetryCount(0)
	hc.token = &Token{
		AccessToken: "tok",
		ExpiresAt:   time.Now().Add(1 * time.Hour).UnixNano(),
	}

	req := NewRequest().
		WithRequestMethod(MethodGet).
		WithOkCodes(200).
		WithJSONResponse(&got).
		WithSkipAuth(true)

	_, err := hc.DoRequest(context.Background(), ts.URL, req)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got.Name != "alice" {
		t.Fatalf("got %q, want %q", got.Name, "alice")
	}
}

func TestDoRequest_JSONRequestBody(t *testing.T) {
	var received map[string]string
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if ct := r.Header.Get("Content-Type"); ct != "application/json" {
			t.Errorf("Content-Type = %q, want application/json", ct)
		}
		body, _ := io.ReadAll(r.Body)
		json.Unmarshal(body, &received)
		w.WriteHeader(http.StatusOK)
	}))
	defer ts.Close()

	hc := NewHTTPClient().WithRetryCount(0)
	req := NewRequest().
		WithRequestMethod(MethodPost).
		WithJSONBody(map[string]string{"key": "value"}).
		WithOkCodes(200).
		WithSkipAuth(true)

	_, err := hc.DoRequest(context.Background(), ts.URL, req)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if received["key"] != "value" {
		t.Fatalf("got %v", received)
	}
}

func TestDoRequest_RetryOnFailure(t *testing.T) {
	var attempts int32
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		n := atomic.AddInt32(&attempts, 1)
		if n < 3 {
			// Close connection abruptly to simulate network error
			hj, ok := w.(http.Hijacker)
			if !ok {
				t.Fatal("server doesn't support hijacking")
			}
			conn, _, _ := hj.Hijack()
			conn.Close()
			return
		}
		w.WriteHeader(http.StatusOK)
	}))
	defer ts.Close()

	hc := NewHTTPClient().WithRetryCount(3).WithSleep(1 * time.Millisecond)
	req := NewRequest().
		WithRequestMethod(MethodGet).
		WithOkCodes(200).
		WithSkipAuth(true)

	_, err := hc.DoRequest(context.Background(), ts.URL, req)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got := atomic.LoadInt32(&attempts); got < 3 {
		t.Fatalf("expected at least 3 attempts, got %d", got)
	}
}

func TestDoRequest_HeaderPropagation(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if got := r.Header.Get("X-Custom"); got != "custom-val" {
			t.Errorf("X-Custom = %q, want %q", got, "custom-val")
		}
		if got := r.Header.Get("X-Default"); got != "default-val" {
			t.Errorf("X-Default = %q, want %q", got, "default-val")
		}
		w.WriteHeader(http.StatusOK)
	}))
	defer ts.Close()

	hc := NewHTTPClient().WithRetryCount(0).WithKvDefaultHeaders("X-Default", "default-val")
	req := NewRequest().
		WithRequestMethod(MethodGet).
		WithHeader("X-Custom", "custom-val").
		WithOkCodes(200).
		WithSkipAuth(true)

	_, err := hc.DoRequest(context.Background(), ts.URL, req)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestDoRequest_Timeout(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(500 * time.Millisecond)
		w.WriteHeader(http.StatusOK)
	}))
	defer ts.Close()

	hc := NewHTTPClient().WithRetryCount(0).WithTimeout(50 * time.Millisecond)
	req := NewRequest().
		WithRequestMethod(MethodGet).
		WithOkCodes(200).
		WithSkipAuth(true)

	_, err := hc.DoRequest(context.Background(), ts.URL, req)
	if err == nil {
		t.Fatal("expected timeout error")
	}
}

func TestDoRequest_JSONErrorUnmarshal(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"message": "bad input"})
	}))
	defer ts.Close()

	type errResp struct {
		Message string `json:"message"`
	}

	var got errResp
	hc := NewHTTPClient().WithRetryCount(0)
	req := NewRequest().
		WithRequestMethod(MethodPost).
		WithOkCodes(200).
		WithJSONError(&got).
		WithSkipAuth(true)

	_, err := hc.DoRequest(context.Background(), ts.URL, req)
	// No SDK error for non-special status codes — the response is returned as-is
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got.Message != "bad input" {
		t.Fatalf("got %q, want %q", got.Message, "bad input")
	}
}

func TestDoRequest_HTTPMethods(t *testing.T) {
	methods := []struct {
		method requestMethod
		want   string
	}{
		{MethodGet, "GET"},
		{MethodPost, "POST"},
		{MethodPut, "PUT"},
		{MethodPatch, "PATCH"},
		{MethodDelete, "DELETE"},
	}

	for _, tt := range methods {
		t.Run(tt.want, func(t *testing.T) {
			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if r.Method != tt.want {
					t.Errorf("method = %q, want %q", r.Method, tt.want)
				}
				w.WriteHeader(http.StatusOK)
			}))
			defer ts.Close()

			hc := NewHTTPClient().WithRetryCount(0)
			req := NewRequest().
				WithRequestMethod(tt.method).
				WithOkCodes(200).
				WithSkipAuth(true)

			_, err := hc.DoRequest(context.Background(), ts.URL, req)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
		})
	}
}

func TestDoRequest_StatusCodeErrors(t *testing.T) {
	tests := []struct {
		name       string
		statusCode int
	}{
		{"internal server error", http.StatusInternalServerError},
		{"service unavailable", http.StatusServiceUnavailable},
		{"forbidden", http.StatusForbidden},
		{"too many requests", http.StatusTooManyRequests},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(tt.statusCode)
			}))
			defer ts.Close()

			hc := NewHTTPClient().WithRetryCount(0)
			req := NewRequest().
				WithRequestMethod(MethodGet).
				WithOkCodes(200).
				WithSkipAuth(true)

			_, err := hc.DoRequest(context.Background(), ts.URL, req)
			if err == nil {
				t.Fatalf("expected error for status %d", tt.statusCode)
			}
		})
	}
}

func TestNewHTTPClient_Defaults(t *testing.T) {
	hc := NewHTTPClient()
	if hc.retryCount != 3 {
		t.Fatalf("retryCount = %d, want 3", hc.retryCount)
	}
	if hc.retryInterval != 10*time.Second {
		t.Fatalf("retryInterval = %v, want 10s", hc.retryInterval)
	}
	if hc.client.Timeout != 120*time.Second {
		t.Fatalf("timeout = %v, want 120s", hc.client.Timeout)
	}
}
