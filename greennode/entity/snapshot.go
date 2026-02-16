package entity

type Snapshot struct {
	ID         string `json:"id"`
	CreatedAt  string `json:"createdAt"`
	VolumeID   string `json:"volumeId"`
	Size       int64  `json:"size"`
	VolumeSize int64  `json:"volumeSize"`
	Status     string `json:"status"`
	Name       string `json:"name"`
}

type ListSnapshots struct {
	Items      []*Snapshot
	TotalPages int
	Page       int
	PageSize   int
	TotalItems int
}
