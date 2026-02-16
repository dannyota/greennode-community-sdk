package v1

type ListZoneResponse struct {
	Data []Zone `json:"data"`
}

func (r *ListZoneResponse) ToEntityListZones() *ListZones {
	listZones := &ListZones{
		Items: make([]*Zone, 0),
	}
	for i := range r.Data {
		listZones.Items = append(listZones.Items, &r.Data[i])
	}

	return listZones
}
