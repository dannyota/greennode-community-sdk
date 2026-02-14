package common

type LoadBalancerCommon struct {
	LoadBalancerID string
}

func (s *LoadBalancerCommon) GetLoadBalancerID() string {
	return s.LoadBalancerID
}

type ListenerCommon struct {
	ListenerID string
}

func (s *ListenerCommon) GetListenerID() string {
	return s.ListenerID
}

type PoolCommon struct {
	PoolID string
}

func (s *PoolCommon) GetPoolID() string {
	return s.PoolID
}

type PolicyCommon struct {
	PolicyID string
}

func (s *PolicyCommon) GetPolicyID() string {
	return s.PolicyID
}

type PoolMemberCommon struct {
	PoolMemberID string
}

func (s *PoolMemberCommon) GetPoolMemberID() string {
	return s.PoolMemberID
}
