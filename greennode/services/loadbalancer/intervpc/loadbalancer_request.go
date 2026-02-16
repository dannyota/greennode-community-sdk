package intervpc

import "github.com/dannyota/greennode-community-sdk/v2/greennode/services/common"

const (
	InterVpcLoadBalancerScheme LoadBalancerScheme = "InterVPC"
	InternalLoadBalancerScheme LoadBalancerScheme = "Internal"
	InternetLoadBalancerScheme LoadBalancerScheme = "Internet"
)

const (
	CreateOptsTypeOptLayer4 LoadBalancerType = "Layer 4"
	CreateOptsTypeOptLayer7 LoadBalancerType = "Layer 7"
)

type (
	LoadBalancerScheme string
	LoadBalancerType   string
)

type CreateLoadBalancerRequest struct {
	Name            string                 `json:"name"`
	PackageID       string                 `json:"packageId"`
	Scheme          LoadBalancerScheme     `json:"scheme"`
	SubnetID        string                 `json:"subnetId"`
	BackEndSubnetID string                 `json:"backendSubnetId,omitempty"`
	ProjectID       string                 `json:"projectId,omitempty"`
	Type            LoadBalancerType       `json:"type"`
	Listener        *CreateListenerRequest `json:"listener,omitempty"`
	Pool            *CreatePoolRequest     `json:"pool,omitempty"`
	Tags            []common.Tag           `json:"tags,omitempty"`
	ZoneID          *common.Zone           `json:"zoneId,omitempty"`

	common.PortalUser
}

// normalizeForAPI delegates to Pool and Listener normalization, clearing
// protocol-irrelevant fields before the API call. This mutates the receiver.
func (r *CreateLoadBalancerRequest) normalizeForAPI() {
	if r.Pool != nil {
		r.Pool.normalizeForAPI()
	}

	if r.Listener != nil {
		r.Listener.normalizeForAPI()
	}
}

func NewCreateLoadBalancerRequest(userID, name, packageID, beSubnetID, subnetID string) *CreateLoadBalancerRequest {
	return &CreateLoadBalancerRequest{
		Name:            name,
		PackageID:       packageID,
		Scheme:          InterVpcLoadBalancerScheme,
		BackEndSubnetID: beSubnetID,
		SubnetID:        subnetID,
		Type:            CreateOptsTypeOptLayer4,
		PortalUser:      common.PortalUser{ID: userID},
	}
}
