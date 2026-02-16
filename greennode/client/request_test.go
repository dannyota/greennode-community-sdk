package client

import (
	"testing"
)

func TestNewRequest_Defaults(t *testing.T) {
	r := NewRequest()
	if r.okCodes == nil {
		t.Fatal("okCodes should be initialized")
	}
	if len(r.okCodes) != 0 {
		t.Fatalf("expected empty okCodes, got %d", len(r.okCodes))
	}
	if r.method != "" {
		t.Fatalf("expected empty method, got %q", r.method)
	}
	if r.skipAuth {
		t.Fatal("skipAuth should default to false")
	}
}

func TestWithOKCodes(t *testing.T) {
	r := NewRequest().WithOKCodes(200, 201)
	if !r.containsOKCode(200) {
		t.Fatal("expected 200")
	}
	if !r.containsOKCode(201) {
		t.Fatal("expected 201")
	}
	if r.containsOKCode(404) {
		t.Fatal("unexpected 404")
	}
}

func TestContainsOKCode_Multiple(t *testing.T) {
	r := NewRequest().WithOKCodes(200)
	// containsOKCode accepts variadic — returns true if ANY match
	if !r.containsOKCode(404, 200) {
		t.Fatal("expected match on 200")
	}
	if r.containsOKCode(404, 500) {
		t.Fatal("expected no match")
	}
}

func TestWithHeader(t *testing.T) {
	r := NewRequest().WithHeader("X-Custom", "val")
	if r.moreHeaders["X-Custom"] != "val" {
		t.Fatalf("got %q", r.moreHeaders["X-Custom"])
	}
}

func TestWithHeader_EmptyKeyOrValue(t *testing.T) {
	r := NewRequest().WithHeader("", "val")
	if r.moreHeaders != nil {
		t.Fatal("empty key should be ignored")
	}

	r2 := NewRequest().WithHeader("key", "")
	if r2.moreHeaders != nil {
		t.Fatal("empty value should be ignored")
	}
}

func TestWithUserID(t *testing.T) {
	r := NewRequest().WithUserID("user-123")
	if r.moreHeaders["portal-user-id"] != "user-123" {
		t.Fatalf("got %q", r.moreHeaders["portal-user-id"])
	}
}

func TestWithSkipAuth(t *testing.T) {
	r := NewRequest().WithSkipAuth(true)
	if !r.skipAuth {
		t.Fatal("expected skipAuth=true")
	}
}

func TestWithJSONBody(t *testing.T) {
	body := map[string]string{"key": "val"}
	r := NewRequest().WithJSONBody(body)
	if r.jsonBody == nil {
		t.Fatal("expected non-nil body")
	}
}

func TestWithJSONResponse(t *testing.T) {
	resp := &struct{ Name string }{}
	r := NewRequest().WithJSONResponse(resp)
	if r.jsonResponse != resp {
		t.Fatal("expected same response pointer")
	}
}

func TestWithJSONError(t *testing.T) {
	errResp := &struct{ Message string }{}
	r := NewRequest().WithJSONError(errResp)
	if r.jsonError != errResp {
		t.Fatal("expected same error pointer")
	}
}

func TestRequestMethod_DirectSet(t *testing.T) {
	r := NewRequest()
	r.method = MethodPost
	if r.method != MethodPost {
		t.Fatalf("got %q", r.method)
	}
}
