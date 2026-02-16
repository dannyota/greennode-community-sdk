package client

import "time"

type SdkAuthentication struct {
	accessToken string
	expiresAt   int64
}

func NewSdkAuthentication() *SdkAuthentication {
	return &SdkAuthentication{}
}

func (a *SdkAuthentication) WithAccessToken(accessToken string) *SdkAuthentication {
	a.accessToken = accessToken
	return a
}

func (a *SdkAuthentication) WithExpiresAt(expiresAt int64) *SdkAuthentication {
	a.expiresAt = expiresAt
	return a
}

func (a *SdkAuthentication) NeedReauth() bool {
	if a.accessToken == "" {
		return true
	}

	ea := time.Unix(0, a.expiresAt)
	return time.Until(ea) < 5*time.Minute
}

func (a *SdkAuthentication) UpdateAuth(auth *SdkAuthentication) {
	a.accessToken = auth.AccessToken()
	a.expiresAt = auth.ExpiresAt()
}

func (a *SdkAuthentication) AccessToken() string {
	return a.accessToken
}

func (a *SdkAuthentication) ExpiresAt() int64 {
	return a.expiresAt
}
