package entity

import "strings"

const (
	ServerStatusActive  = "ACTIVE"
	ServerStatusError   = "ERROR"
	ServerStatusStopped = "STOPPED"
)

type (
	Server struct {
		BootVolumeID       string
		CreatedAt          string
		EncryptionVolume   bool
		Licence            bool
		Location           string
		Metadata           string
		MigrateState       string
		Name               string
		Product            string
		ServerGroupID      string
		ServerGroupName    string
		SshKeyName         string
		Status             string
		StopBeforeMigrate  bool
		User               string
		Uuid               string
		Image              Image
		Flavor             Flavor
		SecGroups          []ServerSecgroup
		ExternalInterfaces []NetworkInterface
		InternalInterfaces []NetworkInterface
		ZoneID             string
	}

	NetworkInterface struct {
		CreatedAt     string
		FixedIp       string
		FloatingIp    string
		FloatingIpID  string
		InterfaceType string
		Mac           string
		NetworkUuid   string
		PortUuid      string
		Product       string
		ServerUuid    string
		Status        string
		SubnetUuid    string
		Type          string
		UpdatedAt     string
		Uuid          string
	}

	Flavor struct {
		Bandwidth              int64
		BandwidthUnit          string
		Cpu                    int64
		CpuPlatformDescription string
		FlavorID               string
		Gpu                    int64
		Group                  string
		Memory                 int64
		MetaData               string
		Name                   string
		RemainingVms           int64
		ZoneID                 string
	}

	Image struct {
		FlavorZoneIDs []string
		ID            string
		ImageType     string
		ImageVersion  string
		Licence       bool
		PackageLimit  PackageLimit
	}

	PackageLimit struct {
		Cpu      int64
		DiskSize int64
		Memory   int64
	}

	ServerSecgroup struct {
		Name string
		Uuid string
	}

	SystemTag struct {
		Key       string
		Value     string
		CreatedAt string
		SystemTag bool
	}
)

type ListServers struct {
	Items []*Server
}

func (sv *Server) CanDelete() bool {
	switch strings.ToUpper(sv.Status) {
	case ServerStatusActive, ServerStatusError, ServerStatusStopped:
		return true
	}

	return false
}

func (sv *Server) IsRunning() bool {
	switch strings.ToUpper(sv.Status) {
	case ServerStatusActive:
		return true
	}

	return false
}

func (sv *Server) GetInternalInterfaceWanInfo() (string, string, string, bool) {
	for _, i := range sv.InternalInterfaces {
		if i.FloatingIp != "" {
			return i.Uuid, i.FloatingIpID, i.FloatingIp, true
		}
	}

	return "", "", "", false
}

func (sv *Server) GetInternalNetworkInterfaceIDs() []string {
	ids := make([]string, 0, len(sv.InternalInterfaces))
	for _, i := range sv.InternalInterfaces {
		ids = append(ids, i.Uuid)
	}

	return ids
}

func (sv *Server) InternalNetworkInterfacePossible() bool {
	return len(sv.InternalInterfaces) > 0
}

func (sv *Server) CanAttachFloatingIp() bool {
	if !sv.InternalNetworkInterfacePossible() {
		return false
	}

	for _, i := range sv.InternalInterfaces {
		if i.FloatingIp != "" {
			return false
		}
	}

	return true
}
