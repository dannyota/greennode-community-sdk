package v2

type ICreateServerRequest interface {
	ToRequestBody() interface{}
	WithRootDiskEncryptionType(encryptionVolume DataDiskEncryptionType) ICreateServerRequest
	WithEncryptionVolume(enabled bool) ICreateServerRequest
	WithAutoRenew(val bool) ICreateServerRequest
	WithUserData(userData string, base64Encode bool) ICreateServerRequest
	WithTags(tags ...string) ICreateServerRequest
	WithAttachFloating(attachFloating bool) ICreateServerRequest
	WithSecgroups(secgroups ...string) ICreateServerRequest
	WithServerGroupID(serverGroupID string) ICreateServerRequest
	WithPoc(isPoc bool) ICreateServerRequest
	WithType(typeVal string) ICreateServerRequest
	WithNetwork(networkID, subnetID string) ICreateServerRequest
	WithProduct(product string) ICreateServerRequest
	WithServerNetworkInterface(projectID, networkID, subnetID string, attachFloating bool) ICreateServerRequest
	WithZone(zone string) ICreateServerRequest
	AddUserAgent(agent ...string) ICreateServerRequest
	ParseUserAgent() string
	ToMap() map[string]interface{}
}

type IGetServerByIDRequest interface {
	GetServerID() string
	AddUserAgent(agent ...string) IGetServerByIDRequest
	ParseUserAgent() string
	ToMap() map[string]interface{}
}

type IDeleteServerByIDRequest interface {
	GetServerID() string
	WithDeleteAllVolume(ok bool) IDeleteServerByIDRequest
	ToRequestBody() interface{}
	AddUserAgent(agent ...string) IDeleteServerByIDRequest
	ParseUserAgent() string
}

type IUpdateServerSecgroupsByServerIDRequest interface {
	GetServerID() string
	ToRequestBody() interface{}
	GetListSecgroupsIDs() []string
	AddUserAgent(agent ...string) IUpdateServerSecgroupsByServerIDRequest
	ParseUserAgent() string
}

type IAttachBlockVolumeRequest interface {
	GetServerID() string
	GetBlockVolumeID() string
}

type IDetachBlockVolumeRequest interface {
	GetServerID() string
	GetBlockVolumeID() string
}

type IAttachFloatingIpRequest interface {
	GetServerID() string
	GetInternalNetworkInterfaceID() string
	ToRequestBody() interface{}
	AddUserAgent(agent ...string) IAttachFloatingIpRequest
	ParseUserAgent() string
	ToMap() map[string]interface{}
}

type IDetachFloatingIpRequest interface {
	GetInternalNetworkInterfaceID() string
	GetWanID() string
	GetServerID() string
	ToRequestBody() interface{}
	AddUserAgent(agent ...string) IDetachFloatingIpRequest
	ParseUserAgent() string
	ToMap() map[string]interface{}
}

type IListServerGroupPoliciesRequest interface {
	AddUserAgent(agent ...string) IListServerGroupPoliciesRequest
	ParseUserAgent() string
}

type IDeleteServerGroupByIDRequest interface {
	AddUserAgent(agent ...string) IDeleteServerGroupByIDRequest
	ParseUserAgent() string
	GetServerGroupID() string
	ToMap() map[string]interface{}
}

type IListServerGroupsRequest interface {
	WithName(name string) IListServerGroupsRequest
	AddUserAgent(agent ...string) IListServerGroupsRequest
	ToListQuery() (string, error)
	ParseUserAgent() string
	GetDefaultQuery() string
	ToMap() map[string]interface{}
}

type ICreateServerGroupRequest interface {
	ParseUserAgent() string
	AddUserAgent(agent ...string) ICreateServerGroupRequest
	ToRequestBody() interface{}
	ToMap() map[string]interface{}
}
