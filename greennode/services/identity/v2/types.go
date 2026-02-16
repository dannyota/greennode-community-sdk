package v2

type AccessToken struct {
	Token     string `json:"access_token"`
	ExpiresAt int64  `json:"expiresAt"`
}
