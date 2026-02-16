package v2

type ListSecgroupRequest struct{}

type DeleteSecgroupByIDRequest struct { //__________________________________________________________________________________
	SecgroupID string
}

type CreateSecgroupRequest struct { //__________________________________________________________________________________
	Name        string `json:"name"`
	Description string `json:"description"`
}

func (r *CreateSecgroupRequest) GetSecgroupName() string {
	return r.Name
}

type GetSecgroupByIDRequest struct { //_________________________________________________________________________________
	SecgroupID string
}

func NewCreateSecgroupRequest(name, description string) *CreateSecgroupRequest {
	return &CreateSecgroupRequest{
		Name:        name,
		Description: description,
	}
}

func NewDeleteSecgroupByIDRequest(secgroupID string) *DeleteSecgroupByIDRequest {
	return &DeleteSecgroupByIDRequest{
		SecgroupID: secgroupID,
	}
}

func NewGetSecgroupByIDRequest(secgroupID string) *GetSecgroupByIDRequest {
	return &GetSecgroupByIDRequest{
		SecgroupID: secgroupID,
	}
}

func NewListSecgroupRequest() *ListSecgroupRequest {
	return &ListSecgroupRequest{}
}
