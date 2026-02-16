package sdkerror

import (
	"errors"
	"fmt"
	"testing"
)

// ---------------------------------------------------------------------------
// Builder chain
// ---------------------------------------------------------------------------

func TestWithErrorCode(t *testing.T) {
	e := &SdkError{}
	e.WithErrorCode(EcInternalServerError)
	if e.ErrorCode() != EcInternalServerError {
		t.Fatalf("got %q, want %q", e.ErrorCode(), EcInternalServerError)
	}
}

func TestWithMessage(t *testing.T) {
	e := &SdkError{}
	e.WithMessage("boom")
	if e.GetMessage() != "boom" {
		t.Fatalf("got %q, want %q", e.GetMessage(), "boom")
	}
}

func TestWithErrors_Single(t *testing.T) {
	inner := fmt.Errorf("inner")
	e := &SdkError{}
	e.WithErrors(inner)
	if e.Err() != inner {
		t.Fatalf("got %v, want %v", e.Err(), inner)
	}
}

func TestWithErrors_Multiple(t *testing.T) {
	e1, e2 := fmt.Errorf("a"), fmt.Errorf("b")
	e := &SdkError{}
	e.WithErrors(e1, e2)
	if e.Err() == nil {
		t.Fatal("expected joined error")
	}
	if !errors.Is(e.Err(), e1) || !errors.Is(e.Err(), e2) {
		t.Fatal("joined error should unwrap both inner errors")
	}
}

func TestWithErrors_Empty(t *testing.T) {
	e := &SdkError{}
	e.WithErrors()
	if e.Err() != nil {
		t.Fatal("expected nil error")
	}
}

func TestAppendCategories_Multiple(t *testing.T) {
	e := &SdkError{}
	e.AppendCategories(ErrCatQuota, ErrCatIAM)
	if !e.IsCategory(ErrCatQuota) {
		t.Fatal("expected quota category")
	}
	if !e.IsCategory(ErrCatIAM) {
		t.Fatal("expected iam category")
	}
}

func TestWithKVparameters(t *testing.T) {
	e := &SdkError{}
	e.WithKVparameters("key1", "val1", "key2", 42)
	params := e.Parameters()
	if params["key1"] != "val1" {
		t.Fatalf("key1: got %v, want val1", params["key1"])
	}
	if params["key2"] != 42 {
		t.Fatalf("key2: got %v, want 42", params["key2"])
	}
}

func TestWithKVparameters_OddCount(t *testing.T) {
	e := &SdkError{}
	e.WithKVparameters("key1", "val1", "key2")
	params := e.Parameters()
	if params["key2"] != nil {
		t.Fatalf("key2: got %v, want nil", params["key2"])
	}
}

func TestWithKVparameters_NonStringKey(t *testing.T) {
	e := &SdkError{}
	e.WithKVparameters(123, "val")
	params := e.Parameters()
	if len(params) != 0 {
		t.Fatalf("expected empty params for non-string key, got %v", params)
	}
}

// ---------------------------------------------------------------------------
// WithParameters — bug: nil map early-returns after initializing sync.Map
// ---------------------------------------------------------------------------

func TestWithParameters_NilMap(t *testing.T) {
	// BUG: When params is nil, WithParameters initializes sync.Map but
	// returns early without storing anything. This is correct for nil
	// but the control flow is surprising. Passing a non-nil empty map
	// should work fine.
	e := &SdkError{}
	e.WithParameters(nil)
	// After this, e.parameters should be non-nil (sync.Map initialized)
	if e.parameters == nil {
		t.Fatal("expected parameters to be initialized even for nil map")
	}
	if len(e.Parameters()) != 0 {
		t.Fatal("expected empty parameters")
	}
}

func TestWithParameters_NonNil(t *testing.T) {
	e := &SdkError{}
	// First call with nil initializes sync.Map
	e.WithParameters(nil)
	// Second call with actual data should store
	e.WithParameters(map[string]any{"a": 1, "b": 2})
	params := e.Parameters()
	if params["a"] != 1 || params["b"] != 2 {
		t.Fatalf("got %v", params)
	}
}

// ---------------------------------------------------------------------------
// Getters
// ---------------------------------------------------------------------------

func TestErrorCode(t *testing.T) {
	e := &SdkError{errorCode: EcPermissionDenied}
	if e.ErrorCode() != EcPermissionDenied {
		t.Fatalf("got %q", e.ErrorCode())
	}
	if e.StringErrorCode() != string(EcPermissionDenied) {
		t.Fatalf("StringErrorCode: got %q", e.StringErrorCode())
	}
}

func TestErrorMessages_NoInner(t *testing.T) {
	e := &SdkError{message: "hello"}
	if e.ErrorMessages() != "hello" {
		t.Fatalf("got %q", e.ErrorMessages())
	}
}

func TestErrorMessages_WithInner(t *testing.T) {
	e := &SdkError{message: "outer", error: fmt.Errorf("inner")}
	want := "outer: inner"
	if e.ErrorMessages() != want {
		t.Fatalf("got %q, want %q", e.ErrorMessages(), want)
	}
}

func TestError_ImplementsErrorInterface(t *testing.T) {
	e := &SdkError{message: "test"}
	var err error = e
	if err.Error() != "test" {
		t.Fatalf("got %q", err.Error())
	}
}

func TestErrorCategories(t *testing.T) {
	e := &SdkError{}
	e.AppendCategories(ErrCatQuota, ErrCatProductVlb)
	cats := e.ErrorCategories()
	if len(cats) != 2 {
		t.Fatalf("expected 2 categories, got %d", len(cats))
	}
}

func TestListParameters(t *testing.T) {
	e := &SdkError{}
	e.WithKVparameters("k", "v")
	lp := e.ListParameters()
	if len(lp) != 2 {
		t.Fatalf("expected 2 items, got %d", len(lp))
	}
}

func TestListParameters_Nil(t *testing.T) {
	e := &SdkError{}
	lp := e.ListParameters()
	if lp != nil {
		t.Fatalf("expected nil, got %v", lp)
	}
}

// ---------------------------------------------------------------------------
// Matching: IsError, IsErrorAny, IsCategory, IsCategories
// ---------------------------------------------------------------------------

func TestIsError(t *testing.T) {
	e := &SdkError{errorCode: EcVServerVolumeNotFound}
	if !e.IsError(EcVServerVolumeNotFound) {
		t.Fatal("expected match")
	}
	if e.IsError(EcUnknownError) {
		t.Fatal("expected no match")
	}
}

func TestIsErrorAny(t *testing.T) {
	e := &SdkError{errorCode: EcVLBPoolNotFound}
	if !e.IsErrorAny(EcVLBPoolNotFound, EcVLBPoolInUse) {
		t.Fatal("expected match")
	}
	if e.IsErrorAny(EcUnknownError, EcPermissionDenied) {
		t.Fatal("expected no match")
	}
}

func TestIsCategory_NilCategories(t *testing.T) {
	e := &SdkError{}
	if e.IsCategory(ErrCatQuota) {
		t.Fatal("nil categories should return false")
	}
}

func TestIsCategories_NilCategories(t *testing.T) {
	e := &SdkError{}
	if e.IsCategories(ErrCatQuota) {
		t.Fatal("nil categories should return false")
	}
}

func TestIsCategories(t *testing.T) {
	e := &SdkError{}
	e.AppendCategories(ErrCatQuota)
	if !e.IsCategories(ErrCatIAM, ErrCatQuota) {
		t.Fatal("expected match on quota")
	}
	if e.IsCategories(ErrCatIAM, ErrCatInfra) {
		t.Fatal("expected no match")
	}
}

// ---------------------------------------------------------------------------
// Category management: AppendCategories, RemoveCategories
// ---------------------------------------------------------------------------

func TestAppendCategories(t *testing.T) {
	e := &SdkError{}
	e.AppendCategories(ErrCatQuota)
	if !e.IsCategory(ErrCatQuota) {
		t.Fatal("expected quota category after append")
	}
	e.AppendCategories(ErrCatIAM)
	if !e.IsCategory(ErrCatIAM) {
		t.Fatal("expected iam category after append")
	}
}

func TestRemoveCategories(t *testing.T) {
	e := &SdkError{}
	e.AppendCategories(ErrCatQuota, ErrCatIAM)
	e.RemoveCategories(ErrCatQuota)
	if e.IsCategory(ErrCatQuota) {
		t.Fatal("expected quota removed")
	}
	if !e.IsCategory(ErrCatIAM) {
		t.Fatal("expected iam still present")
	}
}

func TestRemoveCategories_NilCategories(t *testing.T) {
	e := &SdkError{}
	// Should not panic
	e.RemoveCategories(ErrCatQuota)
}

// ---------------------------------------------------------------------------
// errors.Is / errors.Unwrap
// ---------------------------------------------------------------------------

func TestIs_MatchesByErrorCode(t *testing.T) {
	e1 := &SdkError{errorCode: EcPermissionDenied}
	e2 := &SdkError{errorCode: EcPermissionDenied}
	if !errors.Is(e1, e2) {
		t.Fatal("expected Is to match by error code")
	}
}

func TestIs_NoMatchDifferentCode(t *testing.T) {
	e1 := &SdkError{errorCode: EcPermissionDenied}
	e2 := &SdkError{errorCode: EcUnknownError}
	if errors.Is(e1, e2) {
		t.Fatal("expected Is to not match different codes")
	}
}

func TestIs_NonSdkError(t *testing.T) {
	e := &SdkError{errorCode: EcPermissionDenied}
	other := fmt.Errorf("plain error")
	if errors.Is(e, other) {
		t.Fatal("expected no match against non-SdkError")
	}
}

func TestUnwrap(t *testing.T) {
	inner := fmt.Errorf("root cause")
	e := &SdkError{error: inner}
	if errors.Unwrap(e) != inner {
		t.Fatal("Unwrap should return inner error")
	}
}

func TestUnwrap_ChainedIs(t *testing.T) {
	inner := fmt.Errorf("root cause")
	e := &SdkError{error: inner, errorCode: EcInternalServerError}
	if !errors.Is(e, inner) {
		t.Fatal("errors.Is should find inner error through Unwrap")
	}
}
