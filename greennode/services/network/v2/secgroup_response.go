package v2

import "github.com/dannyota/greennode-community-sdk/v2/greennode/entity"

type CreateSecgroupResponse struct { //_________________________________________________________________________________
	Data struct {
		ID           int     `json:"id"`
		UUID         string  `json:"uuid"`
		CreatedAt    string  `json:"createdAt"`
		DeletedAt    *string `json:"deletedAt,omitempty"`
		Status       string  `json:"status"`
		SecgroupID   int     `json:"secgroupId"`
		SecgroupName string  `json:"secgroupName"`
		ProjectUuid  string  `json:"projectUuid"`
		Description  string  `json:"description"`
		UpdatedAt    *string `json:"updatedAt,omitempty"`
		IsSystem     bool    `json:"isSystem"`
		Type         string  `json:"type"`
	} `json:"data"`
}

func (r *CreateSecgroupResponse) ToEntitySecgroup() *entity.Secgroup {
	return &entity.Secgroup{
		ID:          r.Data.UUID,
		Name:        r.Data.SecgroupName,
		Description: r.Data.Description,
		Status:      r.Data.Status,
	}
}

type GetSecgroupByIDResponse struct { //________________________________________________________________________________
	Data struct {
		ID          string `json:"id"`
		Name        string `json:"name"`
		Description string `json:"description"`
		Status      string `json:"status"`
		CreatedAt   string `json:"createdAt"`
		IsSystem    bool   `json:"isSystem"`
	} `json:"data"`
}

func (r *GetSecgroupByIDResponse) ToEntitySecgroup() *entity.Secgroup {
	return &entity.Secgroup{
		ID:          r.Data.ID,
		Name:        r.Data.Name,
		Description: r.Data.Description,
		Status:      r.Data.Status,
	}
}

type ListSecgroupResponse struct { //_________________________________________________________________________________
	ListData []struct {
		ID          string `json:"id"`
		Name        string `json:"name"`
		Description string `json:"description"`
		Status      string `json:"status"`
		CreatedAt   string `json:"createdAt"`
		IsSystem    bool   `json:"isSystem"`
	} `json:"listData"`
	Page      int `json:"page"`
	PageSize  int `json:"pageSize"`
	TotalPage int `json:"totalPage"`
	TotalItem int `json:"totalItem"`
}

func (r *ListSecgroupResponse) ToListEntitySecgroups() *entity.ListSecgroups {
	items := make([]*entity.Secgroup, 0, len(r.ListData))
	for _, item := range r.ListData {
		items = append(items, &entity.Secgroup{
			ID:          item.ID,
			Name:        item.Name,
			Description: item.Description,
			Status:      item.Status,
		})
	}
	return &entity.ListSecgroups{Items: items}
}
