package v1

type ICreateSystemTagRequest interface {
	ToRequestBody() any
	AddUserAgent(agent ...string) ICreateSystemTagRequest
	ToMap() map[string]any
	AddTag(key, value string) ICreateSystemTagRequest
	ParseUserAgent() string
	GetResourceID() string
	GetResourceType() ResourceType
}
