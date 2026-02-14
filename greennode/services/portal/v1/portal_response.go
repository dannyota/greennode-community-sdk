package v1

import "github.com/dannyota/greennode-community-sdk/v2/greennode/entity"

type GetPortalInfoResponse struct {
	ProjectID string `json:"projectId"`
	UserID    int    `json:"userId"`
}

func (r *GetPortalInfoResponse) ToEntityPortal() *entity.Portal {
	return &entity.Portal{
		ProjectID: r.ProjectID,
		UserID:    r.UserID,
	}
}

type ListProjectsResponse struct {
	Projects []struct {
		ProjectID string `json:"projectId"`
		UserID    int    `json:"userId"`
	}
}

func (r *ListProjectsResponse) ToEntityListPortals() *entity.ListPortals {
	listPortals := entity.NewListPortals()
	for _, p := range r.Projects {
		listPortals.Items = append(listPortals.Items, &entity.Portal{
			ProjectID: p.ProjectID,
			UserID:    p.UserID,
		})
	}

	return listPortals
}
