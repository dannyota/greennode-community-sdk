package v1

import "github.com/dannyota/greennode-community-sdk/v2/greennode/entity"

type GetPortalInfoResponse struct {
	ProjectID string `json:"projectId"`
	UserID    int    `json:"userId"`
}

func (s *GetPortalInfoResponse) ToEntityPortal() *entity.Portal {
	return &entity.Portal{
		ProjectID: s.ProjectID,
		UserID:    s.UserID,
	}
}

type ListProjectsResponse struct {
	Projects []struct {
		ProjectId string `json:"projectId"`
		UserId    int    `json:"userId"`
	}
}

func (s *ListProjectsResponse) ToEntityListPortals() *entity.ListPortals {
	listPortals := entity.NewListPortals()
	for _, p := range s.Projects {
		listPortals.Items = append(listPortals.Items, &entity.Portal{
			ProjectID: p.ProjectId,
			UserID:    p.UserId,
		})
	}

	return listPortals
}
