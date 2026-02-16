package v1

type GetPortalInfoRequest struct {
	BackEndProjectID string
}

type ListProjectsRequest struct{}
func NewGetPortalInfoRequest(backendProjectID string) *GetPortalInfoRequest {
	return &GetPortalInfoRequest{
		BackEndProjectID: backendProjectID,
	}
}

func NewListProjectsRequest() *ListProjectsRequest {
	return &ListProjectsRequest{}
}
