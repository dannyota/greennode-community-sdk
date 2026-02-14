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

func (s Image) toEntityImage() entity.Image {
	return entity.Image{
		FlavorZoneIDs: s.FlavorZoneIDs,
		ID:            s.ID,
		ImageType:     s.ImageType,
		ImageVersion:  s.ImageVersion,
		Licence:       s.Licence,
		PackageLimit: entity.PackageLimit{
			Cpu:      s.PackageLimit.Cpu,
			DiskSize: s.PackageLimit.DiskSize,
			Memory:   s.PackageLimit.Memory,
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

func (s NetworkInterface) toEntityNetworkInterface() entity.NetworkInterface {
	return entity.NetworkInterface{
		CreatedAt:     s.CreatedAt,
		FixedIp:       s.FixedIp,
		FloatingIp:    s.FloatingIp,
		FloatingIpID:  s.FloatingIpID,
		InterfaceType: s.InterfaceType,
		Mac:           s.Mac,
		NetworkUuid:   s.NetworkUuid,
		PortUuid:      s.PortUuid,
		Product:       s.Product,
		ServerUuid:    s.ServerUuid,
		Status:        s.Status,
		SubnetUuid:    s.SubnetUuid,
		Type:          s.Type,
		UpdatedAt:     s.UpdatedAt,
		Uuid:          s.Uuid,
	}
}

func (s ServerSecgroup) toEntityServerSecgroup() entity.ServerSecgroup {
	return entity.ServerSecgroup{
		Name: s.Name,
		Uuid: s.Uuid,
	}
}

func (s Flavor) toEntityFlavor() entity.Flavor {
	return entity.Flavor{
		Bandwidth:              s.Bandwidth,
		BandwidthUnit:          s.BandwidthUnit,
		Cpu:                    s.Cpu,
		CpuPlatformDescription: s.CpuPlatformDescription,
		FlavorID:               s.FlavorID,
		Gpu:                    s.Gpu,
		Group:                  s.Group,
		Memory:                 s.Memory,
		MetaData:               s.MetaData,
		Name:                   s.Name,
		RemainingVms:           s.RemainingVms,
		ZoneID:                 s.ZoneID,
	}
}

func (s Server) toEntityServer() *entity.Server {
	server := new(entity.Server)
	server.BootVolumeID = s.BootVolumeID
	server.CreatedAt = s.CreatedAt
	server.EncryptionVolume = s.EncryptionVolume
	server.Licence = s.Licence
	server.Location = s.Location
	server.Metadata = s.Metadata
	server.MigrateState = s.MigrateState
	server.Name = s.Name
	server.Product = s.Product
	server.ServerGroupID = s.ServerGroupID
	server.ServerGroupName = s.ServerGroupName
	server.SshKeyName = s.SshKeyName
	server.Status = s.Status
	server.StopBeforeMigrate = s.StopBeforeMigrate
	server.User = s.User
	server.Uuid = s.Uuid
	server.Image = s.Image.toEntityImage()
	server.Flavor = s.Flavor.toEntityFlavor()
	server.ZoneID = s.ZoneID

	for _, secGroup := range s.SecGroups {
		server.SecGroups = append(server.SecGroups, secGroup.toEntityServerSecgroup())
	}

	for _, externalInterface := range s.ExternalInterfaces {
		server.ExternalInterfaces = append(server.ExternalInterfaces, externalInterface.toEntityNetworkInterface())
	}

	for _, internalInterface := range s.InternalInterfaces {
		server.InternalInterfaces = append(server.InternalInterfaces, internalInterface.toEntityNetworkInterface())
	}

	return server
}

type CreateServerResponse struct {
	Data Server `json:"data"`
}

func (s *CreateServerResponse) ToEntityServer() *entity.Server {
	return s.Data.toEntityServer()
}

type GetServerByIDResponse struct {
	Data Server `json:"data"`
}

func (s *GetServerByIDResponse) ToEntityServer() *entity.Server {
	return s.Data.toEntityServer()
}

type UpdateServerSecgroupsByServerIDResponse struct {
	Data Server `json:"data"`
}

func (s *UpdateServerSecgroupsByServerIDResponse) ToEntityServer() *entity.Server {
	return s.Data.toEntityServer()
}

func (s *ListServerGroupPoliciesResponse) ToEntityListServerGroupPolicies() *entity.ListServerGroupPolicies {
	serverGroupPolicies := &entity.ListServerGroupPolicies{}
	for _, itemServerGroupPolicy := range s.Data {
		serverGroupPolicies.Add(itemServerGroupPolicy.toEntityServerGroupPolicy())
	}
	return serverGroupPolicies
}

func (s *ListServerGroupsResponse) ToEntityListServerGroups() *entity.ListServerGroups {
	serverGroups := &entity.ListServerGroups{}
	for _, itemServerGroup := range s.ListData {
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

	serverGroups.Page = s.Page
	serverGroups.PageSize = s.PageSize
	serverGroups.TotalPage = s.TotalPage
	serverGroups.TotalItem = s.TotalItem

	return serverGroups
}

func (s *CreateServerGroupResponse) ToEntityServerGroup() *entity.ServerGroup {
	return &entity.ServerGroup{
		UUID:        s.Data.UUID,
		Name:        s.Data.Name,
		Description: s.Data.Description,
		PolicyID:    s.Data.PolicyID,
	}
}
