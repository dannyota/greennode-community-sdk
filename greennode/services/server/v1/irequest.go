package v1

type ICreateSystemTagRequest interface {
	ToRequestBody() interface{}
	AddUserAgent(agent ...string) ICreateSystemTagRequest
	ToMap() map[string]interface{}
	AddTag(key, value string) ICreateSystemTagRequest
	ParseUserAgent() string
	GetResourceId() string
	GetResourceType() ResourceType
}
