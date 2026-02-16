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

type GetServerByIDRequest struct {
	common.ServerCommon
}

type DeleteServerByIDRequest struct {
	DeleteAllVolume bool `json:"deleteAllVolume"`
	common.ServerCommon
}

func (r *DeleteServerByIDRequest) WithDeleteAllVolume(ok bool) *DeleteServerByIDRequest {
	r.DeleteAllVolume = ok
	return r
}

type UpdateServerSecgroupsByServerIDRequest struct {
	Secgroups []string `json:"securityGroup"`

	common.ServerCommon
}

func (r *UpdateServerSecgroupsByServerIDRequest) GetListSecgroupsIDs() []string {
	return r.Secgroups
}

type AttachFloatingIpRequest struct {
	NetworkInterfaceID string `json:"networkInterfaceId"`

	common.InternalNetworkInterfaceCommon
	common.ServerCommon
}

type DetachFloatingIpRequest struct {
	NetworkInterfaceID string `json:"networkInterfaceId"`

	common.ServerCommon
	common.WanCommon
	common.InternalNetworkInterfaceCommon
}

type ListServerGroupPoliciesRequest struct{}

type DeleteServerGroupByIDRequest struct {
	common.ServerGroupCommon
}

type ListServerGroupsRequest struct {
	Name string
	Page int
	Size int
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
		ServerCommon: common.ServerCommon{
			ServerID: serverID,
		},
	}
}

func NewDeleteServerByIDRequest(serverID string) *DeleteServerByIDRequest {
	return &DeleteServerByIDRequest{
		DeleteAllVolume: false,
		ServerCommon: common.ServerCommon{
			ServerID: serverID,
		},
	}
}

func NewUpdateServerSecgroupsRequest(serverID string, secgroups ...string) *UpdateServerSecgroupsByServerIDRequest {
	return &UpdateServerSecgroupsByServerIDRequest{
		Secgroups: secgroups,
		ServerCommon: common.ServerCommon{
			ServerID: serverID,
		},
	}
}

func NewAttachBlockVolumeRequest(serverID, volumeID string) *AttachBlockVolumeRequest {
	return &AttachBlockVolumeRequest{
		BlockVolumeCommon: common.BlockVolumeCommon{
			BlockVolumeID: volumeID,
		},
		ServerCommon: common.ServerCommon{
			ServerID: serverID,
		},
	}
}

func NewDetachBlockVolumeRequest(serverID, volumeID string) *DetachBlockVolumeRequest {
	return &DetachBlockVolumeRequest{
		BlockVolumeCommon: common.BlockVolumeCommon{
			BlockVolumeID: volumeID,
		},
		ServerCommon: common.ServerCommon{
			ServerID: serverID,
		},
	}
}

func NewAttachFloatingIpRequest(serverID, niid string) *AttachFloatingIpRequest {
	return &AttachFloatingIpRequest{
		NetworkInterfaceID: niid,
		InternalNetworkInterfaceCommon: common.InternalNetworkInterfaceCommon{
			InternalNetworkInterfaceID: niid,
		},
		ServerCommon: common.ServerCommon{
			ServerID: serverID,
		},
	}
}

func NewDetachFloatingIpRequest(serverID, wanID, niid string) *DetachFloatingIpRequest {
	return &DetachFloatingIpRequest{
		NetworkInterfaceID: niid,
		ServerCommon: common.ServerCommon{
			ServerID: serverID,
		},
		WanCommon: common.WanCommon{
			WanID: wanID,
		},
		InternalNetworkInterfaceCommon: common.InternalNetworkInterfaceCommon{
			InternalNetworkInterfaceID: niid,
		},
	}
}

func NewListServerGroupPoliciesRequest() *ListServerGroupPoliciesRequest {
	return new(ListServerGroupPoliciesRequest)
}

func NewDeleteServerGroupByIDRequest(serverGroupID string) *DeleteServerGroupByIDRequest {
	return &DeleteServerGroupByIDRequest{
		ServerGroupCommon: common.ServerGroupCommon{
			ServerGroupID: serverGroupID,
		},
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
