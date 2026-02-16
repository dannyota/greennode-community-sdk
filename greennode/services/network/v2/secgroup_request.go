package v2

import "github.com/dannyota/greennode-community-sdk/v2/greennode/services/common"

type ListSecgroupRequest struct{}

type DeleteSecgroupByIDRequest struct { //__________________________________________________________________________________
	common.SecgroupCommon
}

type CreateSecgroupRequest struct { //__________________________________________________________________________________
	Name        string `json:"name"`
	Description string `json:"description"`
}

func (r *CreateSecgroupRequest) GetSecgroupName() string {
	return r.Name
}

type GetSecgroupByIDRequest struct { //_________________________________________________________________________________
	common.SecgroupCommon
}

func NewCreateSecgroupRequest(name, description string) *CreateSecgroupRequest {
	return &CreateSecgroupRequest{
		Name:        name,
		Description: description,
	}
}

func NewDeleteSecgroupByIDRequest(secgroupID string) *DeleteSecgroupByIDRequest {
	opts := new(DeleteSecgroupByIDRequest)
	opts.SecgroupID = secgroupID
	return opts
}

func NewGetSecgroupByIDRequest(secgroupID string) *GetSecgroupByIDRequest {
	opt := new(GetSecgroupByIDRequest)
	opt.SecgroupID = secgroupID
	return opt
}

func NewListSecgroupRequest() *ListSecgroupRequest {
	return &ListSecgroupRequest{}
}
