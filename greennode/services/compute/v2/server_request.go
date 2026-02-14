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
	BackupInstancePointId  string                   `json:"backupInstancePointId,omitempty"`
	DataDiskEncryptionType DataDiskEncryptionType   `json:"dataDiskEncryptionType,omitempty"`
	DataDiskName           string                   `json:"dataDiskName,omitempty"`
	DataDiskSize           int                      `json:"dataDiskSize,omitempty"`
	DataDiskTypeId         string                   `json:"dataDiskTypeId,omitempty"`
	EnableBackup           bool                     `json:"enableBackup,omitempty"`
	EncryptionVolume       bool                     `json:"encryptionVolume"`
	ExpirePassword         bool                     `json:"expirePassword,omitempty"`
	FlavorId               string                   `json:"flavorId"`
	ImageId                string                   `json:"imageId"`
	Name                   string                   `json:"name"`
	NetworkId              string                   `json:"networkId,omitempty"`
	SubnetId               string                   `json:"subnetId,omitempty"`
	OsLicence              bool                     `json:"osLicence,omitempty"`
	RestoreBackup          bool                     `json:"restoreBackup,omitempty"`
	RootDiskEncryptionType DataDiskEncryptionType   `json:"rootDiskEncryptionType,omitempty"`
	RootDiskSize           int                      `json:"rootDiskSize"`
	RootDiskTypeId         string                   `json:"rootDiskTypeId"`
	SecurityGroup          []string                 `json:"securityGroup,omitempty"`
	ServerGroupId          string                   `json:"serverGroupId,omitempty"`
	SshKeyId               string                   `json:"sshKeyId,omitempty"`
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
	ProjectId      string `json:"projectId"`
	NetworkId      string `json:"networkId"`
	SubnetId       string `json:"subnetId"`
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

func (s *CreateServerRequest) WithZone(zoneId string) ICreateServerRequest {
	s.Zone = zoneId
	return s
}

func (s *CreateServerRequest) WithServerNetworkInterface(projectId, networkId, subnetId string, attachFloating bool) ICreateServerRequest {
	s.Networks = append(s.Networks, ServerNetworkInterface{
		ProjectId:      projectId,
		NetworkId:      networkId,
		SubnetId:       subnetId,
		AttachFloating: attachFloating,
	})

	return s.WithNetwork(s.Networks[0].NetworkId, s.Networks[0].SubnetId)
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

func (s *CreateServerRequest) WithServerGroupId(serverGroupId string) ICreateServerRequest {
	s.ServerGroupId = serverGroupId
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

func (s *CreateServerRequest) WithNetwork(networkId, subnetId string) ICreateServerRequest {
	s.NetworkId = networkId
	s.SubnetId = subnetId

	return s
}

func (s *CreateServerRequest) AddUserAgent(agent ...string) ICreateServerRequest {
	s.UserAgent.AddUserAgent(agent...)
	return s
}

func (s *CreateServerRequest) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"attachFloating":         s.AttachFloating,
		"backupInstancePointId":  s.BackupInstancePointId,
		"dataDiskEncryptionType": s.DataDiskEncryptionType,
		"dataDiskName":           s.DataDiskName,
		"dataDiskSize":           s.DataDiskSize,
		"dataDiskTypeId":         s.DataDiskTypeId,
		"enableBackup":           s.EnableBackup,
		"encryptionVolume":       s.EncryptionVolume,
		"expirePassword":         s.ExpirePassword,
		"flavorId":               s.FlavorId,
		"imageId":                s.ImageId,
		"name":                   s.Name,
		"networkId":              s.NetworkId,
		"subnetId":               s.SubnetId,
		"osLicence":              s.OsLicence,
		"restoreBackup":          s.RestoreBackup,
		"rootDiskEncryptionType": s.RootDiskEncryptionType,
		"rootDiskSize":           s.RootDiskSize,
		"rootDiskTypeId":         s.RootDiskTypeId,
		"securityGroup":          s.SecurityGroup,
		"serverGroupId":          s.ServerGroupId,
		"sshKeyId":               s.SshKeyId,
		"userName":               s.UserName,
		"isPoc":                  s.IsPoc,
		"product":                s.Product,
		"type":                   s.Type,
		"tags":                   s.Tags,
		"autoRenew":              s.AutoRenew,
		"networks":               s.Networks,
	}
}

type GetServerByIdRequest struct {
	common.ServerCommon
	common.UserAgent
}

func (s *GetServerByIdRequest) AddUserAgent(agent ...string) IGetServerByIdRequest {
	s.UserAgent.AddUserAgent(agent...)
	return s
}

func (s *GetServerByIdRequest) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"serverId": s.ServerId,
	}
}

type DeleteServerByIdRequest struct {
	DeleteAllVolume bool `json:"deleteAllVolume"`
	common.ServerCommon
	common.UserAgent
}

func (s *DeleteServerByIdRequest) WithDeleteAllVolume(ok bool) IDeleteServerByIdRequest {
	s.DeleteAllVolume = ok
	return s
}

func (s *DeleteServerByIdRequest) AddUserAgent(agent ...string) IDeleteServerByIdRequest {
	s.UserAgent.AddUserAgent(agent...)
	return s
}

func (s *DeleteServerByIdRequest) ToRequestBody() interface{} {
	return s
}

type UpdateServerSecgroupsByServerIdRequest struct {
	Secgroups []string `json:"securityGroup"`

	common.ServerCommon
	common.UserAgent
}

func (s *UpdateServerSecgroupsByServerIdRequest) AddUserAgent(agent ...string) IUpdateServerSecgroupsByServerIdRequest {
	s.UserAgent.AddUserAgent(agent...)
	return s
}

func (s *UpdateServerSecgroupsByServerIdRequest) ToRequestBody() interface{} {
	return s
}

func (s *UpdateServerSecgroupsByServerIdRequest) GetListSecgroupsIds() []string {
	return s.Secgroups
}

type AttachFloatingIpRequest struct {
	NetworkInterfaceId string `json:"networkInterfaceId"`

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
		"serverId":                   s.ServerId,
		"internalNetworkInterfaceId": s.InternalNetworkInterfaceId,
		"networkId":                  s.NetworkInterfaceId,
	}
}

type DetachFloatingIpRequest struct {
	NetworkInterfaceId string `json:"networkInterfaceId"`

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
		"serverId":                   s.ServerId,
		"internalNetworkInterfaceId": s.InternalNetworkInterfaceId,
		"networkId":                  s.NetworkInterfaceId,
		"wanId":                      s.WanId,
	}
}

type ListServerGroupPoliciesRequest struct {
	common.UserAgent
}

func (s *ListServerGroupPoliciesRequest) AddUserAgent(agent ...string) IListServerGroupPoliciesRequest {
	s.Agent = append(s.Agent, agent...)
	return s
}

type DeleteServerGroupByIdRequest struct {
	common.ServerGroupCommon
	common.UserAgent
}

func (s *DeleteServerGroupByIdRequest) AddUserAgent(agent ...string) IDeleteServerGroupByIdRequest {
	s.Agent = append(s.Agent, agent...)
	return s
}

func (s *DeleteServerGroupByIdRequest) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"serverGroupId": s.ServerGroupId,
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
	PolicyId    string `json:"policyId,omitempty"`

	common.UserAgent
}

func (s *CreateServerGroupRequest) ToRequestBody() interface{} {
	return s
}

func (s *CreateServerGroupRequest) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"name":        s.Name,
		"description": s.Description,
		"policyId":    s.PolicyId,
	}
}

func (s *CreateServerGroupRequest) AddUserAgent(agent ...string) ICreateServerGroupRequest {
	s.UserAgent.AddUserAgent(agent...)
	return s
}
