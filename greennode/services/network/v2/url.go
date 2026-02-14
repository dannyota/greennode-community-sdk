package v2

import "github.com/dannyota/greennode-community-sdk/v2/greennode/client"

func getSecgroupByIdUrl(psc client.IServiceClient, popts IGetSecgroupByIdRequest) string {
	return psc.ServiceURL(
		psc.GetProjectId(),
		"secgroups",
		popts.GetSecgroupId())
}

func createSecgroupUrl(psc client.IServiceClient) string {
	return psc.ServiceURL(
		psc.GetProjectId(),
		"secgroups")
}

func listSecgroupUrl(psc client.IServiceClient, _ IListSecgroupRequest) string {
	return psc.ServiceURL(
		psc.GetProjectId(),
		"secgroups")
}

func deleteSecgroupByIdUrl(psc client.IServiceClient, popts IDeleteSecgroupByIdRequest) string {
	return psc.ServiceURL(
		psc.GetProjectId(),
		"secgroups",
		popts.GetSecgroupId())
}

func createSecgroupRuleUrl(psc client.IServiceClient, popts ICreateSecgroupRuleRequest) string {
	return psc.ServiceURL(
		psc.GetProjectId(),
		"secgroups",
		popts.GetSecgroupId(),
		"secgroupRules")
}

func deleteSecgroupRuleByIdUrl(psc client.IServiceClient, popts IDeleteSecgroupRuleByIdRequest) string {
	return psc.ServiceURL(
		psc.GetProjectId(),
		"secgroups",
		popts.GetSecgroupId(),
		"secgroupRules",
		popts.GetSecgroupRuleId())
}

func listSecgroupRulesBySecgroupIdUrl(psc client.IServiceClient, popts IListSecgroupRulesBySecgroupIdRequest) string {
	return psc.ServiceURL(
		psc.GetProjectId(),
		"secgroups",
		popts.GetSecgroupId(),
		"secGroupRules")
}

func getNetworkByIdUrl(psc client.IServiceClient, popts IGetNetworkByIdRequest) string {
	return psc.ServiceURL(
		psc.GetProjectId(),
		"networks",
		popts.GetNetworkId())
}

func getSubnetByIdUrl(psc client.IServiceClient, popts IGetSubnetByIdRequest) string {
	return psc.ServiceURL(
		psc.GetProjectId(),
		"networks",
		popts.GetNetworkId(),
		"subnets",
		popts.GetSubnetId())
}

func updateSubnetByIdUrl(psc client.IServiceClient, popts IUpdateSubnetByIdRequest) string {
	return psc.ServiceURL(
		psc.GetProjectId(),
		"networks",
		popts.GetNetworkId(),
		"subnets",
		popts.GetSubnetId())
}

func getAllAddressPairByVirtualSubnetIdUrl(psc client.IServiceClient, popts IGetAllAddressPairByVirtualSubnetIdRequest) string {
	return psc.ServiceURL(
		psc.GetProjectId(),
		"virtual-subnets",
		popts.GetVirtualSubnetId(),
		"addressPairs")
}

func setAddressPairInVirtualSubnetUrl(psc client.IServiceClient, popts IGetAllAddressPairByVirtualSubnetIdRequest) string {
	return psc.ServiceURL(
		psc.GetProjectId(),
		"virtual-subnets",
		popts.GetVirtualSubnetId(),
		"addressPairs")
}

func deleteAddressPairUrl(psc client.IServiceClient, popts IDeleteAddressPairRequest) string {
	return psc.ServiceURL(
		psc.GetProjectId(),
		"virtual-subnets",
		"addressPairs",
		popts.GetAddressPairID())
}

func createAddressPairUrl(psc client.IServiceClient, popts ICreateAddressPairRequest) string {
	return psc.ServiceURL(
		psc.GetProjectId(),
		"virtualIpAddress",
		popts.GetVirtualAddressId(),
		"addressPairs")
}

func listAllServersBySecgroupIdUrl(psc client.IServiceClient, popts IListAllServersBySecgroupIdRequest) string {
	return psc.ServiceURL(
		psc.GetProjectId(),
		"secgroups",
		popts.GetSecgroupId(),
		"servers")
}

func createVirtualAddressCrossProjectUrl(psc client.IServiceClient) string {
	return psc.ServiceURL(
		psc.GetProjectId(),
		"virtualIpAddress")
}

func deleteVirtualAddressByIdUrl(psc client.IServiceClient, popts IDeleteVirtualAddressByIdRequest) string {
	return psc.ServiceURL(
		psc.GetProjectId(),
		"virtualIpAddress",
		popts.GetVirtualAddressId())
}

func getVirtualAddressByIdUrl(psc client.IServiceClient, popts IGetVirtualAddressByIdRequest) string {
	return psc.ServiceURL(
		psc.GetProjectId(),
		"virtualIpAddress",
		popts.GetVirtualAddressId())
}

func listAddressPairsByVirtualAddressIdUrl(psc client.IServiceClient, popts IListAddressPairsByVirtualAddressIdRequest) string {
	return psc.ServiceURL(
		psc.GetProjectId(),
		"virtualIpAddress",
		popts.GetVirtualAddressId(),
		"addressPairs")
}
