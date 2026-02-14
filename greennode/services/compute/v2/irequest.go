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
	WithServerGroupId(serverGroupId string) ICreateServerRequest
	WithPoc(isPoc bool) ICreateServerRequest
	WithType(typeVal string) ICreateServerRequest
	WithNetwork(networkId, subnetId string) ICreateServerRequest
	WithProduct(product string) ICreateServerRequest
	WithServerNetworkInterface(projectId, networkId, subnetId string, attachFloating bool) ICreateServerRequest
	WithZone(zone string) ICreateServerRequest
	AddUserAgent(agent ...string) ICreateServerRequest
	ParseUserAgent() string
	ToMap() map[string]interface{}
}

type IGetServerByIdRequest interface {
	GetServerId() string
	AddUserAgent(agent ...string) IGetServerByIdRequest
	ParseUserAgent() string
	ToMap() map[string]interface{}
}

type IDeleteServerByIdRequest interface {
	GetServerId() string
	WithDeleteAllVolume(ok bool) IDeleteServerByIdRequest
	ToRequestBody() interface{}
	AddUserAgent(agent ...string) IDeleteServerByIdRequest
	ParseUserAgent() string
}

type IUpdateServerSecgroupsByServerIdRequest interface {
	GetServerId() string
	ToRequestBody() interface{}
	GetListSecgroupsIds() []string
	AddUserAgent(agent ...string) IUpdateServerSecgroupsByServerIdRequest
	ParseUserAgent() string
}

type IAttachBlockVolumeRequest interface {
	GetServerId() string
	GetBlockVolumeId() string
}

type IDetachBlockVolumeRequest interface {
	GetServerId() string
	GetBlockVolumeId() string
}

type IAttachFloatingIpRequest interface {
	GetServerId() string
	GetInternalNetworkInterfaceId() string
	ToRequestBody() interface{}
	AddUserAgent(agent ...string) IAttachFloatingIpRequest
	ParseUserAgent() string
	ToMap() map[string]interface{}
}

type IDetachFloatingIpRequest interface {
	GetInternalNetworkInterfaceId() string
	GetWanId() string
	GetServerId() string
	ToRequestBody() interface{}
	AddUserAgent(agent ...string) IDetachFloatingIpRequest
	ParseUserAgent() string
	ToMap() map[string]interface{}
}

type IListServerGroupPoliciesRequest interface {
	AddUserAgent(agent ...string) IListServerGroupPoliciesRequest
	ParseUserAgent() string
}

type IDeleteServerGroupByIdRequest interface {
	AddUserAgent(agent ...string) IDeleteServerGroupByIdRequest
	ParseUserAgent() string
	GetServerGroupId() string
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
