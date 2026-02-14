package v2

import "github.com/dannyota/greennode-community-sdk/v2/greennode/client"

func getSecgroupByIdUrl(psc client.ServiceClient, popts IGetSecgroupByIdRequest) string {
	return psc.ServiceURL(
		psc.GetProjectId(),
		"secgroups",
		popts.GetSecgroupId())
}

func createSecgroupUrl(psc client.ServiceClient) string {
	return psc.ServiceURL(
		psc.GetProjectId(),
		"secgroups")
}

func listSecgroupUrl(psc client.ServiceClient, _ IListSecgroupRequest) string {
	return psc.ServiceURL(
		psc.GetProjectId(),
		"secgroups")
}

func deleteSecgroupByIdUrl(psc client.ServiceClient, popts IDeleteSecgroupByIdRequest) string {
	return psc.ServiceURL(
		psc.GetProjectId(),
		"secgroups",
		popts.GetSecgroupId())
}

func createSecgroupRuleUrl(psc client.ServiceClient, popts ICreateSecgroupRuleRequest) string {
	return psc.ServiceURL(
		psc.GetProjectId(),
		"secgroups",
		popts.GetSecgroupId(),
		"secgroupRules")
}

func deleteSecgroupRuleByIdUrl(psc client.ServiceClient, popts IDeleteSecgroupRuleByIdRequest) string {
	return psc.ServiceURL(
		psc.GetProjectId(),
		"secgroups",
		popts.GetSecgroupId(),
		"secgroupRules",
		popts.GetSecgroupRuleId())
}

func listSecgroupRulesBySecgroupIdUrl(psc client.ServiceClient, popts IListSecgroupRulesBySecgroupIdRequest) string {
	return psc.ServiceURL(
		psc.GetProjectId(),
		"secgroups",
		popts.GetSecgroupId(),
		"secGroupRules")
}

func getNetworkByIdUrl(psc client.ServiceClient, popts IGetNetworkByIdRequest) string {
	return psc.ServiceURL(
		psc.GetProjectId(),
		"networks",
		popts.GetNetworkId())
}

func getSubnetByIdUrl(psc client.ServiceClient, popts IGetSubnetByIdRequest) string {
	return psc.ServiceURL(
		psc.GetProjectId(),
		"networks",
		popts.GetNetworkId(),
		"subnets",
		popts.GetSubnetId())
}

func updateSubnetByIdUrl(psc client.ServiceClient, popts IUpdateSubnetByIdRequest) string {
	return psc.ServiceURL(
		psc.GetProjectId(),
		"networks",
		popts.GetNetworkId(),
		"subnets",
		popts.GetSubnetId())
}

func getAllAddressPairByVirtualSubnetIdUrl(psc client.ServiceClient, popts IGetAllAddressPairByVirtualSubnetIdRequest) string {
	return psc.ServiceURL(
		psc.GetProjectId(),
		"virtual-subnets",
		popts.GetVirtualSubnetId(),
		"addressPairs")
}

func setAddressPairInVirtualSubnetUrl(psc client.ServiceClient, popts IGetAllAddressPairByVirtualSubnetIdRequest) string {
	return psc.ServiceURL(
		psc.GetProjectId(),
		"virtual-subnets",
		popts.GetVirtualSubnetId(),
		"addressPairs")
}

func deleteAddressPairUrl(psc client.ServiceClient, popts IDeleteAddressPairRequest) string {
	return psc.ServiceURL(
		psc.GetProjectId(),
		"virtual-subnets",
		"addressPairs",
		popts.GetAddressPairID())
}

func createAddressPairUrl(psc client.ServiceClient, popts ICreateAddressPairRequest) string {
	return psc.ServiceURL(
		psc.GetProjectId(),
		"virtualIpAddress",
		popts.GetVirtualAddressId(),
		"addressPairs")
}

func listAllServersBySecgroupIdUrl(psc client.ServiceClient, popts IListAllServersBySecgroupIdRequest) string {
	return psc.ServiceURL(
		psc.GetProjectId(),
		"secgroups",
		popts.GetSecgroupId(),
		"servers")
}

func createVirtualAddressCrossProjectUrl(psc client.ServiceClient) string {
	return psc.ServiceURL(
		psc.GetProjectId(),
		"virtualIpAddress")
}

func deleteVirtualAddressByIdUrl(psc client.ServiceClient, popts IDeleteVirtualAddressByIdRequest) string {
	return psc.ServiceURL(
		psc.GetProjectId(),
		"virtualIpAddress",
		popts.GetVirtualAddressId())
}

func getVirtualAddressByIdUrl(psc client.ServiceClient, popts IGetVirtualAddressByIdRequest) string {
	return psc.ServiceURL(
		psc.GetProjectId(),
		"virtualIpAddress",
		popts.GetVirtualAddressId())
}

func listAddressPairsByVirtualAddressIdUrl(psc client.ServiceClient, popts IListAddressPairsByVirtualAddressIdRequest) string {
	return psc.ServiceURL(
		psc.GetProjectId(),
		"virtualIpAddress",
		popts.GetVirtualAddressId(),
		"addressPairs")
}
