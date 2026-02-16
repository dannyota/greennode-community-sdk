package entity

import (
	"github.com/dannyota/greennode-community-sdk/v2/greennode/client"
)

type AccessToken struct {
	Token     string `json:"access_token"`
	ExpiresAt int64  `json:"expiresAt"`
}

func (t AccessToken) ToSdkAuthentication() *client.SdkAuthentication {
	return client.NewSdkAuthentication().
		WithAccessToken(t.Token).
		WithExpiresAt(t.ExpiresAt)
}
