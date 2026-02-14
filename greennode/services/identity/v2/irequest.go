package v2

type IGetAccessTokenRequest interface {
	WithClientId(clientId string) IGetAccessTokenRequest
	WithClientSecret(clientSecret string) IGetAccessTokenRequest
	GetClientId() string
	GetClientSecret() string
	ToRequestBody() interface{}
}
