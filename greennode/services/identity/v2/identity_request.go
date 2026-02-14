package v2

type IGetAccessTokenRequest interface {
	WithClientID(clientID string) IGetAccessTokenRequest
	WithClientSecret(clientSecret string) IGetAccessTokenRequest
	GetClientID() string
	GetClientSecret() string
	ToRequestBody() any
}

type GetAccessTokenRequest struct {
	ClientID     string
	ClientSecret string

	GrantType string `json:"grant_type"`
}

func NewGetAccessTokenRequest(clientID, clientSecret string) IGetAccessTokenRequest {
	return &GetAccessTokenRequest{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		GrantType:    "client_credentials",
	}
}

func (r *GetAccessTokenRequest) WithClientID(clientID string) IGetAccessTokenRequest {
	r.ClientID = clientID
	return r
}

func (r *GetAccessTokenRequest) WithClientSecret(clientSecret string) IGetAccessTokenRequest {
	r.ClientSecret = clientSecret
	return r
}

func (r *GetAccessTokenRequest) GetClientID() string {
	return r.ClientID
}

func (r *GetAccessTokenRequest) GetClientSecret() string {
	return r.ClientSecret
}

func (r *GetAccessTokenRequest) ToRequestBody() any {
	return r
}
