package v1

import "github.com/dannyota/greennode-community-sdk/v2/greennode/entity"

type ListProjectsResponse struct {
	Projects []entity.Portal `json:"projects"`
}

func (r *ListProjectsResponse) ToEntityListPortals() *entity.ListPortals {
	listPortals := entity.NewListPortals()
	for i := range r.Projects {
		listPortals.Items = append(listPortals.Items, &r.Projects[i])
	}

	return listPortals
}
