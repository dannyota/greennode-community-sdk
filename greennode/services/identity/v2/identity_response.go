package v2

import (
	"time"

	"github.com/dannyota/greennode-community-sdk/v2/greennode/entity"
)

type GetAccessTokenResponse struct {
	AccessToken      string `json:"access_token"`
	ExpiresIn        int    `json:"expires_in"`
	TokenType        string `json:"token_type"`
	RefreshExpiresIn int    `json:"refresh_expires_in"`
}

func (s *GetAccessTokenResponse) ToEntityAccessToken() *entity.AccessToken {
	return &entity.AccessToken{
		Token:     s.AccessToken,
		ExpiresAt: time.Now().Add(time.Duration(s.ExpiresIn) * time.Second).UnixNano(),
	}
}
