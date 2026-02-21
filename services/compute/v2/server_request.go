package v2

import (
	"fmt"
	"net/url"
	"strconv"
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
}

type ServerNetworkInterface struct {
	ProjectID      string `json:"projectId"`
	NetworkID      string `json:"networkId"`
	SubnetID       string `json:"subnetId"`
	AttachFloating bool   `json:"attachFloating"`
}

type AttachBlockVolumeRequest struct {
	BlockVolumeID string
	ServerID      string
}

type DetachBlockVolumeRequest struct {
	BlockVolumeID string
	ServerID      string
}

type DataDiskEncryptionType string

type ServerTag struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

// NewServerTags creates a []ServerTag from variadic key-value string pairs.
// If an odd number of strings is provided, the last value defaults to "none".
func NewServerTags(kvPairs ...string) []ServerTag {
	if len(kvPairs)%2 != 0 {
		kvPairs = append(kvPairs, "none")
	}
	tags := make([]ServerTag, 0, len(kvPairs)/2)
	for i := 0; i < len(kvPairs); i += 2 {
		tags = append(tags, ServerTag{Key: kvPairs[i], Value: kvPairs[i+1]})
	}
	return tags
}

type GetServerByIDRequest struct {
	ServerID string
}

type DeleteServerByIDRequest struct {
	DeleteAllVolume bool `json:"deleteAllVolume"`
	ServerID        string
}

type UpdateServerSecgroupsByServerIDRequest struct {
	Secgroups []string `json:"securityGroup"`
	ServerID  string
}

type AttachFloatingIpRequest struct {
	NetworkInterfaceID            string `json:"networkInterfaceId"`
	InternalNetworkInterfaceID    string
	ServerID                      string
}

type DetachFloatingIpRequest struct {
	NetworkInterfaceID            string `json:"networkInterfaceId"`
	ServerID                      string
	WanID                         string
	InternalNetworkInterfaceID    string
}

type ListServerGroupPoliciesRequest struct{}

type DeleteServerGroupByIDRequest struct {
	ServerGroupID string
}

type ListServerGroupsRequest struct {
	Name string
	Page int
	Size int
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

func (r *ListServerGroupsRequest) getDefaultQuery() string {
	return fmt.Sprintf("offset=%d&limit=%d&name=", defaultOffsetListServerGroups, defaultLimitListServerGroups)
}

type CreateServerGroupRequest struct {
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
	PolicyID    string `json:"policyId,omitempty"`
}
func NewCreateServerRequest(name, imageID, flavorID, rootDiskType string, rootDiskSize int) *CreateServerRequest {
	return &CreateServerRequest{
		Name:           name,
		ImageID:        imageID,
		FlavorID:       flavorID,
		RootDiskTypeID: rootDiskType,
		RootDiskSize:   rootDiskSize,
	}
}

func NewGetServerByIDRequest(serverID string) *GetServerByIDRequest {
	return &GetServerByIDRequest{
		ServerID: serverID,
	}
}

func NewDeleteServerByIDRequest(serverID string) *DeleteServerByIDRequest {
	return &DeleteServerByIDRequest{
		DeleteAllVolume: false,
		ServerID:        serverID,
	}
}

func NewUpdateServerSecgroupsRequest(serverID string, secgroups ...string) *UpdateServerSecgroupsByServerIDRequest {
	return &UpdateServerSecgroupsByServerIDRequest{
		Secgroups: secgroups,
		ServerID:  serverID,
	}
}

func NewAttachBlockVolumeRequest(serverID, volumeID string) *AttachBlockVolumeRequest {
	return &AttachBlockVolumeRequest{
		BlockVolumeID: volumeID,
		ServerID:      serverID,
	}
}

func NewDetachBlockVolumeRequest(serverID, volumeID string) *DetachBlockVolumeRequest {
	return &DetachBlockVolumeRequest{
		BlockVolumeID: volumeID,
		ServerID:      serverID,
	}
}

func NewAttachFloatingIpRequest(serverID, niid string) *AttachFloatingIpRequest {
	return &AttachFloatingIpRequest{
		NetworkInterfaceID:         niid,
		InternalNetworkInterfaceID: niid,
		ServerID:                   serverID,
	}
}

func NewDetachFloatingIpRequest(serverID, wanID, niid string) *DetachFloatingIpRequest {
	return &DetachFloatingIpRequest{
		NetworkInterfaceID:         niid,
		ServerID:                   serverID,
		WanID:                      wanID,
		InternalNetworkInterfaceID: niid,
	}
}

func NewListServerGroupPoliciesRequest() *ListServerGroupPoliciesRequest {
	return new(ListServerGroupPoliciesRequest)
}

func NewDeleteServerGroupByIDRequest(serverGroupID string) *DeleteServerGroupByIDRequest {
	return &DeleteServerGroupByIDRequest{
		ServerGroupID: serverGroupID,
	}
}

func NewListServerGroupsRequest(page, size int) *ListServerGroupsRequest {
	return &ListServerGroupsRequest{
		Page: page,
		Size: size,
		Name: "",
	}
}

func NewCreateServerGroupRequest(name, description, policyID string) *CreateServerGroupRequest {
	return &CreateServerGroupRequest{
		Name:        name,
		Description: description,
		PolicyID:    policyID,
	}
}
