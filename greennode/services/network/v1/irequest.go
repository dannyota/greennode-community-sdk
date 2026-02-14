package v1

import "github.com/dannyota/greennode-community-sdk/v2/greennode/client"

type IGetEndpointByIdRequest interface {
	GetEndpointId() string
	AddUserAgent(agent ...string) IGetEndpointByIdRequest
}

type ICreateEndpointRequest interface {
	ToRequestBody(svc client.ServiceClient) interface{}
	WithDescription(desp string) ICreateEndpointRequest
	WithSubnetUuid(subnetUuid string) ICreateEndpointRequest
	WithVpcUuid(vpcUuid string) ICreateEndpointRequest
	GetPortalUserId() string
	WithPortalUserId(portalUserId string) ICreateEndpointRequest
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

type IDeleteEndpointByIdRequest interface {
	GetEndpointId() string
	AddUserAgent(agent ...string) IDeleteEndpointByIdRequest
	ParseUserAgent() string
	ToRequestBody(svc client.ServiceClient) interface{}
	ToMap() map[string]interface{}
}

type IListEndpointsRequest interface {
	WithPage(page int) IListEndpointsRequest
	WithSize(size int) IListEndpointsRequest
	WithVpcId(vpcId string) IListEndpointsRequest
	WithUuid(uuid string) IListEndpointsRequest
	ToListQuery() (string, error)
	GetDefaultQuery() string
	AddUserAgent(agent ...string) IListEndpointsRequest
	ParseUserAgent() string
	ToMap() map[string]interface{}
}

type IListTagsByEndpointIdRequest interface {
	ToListQuery() (string, error)
	GetDefaultQuery() string
	ToMap() map[string]interface{}
	GetMapHeaders() map[string]string
	ParseUserAgent() string
	GetProjectId() string
	AddUserAgent(agent ...string) IListTagsByEndpointIdRequest
}

type ICreateTagsWithEndpointIdRequest interface {
	ToMap() map[string]interface{}
	AddUserAgent(agent ...string) ICreateTagsWithEndpointIdRequest
	GetMapHeaders() map[string]string
	AddTag(key, value string) ICreateTagsWithEndpointIdRequest
	ParseUserAgent() string
	GetProjectId() string
	ToRequestBody() interface{}
}

type IDeleteTagOfEndpointRequest interface {
	ToMap() map[string]interface{}
	AddUserAgent(agent ...string) IDeleteTagOfEndpointRequest
	GetMapHeaders() map[string]string
	ParseUserAgent() string
	GetTagId() string
	GetProjectId() string
}

type IUpdateTagValueOfEndpointRequest interface {
	ToMap() map[string]interface{}
	AddUserAgent(agent ...string) IUpdateTagValueOfEndpointRequest
	GetMapHeaders() map[string]string
	ParseUserAgent() string
	GetTagId() string
	GetProjectId() string
	ToRequestBody() interface{}
}
