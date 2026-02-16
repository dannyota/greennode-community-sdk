package entity

type VolumeTypeZone struct {
	ID       string   `json:"id"`
	Name     string   `json:"name"`
	PoolName []string `json:"poolName"`
}

type ListVolumeTypeZones struct {
	VolumeTypeZones []*VolumeTypeZone
}
