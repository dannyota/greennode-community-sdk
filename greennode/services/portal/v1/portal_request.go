package v1

type GetPortalInfoRequest struct {
	BackEndProjectID string
}

func (r *GetPortalInfoRequest) GetBackEndProjectID() string {
	return r.BackEndProjectID
}

type ListProjectsRequest struct{}
