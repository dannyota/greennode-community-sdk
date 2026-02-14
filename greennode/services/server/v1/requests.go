package v1

func NewSystemTagRequest(resourceID string, resourceType ResourceType) *CreateSystemTagRequest {
	opt := new(CreateSystemTagRequest)
	opt.ResourceID = resourceID
	opt.ResourceType = resourceType
	return opt
}
