package v2

import "github.com/dannyota/greennode-community-sdk/v2/greennode/client"

func getSecgroupByIDURL(sc client.ServiceClient, opts IGetSecgroupByIDRequest) string {
	return sc.ServiceURL(
		sc.GetProjectID(),
		"secgroups",
		opts.GetSecgroupID())
}

func createSecgroupURL(sc client.ServiceClient) string {
	return sc.ServiceURL(
		sc.GetProjectID(),
		"secgroups")
}

func listSecgroupURL(sc client.ServiceClient, _ IListSecgroupRequest) string {
	return sc.ServiceURL(
		sc.GetProjectID(),
		"secgroups")
}

func deleteSecgroupByIDURL(sc client.ServiceClient, opts IDeleteSecgroupByIDRequest) string {
	return sc.ServiceURL(
		sc.GetProjectID(),
		"secgroups",
		opts.GetSecgroupID())
}

func createSecgroupRuleURL(sc client.ServiceClient, opts ICreateSecgroupRuleRequest) string {
	return sc.ServiceURL(
		sc.GetProjectID(),
		"secgroups",
		opts.GetSecgroupID(),
		"secgroupRules")
}

func deleteSecgroupRuleByIDURL(sc client.ServiceClient, opts IDeleteSecgroupRuleByIDRequest) string {
	return sc.ServiceURL(
		sc.GetProjectID(),
		"secgroups",
		opts.GetSecgroupID(),
		"secgroupRules",
		opts.GetSecgroupRuleID())
}

func listSecgroupRulesBySecgroupIDURL(sc client.ServiceClient, opts IListSecgroupRulesBySecgroupIDRequest) string {
	return sc.ServiceURL(
		sc.GetProjectID(),
		"secgroups",
		opts.GetSecgroupID(),
		"secGroupRules")
}

func getNetworkByIDURL(sc client.ServiceClient, opts IGetNetworkByIDRequest) string {
	return sc.ServiceURL(
		sc.GetProjectID(),
		"networks",
		opts.GetNetworkID())
}

func getSubnetByIDURL(sc client.ServiceClient, opts IGetSubnetByIDRequest) string {
	return sc.ServiceURL(
		sc.GetProjectID(),
		"networks",
		opts.GetNetworkID(),
		"subnets",
		opts.GetSubnetID())
}

func updateSubnetByIDURL(sc client.ServiceClient, opts IUpdateSubnetByIDRequest) string {
	return sc.ServiceURL(
		sc.GetProjectID(),
		"networks",
		opts.GetNetworkID(),
		"subnets",
		opts.GetSubnetID())
}

func getAllAddressPairByVirtualSubnetIDURL(sc client.ServiceClient, opts IGetAllAddressPairByVirtualSubnetIDRequest) string {
	return sc.ServiceURL(
		sc.GetProjectID(),
		"virtual-subnets",
		opts.GetVirtualSubnetID(),
		"addressPairs")
}

func setAddressPairInVirtualSubnetURL(sc client.ServiceClient, opts IGetAllAddressPairByVirtualSubnetIDRequest) string {
	return sc.ServiceURL(
		sc.GetProjectID(),
		"virtual-subnets",
		opts.GetVirtualSubnetID(),
		"addressPairs")
}

func deleteAddressPairURL(sc client.ServiceClient, opts IDeleteAddressPairRequest) string {
	return sc.ServiceURL(
		sc.GetProjectID(),
		"virtual-subnets",
		"addressPairs",
		opts.GetAddressPairID())
}

func createAddressPairURL(sc client.ServiceClient, opts ICreateAddressPairRequest) string {
	return sc.ServiceURL(
		sc.GetProjectID(),
		"virtualIpAddress",
		opts.GetVirtualAddressID(),
		"addressPairs")
}

func listAllServersBySecgroupIDURL(sc client.ServiceClient, opts IListAllServersBySecgroupIDRequest) string {
	return sc.ServiceURL(
		sc.GetProjectID(),
		"secgroups",
		opts.GetSecgroupID(),
		"servers")
}

func createVirtualAddressCrossProjectURL(sc client.ServiceClient) string {
	return sc.ServiceURL(
		sc.GetProjectID(),
		"virtualIpAddress")
}

func deleteVirtualAddressByIDURL(sc client.ServiceClient, opts IDeleteVirtualAddressByIDRequest) string {
	return sc.ServiceURL(
		sc.GetProjectID(),
		"virtualIpAddress",
		opts.GetVirtualAddressID())
}

func getVirtualAddressByIDURL(sc client.ServiceClient, opts IGetVirtualAddressByIDRequest) string {
	return sc.ServiceURL(
		sc.GetProjectID(),
		"virtualIpAddress",
		opts.GetVirtualAddressID())
}

func listAddressPairsByVirtualAddressIDURL(sc client.ServiceClient, opts IListAddressPairsByVirtualAddressIDRequest) string {
	return sc.ServiceURL(
		sc.GetProjectID(),
		"virtualIpAddress",
		opts.GetVirtualAddressID(),
		"addressPairs")
}
