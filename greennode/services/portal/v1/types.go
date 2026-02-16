package v1

type Portal struct {
	ProjectID string `json:"projectId"`
	UserID    int    `json:"userId"`
}

type ListPortals struct {
	Items []*Portal
}

func NewListPortals() *ListPortals {
	return &ListPortals{
		Items: make([]*Portal, 0),
	}
}

func (l ListPortals) At(index int) *Portal {
	if index < 0 || index >= len(l.Items) {
		return nil
	}

	return l.Items[index]
}

type Zone struct {
	Uuid          string `json:"uuid"`
	Name          string `json:"name"`
	OpenstackZone string `json:"openstackZone"`
}

type ListZones struct {
	Items []*Zone
}
