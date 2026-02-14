package v1

type IGetPortalInfoRequest interface {
	GetBackEndProjectID() string
}

type IListProjectsRequest interface {
	AddUserAgent(agent ...string) IListProjectsRequest
	ParseUserAgent() string
}
