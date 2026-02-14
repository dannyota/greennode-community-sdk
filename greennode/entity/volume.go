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

func (s *ListVolumes) Len() int {
	return len(s.Items)
}

func (s *Volume) AttachedTheInstance(instanceID string) bool {
	if s.VmID == instanceID {
		return true
	}

	for _, machineID := range s.AttachedMachine {
		if machineID == instanceID {
			return true
		}
	}

	return false
}

func (s *Volume) IsAvailable() bool {
	return s.Status == "AVAILABLE"
}

func (s *Volume) IsError() bool {
	return s.Status == ServerStatusError
}

func (s *Volume) IsInUse() bool {
	return s.Status == "IN-USE"
}

func (s *Volume) CanDelete() bool {
	if len(s.AttachedMachine) < 1 && s.VmID == "" && s.Status == "AVAILABLE" {
		return true
	}

	if s.Status == "ERROR" {
		return true
	}

	return false
}
