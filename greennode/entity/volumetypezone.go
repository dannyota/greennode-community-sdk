package entity

type VolumeTypeZone struct {
	ID       string
	Name     string
	PoolName []string
}

type ListVolumeTypeZones struct {
	VolumeTypeZones []*VolumeTypeZone
}
