package v1

import "github.com/dannyota/greennode-community-sdk/v2/greennode/client"

type IGetEndpointByIDRequest interface {
	GetEndpointID() string
	AddUserAgent(agent ...string) IGetEndpointByIDRequest
}

type ICreateEndpointRequest interface {
	ToRequestBody(svc client.ServiceClient) interface{}
	WithDescription(desp string) ICreateEndpointRequest
	WithSubnetUuid(subnetUuid string) ICreateEndpointRequest
	WithVpcUuid(vpcUuid string) ICreateEndpointRequest
	GetPortalUserID() string
	WithPortalUserID(portalUserID string) ICreateEndpointRequest
	WithPackageUuid(packageUuid string) ICreateEndpointRequest
	WithServiceUuid(serviceUuid string) ICreateEndpointRequest
	WithCategoryUuid(categoryUuid string) ICreateEndpointRequest
	WithEndpointName(endpointName string) ICreateEndpointRequest
	WithPoc(yes bool) ICreateEndpointRequest
	WithEnableDnsName(yes bool) ICreateEndpointRequest
	WithBuyMorePoc(yes bool) ICreateEndpointRequest
	WithEnableAutoRenew(yes bool) ICreateEndpointRequest
	AddNetworking(zone string, subnetUuid string) ICreateEndpointRequest
	WithScaling(minSize int, maxSize int) ICreateEndpointRequest
	AddUserAgent(agent ...string) ICreateEndpointRequest
	ToMap() map[string]interface{}
}

type IDeleteEndpointByIDRequest interface {
	GetEndpointID() string
	AddUserAgent(agent ...string) IDeleteEndpointByIDRequest
	ParseUserAgent() string
	ToRequestBody(svc client.ServiceClient) interface{}
	ToMap() map[string]interface{}
}

type IListEndpointsRequest interface {
	WithPage(page int) IListEndpointsRequest
	WithSize(size int) IListEndpointsRequest
	WithVpcID(vpcID string) IListEndpointsRequest
	WithUuid(uuid string) IListEndpointsRequest
	ToListQuery() (string, error)
	GetDefaultQuery() string
	AddUserAgent(agent ...string) IListEndpointsRequest
	ParseUserAgent() string
	ToMap() map[string]interface{}
}

type IListTagsByEndpointIDRequest interface {
	ToListQuery() (string, error)
	GetDefaultQuery() string
	ToMap() map[string]interface{}
	GetMapHeaders() map[string]string
	ParseUserAgent() string
	GetProjectID() string
	AddUserAgent(agent ...string) IListTagsByEndpointIDRequest
}

type ICreateTagsWithEndpointIDRequest interface {
	ToMap() map[string]interface{}
	AddUserAgent(agent ...string) ICreateTagsWithEndpointIDRequest
	GetMapHeaders() map[string]string
	AddTag(key, value string) ICreateTagsWithEndpointIDRequest
	ParseUserAgent() string
	GetProjectID() string
	ToRequestBody() interface{}
}

type IDeleteTagOfEndpointRequest interface {
	ToMap() map[string]interface{}
	AddUserAgent(agent ...string) IDeleteTagOfEndpointRequest
	GetMapHeaders() map[string]string
	ParseUserAgent() string
	GetTagID() string
	GetProjectID() string
}

type IUpdateTagValueOfEndpointRequest interface {
	ToMap() map[string]interface{}
	AddUserAgent(agent ...string) IUpdateTagValueOfEndpointRequest
	GetMapHeaders() map[string]string
	ParseUserAgent() string
	GetTagID() string
	GetProjectID() string
	ToRequestBody() interface{}
}
