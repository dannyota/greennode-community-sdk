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
	if r.RequestMethod() != "" {
		t.Fatalf("expected empty method, got %q", r.RequestMethod())
	}
	if r.SkipAuthentication() {
		t.Fatal("skipAuth should default to false")
	}
}

func TestWithOkCodes(t *testing.T) {
	r := NewRequest().WithOkCodes(200, 201)
	if !r.ContainsOkCode(200) {
		t.Fatal("expected 200")
	}
	if !r.ContainsOkCode(201) {
		t.Fatal("expected 201")
	}
	if r.ContainsOkCode(404) {
		t.Fatal("unexpected 404")
	}
}

func TestContainsOkCode_Multiple(t *testing.T) {
	r := NewRequest().WithOkCodes(200)
	// ContainsOkCode accepts variadic — returns true if ANY match
	if !r.ContainsOkCode(404, 200) {
		t.Fatal("expected match on 200")
	}
	if r.ContainsOkCode(404, 500) {
		t.Fatal("expected no match")
	}
}

func TestWithHeader(t *testing.T) {
	r := NewRequest().WithHeader("X-Custom", "val")
	h := r.MoreHeaders()
	if h["X-Custom"] != "val" {
		t.Fatalf("got %q", h["X-Custom"])
	}
}

func TestWithHeader_EmptyKeyOrValue(t *testing.T) {
	r := NewRequest().WithHeader("", "val")
	if r.MoreHeaders() != nil {
		t.Fatal("empty key should be ignored")
	}

	r2 := NewRequest().WithHeader("key", "")
	if r2.MoreHeaders() != nil {
		t.Fatal("empty value should be ignored")
	}
}

func TestWithMapHeaders(t *testing.T) {
	r := NewRequest().WithMapHeaders(map[string]string{
		"A": "1",
		"B": "2",
	})
	h := r.MoreHeaders()
	if h["A"] != "1" || h["B"] != "2" {
		t.Fatalf("got %v", h)
	}
}

func TestWithMapHeaders_NilInit(t *testing.T) {
	r := NewRequest()
	if r.MoreHeaders() != nil {
		t.Fatal("headers should be nil initially")
	}
	r.WithMapHeaders(map[string]string{"K": "V"})
	if r.MoreHeaders()["K"] != "V" {
		t.Fatal("expected K=V after WithMapHeaders")
	}
}

func TestWithUserID(t *testing.T) {
	r := NewRequest().WithUserID("user-123")
	h := r.MoreHeaders()
	if h["portal-user-id"] != "user-123" {
		t.Fatalf("got %q", h["portal-user-id"])
	}
}

func TestWithSkipAuth(t *testing.T) {
	r := NewRequest().WithSkipAuth(true)
	if !r.SkipAuthentication() {
		t.Fatal("expected skipAuth=true")
	}
}

func TestWithJSONBody(t *testing.T) {
	body := map[string]string{"key": "val"}
	r := NewRequest().WithJSONBody(body)
	if r.RequestBody() == nil {
		t.Fatal("expected non-nil body")
	}
}

func TestWithJSONResponse(t *testing.T) {
	resp := &struct{ Name string }{}
	r := NewRequest().WithJSONResponse(resp)
	if r.JSONResponse() != resp {
		t.Fatal("expected same response pointer")
	}
}

func TestWithJSONError(t *testing.T) {
	errResp := &struct{ Message string }{}
	r := NewRequest().WithJSONError(errResp)
	if r.JSONError() != errResp {
		t.Fatal("expected same error pointer")
	}
}

func TestWithJSONResponseOverwrite(t *testing.T) {
	r := NewRequest()
	resp := &struct{ X int }{}
	r.WithJSONResponse(resp)
	if r.JSONResponse() != resp {
		t.Fatal("expected same pointer after WithJSONResponse")
	}
}

func TestWithJSONErrorOverwrite(t *testing.T) {
	r := NewRequest()
	errResp := &struct{ X int }{}
	r.WithJSONError(errResp)
	if r.JSONError() != errResp {
		t.Fatal("expected same pointer after WithJSONError")
	}
}

func TestWithRequestMethod(t *testing.T) {
	r := NewRequest().WithRequestMethod(MethodPost)
	if r.RequestMethod() != "POST" {
		t.Fatalf("got %q", r.RequestMethod())
	}
}
