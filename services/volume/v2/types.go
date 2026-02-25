package v2

type Volume struct {
	Name             string   `json:"name"`
	ID               string   `json:"uuid"`
	VolumeTypeID     string   `json:"volumeTypeId"`
	ClusterID        *string  `json:"clusterId"`
	VmID             string   `json:"serverId"`
	Size             uint64   `json:"size"`
	IopsID           uint64   `json:"iopsId"`
	Status           string   `json:"status"`
	CreatedAt        string   `json:"createdAt"`
	UpdatedAt        *string  `json:"updatedAt"`
	PersistentVolume bool     `json:"persistentVolume"`
	AttachedMachine  []string `json:"serverIdList"`
	UnderID          string   `json:"underId"`
	MigrateState     string   `json:"migrateState"`
	MultiAttach      bool     `json:"multiAttach"`
	ZoneID           string   `json:"zoneId"`
}

type ListVolumes struct {
	Items      []*Volume
	TotalPage  int
	Page       int
	PageSize   int
	TotalItems int
}

func (l ListVolumes) Len() int {
	return len(l.Items)
}

func (v Volume) AttachedTheInstance(instanceID string) bool {
	if v.VmID == instanceID {
		return true
	}

	for _, machineID := range v.AttachedMachine {
		if machineID == instanceID {
			return true
		}
	}

	return false
}

func (v Volume) IsAvailable() bool {
	return v.Status == "AVAILABLE"
}

func (v Volume) IsError() bool {
	return v.Status == "ERROR"
}

func (v Volume) IsInUse() bool {
	return v.Status == "IN-USE"
}

func (v Volume) CanDelete() bool {
	if len(v.AttachedMachine) < 1 && v.VmID == "" && v.Status == "AVAILABLE" {
		return true
	}

	if v.Status == "ERROR" {
		return true
	}

	return false
}

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
