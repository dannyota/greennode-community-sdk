package v1

func NewGetPortalInfoRequest(backendProjectID string) IGetPortalInfoRequest {
	return &GetPortalInfoRequest{
		BackEndProjectID: backendProjectID,
	}
}

func NewListProjectsRequest() IListProjectsRequest {
	return &ListProjectsRequest{}
}
