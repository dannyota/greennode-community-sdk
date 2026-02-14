package v2

func NewResizeLoadBalancerByIdRequest(lbId, packageId string) IResizeLoadBalancerByIdRequest {
	opts := new(ResizeLoadBalancerByIdRequest)
	opts.LoadBalancerId = lbId
	opts.PackageId = packageId
	return opts
}
