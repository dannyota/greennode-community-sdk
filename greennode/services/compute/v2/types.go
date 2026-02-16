package v2

import "strings"

const (
	ServerStatusActive  = "ACTIVE"
	ServerStatusError   = "ERROR"
	ServerStatusStopped = "STOPPED"
)

type (
	Server struct {
		BootVolumeID       string             `json:"bootVolumeId"`
		CreatedAt          string             `json:"createdAt"`
		EncryptionVolume   bool               `json:"encryptionVolume"`
		Licence            bool               `json:"licence"`
		Location           string             `json:"location"`
		Metadata           string             `json:"metadata"`
		MigrateState       string             `json:"migrateState"`
		Name               string             `json:"name"`
		Product            string             `json:"product"`
		ServerGroupID      string             `json:"serverGroupId"`
		ServerGroupName    string             `json:"serverGroupName"`
		SshKeyName         string             `json:"sshKeyName"`
		Status             string             `json:"status"`
		StopBeforeMigrate  bool               `json:"stopBeforeMigrate"`
		User               string             `json:"user"`
		Uuid               string             `json:"uuid"`
		Image              Image              `json:"image"`
		Flavor             Flavor             `json:"flavor"`
		SecGroups          []ServerSecgroup   `json:"secGroups"`
		ExternalInterfaces []NetworkInterface `json:"externalInterfaces"`
		InternalInterfaces []NetworkInterface `json:"internalInterfaces"`
		ZoneID             string             `json:"zoneId"`
	}

	NetworkInterface struct {
		CreatedAt     string `json:"createdAt"`
		FixedIp       string `json:"fixedIp"`
		FloatingIp    string `json:"floatingIp"`
		FloatingIpID  string `json:"floatingIpId"`
		InterfaceType string `json:"interfaceType"`
		Mac           string `json:"mac"`
		NetworkUuid   string `json:"networkUuid"`
		PortUuid      string `json:"portUuid"`
		Product       string `json:"product"`
		ServerUuid    string `json:"serverUuid"`
		Status        string `json:"status"`
		SubnetUuid    string `json:"subnetUuid"`
		Type          string `json:"type"`
		UpdatedAt     string `json:"updatedAt"`
		Uuid          string `json:"uuid"`
	}

	Flavor struct {
		Bandwidth              int64  `json:"bandwidth"`
		BandwidthUnit          string `json:"bandwidthUnit"`
		Cpu                    int64  `json:"cpu"`
		CpuPlatformDescription string `json:"cpuPlatformDescription"`
		FlavorID               string `json:"flavorId"`
		Gpu                    int64  `json:"gpu"`
		Group                  string `json:"group"`
		Memory                 int64  `json:"memory"`
		MetaData               string `json:"metaData"`
		Name                   string `json:"name"`
		RemainingVms           int64  `json:"remainingVms"`
		ZoneID                 string `json:"zoneId"`
	}

	Image struct {
		FlavorZoneIDs []string     `json:"flavorZoneIds"`
		ID            string       `json:"id"`
		ImageType     string       `json:"imageType"`
		ImageVersion  string       `json:"imageVersion"`
		Licence       bool         `json:"licence"`
		PackageLimit  PackageLimit `json:"packageLimit"`
	}

	PackageLimit struct {
		Cpu      int64 `json:"cpu"`
		DiskSize int64 `json:"diskSize"`
		Memory   int64 `json:"memory"`
	}

	ServerSecgroup struct {
		Name string `json:"name"`
		Uuid string `json:"uuid"`
	}
)

type ListServers struct {
	Items []*Server
}

func (sv Server) CanDelete() bool {
	switch strings.ToUpper(sv.Status) {
	case ServerStatusActive, ServerStatusError, ServerStatusStopped:
		return true
	}

	return false
}

func (sv Server) IsRunning() bool {
	switch strings.ToUpper(sv.Status) {
	case ServerStatusActive:
		return true
	}

	return false
}

func (sv Server) GetInternalInterfaceWanInfo() (string, string, string, bool) {
	for _, i := range sv.InternalInterfaces {
		if i.FloatingIp != "" {
			return i.Uuid, i.FloatingIpID, i.FloatingIp, true
		}
	}

	return "", "", "", false
}

func (sv Server) GetInternalNetworkInterfaceIDs() []string {
	ids := make([]string, 0, len(sv.InternalInterfaces))
	for _, i := range sv.InternalInterfaces {
		ids = append(ids, i.Uuid)
	}

	return ids
}

func (sv Server) InternalNetworkInterfacePossible() bool {
	return len(sv.InternalInterfaces) > 0
}

func (sv Server) CanAttachFloatingIp() bool {
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

type ServerGroup struct {
	UUID        string              `json:"uuid"`
	Name        string              `json:"name"`
	Description string              `json:"description"`
	PolicyID    string              `json:"policyId"`
	PolicyName  string              `json:"policyName"`
	Servers     []ServerGroupMember `json:"servers"`
}

type ServerGroupMember struct {
	Name string `json:"name"`
	UUID string `json:"uuid"`
}

type ListServerGroups struct {
	Items     []*ServerGroup `json:"listData"`
	Page      int            `json:"page"`
	PageSize  int            `json:"pageSize"`
	TotalPage int            `json:"totalPage"`
	TotalItem int            `json:"totalItem"`
}

func (l *ListServerGroups) Add(item *ServerGroup) {
	l.Items = append(l.Items, item)
}

func (l *ListServerGroups) FindServerGroupByServerGroupID(serverGroupID string) (*ServerGroup, bool) {
	for _, item := range l.Items {
		if item.UUID == serverGroupID {
			return item, true
		}
	}

	return nil, false
}

type ServerGroupPolicy struct {
	Name         string            `json:"name"`
	UUID         string            `json:"uuid"`
	Status       string            `json:"status"`
	Descriptions map[string]string `json:"descriptions"`
}

type ListServerGroupPolicies struct {
	Items []*ServerGroupPolicy
}

func (l *ListServerGroupPolicies) Add(item *ServerGroupPolicy) {
	l.Items = append(l.Items, item)
}

func (l *ListServerGroupPolicies) At(idx int) *ServerGroupPolicy {
	if idx < 0 || idx >= l.Len() {
		return nil
	}

	return l.Items[idx]
}

func (l *ListServerGroupPolicies) Len() int {
	return len(l.Items)
}
