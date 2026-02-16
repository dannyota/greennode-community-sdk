package v1

type ListProjectsResponse struct {
	Projects []Portal `json:"projects"`
}

func (r *ListProjectsResponse) ToEntityListPortals() *ListPortals {
	listPortals := NewListPortals()
	for i := range r.Projects {
		listPortals.Items = append(listPortals.Items, &r.Projects[i])
	}

	return listPortals
}
