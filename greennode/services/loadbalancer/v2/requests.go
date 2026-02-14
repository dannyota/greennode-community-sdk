package v2

func NewResizeLoadBalancerByIDRequest(lbID, packageID string) *ResizeLoadBalancerByIDRequest {
	opts := new(ResizeLoadBalancerByIDRequest)
	opts.LoadBalancerID = lbID
	opts.PackageID = packageID
	return opts
}
