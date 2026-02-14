package v2

type GetAccessTokenRequest struct {
	ClientId     string
	ClientSecret string

	GrantType string `json:"grant_type"`
}

func NewGetAccessTokenRequest(clientId, clientSecret string) IGetAccessTokenRequest {
	return &GetAccessTokenRequest{
		ClientId:     clientId,
		ClientSecret: clientSecret,
		GrantType:    "client_credentials",
	}
}

func (s *GetAccessTokenRequest) WithClientId(clientId string) IGetAccessTokenRequest {
	s.ClientId = clientId
	return s
}

func (s *GetAccessTokenRequest) WithClientSecret(clientSecret string) IGetAccessTokenRequest {
	s.ClientSecret = clientSecret
	return s
}

func (s *GetAccessTokenRequest) GetClientId() string {
	return s.ClientId
}

func (s *GetAccessTokenRequest) GetClientSecret() string {
	return s.ClientSecret
}

func (s *GetAccessTokenRequest) ToRequestBody() interface{} {
	return s
}
