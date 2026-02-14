package common

import (
	"slices"
	"strings"
)

func Ptr[T any](v T) *T { return &v }

type Project struct {
	Id string
}

func (s *Project) GetProjectId() string {
	return s.Id
}

func (s *Project) SetProjectId(id string) {
	s.Id = id
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

func (s *Paging) GetPage() int {
	return s.Page
}

func (s *Paging) GetSize() int {
	return s.Size
}

func (s *Paging) SetPage(page int) *Paging {
	s.Page = page
	return s
}

func (s *Paging) SetSize(size int) *Paging {
	s.Size = size
	return s
}

type Tag struct {
	Key      string `json:"key"`
	Value    string `json:"value"`
	IsEdited bool   `json:"isEdited,omitempty"`
}

type UserAgent struct {
	Agent []string
}

func (s *UserAgent) ParseUserAgent() string {
	// Parse the array into string
	return strings.Join(s.Agent, "; ")
}

func (s *UserAgent) AddUserAgent(pagent ...string) *UserAgent {
	for _, agent := range pagent {
		if !slices.Contains(s.Agent, agent) {
			s.Agent = append(s.Agent, agent)
		}
	}
	return s
}

type PortalUser struct {
	Id string
}

func (s *PortalUser) GetPortalUserId() string {
	return s.Id
}

func (s *PortalUser) SetPortalUserId(id string) {
	s.Id = id
}

func (s *PortalUser) GetMapHeaders() map[string]string {
	return map[string]string{
		"portal-user-id": s.Id,
	}
}
