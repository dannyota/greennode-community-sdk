package sdkerror

import (
	"fmt"
	"testing"
)

// ---------------------------------------------------------------------------
// SdkErrorHandler
// ---------------------------------------------------------------------------

func TestSdkErrorHandler_NilNil(t *testing.T) {
	result := SdkErrorHandler(nil, nil)
	if result != nil {
		t.Fatalf("expected nil, got %v", result)
	}
}

func TestSdkErrorHandler_MatchingMessage(t *testing.T) {
	resp := &NormalErrorResponse{Message: "cannot get security group with id abc"}
	result := SdkErrorHandler(nil, resp, EcVServerSecgroupNotFound)
	if result == nil {
		t.Fatal("expected non-nil result")
	}
	if result.ErrorCode() != EcVServerSecgroupNotFound {
		t.Fatalf("got %q, want %q", result.ErrorCode(), EcVServerSecgroupNotFound)
	}
}

func TestSdkErrorHandler_NoMatch(t *testing.T) {
	resp := &NormalErrorResponse{Message: "something completely different"}
	result := SdkErrorHandler(nil, resp, EcVServerSecgroupNotFound)
	if result == nil {
		t.Fatal("expected non-nil result")
	}
	if result.ErrorCode() != EcUnknownError {
		t.Fatalf("got %q, want %q", result.ErrorCode(), EcUnknownError)
	}
}

func TestSdkErrorHandler_AlreadyClassified(t *testing.T) {
	existing := &SdkError{errorCode: EcPermissionDenied, message: "denied"}
	result := SdkErrorHandler(existing, nil)
	if result.ErrorCode() != EcPermissionDenied {
		t.Fatalf("got %q, want %q", result.ErrorCode(), EcPermissionDenied)
	}
}

func TestSdkErrorHandler_FirstMatchWins(t *testing.T) {
	resp := &NormalErrorResponse{Message: "exceeded secgroup quota"}
	result := SdkErrorHandler(nil, resp,
		EcVServerSecgroupNotFound, // won't match
		EcVServerSecgroupExceedQuota, // will match
		EcVServerSecgroupInUse,       // won't reach
	)
	if result.ErrorCode() != EcVServerSecgroupExceedQuota {
		t.Fatalf("got %q, want %q", result.ErrorCode(), EcVServerSecgroupExceedQuota)
	}
	if !result.IsCategory(ErrCatQuota) {
		t.Fatal("expected quota category")
	}
}

// ---------------------------------------------------------------------------
// ErrorHandler
// ---------------------------------------------------------------------------

func TestErrorHandler_Nil(t *testing.T) {
	result := ErrorHandler(nil)
	if result == nil {
		t.Fatal("expected non-nil for nil error")
	}
	if result.ErrorCode() != EcUnknownError {
		t.Fatalf("got %q", result.ErrorCode())
	}
}

func TestErrorHandler_ServiceMaintenance(t *testing.T) {
	err := fmt.Errorf("This service is in maintenance right now")
	result := ErrorHandler(err)
	if result.ErrorCode() != EcServiceMaintenance {
		t.Fatalf("got %q, want %q", result.ErrorCode(), EcServiceMaintenance)
	}
}

func TestErrorHandler_RegularError(t *testing.T) {
	err := fmt.Errorf("some random error")
	result := ErrorHandler(err)
	if result.ErrorCode() != EcUnknownError {
		t.Fatalf("got %q, want %q", result.ErrorCode(), EcUnknownError)
	}
}

// ---------------------------------------------------------------------------
// catchAll classifier (identity's EcUnknownAuthFailure)
// ---------------------------------------------------------------------------

func TestSdkErrorHandler_CatchAll(t *testing.T) {
	resp := &IAMErrorResponse{
		Errors: []struct {
			Code    string `json:"code,omitempty"`
			Message string `json:"message,omitempty"`
		}{
			{Code: "SOME_UNKNOWN_IAM_ERROR", Message: "something new"},
		},
	}
	result := SdkErrorHandler(nil, resp,
		EcAuthenticationFailed, // won't match
		EcTooManyFailedLogins,  // won't match
		EcUnknownAuthFailure,   // catchAll: matches when still EcUnknownError
	)
	if result.ErrorCode() != EcUnknownAuthFailure {
		t.Fatalf("got %q, want %q", result.ErrorCode(), EcUnknownAuthFailure)
	}
}

// ---------------------------------------------------------------------------
// Domain classifier registrations — representative samples
// ---------------------------------------------------------------------------

func TestClassifier_VolumeNotFound_Regex(t *testing.T) {
	resp := &NormalErrorResponse{Message: "volume with id abc-123 is not found"}
	result := SdkErrorHandler(nil, resp, EcVServerVolumeNotFound)
	if result.ErrorCode() != EcVServerVolumeNotFound {
		t.Fatalf("got %q", result.ErrorCode())
	}
}

func TestClassifier_VolumeNotFound_Contains(t *testing.T) {
	resp := &NormalErrorResponse{Message: "cannot get volume with id abc-123"}
	result := SdkErrorHandler(nil, resp, EcVServerVolumeNotFound)
	if result.ErrorCode() != EcVServerVolumeNotFound {
		t.Fatalf("got %q", result.ErrorCode())
	}
}

func TestClassifier_LoadBalancerNotReady_Regex(t *testing.T) {
	resp := &NormalErrorResponse{Message: "load balancer id lb-xyz is not ready"}
	result := SdkErrorHandler(nil, resp, EcVLBLoadBalancerNotReady)
	if result.ErrorCode() != EcVLBLoadBalancerNotReady {
		t.Fatalf("got %q", result.ErrorCode())
	}
}

func TestClassifier_LoadBalancerIsDeleting(t *testing.T) {
	resp := &NormalErrorResponse{Message: "load balancer id lb-xyz is deleting"}
	result := SdkErrorHandler(nil, resp, EcVLBLoadBalancerIsDeleting)
	if result.ErrorCode() != EcVLBLoadBalancerIsDeleting {
		t.Fatalf("got %q", result.ErrorCode())
	}
}

func TestClassifier_AuthenticationFailed(t *testing.T) {
	resp := &IAMErrorResponse{
		Errors: []struct {
			Code    string `json:"code,omitempty"`
			Message string `json:"message,omitempty"`
		}{
			{Code: "AUTHENTICATION_FAILED", Message: "bad key pair"},
		},
	}
	result := SdkErrorHandler(nil, resp, EcAuthenticationFailed)
	if result.ErrorCode() != EcAuthenticationFailed {
		t.Fatalf("got %q", result.ErrorCode())
	}
	if !result.IsCategory(ErrCatIAM) {
		t.Fatal("expected IAM category")
	}
	// Verify msgFmt is applied
	if result.GetMessage() == "bad key pair" {
		t.Fatal("expected formatted message, got raw")
	}
}

func TestClassifier_PurchaseIssue(t *testing.T) {
	resp := &NormalErrorResponse{
		Message: "you do not have sufficient credits to complete the purchase",
	}
	result := SdkErrorHandler(nil, resp, EcPurchaseIssue)
	if result.ErrorCode() != EcPurchaseIssue {
		t.Fatalf("got %q", result.ErrorCode())
	}
	if !result.IsCategory(ErrCatPurchase) {
		t.Fatal("expected purchase category")
	}
}

func TestClassifier_GlobalLoadBalancerNotFound(t *testing.T) {
	resp := &GlobalLoadBalancerErrorResponse{
		Code:  "global_load_balancer_not_found",
		Error: "Global load balancer is not found",
	}
	result := SdkErrorHandler(nil, resp, EcGlobalLoadBalancerNotFound)
	if result.ErrorCode() != EcGlobalLoadBalancerNotFound {
		t.Fatalf("got %q", result.ErrorCode())
	}
}

// ---------------------------------------------------------------------------
// Constructor functions
// ---------------------------------------------------------------------------

func TestNewInternalServerError(t *testing.T) {
	e := NewInternalServerError()
	if e.ErrorCode() != EcInternalServerError {
		t.Fatalf("got %q", e.ErrorCode())
	}
	if e.Err() == nil {
		t.Fatal("expected inner error")
	}
}

func TestNewServiceMaintenance(t *testing.T) {
	e := NewServiceMaintenance()
	if e.ErrorCode() != EcServiceMaintenance {
		t.Fatalf("got %q", e.ErrorCode())
	}
}

func TestNewPermissionDenied(t *testing.T) {
	e := NewPermissionDenied()
	if e.ErrorCode() != EcPermissionDenied {
		t.Fatalf("got %q", e.ErrorCode())
	}
}

func TestNewReauthFuncNotSet(t *testing.T) {
	e := NewReauthFuncNotSet()
	if e.ErrorCode() != EcReauthFuncNotSet {
		t.Fatalf("got %q", e.ErrorCode())
	}
}

func TestNewQuotaNotFound(t *testing.T) {
	e := NewQuotaNotFound()
	if e.ErrorCode() != EcVServerQuotaNotFound {
		t.Fatalf("got %q", e.ErrorCode())
	}
}

// ---------------------------------------------------------------------------
// Classifier helpers (unit-level)
// ---------------------------------------------------------------------------

func TestContainsAny(t *testing.T) {
	fn := containsAny("foo", "bar")
	if !fn("something foo here", nil) {
		t.Fatal("expected match")
	}
	if fn("something else", nil) {
		t.Fatal("expected no match")
	}
}

func TestMatchErrCode(t *testing.T) {
	fn := matchErrCode("AUTH_FAILED")
	resp := &IAMErrorResponse{
		Errors: []struct {
			Code    string `json:"code,omitempty"`
			Message string `json:"message,omitempty"`
		}{
			{Code: "auth_failed", Message: "test"},
		},
	}
	if !fn("", resp) {
		t.Fatal("expected match (case insensitive)")
	}

	resp2 := &NormalErrorResponse{Message: "test"}
	if fn("", resp2) {
		t.Fatal("expected no match for NormalErrorResponse Err()")
	}
}

func TestMatchAnyOf(t *testing.T) {
	fn := matchAnyOf(
		containsAny("nope"),
		containsAny("found"),
	)
	if !fn("we found it", nil) {
		t.Fatal("expected match on second func")
	}
	if fn("nothing here", nil) {
		t.Fatal("expected no match")
	}
}
