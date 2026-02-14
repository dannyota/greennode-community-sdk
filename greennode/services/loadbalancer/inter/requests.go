package inter

func NewCreateLoadBalancerRequest(userId, name, packageId, beSubnetId, subnetId string) ICreateLoadBalancerRequest {
	opt := new(CreateLoadBalancerRequest)
	opt.SetPortalUserId(userId)
	opt.Name = name
	opt.PackageID = packageId
	opt.Scheme = InterVpcLoadBalancerScheme
	opt.BackEndSubnetId = beSubnetId
	opt.SubnetID = subnetId
	opt.Type = CreateOptsTypeOptLayer4
	return opt
}
