package v1

// ListGlobalListenersResponse wraps a bare JSON array of listeners into
// the entity list type.
type ListGlobalListenersResponse []GlobalListener

func (r *ListGlobalListenersResponse) ToEntityListGlobalListeners() *ListGlobalListeners {
	items := make([]*GlobalListener, len(*r))
	for i := range *r {
		items[i] = &(*r)[i]
	}
	return &ListGlobalListeners{Items: items}
}
