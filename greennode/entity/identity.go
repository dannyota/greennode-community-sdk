package entity

import (
	"github.com/dannyota/greennode-community-sdk/v2/greennode/client"
)

type AccessToken struct {
	Token     string
	ExpiresAt int64
}

func (s *AccessToken) ToSdkAuthentication() client.ISdkAuthentication {
	return new(client.SdkAuthentication).
		WithAccessToken(s.Token).
		WithExpiresAt(s.ExpiresAt)
}
