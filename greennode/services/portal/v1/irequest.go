package v1

type IGetPortalInfoRequest interface {
	GetBackEndProjectId() string
}

type IListProjectsRequest interface {
	AddUserAgent(agent ...string) IListProjectsRequest
	ParseUserAgent() string
}
