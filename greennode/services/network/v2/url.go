package v2

import "github.com/dannyota/greennode-community-sdk/v2/greennode/client"

func getSecgroupByIdUrl(sc client.ServiceClient, opts IGetSecgroupByIdRequest) string {
	return sc.ServiceURL(
		sc.GetProjectId(),
		"secgroups",
		opts.GetSecgroupId())
}

func createSecgroupUrl(sc client.ServiceClient) string {
	return sc.ServiceURL(
		sc.GetProjectId(),
		"secgroups")
}

func listSecgroupUrl(sc client.ServiceClient, _ IListSecgroupRequest) string {
	return sc.ServiceURL(
		sc.GetProjectId(),
		"secgroups")
}

func deleteSecgroupByIdUrl(sc client.ServiceClient, opts IDeleteSecgroupByIdRequest) string {
	return sc.ServiceURL(
		sc.GetProjectId(),
		"secgroups",
		opts.GetSecgroupId())
}

func createSecgroupRuleUrl(sc client.ServiceClient, opts ICreateSecgroupRuleRequest) string {
	return sc.ServiceURL(
		sc.GetProjectId(),
		"secgroups",
		opts.GetSecgroupId(),
		"secgroupRules")
}

func deleteSecgroupRuleByIdUrl(sc client.ServiceClient, opts IDeleteSecgroupRuleByIdRequest) string {
	return sc.ServiceURL(
		sc.GetProjectId(),
		"secgroups",
		opts.GetSecgroupId(),
		"secgroupRules",
		opts.GetSecgroupRuleId())
}

func listSecgroupRulesBySecgroupIdUrl(sc client.ServiceClient, opts IListSecgroupRulesBySecgroupIdRequest) string {
	return sc.ServiceURL(
		sc.GetProjectId(),
		"secgroups",
		opts.GetSecgroupId(),
		"secGroupRules")
}

func getNetworkByIdUrl(sc client.ServiceClient, opts IGetNetworkByIdRequest) string {
	return sc.ServiceURL(
		sc.GetProjectId(),
		"networks",
		opts.GetNetworkId())
}

func getSubnetByIdUrl(sc client.ServiceClient, opts IGetSubnetByIdRequest) string {
	return sc.ServiceURL(
		sc.GetProjectId(),
		"networks",
		opts.GetNetworkId(),
		"subnets",
		opts.GetSubnetId())
}

func updateSubnetByIdUrl(sc client.ServiceClient, opts IUpdateSubnetByIdRequest) string {
	return sc.ServiceURL(
		sc.GetProjectId(),
		"networks",
		opts.GetNetworkId(),
		"subnets",
		opts.GetSubnetId())
}

func getAllAddressPairByVirtualSubnetIdUrl(sc client.ServiceClient, opts IGetAllAddressPairByVirtualSubnetIdRequest) string {
	return sc.ServiceURL(
		sc.GetProjectId(),
		"virtual-subnets",
		opts.GetVirtualSubnetId(),
		"addressPairs")
}

func setAddressPairInVirtualSubnetUrl(sc client.ServiceClient, opts IGetAllAddressPairByVirtualSubnetIdRequest) string {
	return sc.ServiceURL(
		sc.GetProjectId(),
		"virtual-subnets",
		opts.GetVirtualSubnetId(),
		"addressPairs")
}

func deleteAddressPairUrl(sc client.ServiceClient, opts IDeleteAddressPairRequest) string {
	return sc.ServiceURL(
		sc.GetProjectId(),
		"virtual-subnets",
		"addressPairs",
		opts.GetAddressPairID())
}

func createAddressPairUrl(sc client.ServiceClient, opts ICreateAddressPairRequest) string {
	return sc.ServiceURL(
		sc.GetProjectId(),
		"virtualIpAddress",
		opts.GetVirtualAddressId(),
		"addressPairs")
}

func listAllServersBySecgroupIdUrl(sc client.ServiceClient, opts IListAllServersBySecgroupIdRequest) string {
	return sc.ServiceURL(
		sc.GetProjectId(),
		"secgroups",
		opts.GetSecgroupId(),
		"servers")
}

func createVirtualAddressCrossProjectUrl(sc client.ServiceClient) string {
	return sc.ServiceURL(
		sc.GetProjectId(),
		"virtualIpAddress")
}

func deleteVirtualAddressByIdUrl(sc client.ServiceClient, opts IDeleteVirtualAddressByIdRequest) string {
	return sc.ServiceURL(
		sc.GetProjectId(),
		"virtualIpAddress",
		opts.GetVirtualAddressId())
}

func getVirtualAddressByIdUrl(sc client.ServiceClient, opts IGetVirtualAddressByIdRequest) string {
	return sc.ServiceURL(
		sc.GetProjectId(),
		"virtualIpAddress",
		opts.GetVirtualAddressId())
}

func listAddressPairsByVirtualAddressIdUrl(sc client.ServiceClient, opts IListAddressPairsByVirtualAddressIdRequest) string {
	return sc.ServiceURL(
		sc.GetProjectId(),
		"virtualIpAddress",
		opts.GetVirtualAddressId(),
		"addressPairs")
}
