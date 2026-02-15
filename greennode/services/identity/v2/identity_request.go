package v2

type GetAccessTokenRequest struct {
	ClientID     string
	ClientSecret string

	GrantType string `json:"grant_type"`
}

func NewGetAccessTokenRequest(clientID, clientSecret string) *GetAccessTokenRequest {
	return &GetAccessTokenRequest{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		GrantType:    "client_credentials",
	}
}

func (r *GetAccessTokenRequest) WithClientID(clientID string) *GetAccessTokenRequest {
	r.ClientID = clientID
	return r
}

func (r *GetAccessTokenRequest) WithClientSecret(clientSecret string) *GetAccessTokenRequest {
	r.ClientSecret = clientSecret
	return r
}

func (r *GetAccessTokenRequest) GetClientID() string {
	return r.ClientID
}

func (r *GetAccessTokenRequest) GetClientSecret() string {
	return r.ClientSecret
}

