package v1

func NewSystemTagRequest(resourceID string, resourceType ResourceType) ICreateSystemTagRequest {
	opt := new(CreateSystemTagRequest)
	opt.ResourceID = resourceID
	opt.ResourceType = resourceType
	return opt
}
