package entity

type Volume struct {
	Name             string
	ID               string
	VolumeTypeID     string
	ClusterID        *string
	VmID             string
	Size             uint64
	IopsID           uint64
	Status           string
	CreatedAt        string
	UpdatedAt        *string
	PersistentVolume bool
	AttachedMachine  []string
	UnderID          string
	MigrateState     string
	MultiAttach      bool
	ZoneID           string
}

type ListVolumes struct {
	Items []*Volume
}

func (l *ListVolumes) Len() int {
	return len(l.Items)
}

func (v *Volume) AttachedTheInstance(instanceID string) bool {
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

func (v *Volume) IsAvailable() bool {
	return v.Status == "AVAILABLE"
}

func (v *Volume) IsError() bool {
	return v.Status == ServerStatusError
}

func (v *Volume) IsInUse() bool {
	return v.Status == "IN-USE"
}

func (v *Volume) CanDelete() bool {
	if len(v.AttachedMachine) < 1 && v.VmID == "" && v.Status == "AVAILABLE" {
		return true
	}

	if v.Status == "ERROR" {
		return true
	}

	return false
}
