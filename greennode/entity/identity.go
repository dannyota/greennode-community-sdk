package entity

import (
	"github.com/dannyota/greennode-community-sdk/v2/greennode/client"
)

type AccessToken struct {
	Token     string
	ExpiresAt int64
}

func (t AccessToken) ToSdkAuthentication() client.SdkAuthentication {
	return client.NewSdkAuthentication().
		WithAccessToken(t.Token).
		WithExpiresAt(t.ExpiresAt)
}
