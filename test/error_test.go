package test

import (
	"fmt"
	"testing"

	sdkerror "github.com/dannyota/greennode-community-sdk/v2/greennode/sdkerror"
)

func TestDeleteListener(t *testing.T) {
	errorMsg := "Listener id lis-760f69b4-0b24-4813-9854-9e2c6a85e214 is not belong to load balancer id lb-03410465-5354-4a45-adb0-32c9a75d2a06"
	errResp := &sdkerror.NormalErrorResponse{
		Message: errorMsg,
	}

	serr := sdkerror.ErrorHandler(fmt.Errorf("haha"), sdkerror.WithErrorListenerNotFound(errResp))
	t.Log(serr)
}
