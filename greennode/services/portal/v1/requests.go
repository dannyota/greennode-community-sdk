package v1

func NewGetPortalInfoRequest(backendProjectID string) *GetPortalInfoRequest {
	return &GetPortalInfoRequest{
		BackEndProjectID: backendProjectID,
	}
}

func NewListProjectsRequest() *ListProjectsRequest {
	return &ListProjectsRequest{}
}
