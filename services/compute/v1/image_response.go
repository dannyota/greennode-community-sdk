package v1

type listImagesResponse struct {
	Images []OSImage `json:"images"`
}

func (r *listImagesResponse) ToEntityListOSImages() *ListOSImages {
	list := &ListOSImages{
		Items: make([]*OSImage, 0, len(r.Images)),
	}
	for i := range r.Images {
		list.Items = append(list.Items, &r.Images[i])
	}
	return list
}
