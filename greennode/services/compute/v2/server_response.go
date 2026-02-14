package v2

import "github.com/dannyota/greennode-community-sdk/v2/greennode/entity"

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
		ZoneID             string             `json:"zoneId"`
		ExternalInterfaces []NetworkInterface `json:"externalInterfaces"`
		InternalInterfaces []NetworkInterface `json:"internalInterfaces"`
	}

	NetworkInterface struct {
		CreatedAt     string `json:"createdAt"`
		FixedIp       string `json:"fixedIp"`
		FloatingIp    string `json:"floatingIp,omitempty"`
		FloatingIpID  string `json:"floatingIpId,omitempty"`
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

	ServerSecgroupPolicy struct {
		Name          string `json:"name"`
		UUID          string `json:"uuid"`
		Status        string `json:"status"`
		Description   string `json:"description"`
		DescriptionVi string `json:"descriptionVi"`
	}

	ListServerGroupPoliciesResponse struct {
		Data []ServerSecgroupPolicy `json:"data"`
	}

	ListServerGroupsResponse struct {
		ListData  []ServerGroup `json:"listData"`
		Page      int           `json:"page"`
		PageSize  int           `json:"pageSize"`
		TotalPage int           `json:"totalPage"`
		TotalItem int           `json:"totalItem"`
	}

	ServerGroup struct {
		UUID        string `json:"uuid"`
		Name        string `json:"name"`
		Description string `json:"description"`
		PolicyID    string `json:"policyId"`
		PolicyName  string `json:"policyName"`
		Servers     []struct {
			Name string `json:"name"`
			UUID string `json:"uuid"`
		} `json:"servers"`
	}

	CreateServerGroupResponse struct {
		Data struct {
			UUID          string `json:"uuid"`
			ServerGroupID int    `json:"serverGroupId"`
			Name          string `json:"name"`
			Description   string `json:"description"`
			PolicyID      string `json:"policyId"`
			CreatedAt     string `json:"createdAt"`
		} `json:"data"`
	}
)

func (i Image) toEntityImage() entity.Image {
	return entity.Image{
		FlavorZoneIDs: i.FlavorZoneIDs,
		ID:            i.ID,
		ImageType:     i.ImageType,
		ImageVersion:  i.ImageVersion,
		Licence:       i.Licence,
		PackageLimit: entity.PackageLimit{
			Cpu:      i.PackageLimit.Cpu,
			DiskSize: i.PackageLimit.DiskSize,
			Memory:   i.PackageLimit.Memory,
		},
	}
}

func (s ServerSecgroupPolicy) toEntityServerGroupPolicy() *entity.ServerGroupPolicy {
	return &entity.ServerGroupPolicy{
		Name:   s.Name,
		UUID:   s.UUID,
		Status: s.Status,
		Descriptions: map[string]string{
			"en": s.Description,
			"vi": s.DescriptionVi,
		},
	}
}

func (n NetworkInterface) toEntityNetworkInterface() entity.NetworkInterface {
	return entity.NetworkInterface{
		CreatedAt:     n.CreatedAt,
		FixedIp:       n.FixedIp,
		FloatingIp:    n.FloatingIp,
		FloatingIpID:  n.FloatingIpID,
		InterfaceType: n.InterfaceType,
		Mac:           n.Mac,
		NetworkUuid:   n.NetworkUuid,
		PortUuid:      n.PortUuid,
		Product:       n.Product,
		ServerUuid:    n.ServerUuid,
		Status:        n.Status,
		SubnetUuid:    n.SubnetUuid,
		Type:          n.Type,
		UpdatedAt:     n.UpdatedAt,
		Uuid:          n.Uuid,
	}
}

func (s ServerSecgroup) toEntityServerSecgroup() entity.ServerSecgroup {
	return entity.ServerSecgroup{
		Name: s.Name,
		Uuid: s.Uuid,
	}
}

func (f Flavor) toEntityFlavor() entity.Flavor {
	return entity.Flavor{
		Bandwidth:              f.Bandwidth,
		BandwidthUnit:          f.BandwidthUnit,
		Cpu:                    f.Cpu,
		CpuPlatformDescription: f.CpuPlatformDescription,
		FlavorID:               f.FlavorID,
		Gpu:                    f.Gpu,
		Group:                  f.Group,
		Memory:                 f.Memory,
		MetaData:               f.MetaData,
		Name:                   f.Name,
		RemainingVms:           f.RemainingVms,
		ZoneID:                 f.ZoneID,
	}
}

func (sv Server) toEntityServer() *entity.Server {
	server := new(entity.Server)
	server.BootVolumeID = sv.BootVolumeID
	server.CreatedAt = sv.CreatedAt
	server.EncryptionVolume = sv.EncryptionVolume
	server.Licence = sv.Licence
	server.Location = sv.Location
	server.Metadata = sv.Metadata
	server.MigrateState = sv.MigrateState
	server.Name = sv.Name
	server.Product = sv.Product
	server.ServerGroupID = sv.ServerGroupID
	server.ServerGroupName = sv.ServerGroupName
	server.SshKeyName = sv.SshKeyName
	server.Status = sv.Status
	server.StopBeforeMigrate = sv.StopBeforeMigrate
	server.User = sv.User
	server.Uuid = sv.Uuid
	server.Image = sv.Image.toEntityImage()
	server.Flavor = sv.Flavor.toEntityFlavor()
	server.ZoneID = sv.ZoneID

	for _, secGroup := range sv.SecGroups {
		server.SecGroups = append(server.SecGroups, secGroup.toEntityServerSecgroup())
	}

	for _, externalInterface := range sv.ExternalInterfaces {
		server.ExternalInterfaces = append(server.ExternalInterfaces, externalInterface.toEntityNetworkInterface())
	}

	for _, internalInterface := range sv.InternalInterfaces {
		server.InternalInterfaces = append(server.InternalInterfaces, internalInterface.toEntityNetworkInterface())
	}

	return server
}

type CreateServerResponse struct {
	Data Server `json:"data"`
}

func (r *CreateServerResponse) ToEntityServer() *entity.Server {
	return r.Data.toEntityServer()
}

type GetServerByIDResponse struct {
	Data Server `json:"data"`
}

func (r *GetServerByIDResponse) ToEntityServer() *entity.Server {
	return r.Data.toEntityServer()
}

type UpdateServerSecgroupsByServerIDResponse struct {
	Data Server `json:"data"`
}

func (r *UpdateServerSecgroupsByServerIDResponse) ToEntityServer() *entity.Server {
	return r.Data.toEntityServer()
}

func (r *ListServerGroupPoliciesResponse) ToEntityListServerGroupPolicies() *entity.ListServerGroupPolicies {
	serverGroupPolicies := &entity.ListServerGroupPolicies{}
	for _, itemServerGroupPolicy := range r.Data {
		serverGroupPolicies.Add(itemServerGroupPolicy.toEntityServerGroupPolicy())
	}
	return serverGroupPolicies
}

func (r *ListServerGroupsResponse) ToEntityListServerGroups() *entity.ListServerGroups {
	serverGroups := &entity.ListServerGroups{}
	for _, itemServerGroup := range r.ListData {
		serverGroup := &entity.ServerGroup{
			UUID:        itemServerGroup.UUID,
			Name:        itemServerGroup.Name,
			Description: itemServerGroup.Description,
			PolicyID:    itemServerGroup.PolicyID,
			PolicyName:  itemServerGroup.PolicyName,
		}

		for _, server := range itemServerGroup.Servers {
			serverGroup.Servers = append(serverGroup.Servers, entity.ServerGroupMember{
				Name: server.Name,
				UUID: server.UUID,
			})
		}
		serverGroups.Add(serverGroup)
	}

	serverGroups.Page = r.Page
	serverGroups.PageSize = r.PageSize
	serverGroups.TotalPage = r.TotalPage
	serverGroups.TotalItem = r.TotalItem

	return serverGroups
}

func (r *CreateServerGroupResponse) ToEntityServerGroup() *entity.ServerGroup {
	return &entity.ServerGroup{
		UUID:        r.Data.UUID,
		Name:        r.Data.Name,
		Description: r.Data.Description,
		PolicyID:    r.Data.PolicyID,
	}
}
