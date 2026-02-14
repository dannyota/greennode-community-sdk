package inter

func NewCreateLoadBalancerRequest(userID, name, packageID, beSubnetID, subnetID string) ICreateLoadBalancerRequest {
	opt := new(CreateLoadBalancerRequest)
	opt.SetPortalUserID(userID)
	opt.Name = name
	opt.PackageID = packageID
	opt.Scheme = InterVpcLoadBalancerScheme
	opt.BackEndSubnetID = beSubnetID
	opt.SubnetID = subnetID
	opt.Type = CreateOptsTypeOptLayer4
	return opt
}
