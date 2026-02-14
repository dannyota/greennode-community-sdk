package v2

// Secgroup

type IGetSecgroupByIDRequest interface {
	AddUserAgent(agent ...string) IGetSecgroupByIDRequest
	ParseUserAgent() string
	GetSecgroupID() string
}

type ICreateSecgroupRequest interface {
	ToRequestBody() interface{}
	GetSecgroupName() string
	AddUserAgent(agent ...string) ICreateSecgroupRequest
	ParseUserAgent() string
}

type IDeleteSecgroupByIDRequest interface {
	GetSecgroupID() string
	AddUserAgent(agent ...string) IDeleteSecgroupByIDRequest
	ParseUserAgent() string
}

type IListSecgroupRequest interface {
	AddUserAgent(agent ...string) IListSecgroupRequest
	ParseUserAgent() string
}

// Secgroup Rule
type ICreateSecgroupRuleRequest interface {
	GetSecgroupID() string
	ToRequestBody() interface{}
	AddUserAgent(agent ...string) ICreateSecgroupRuleRequest
	ParseUserAgent() string
	ToMap() map[string]interface{}
}

type IDeleteSecgroupRuleByIDRequest interface {
	GetSecgroupID() string
	GetSecgroupRuleID() string
	AddUserAgent(agent ...string) IDeleteSecgroupRuleByIDRequest
	ParseUserAgent() string
}

type IListSecgroupRulesBySecgroupIDRequest interface {
	GetSecgroupID() string
	AddUserAgent(agent ...string) IListSecgroupRulesBySecgroupIDRequest
	ParseUserAgent() string
}

// Network

type IGetNetworkByIDRequest interface {
	GetNetworkID() string
	AddUserAgent(agent ...string) IGetNetworkByIDRequest
	ParseUserAgent() string
}

// Subnet

type IGetSubnetByIDRequest interface {
	AddUserAgent(agent ...string) IGetSubnetByIDRequest
	ParseUserAgent() string
	GetNetworkID() string
	GetSubnetID() string
}

type IUpdateSubnetByIDRequest interface {
	AddUserAgent(agent ...string) IUpdateSubnetByIDRequest
	ParseUserAgent() string
	GetNetworkID() string
	GetSubnetID() string
	ToRequestBody() interface{}
}

/**
 * The interface request group of Address Pair API
 */

type IGetAllAddressPairByVirtualSubnetIDRequest interface {
	GetVirtualSubnetID() string
	ParseUserAgent() string
}

type ISetAddressPairInVirtualSubnetRequest interface {
	GetVirtualSubnetID() string
	ParseUserAgent() string
	ToRequestBody() interface{}
}

type IDeleteAddressPairRequest interface {
	ParseUserAgent() string
	GetAddressPairID() string
	AddUserAgent(agent ...string) IDeleteAddressPairRequest
}

type IListAllServersBySecgroupIDRequest interface {
	GetSecgroupID() string
	ParseUserAgent() string
	AddUserAgent(agent ...string) IListAllServersBySecgroupIDRequest
}

type ICreateAddressPairRequest interface {
	GetVirtualAddressID() string
	ToRequestBody() interface{}
	ParseUserAgent() string
	ToMap() map[string]interface{}
	AddUserAgent(agent ...string) ICreateAddressPairRequest
	WithMode(mode AddressPairMode) ICreateAddressPairRequest
}

/**
 * The interface request group of Virtual Address API
 */

// Request interface for creating virtual address cross project
type ICreateVirtualAddressCrossProjectRequest interface {
	ToRequestBody() interface{}
	ParseUserAgent() string
	ToMap() map[string]interface{}
	AddUserAgent(agent ...string) ICreateVirtualAddressCrossProjectRequest
	WithDescription(description string) ICreateVirtualAddressCrossProjectRequest
}

// Request interface for deleting virtual address by ID
type IDeleteVirtualAddressByIDRequest interface {
	GetVirtualAddressID() string
	ParseUserAgent() string
	AddUserAgent(agent ...string) IDeleteVirtualAddressByIDRequest
	ToMap() map[string]interface{}
}

// Request interface for getting virtual address by ID
type IGetVirtualAddressByIDRequest interface {
	GetVirtualAddressID() string
	ParseUserAgent() string
	AddUserAgent(agent ...string) IGetVirtualAddressByIDRequest
	ToMap() map[string]interface{}
}

// Request interface for listing address pairs of virtual address by ID

type IListAddressPairsByVirtualAddressIDRequest interface {
	GetVirtualAddressID() string
	ParseUserAgent() string
	AddUserAgent(agent ...string) IListAddressPairsByVirtualAddressIDRequest
	ToMap() map[string]interface{}
}
