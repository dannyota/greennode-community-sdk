package entity

import (
	lsvcClient "github.com/dannyota/greennode-community-sdk/v2/greennode/client"
)

type AccessToken struct {
	Token     string
	ExpiresAt int64
}

func (s *AccessToken) ToSdkAuthentication() lsvcClient.ISdkAuthentication {
	return new(lsvcClient.SdkAuthentication).
		WithAccessToken(s.Token).
		WithExpiresAt(s.ExpiresAt)
}
