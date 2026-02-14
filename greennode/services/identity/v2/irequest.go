package v2

type IGetAccessTokenRequest interface {
	WithClientID(clientID string) IGetAccessTokenRequest
	WithClientSecret(clientSecret string) IGetAccessTokenRequest
	GetClientID() string
	GetClientSecret() string
	ToRequestBody() any
}
