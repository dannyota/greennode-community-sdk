package v2

import (
	"fmt"
	"net/url"
	"strconv"

	"github.com/dannyota/greennode-community-sdk/v2/greennode/services/common"
)

const (
	DataDiskEncryptionAesXts256Type DataDiskEncryptionType = "aes-xts-plain64_256"
	DataDiskEncryptionAesXts128Type DataDiskEncryptionType = "aes-xts-plain64_128"
)

type CreateServerRequest struct {
	AttachFloating         bool                     `json:"attachFloating,omitempty"`
	BackupInstancePointID  string                   `json:"backupInstancePointId,omitempty"`
	DataDiskEncryptionType DataDiskEncryptionType   `json:"dataDiskEncryptionType,omitempty"`
	DataDiskName           string                   `json:"dataDiskName,omitempty"`
	DataDiskSize           int                      `json:"dataDiskSize,omitempty"`
	DataDiskTypeID         string                   `json:"dataDiskTypeId,omitempty"`
	EnableBackup           bool                     `json:"enableBackup,omitempty"`
	EncryptionVolume       bool                     `json:"encryptionVolume"`
	ExpirePassword         bool                     `json:"expirePassword,omitempty"`
	FlavorID               string                   `json:"flavorId"`
	ImageID                string                   `json:"imageId"`
	Name                   string                   `json:"name"`
	NetworkID              string                   `json:"networkId,omitempty"`
	SubnetID               string                   `json:"subnetId,omitempty"`
	OsLicence              bool                     `json:"osLicence,omitempty"`
	RestoreBackup          bool                     `json:"restoreBackup,omitempty"`
	RootDiskEncryptionType DataDiskEncryptionType   `json:"rootDiskEncryptionType,omitempty"`
	RootDiskSize           int                      `json:"rootDiskSize"`
	RootDiskTypeID         string                   `json:"rootDiskTypeId"`
	SecurityGroup          []string                 `json:"securityGroup,omitempty"`
	ServerGroupID          string                   `json:"serverGroupId,omitempty"`
	SshKeyID               string                   `json:"sshKeyId,omitempty"`
	UserData               string                   `json:"userData,omitempty"`
	UserDataBase64Encoded  bool                     `json:"userDataBase64Encoded,omitempty"`
	UserName               string                   `json:"userName,omitempty"`
	UserPassword           string                   `json:"userPassword,omitempty"`
	IsPoc                  bool                     `json:"isPoc,omitempty"`
	Product                string                   `json:"product,omitempty"`
	Type                   string                   `json:"type,omitempty"`
	Tags                   []ServerTag              `json:"tags,omitempty"`
	AutoRenew              bool                     `json:"isEnableAutoRenew,omitempty"`
	Networks               []ServerNetworkInterface `json:"networks,omitempty"`
	Zone                   string                   `json:"zoneId,omitempty"`
	common.UserAgent
}

type ServerNetworkInterface struct {
	ProjectID      string `json:"projectId"`
	NetworkID      string `json:"networkId"`
	SubnetID       string `json:"subnetId"`
	AttachFloating bool   `json:"attachFloating"`
}

type AttachBlockVolumeRequest struct {
	common.BlockVolumeCommon
	common.ServerCommon
}

type DetachBlockVolumeRequest struct {
	common.BlockVolumeCommon
	common.ServerCommon
}

type DataDiskEncryptionType string

type ServerTag struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func (s *CreateServerRequest) ToRequestBody() interface{} {
	return s
}

func (s *CreateServerRequest) WithZone(zoneID string) ICreateServerRequest {
	s.Zone = zoneID
	return s
}

func (s *CreateServerRequest) WithServerNetworkInterface(projectID, networkID, subnetID string, attachFloating bool) ICreateServerRequest {
	s.Networks = append(s.Networks, ServerNetworkInterface{
		ProjectID:      projectID,
		NetworkID:      networkID,
		SubnetID:       subnetID,
		AttachFloating: attachFloating,
	})

	return s.WithNetwork(s.Networks[0].NetworkID, s.Networks[0].SubnetID)
}

func (s *CreateServerRequest) WithRootDiskEncryptionType(dataDisk DataDiskEncryptionType) ICreateServerRequest {
	s.EncryptionVolume = true
	s.RootDiskEncryptionType = dataDisk
	return s
}

func (s *CreateServerRequest) WithEncryptionVolume(enabled bool) ICreateServerRequest {
	s.EncryptionVolume = enabled
	return s
}

func (s *CreateServerRequest) WithUserData(userData string, base64Encode bool) ICreateServerRequest {
	s.UserData = userData
	s.UserDataBase64Encoded = base64Encode
	return s
}

func (s *CreateServerRequest) WithAutoRenew(val bool) ICreateServerRequest {
	s.AutoRenew = val
	return s
}

func (s *CreateServerRequest) WithTags(tags ...string) ICreateServerRequest {
	if s.Tags == nil {
		s.Tags = make([]ServerTag, 0)
	}

	if len(tags)%2 != 0 {
		tags = append(tags, "none")
	}

	for i := 0; i < len(tags); i += 2 {
		s.Tags = append(s.Tags, ServerTag{Key: tags[i], Value: tags[i+1]})
	}
	return s
}

func (s *CreateServerRequest) WithAttachFloating(attachFloating bool) ICreateServerRequest {
	s.AttachFloating = attachFloating
	return s
}

func (s *CreateServerRequest) WithSecgroups(secgroups ...string) ICreateServerRequest {
	s.SecurityGroup = append(s.SecurityGroup, secgroups...)
	return s
}

func (s *CreateServerRequest) WithServerGroupID(serverGroupID string) ICreateServerRequest {
	s.ServerGroupID = serverGroupID
	return s
}

func (s *CreateServerRequest) WithPoc(isPoc bool) ICreateServerRequest {
	s.IsPoc = isPoc
	return s
}

func (s *CreateServerRequest) WithType(typeVal string) ICreateServerRequest {
	s.Type = typeVal
	return s
}

func (s *CreateServerRequest) WithProduct(product string) ICreateServerRequest {
	s.Product = product
	return s
}

func (s *CreateServerRequest) WithNetwork(networkID, subnetID string) ICreateServerRequest {
	s.NetworkID = networkID
	s.SubnetID = subnetID

	return s
}

func (s *CreateServerRequest) AddUserAgent(agent ...string) ICreateServerRequest {
	s.UserAgent.AddUserAgent(agent...)
	return s
}

func (s *CreateServerRequest) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"attachFloating":         s.AttachFloating,
		"backupInstancePointId":  s.BackupInstancePointID,
		"dataDiskEncryptionType": s.DataDiskEncryptionType,
		"dataDiskName":           s.DataDiskName,
		"dataDiskSize":           s.DataDiskSize,
		"dataDiskTypeId":         s.DataDiskTypeID,
		"enableBackup":           s.EnableBackup,
		"encryptionVolume":       s.EncryptionVolume,
		"expirePassword":         s.ExpirePassword,
		"flavorId":               s.FlavorID,
		"imageId":                s.ImageID,
		"name":                   s.Name,
		"networkId":              s.NetworkID,
		"subnetId":               s.SubnetID,
		"osLicence":              s.OsLicence,
		"restoreBackup":          s.RestoreBackup,
		"rootDiskEncryptionType": s.RootDiskEncryptionType,
		"rootDiskSize":           s.RootDiskSize,
		"rootDiskTypeId":         s.RootDiskTypeID,
		"securityGroup":          s.SecurityGroup,
		"serverGroupId":          s.ServerGroupID,
		"sshKeyId":               s.SshKeyID,
		"userName":               s.UserName,
		"isPoc":                  s.IsPoc,
		"product":                s.Product,
		"type":                   s.Type,
		"tags":                   s.Tags,
		"autoRenew":              s.AutoRenew,
		"networks":               s.Networks,
	}
}

type GetServerByIDRequest struct {
	common.ServerCommon
	common.UserAgent
}

func (s *GetServerByIDRequest) AddUserAgent(agent ...string) IGetServerByIDRequest {
	s.UserAgent.AddUserAgent(agent...)
	return s
}

func (s *GetServerByIDRequest) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"serverId": s.ServerID,
	}
}

type DeleteServerByIDRequest struct {
	DeleteAllVolume bool `json:"deleteAllVolume"`
	common.ServerCommon
	common.UserAgent
}

func (s *DeleteServerByIDRequest) WithDeleteAllVolume(ok bool) IDeleteServerByIDRequest {
	s.DeleteAllVolume = ok
	return s
}

func (s *DeleteServerByIDRequest) AddUserAgent(agent ...string) IDeleteServerByIDRequest {
	s.UserAgent.AddUserAgent(agent...)
	return s
}

func (s *DeleteServerByIDRequest) ToRequestBody() interface{} {
	return s
}

type UpdateServerSecgroupsByServerIDRequest struct {
	Secgroups []string `json:"securityGroup"`

	common.ServerCommon
	common.UserAgent
}

func (s *UpdateServerSecgroupsByServerIDRequest) AddUserAgent(agent ...string) IUpdateServerSecgroupsByServerIDRequest {
	s.UserAgent.AddUserAgent(agent...)
	return s
}

func (s *UpdateServerSecgroupsByServerIDRequest) ToRequestBody() interface{} {
	return s
}

func (s *UpdateServerSecgroupsByServerIDRequest) GetListSecgroupsIDs() []string {
	return s.Secgroups
}

type AttachFloatingIpRequest struct {
	NetworkInterfaceID string `json:"networkInterfaceId"`

	common.InternalNetworkInterfaceCommon
	common.ServerCommon
	common.UserAgent
}

func (s *AttachFloatingIpRequest) ToRequestBody() interface{} {
	return s
}

func (s *AttachFloatingIpRequest) AddUserAgent(agent ...string) IAttachFloatingIpRequest {
	s.Agent = append(s.Agent, agent...)
	return s
}

func (s *AttachFloatingIpRequest) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"serverId":                   s.ServerID,
		"internalNetworkInterfaceId": s.InternalNetworkInterfaceID,
		"networkId":                  s.NetworkInterfaceID,
	}
}

type DetachFloatingIpRequest struct {
	NetworkInterfaceID string `json:"networkInterfaceId"`

	common.ServerCommon
	common.WanCommon
	common.InternalNetworkInterfaceCommon
	common.UserAgent
}

func (s *DetachFloatingIpRequest) ToRequestBody() interface{} {
	return s
}

func (s *DetachFloatingIpRequest) AddUserAgent(agent ...string) IDetachFloatingIpRequest {
	s.Agent = append(s.Agent, agent...)
	return s
}

func (s *DetachFloatingIpRequest) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"serverId":                   s.ServerID,
		"internalNetworkInterfaceId": s.InternalNetworkInterfaceID,
		"networkId":                  s.NetworkInterfaceID,
		"wanId":                      s.WanID,
	}
}

type ListServerGroupPoliciesRequest struct {
	common.UserAgent
}

func (s *ListServerGroupPoliciesRequest) AddUserAgent(agent ...string) IListServerGroupPoliciesRequest {
	s.Agent = append(s.Agent, agent...)
	return s
}

type DeleteServerGroupByIDRequest struct {
	common.ServerGroupCommon
	common.UserAgent
}

func (s *DeleteServerGroupByIDRequest) AddUserAgent(agent ...string) IDeleteServerGroupByIDRequest {
	s.Agent = append(s.Agent, agent...)
	return s
}

func (s *DeleteServerGroupByIDRequest) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"serverGroupId": s.ServerGroupID,
	}
}

type ListServerGroupsRequest struct {
	Name string
	Page int
	Size int

	common.UserAgent
}

func (s *ListServerGroupsRequest) WithName(name string) IListServerGroupsRequest {
	s.Name = name
	return s
}

func (s *ListServerGroupsRequest) ToListQuery() (string, error) {
	v := url.Values{}
	v.Set("name", s.Name)
	if s.Page > 0 {
		v.Set("page", strconv.Itoa(s.Page))
	}
	if s.Size > 0 {
		v.Set("size", strconv.Itoa(s.Size))
	}
	return v.Encode(), nil
}

func (s *ListServerGroupsRequest) GetDefaultQuery() string {
	return fmt.Sprintf("offset=%d&limit=%d&name=", defaultOffsetListServerGroups, defaultLimitListServerGroups)
}

func (s *ListServerGroupsRequest) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"page": s.Page,
		"size": s.Size,
		"name": s.Name,
	}
}

func (s *ListServerGroupsRequest) AddUserAgent(agent ...string) IListServerGroupsRequest {
	s.UserAgent.AddUserAgent(agent...)
	return s
}

type CreateServerGroupRequest struct {
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
	PolicyID    string `json:"policyId,omitempty"`

	common.UserAgent
}

func (s *CreateServerGroupRequest) ToRequestBody() interface{} {
	return s
}

func (s *CreateServerGroupRequest) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"name":        s.Name,
		"description": s.Description,
		"policyId":    s.PolicyID,
	}
}

func (s *CreateServerGroupRequest) AddUserAgent(agent ...string) ICreateServerGroupRequest {
	s.UserAgent.AddUserAgent(agent...)
	return s
}
