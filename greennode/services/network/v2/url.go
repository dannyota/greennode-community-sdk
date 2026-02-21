package v2

import "danny.vn/greennode/greennode/client"

func getSecgroupByIDURL(sc *client.ServiceClient, opts *GetSecgroupByIDRequest) string {
	return sc.ServiceURL(
		sc.ProjectID,
		"secgroups",
		opts.SecgroupID)
}

func createSecgroupURL(sc *client.ServiceClient) string {
	return sc.ServiceURL(
		sc.ProjectID,
		"secgroups")
}

func listSecgroupURL(sc *client.ServiceClient, _ *ListSecgroupRequest) string {
	return sc.ServiceURL(
		sc.ProjectID,
		"secgroups")
}

func deleteSecgroupByIDURL(sc *client.ServiceClient, opts *DeleteSecgroupByIDRequest) string {
	return sc.ServiceURL(
		sc.ProjectID,
		"secgroups",
		opts.SecgroupID)
}

func createSecgroupRuleURL(sc *client.ServiceClient, opts *CreateSecgroupRuleRequest) string {
	return sc.ServiceURL(
		sc.ProjectID,
		"secgroups",
		opts.SecgroupID,
		"secgroupRules")
}

func deleteSecgroupRuleByIDURL(sc *client.ServiceClient, opts *DeleteSecgroupRuleByIDRequest) string {
	return sc.ServiceURL(
		sc.ProjectID,
		"secgroups",
		opts.SecgroupID,
		"secgroupRules",
		opts.SecgroupRuleID)
}

func listSecgroupRulesBySecgroupIDURL(sc *client.ServiceClient, opts *ListSecgroupRulesBySecgroupIDRequest) string {
	return sc.ServiceURL(
		sc.ProjectID,
		"secgroups",
		opts.SecgroupID,
		"secGroupRules")
}

func getNetworkByIDURL(sc *client.ServiceClient, opts *GetNetworkByIDRequest) string {
	return sc.ServiceURL(
		sc.ProjectID,
		"networks",
		opts.NetworkID)
}

func getSubnetByIDURL(sc *client.ServiceClient, opts *GetSubnetByIDRequest) string {
	return sc.ServiceURL(
		sc.ProjectID,
		"networks",
		opts.NetworkID,
		"subnets",
		opts.SubnetID)
}

func updateSubnetByIDURL(sc *client.ServiceClient, opts *UpdateSubnetByIDRequest) string {
	return sc.ServiceURL(
		sc.ProjectID,
		"networks",
		opts.NetworkID,
		"subnets",
		opts.SubnetID)
}

func getAllAddressPairByVirtualSubnetIDURL(sc *client.ServiceClient, opts *GetAllAddressPairByVirtualSubnetIDRequest) string {
	return sc.ServiceURL(
		sc.ProjectID,
		"virtual-subnets",
		opts.VirtualSubnetID,
		"addressPairs")
}

func setAddressPairInVirtualSubnetURL(sc *client.ServiceClient, opts *SetAddressPairInVirtualSubnetRequest) string {
	return sc.ServiceURL(
		sc.ProjectID,
		"virtual-subnets",
		opts.VirtualSubnetID,
		"addressPairs")
}

func deleteAddressPairURL(sc *client.ServiceClient, opts *DeleteAddressPairRequest) string {
	return sc.ServiceURL(
		sc.ProjectID,
		"virtual-subnets",
		"addressPairs",
		opts.AddressPairID)
}

func createAddressPairURL(sc *client.ServiceClient, opts *CreateAddressPairRequest) string {
	return sc.ServiceURL(
		sc.ProjectID,
		"virtualIpAddress",
		opts.VirtualAddressID,
		"addressPairs")
}

func listAllServersBySecgroupIDURL(sc *client.ServiceClient, opts *ListAllServersBySecgroupIDRequest) string {
	return sc.ServiceURL(
		sc.ProjectID,
		"secgroups",
		opts.SecgroupID,
		"servers")
}

func createVirtualAddressCrossProjectURL(sc *client.ServiceClient) string {
	return sc.ServiceURL(
		sc.ProjectID,
		"virtualIpAddress")
}

func deleteVirtualAddressByIDURL(sc *client.ServiceClient, opts *DeleteVirtualAddressByIDRequest) string {
	return sc.ServiceURL(
		sc.ProjectID,
		"virtualIpAddress",
		opts.VirtualAddressID)
}

func getVirtualAddressByIDURL(sc *client.ServiceClient, opts *GetVirtualAddressByIDRequest) string {
	return sc.ServiceURL(
		sc.ProjectID,
		"virtualIpAddress",
		opts.VirtualAddressID)
}

func listAddressPairsByVirtualAddressIDURL(sc *client.ServiceClient, opts *ListAddressPairsByVirtualAddressIDRequest) string {
	return sc.ServiceURL(
		sc.ProjectID,
		"virtualIpAddress",
		opts.VirtualAddressID,
		"addressPairs")
}
