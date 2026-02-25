package v2

type interconnectResp struct {
	ID           string `json:"uuid"`
	ProjectID    string `json:"projectId"`
	Name         string `json:"name"`
	Description  string `json:"description"`
	EnableGw2    bool   `json:"enableGw2"`
	CircuitID    int    `json:"circuitId"`
	Gw01IP       string `json:"gw01Ip"`
	Gw02IP       string `json:"gw02Ip"`
	GwVIP        string `json:"gwVip"`
	RemoteGw01IP string `json:"remoteGw01Ip"`
	RemoteGw02IP string `json:"remoteGw02Ip"`
	PackageID    string `json:"packageId"`
	Status       string `json:"status"`
	TypeID       string `json:"typeId"`
	TypeName     string `json:"typeName"`
	CreatedAt    string `json:"createdAt"`
}

func (r *interconnectResp) toEntity() *Interconnect {
	return &Interconnect{
		UUID:         r.ID,
		ProjectID:    r.ProjectID,
		Name:         r.Name,
		Description:  r.Description,
		EnableGw2:    r.EnableGw2,
		CircuitID:    r.CircuitID,
		Gw01IP:       r.Gw01IP,
		Gw02IP:       r.Gw02IP,
		GwVIP:        r.GwVIP,
		RemoteGw01IP: r.RemoteGw01IP,
		RemoteGw02IP: r.RemoteGw02IP,
		PackageID:    r.PackageID,
		Status:       r.Status,
		TypeID:       r.TypeID,
		TypeName:     r.TypeName,
		CreatedAt:    r.CreatedAt,
	}
}

type ListInterconnectsResponse struct {
	ListData  []interconnectResp `json:"listData"`
	Page      int                `json:"page"`
	PageSize  int                `json:"pageSize"`
	TotalPage int                `json:"totalPage"`
	TotalItem int                `json:"totalItem"`
}

func (r *ListInterconnectsResponse) ToEntityListInterconnects() *ListInterconnects {
	result := &ListInterconnects{
		Page:      r.Page,
		PageSize:  r.PageSize,
		TotalPage: r.TotalPage,
		TotalItem: r.TotalItem,
	}
	for i := range r.ListData {
		result.Items = append(result.Items, r.ListData[i].toEntity())
	}
	return result
}
