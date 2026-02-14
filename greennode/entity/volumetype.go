package entity

type VolumeType struct {
	ID         string
	Name       string
	Iops       int
	MaxSize    int
	MinSize    int
	ThroughPut int
	ZoneID     string
}

type ListVolumeType struct {
	VolumeTypes []*VolumeType
}
