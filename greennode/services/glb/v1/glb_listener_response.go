package v1

import "github.com/dannyota/greennode-community-sdk/v2/greennode/entity"

// ListGlobalListenersResponse wraps a bare JSON array of listeners into
// the entity list type.
type ListGlobalListenersResponse []entity.GlobalListener

func (r *ListGlobalListenersResponse) ToEntityListGlobalListeners() *entity.ListGlobalListeners {
	items := make([]*entity.GlobalListener, len(*r))
	for i := range *r {
		items[i] = &(*r)[i]
	}
	return &entity.ListGlobalListeners{Items: items}
}
