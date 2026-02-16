package entity

type VolumeType struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	Iops       int    `json:"iops"`
	MaxSize    int    `json:"maxSize"`
	MinSize    int    `json:"minSize"`
	ThroughPut int    `json:"throughPut"`
	ZoneID     string `json:"zoneId"`
}

type ListVolumeTypes struct {
	VolumeTypes []*VolumeType
}
