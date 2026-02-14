package sdkerror

import "fmt"

type (
	IAMErrorResponse struct {
		Errors []struct {
			Code    string `json:"code,omitempty"`
			Message string `json:"message,omitempty"`
		} `json:"errors,omitempty"`
	}

	NormalErrorResponse struct {
		Message string `json:"message,omitempty"`
	}

	NetworkGatewayErrorResponse struct {
		Message   string `json:"message,omitempty"`
		Code      int    `json:"code,omitempty"`
		ErrorCode string `json:"errorCode,omitempty"`
		Success   bool   `json:"success,omitempty"`
	}

	// {"code":"global_load_balancer_not_found","error":"Global load balancer is not found"}
	GlobalLoadBalancerErrorResponse struct {
		Code  string `json:"code,omitempty"`
		Error string `json:"error,omitempty"`
	}
)

const (
	NormalErrorType = iota
	IAMErrorType
	NetworkGatewayErrorType
	GlobalLoadBalancerErrorType
)

func NewErrorResponse(typeVal int) ErrorResponse {
	switch typeVal {
	case IAMErrorType:
		return new(IAMErrorResponse)
	case NetworkGatewayErrorType:
		return new(NetworkGatewayErrorResponse)
	case GlobalLoadBalancerErrorType:
		return new(GlobalLoadBalancerErrorResponse)
	default:
		return new(NormalErrorResponse)
	}
}

func (r *IAMErrorResponse) GetMessage() string {
	if len(r.Errors) < 1 {
		return ""
	}

	return r.Errors[0].Message
}

func (r *IAMErrorResponse) Err() error {
	if len(r.Errors) < 1 {
		return nil
	}

	return fmt.Errorf("%s", r.Errors[0].Code)
}

func (r *NormalErrorResponse) GetMessage() string {
	return r.Message
}

func (r *NormalErrorResponse) Err() error {
	return fmt.Errorf("%s", r.Message)
}

func (r *NetworkGatewayErrorResponse) GetMessage() string {
	return fmt.Sprintf("%s/%d/%s", r.ErrorCode, r.Code, r.Message)
}

func (r *NetworkGatewayErrorResponse) Err() error {
	return fmt.Errorf("%s", r.ErrorCode)
}

func (r *GlobalLoadBalancerErrorResponse) GetMessage() string {
	return r.Error
}

func (r *GlobalLoadBalancerErrorResponse) Err() error {
	return fmt.Errorf("%s", r.Code)
}
