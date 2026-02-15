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

func (r *CreateServerRequest) WithZone(zoneID string) *CreateServerRequest {
	r.Zone = zoneID
	return r
}

func (r *CreateServerRequest) WithServerNetworkInterface(projectID, networkID, subnetID string, attachFloating bool) *CreateServerRequest {
	r.Networks = append(r.Networks, ServerNetworkInterface{
		ProjectID:      projectID,
		NetworkID:      networkID,
		SubnetID:       subnetID,
		AttachFloating: attachFloating,
	})

	return r.WithNetwork(r.Networks[0].NetworkID, r.Networks[0].SubnetID)
}

func (r *CreateServerRequest) WithRootDiskEncryptionType(dataDisk DataDiskEncryptionType) *CreateServerRequest {
	r.EncryptionVolume = true
	r.RootDiskEncryptionType = dataDisk
	return r
}

func (r *CreateServerRequest) WithEncryptionVolume(enabled bool) *CreateServerRequest {
	r.EncryptionVolume = enabled
	return r
}

func (r *CreateServerRequest) WithUserData(userData string, base64Encode bool) *CreateServerRequest {
	r.UserData = userData
	r.UserDataBase64Encoded = base64Encode
	return r
}

func (r *CreateServerRequest) WithAutoRenew(val bool) *CreateServerRequest {
	r.AutoRenew = val
	return r
}

func (r *CreateServerRequest) WithTags(tags ...string) *CreateServerRequest {
	if r.Tags == nil {
		r.Tags = make([]ServerTag, 0)
	}

	if len(tags)%2 != 0 {
		tags = append(tags, "none")
	}

	for i := 0; i < len(tags); i += 2 {
		r.Tags = append(r.Tags, ServerTag{Key: tags[i], Value: tags[i+1]})
	}
	return r
}

func (r *CreateServerRequest) WithAttachFloating(attachFloating bool) *CreateServerRequest {
	r.AttachFloating = attachFloating
	return r
}

func (r *CreateServerRequest) WithSecgroups(secgroups ...string) *CreateServerRequest {
	r.SecurityGroup = append(r.SecurityGroup, secgroups...)
	return r
}

func (r *CreateServerRequest) WithServerGroupID(serverGroupID string) *CreateServerRequest {
	r.ServerGroupID = serverGroupID
	return r
}

func (r *CreateServerRequest) WithPoc(isPoc bool) *CreateServerRequest {
	r.IsPoc = isPoc
	return r
}

func (r *CreateServerRequest) WithType(typeVal string) *CreateServerRequest {
	r.Type = typeVal
	return r
}

func (r *CreateServerRequest) WithProduct(product string) *CreateServerRequest {
	r.Product = product
	return r
}

func (r *CreateServerRequest) WithNetwork(networkID, subnetID string) *CreateServerRequest {
	r.NetworkID = networkID
	r.SubnetID = subnetID

	return r
}

func (r *CreateServerRequest) AddUserAgent(agent ...string) *CreateServerRequest {
	r.UserAgent.AddUserAgent(agent...)
	return r
}

func (r *CreateServerRequest) ToMap() map[string]any {
	return map[string]any{
		"attachFloating":         r.AttachFloating,
		"backupInstancePointId":  r.BackupInstancePointID,
		"dataDiskEncryptionType": r.DataDiskEncryptionType,
		"dataDiskName":           r.DataDiskName,
		"dataDiskSize":           r.DataDiskSize,
		"dataDiskTypeId":         r.DataDiskTypeID,
		"enableBackup":           r.EnableBackup,
		"encryptionVolume":       r.EncryptionVolume,
		"expirePassword":         r.ExpirePassword,
		"flavorId":               r.FlavorID,
		"imageId":                r.ImageID,
		"name":                   r.Name,
		"networkId":              r.NetworkID,
		"subnetId":               r.SubnetID,
		"osLicence":              r.OsLicence,
		"restoreBackup":          r.RestoreBackup,
		"rootDiskEncryptionType": r.RootDiskEncryptionType,
		"rootDiskSize":           r.RootDiskSize,
		"rootDiskTypeId":         r.RootDiskTypeID,
		"securityGroup":          r.SecurityGroup,
		"serverGroupId":          r.ServerGroupID,
		"sshKeyId":               r.SshKeyID,
		"userName":               r.UserName,
		"isPoc":                  r.IsPoc,
		"product":                r.Product,
		"type":                   r.Type,
		"tags":                   r.Tags,
		"autoRenew":              r.AutoRenew,
		"networks":               r.Networks,
	}
}

type GetServerByIDRequest struct {
	common.ServerCommon
	common.UserAgent
}

func (r *GetServerByIDRequest) AddUserAgent(agent ...string) *GetServerByIDRequest {
	r.UserAgent.AddUserAgent(agent...)
	return r
}

func (r *GetServerByIDRequest) ToMap() map[string]any {
	return map[string]any{
		"serverId": r.ServerID,
	}
}

type DeleteServerByIDRequest struct {
	DeleteAllVolume bool `json:"deleteAllVolume"`
	common.ServerCommon
	common.UserAgent
}

func (r *DeleteServerByIDRequest) WithDeleteAllVolume(ok bool) *DeleteServerByIDRequest {
	r.DeleteAllVolume = ok
	return r
}

func (r *DeleteServerByIDRequest) AddUserAgent(agent ...string) *DeleteServerByIDRequest {
	r.UserAgent.AddUserAgent(agent...)
	return r
}

type UpdateServerSecgroupsByServerIDRequest struct {
	Secgroups []string `json:"securityGroup"`

	common.ServerCommon
	common.UserAgent
}

func (r *UpdateServerSecgroupsByServerIDRequest) AddUserAgent(agent ...string) *UpdateServerSecgroupsByServerIDRequest {
	r.UserAgent.AddUserAgent(agent...)
	return r
}

func (r *UpdateServerSecgroupsByServerIDRequest) GetListSecgroupsIDs() []string {
	return r.Secgroups
}

type AttachFloatingIpRequest struct {
	NetworkInterfaceID string `json:"networkInterfaceId"`

	common.InternalNetworkInterfaceCommon
	common.ServerCommon
	common.UserAgent
}

func (r *AttachFloatingIpRequest) AddUserAgent(agent ...string) *AttachFloatingIpRequest {
	r.Agent = append(r.Agent, agent...)
	return r
}

func (r *AttachFloatingIpRequest) ToMap() map[string]any {
	return map[string]any{
		"serverId":                   r.ServerID,
		"internalNetworkInterfaceId": r.InternalNetworkInterfaceID,
		"networkId":                  r.NetworkInterfaceID,
	}
}

type DetachFloatingIpRequest struct {
	NetworkInterfaceID string `json:"networkInterfaceId"`

	common.ServerCommon
	common.WanCommon
	common.InternalNetworkInterfaceCommon
	common.UserAgent
}

func (r *DetachFloatingIpRequest) AddUserAgent(agent ...string) *DetachFloatingIpRequest {
	r.Agent = append(r.Agent, agent...)
	return r
}

func (r *DetachFloatingIpRequest) ToMap() map[string]any {
	return map[string]any{
		"serverId":                   r.ServerID,
		"internalNetworkInterfaceId": r.InternalNetworkInterfaceID,
		"networkId":                  r.NetworkInterfaceID,
		"wanId":                      r.WanID,
	}
}

type ListServerGroupPoliciesRequest struct {
	common.UserAgent
}

func (r *ListServerGroupPoliciesRequest) AddUserAgent(agent ...string) *ListServerGroupPoliciesRequest {
	r.Agent = append(r.Agent, agent...)
	return r
}

type DeleteServerGroupByIDRequest struct {
	common.ServerGroupCommon
	common.UserAgent
}

func (r *DeleteServerGroupByIDRequest) AddUserAgent(agent ...string) *DeleteServerGroupByIDRequest {
	r.Agent = append(r.Agent, agent...)
	return r
}

func (r *DeleteServerGroupByIDRequest) ToMap() map[string]any {
	return map[string]any{
		"serverGroupId": r.ServerGroupID,
	}
}

type ListServerGroupsRequest struct {
	Name string
	Page int
	Size int

	common.UserAgent
}

func (r *ListServerGroupsRequest) WithName(name string) *ListServerGroupsRequest {
	r.Name = name
	return r
}

func (r *ListServerGroupsRequest) ToListQuery() (string, error) {
	v := url.Values{}
	v.Set("name", r.Name)
	if r.Page > 0 {
		v.Set("page", strconv.Itoa(r.Page))
	}
	if r.Size > 0 {
		v.Set("size", strconv.Itoa(r.Size))
	}
	return v.Encode(), nil
}

func (r *ListServerGroupsRequest) GetDefaultQuery() string {
	return fmt.Sprintf("offset=%d&limit=%d&name=", defaultOffsetListServerGroups, defaultLimitListServerGroups)
}

func (r *ListServerGroupsRequest) ToMap() map[string]any {
	return map[string]any{
		"page": r.Page,
		"size": r.Size,
		"name": r.Name,
	}
}

func (r *ListServerGroupsRequest) AddUserAgent(agent ...string) *ListServerGroupsRequest {
	r.UserAgent.AddUserAgent(agent...)
	return r
}

type CreateServerGroupRequest struct {
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
	PolicyID    string `json:"policyId,omitempty"`

	common.UserAgent
}

func (r *CreateServerGroupRequest) ToMap() map[string]any {
	return map[string]any{
		"name":        r.Name,
		"description": r.Description,
		"policyId":    r.PolicyID,
	}
}

func (r *CreateServerGroupRequest) AddUserAgent(agent ...string) *CreateServerGroupRequest {
	r.UserAgent.AddUserAgent(agent...)
	return r
}
