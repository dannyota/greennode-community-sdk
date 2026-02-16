package common

import (
	"encoding/json"
)

// StructToMap converts a struct to map[string]any using its JSON tags.
func StructToMap(v any) map[string]any {
	b, _ := json.Marshal(v)
	var m map[string]any
	_ = json.Unmarshal(b, &m)
	return m
}

func Ptr[T any](v T) *T { return &v }

type Project struct {
	ID string
}

func (pr *Project) GetProjectID() string {
	return pr.ID
}


type Paging struct {
	Page int
	Size int
}

type Zone string

const (
	HCM_03_1A_ZONE     Zone = "HCM03-1A"
	HCM_03_1B_ZONE     Zone = "HCM03-1B"
	HCM_03_1C_ZONE     Zone = "HCM03-1C"
	HCM_03_BKK_01_ZONE Zone = "HCM03-BKK-01"
	HAN_01_1A_ZONE     Zone = "HAN01-1A"
)


type Tag struct {
	Key      string `json:"key"`
	Value    string `json:"value"`
	IsEdited bool   `json:"isEdited,omitempty"`
}

type PortalUser struct {
	ID string
}

func (pu *PortalUser) GetPortalUserID() string {
	return pu.ID
}


func (pu *PortalUser) GetMapHeaders() map[string]string {
	return map[string]string{
		"portal-user-id": pu.ID,
	}
}
