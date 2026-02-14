package v2

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

func (s *GetAccessTokenRequest) WithClientID(clientID string) IGetAccessTokenRequest {
	s.ClientID = clientID
	return s
}

func (s *GetAccessTokenRequest) WithClientSecret(clientSecret string) IGetAccessTokenRequest {
	s.ClientSecret = clientSecret
	return s
}

func (s *GetAccessTokenRequest) GetClientID() string {
	return s.ClientID
}

func (s *GetAccessTokenRequest) GetClientSecret() string {
	return s.ClientSecret
}

func (s *GetAccessTokenRequest) ToRequestBody() interface{} {
	return s
}
