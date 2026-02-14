package v1

func NewGetPortalInfoRequest(backendProjectId string) IGetPortalInfoRequest {
	return &GetPortalInfoRequest{
		BackEndProjectId: backendProjectId,
	}
}

func NewListProjectsRequest() IListProjectsRequest {
	return &ListProjectsRequest{}
}
