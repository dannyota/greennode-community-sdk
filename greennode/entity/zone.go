package entity

type Zone struct {
	Uuid          string `json:"uuid"`
	Name          string `json:"name"`
	OpenstackZone string `json:"openstackZone"`
}

type ListZones struct {
	Items []*Zone
}
