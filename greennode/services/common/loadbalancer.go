package common

type LoadBalancerCommon struct {
	LoadBalancerID string
}

func (l *LoadBalancerCommon) GetLoadBalancerID() string {
	return l.LoadBalancerID
}

type ListenerCommon struct {
	ListenerID string
}

func (l *ListenerCommon) GetListenerID() string {
	return l.ListenerID
}

type PoolCommon struct {
	PoolID string
}

func (p *PoolCommon) GetPoolID() string {
	return p.PoolID
}

type PolicyCommon struct {
	PolicyID string
}

func (p *PolicyCommon) GetPolicyID() string {
	return p.PolicyID
}

type PoolMemberCommon struct {
	PoolMemberID string
}

func (p *PoolMemberCommon) GetPoolMemberID() string {
	return p.PoolMemberID
}
