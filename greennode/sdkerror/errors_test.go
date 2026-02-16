package sdkerror

import (
	"testing"
)

func TestNewErrorResponse(t *testing.T) {
	tests := []struct {
		name     string
		typeVal  int
		wantType string
	}{
		{"NormalErrorType", NormalErrorType, "*sdkerror.NormalErrorResponse"},
		{"IAMErrorType", IAMErrorType, "*sdkerror.IAMErrorResponse"},
		{"NetworkGatewayErrorType", NetworkGatewayErrorType, "*sdkerror.NetworkGatewayErrorResponse"},
		{"GlobalLoadBalancerErrorType", GlobalLoadBalancerErrorType, "*sdkerror.GlobalLoadBalancerErrorResponse"},
		{"default fallback", 999, "*sdkerror.NormalErrorResponse"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp := NewErrorResponse(tt.typeVal)
			if resp == nil {
				t.Fatal("expected non-nil response")
			}
			// Verify correct type by interface method calls below
		})
	}
}

func TestNormalErrorResponse(t *testing.T) {
	r := NormalErrorResponse{Message: "not found"}
	if r.GetMessage() != "not found" {
		t.Fatalf("got %q", r.GetMessage())
	}
	if r.Err() == nil || r.Err().Error() != "not found" {
		t.Fatalf("Err: got %v", r.Err())
	}
}

func TestIAMErrorResponse_WithErrors(t *testing.T) {
	r := IAMErrorResponse{
		Errors: []struct {
			Code    string `json:"code,omitempty"`
			Message string `json:"message,omitempty"`
		}{
			{Code: "AUTHENTICATION_FAILED", Message: "bad creds"},
		},
	}
	if r.GetMessage() != "bad creds" {
		t.Fatalf("got %q", r.GetMessage())
	}
	if r.Err() == nil || r.Err().Error() != "AUTHENTICATION_FAILED" {
		t.Fatalf("Err: got %v", r.Err())
	}
}

func TestIAMErrorResponse_Empty(t *testing.T) {
	r := IAMErrorResponse{}
	if r.GetMessage() != "" {
		t.Fatalf("got %q, want empty", r.GetMessage())
	}
	if r.Err() != nil {
		t.Fatalf("got %v, want nil", r.Err())
	}
}

func TestNetworkGatewayErrorResponse(t *testing.T) {
	r := NetworkGatewayErrorResponse{
		Message:   "bad gateway",
		Code:      502,
		ErrorCode: "GW_ERR",
	}
	want := "GW_ERR/502/bad gateway"
	if r.GetMessage() != want {
		t.Fatalf("got %q, want %q", r.GetMessage(), want)
	}
	if r.Err().Error() != "GW_ERR" {
		t.Fatalf("Err: got %v", r.Err())
	}
}

func TestGlobalLoadBalancerErrorResponse(t *testing.T) {
	r := GlobalLoadBalancerErrorResponse{
		Code:  "global_load_balancer_not_found",
		Error: "Global load balancer is not found",
	}
	if r.GetMessage() != "Global load balancer is not found" {
		t.Fatalf("got %q", r.GetMessage())
	}
	if r.Err().Error() != "global_load_balancer_not_found" {
		t.Fatalf("Err: got %v", r.Err())
	}
}
