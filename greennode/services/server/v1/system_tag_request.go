package v1

const (
	Volume       ResourceType = "VOLUME"
	Server       ResourceType = "SERVER"
	LoadBalancer ResourceType = "LOAD-BALANCER"
)

type ResourceType string

type SystemTagInput struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type CreateSystemTagRequest struct {
	ResourceID   string           `json:"resourceId"`
	ResourceType ResourceType     `json:"resourceType"`
	Tags         []SystemTagInput `json:"tagRequestList"`
}

func NewSystemTagRequest(resourceID string, resourceType ResourceType) *CreateSystemTagRequest {
	return &CreateSystemTagRequest{
		ResourceID:   resourceID,
		ResourceType: resourceType,
	}
}
