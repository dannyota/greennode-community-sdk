//go:build integration

package test

import (
	"context"
	"testing"

	"github.com/dannyota/greennode-community-sdk/v2/greennode"
	"github.com/dannyota/greennode-community-sdk/v2/greennode/entity"
	"github.com/dannyota/greennode-community-sdk/v2/greennode/services/common"
	"github.com/dannyota/greennode-community-sdk/v2/greennode/services/loadbalancer/inter"
	lbv2 "github.com/dannyota/greennode-community-sdk/v2/greennode/services/loadbalancer/v2"
)

func TestCreateInterLoadBalancerSuccess1(t *testing.T) {
	vngcloud := validSuperSdkConfig()
	opt := inter.NewCreateLoadBalancerRequest(
		"53461",
		"test-intervpc-2",
		"lbp-96b6b072-aadb-4b58-9d5f-c16ad69d36aa",
		"sub-6d9aa273-713a-47a4-b4c2-38e150bd809c", // demo
		"sub-fc1461fa-d8d6-4afd-b8c4-07fd6245dff6", // qc2
	)
	lb, sdkerr := vngcloud.LoadBalancerInternal.CreateLoadBalancer(context.Background(), opt)
	if sdkerr != nil {
		t.Fatalf("Expect nil but got %+v", sdkerr)
	}

	if lb == nil {
		t.Fatalf("Expect not nil but got nil")
	}

	t.Log("Result: ", lb)
	t.Log("PASS")
}

func TestCreateInterLoadBalancerV2(t *testing.T) {
	vngcloud := validSuperSdkConfig()
	opt := inter.NewCreateLoadBalancerRequest(
		getValueOfEnv("VNGCLOUD_PORTAL_USER_ID"),
		"test-intervpc",
		"lbp-96b6b072-aadb-4b58-9d5f-c16ad69d36aa",
		"sub-403b36d2-39fc-47c4-b40b-8df0ecb71045",
		"sub-f7770744-6aa4-4292-9ff9-b43b44716ede",
	)
	lb, sdkerr := vngcloud.LoadBalancerInternal.CreateLoadBalancer(context.Background(), opt)
	if sdkerr != nil {
		t.Fatalf("Expect nil but got %+v", sdkerr)
	}

	if lb == nil {
		t.Fatalf("Expect not nil but got nil")
	}

	t.Log("Result: ", lb)
	t.Log("PASS")
}

func TestCreateInterLoadBalancerForProject(t *testing.T) {
	vngcloud := validSuperSdkConfig()
	opt := inter.NewCreateLoadBalancerRequest(
		getValueOfEnv("VNGCLOUD_PORTAL_USER_ID"),
		"test-lb-cross-project",
		"lbp-96b6b072-aadb-4b58-9d5f-c16ad69d36aa",
		"sub-a5e4f8e9-e99e-498c-b99a-6b4720cf5c6f",
		"sub-0f20f37a-602c-4b17-b5f8-f81d4c36aab1",
	)
	lb, sdkerr := vngcloud.LoadBalancerInternal.CreateLoadBalancer(context.Background(), opt)
	if sdkerr != nil {
		t.Fatalf("Expect nil but got %+v", sdkerr)
	}

	if lb == nil {
		t.Fatalf("Expect not nil but got nil")
	}

	t.Log("Result: ", lb)
	t.Log("PASS")
}

func TestCreateInterLoadBalancerSuccess2(t *testing.T) {
	vngcloud := validSecondaryUserSdkConfig()
	opt := inter.NewCreateLoadBalancerRequest(
		getValueOfEnv("SECONDARY_USER_ID"),
		"lb-overlap-private-2",
		"lbp-96b6b072-aadb-4b58-9d5f-c16ad69d36aa",
		"sub-0f20f37a-602c-4b17-b5f8-f81d4c36aab1",
		"sub-d1c1e1bf-2364-4a6a-b300-7bc25785a634",
	)
	lb, sdkerr := vngcloud.LoadBalancerInternal.CreateLoadBalancer(context.Background(), opt)
	if sdkerr != nil {
		t.Fatalf("Expect nil but got %+v", sdkerr)
	}

	if lb == nil {
		t.Fatalf("Expect not nil but got nil")
	}

	t.Log("Result: ", lb)
	t.Log("PASS")
}

func TestCreateInterVpcLbHcm3b(t *testing.T) {
	vngcloud := validHcm3bSuperSdkConfig()
	opt := inter.NewCreateLoadBalancerRequest(
		getValueOfEnv("HCM3B_USER_ID"),
		"test-hcm04-vstorage",
		"lbp-96b6b072-aadb-4b58-9d5f-c16ad69d36aa",
		"sub-0f20f37a-602c-4b17-b5f8-f81d4c36aab1",
		"sub-511ef030-c961-45b5-baac-9d2dadf7e44c",
	)
	lb, sdkerr := vngcloud.LoadBalancerInternal.CreateLoadBalancer(context.Background(), opt)
	if sdkerr != nil {
		t.Fatalf("Expect nil but got %+v", sdkerr)
	}

	if lb == nil {
		t.Fatalf("Expect not nil but got nil")
	}

	t.Log("Result: ", lb)
	t.Log("PASS")
}

func TestCreateInterLoadBalancerSuccess3(t *testing.T) {
	vngcloud, _ := greennode.NewClient(context.Background(), greennode.Config{
		ClientID:        getValueOfEnv("SECONDARY_CLIENT_ID"),
		ClientSecret:    getValueOfEnv("SECONDARY_CLIENT_SECRET"),
		ZoneID:          getValueOfEnv("VNGCLOUD_ZONE_ID"),
		ProjectID:       "pro-c8e87532-dc1a-421c-8c5e-4604d772829f",
		IAMEndpoint:     "https://iamapis.vngcloud.vn/accounts-api",
		VServerEndpoint: "https://hcm-3.api.vngcloud.vn/vserver/vserver-gateway",
		VLBEndpoint:     "https://hcm-3.api.vngcloud.vn/vserver/vlb-gateway",
		RetryCount:      1,
		SleepDuration:   10,
	})
	opt := inter.NewCreateLoadBalancerRequest(
		getValueOfEnv("SECONDARY_USER_ID"),
		"test-15percent-2",
		"lbp-96b6b072-aadb-4b58-9d5f-c16ad69d36aa",
		"sub-0f20f37a-602c-4b17-b5f8-f81d4c36aab1",
		"sub-0725ef54-a32e-404c-96f2-34745239c28d",
	)
	lb, sdkerr := vngcloud.LoadBalancerInternal.CreateLoadBalancer(context.Background(), opt)
	if sdkerr != nil {
		t.Fatalf("Expect nil but got %+v", sdkerr)
	}

	if lb == nil {
		t.Fatalf("Expect not nil but got nil")
	}

	t.Log("Result: ", lb)
	t.Log("PASS")
}

func TestCreateLoadBalancerSuccess(t *testing.T) {
	vngcloud := validSdkConfig()
	opt := lbv2.NewCreateLoadBalancerRequest("test-lb-tags", "", "")
	opt.PackageID = "lbp-f60d5354-0600-11f0-a0a4-ec2a72332f83"
	opt.SubnetID = "sub-3409c1f9-6b03-47bd-979f-251c6e4bee97"
	opt.Tags = common.NewTags("test-key", "test-value", "env", "staging")
	zone := common.HCM_03_1C_ZONE
	opt.ZoneID = &zone
	opt.Listener = lbv2.NewCreateListenerRequest("test-listener", lbv2.ListenerProtocolTCP, 80)
	pool := lbv2.NewCreatePoolRequest("test-pool", lbv2.PoolProtocolTCP)
	pool.Members = append(pool.Members, lbv2.NewMember("test-member-1", "10.84.0.22", 80, 80))
	pool.HealthMonitor = lbv2.NewHealthMonitor(lbv2.HealthCheckProtocolTCP)
	opt.Pool = pool

	lb, sdkerr := vngcloud.LoadBalancer.CreateLoadBalancer(context.Background(), opt)
	if sdkerr != nil {
		t.Fatalf("Expect nil but got %+v", sdkerr)
	}

	if lb == nil {
		t.Fatalf("Expect not nil but got nil")
	}

	t.Log("Result: ", lb)
	t.Log("PASS")
}

func TestCreateLoadBalancerEmptyMemberSuccess(t *testing.T) {
	vngcloud := validSdkConfig()
	opt := lbv2.NewCreateLoadBalancerRequest("test-lb-empty-members", "", "")
	opt.PackageID = "lbp-96b6b072-aadb-4b58-9d5f-c16ad69d36aa"
	opt.SubnetID = "sub-27a0562d-07f9-4e87-81fd-e0ba9658f156"
	opt.Tags = common.NewTags("test-key", "test-value", "env", "staging")
	opt.Listener = lbv2.NewCreateListenerRequest("test-listener", lbv2.ListenerProtocolTCP, 80)
	pool := lbv2.NewCreatePoolRequest("test-pool", lbv2.PoolProtocolTCP)
	pool.HealthMonitor = lbv2.NewHealthMonitor(lbv2.HealthCheckProtocolTCP)
	opt.Pool = pool

	lb, sdkerr := vngcloud.LoadBalancer.CreateLoadBalancer(context.Background(), opt)
	if sdkerr != nil {
		t.Fatalf("Expect nil but got %+v", sdkerr)
	}

	if lb == nil {
		t.Fatalf("Expect not nil but got nil")
	}

	t.Log("Result: ", lb)
	t.Log("PASS")
}

func TestCreateInterVPCLoadBalancerWithPoolAndListenerSuccess(t *testing.T) {
	vngcloud := validSuperSdkConfig()
	hcMethod := inter.HealthCheckMethodGET
	httpVer := inter.HealthCheckHTTPVersionHTTP1
	hcPath := "/health"
	domainName := "vngcloud.com"
	successCode := "200"
	opt := inter.NewCreateLoadBalancerRequest(
		getValueOfEnv("VNGCLOUD_PORTAL_USER_ID"),
		"test-intervpc",
		"lbp-96b6b072-aadb-4b58-9d5f-c16ad69d36aa",
		"sub-27a0562d-07f9-4e87-81fd-e0ba9658f156",
		"sub-888d8fd2-3fed-4aaa-a62d-8554c0aff651")
	opt.Tags = common.NewTags("test-key", "test-value", "env", "staging")
	opt.Listener = inter.NewCreateListenerRequest("test-listener", inter.ListenerProtocolTCP, 80)
	opt.Listener.AllowedCidrs = "0.0.0.0/0"
	opt.Listener.TimeoutClient = 50
	opt.Listener.TimeoutMember = 50
	opt.Listener.TimeoutConnection = 5
	pool := inter.NewCreatePoolRequest("test-pool", inter.PoolProtocolTCP)
	pool.Members = append(pool.Members, inter.NewMember("test-member-1", "10.84.0.22", 80, 80))
	pool.HealthMonitor = &inter.HealthMonitor{
		HealthCheckProtocol: inter.HealthCheckProtocolTCP,
		HealthyThreshold:    3,
		UnhealthyThreshold:  3,
		Timeout:             5,
		Interval:            30,
		HealthCheckMethod:   &hcMethod,
		HTTPVersion:         &httpVer,
		HealthCheckPath:     &hcPath,
		DomainName:          &domainName,
		SuccessCode:         &successCode,
	}
	opt.Pool = pool

	lb, sdkerr := vngcloud.LoadBalancerInternal.CreateLoadBalancer(context.Background(), opt)
	if sdkerr != nil {
		t.Fatalf("Expect nil but got %+v", sdkerr)
	}

	if lb == nil {
		t.Fatalf("Expect not nil but got nil")
	}

	t.Log("Result: ", lb)
	t.Log("PASS")
}

func TestResizeLoadBalancerSuccess(t *testing.T) {
	vngcloud := validSdkConfig()
	opt := lbv2.NewResizeLoadBalancerRequest(
		"lb-4d1508f9-8bb0-45a6-b55b-21a7412b4658",
		"")
	opt.PackageID = "lbp-71cc3022-5fee-426d-9509-3341053e2477"

	lb, sdkerr := vngcloud.LoadBalancer.ResizeLoadBalancer(context.Background(), opt)
	if sdkerr != nil {
		t.Fatalf("Expect nil but got %+v", sdkerr)
	}

	if lb == nil {
		t.Fatalf("Expect not nil but got nil")
	}

	t.Log("Result: ", lb)
	t.Log("PASS")
}

func TestListLoadBalancerPackagesSuccess(t *testing.T) {
	vngcloud := validSdkConfig()
	opt := lbv2.NewListLoadBalancerPackagesRequest()
	opt.ZoneID = common.HCM_03_BKK_01_ZONE
	packages, sdkerr := vngcloud.LoadBalancer.ListLoadBalancerPackages(
		context.Background(), opt)
	if sdkerr != nil {
		t.Fatalf("Expect nil but got %+v", sdkerr)
	}

	if packages == nil {
		t.Fatalf("Expect not nil but got nil")
	}

	t.Log("Result: ", packages)
	for _, pkg := range packages.Items {
		t.Logf("Package: %+v", pkg)
	}
	t.Log("PASS")
}

func TestGetLoadBalancerSuccess(t *testing.T) {
	vngcloud := validSdkConfig()
	opt := lbv2.NewGetLoadBalancerByIDRequest("lb-8f54cbd4-b8ee-4b86-aa9b-d365c468a902")
	lb, sdkerr := vngcloud.LoadBalancer.GetLoadBalancerByID(context.Background(), opt)
	if sdkerr != nil {
		t.Fatalf("Expect nil but got %+v", sdkerr)
	}

	if lb == nil {
		t.Fatalf("Expect not nil but got nil")
	}

	t.Log("Result: ", lb)
	t.Log("PASS")
}

func TestGetLoadBalancerFailure(t *testing.T) {
	vngcloud := validSdkConfig()
	opt := lbv2.NewGetLoadBalancerByIDRequest("lb-f7adf4ba-7734-45f3-8cb5-9b0c3850cc6f")
	lb, sdkerr := vngcloud.LoadBalancer.GetLoadBalancerByID(context.Background(), opt)
	if sdkerr != nil {
		t.Fatalf("Expect nil but got %+v", sdkerr)
	}

	if lb == nil {
		t.Fatalf("Expect not nil but got nil")
	}

	t.Log("Result: ", lb)
	t.Log("PASS")
}

func TestListLoadBalancer(t *testing.T) {
	vngcloud := validSdkConfig()
	opt := lbv2.NewListLoadBalancersRequest(1, 10)
	opt.Name = "test-lb-empty-members"
	lbs, sdkerr := vngcloud.LoadBalancer.ListLoadBalancers(context.Background(), opt)
	if sdkerr != nil {
		t.Fatalf("Expect nil but got %+v", sdkerr)
	}

	if lbs == nil {
		t.Fatalf("Expect not nil but got nil")
	}

	t.Log("Result: ", lbs)
	t.Log("PASS")
}

func TestListLoadBalancerByTagsSuccess(t *testing.T) {
	vngcloud := validSdkConfig()
	opt := lbv2.NewListLoadBalancersRequest(1, 10)
	opt.Tags = common.NewTags("vks-owned-cluster", "hcm03a_user-11412_k8s-a3c03d8e-344c-4a1e-98e0-6d9999ac8077")

	lbs, sdkerr := vngcloud.LoadBalancer.ListLoadBalancers(context.Background(), opt)
	if sdkerr != nil {
		t.Fatalf("Expect nil but got %+v", sdkerr)
	}

	if lbs == nil {
		t.Fatalf("Expect not nil but got nil")
	}

	t.Log("Result: ", lbs)
	t.Log("PASS")
}

func TestCreatePoolWithoutMembersSuccess(t *testing.T) {
	vngcloud := validTestProjectSdkConfig()
	opt := lbv2.NewCreatePoolRequest("test-pool-7", lbv2.PoolProtocolTCP)
	opt.LoadBalancerID = "lb-fb378dc6-71c5-417a-9466-677c03885d6f"
	opt.HealthMonitor = lbv2.NewHealthMonitor(lbv2.HealthCheckProtocolTCP)
	pool, sdkerr := vngcloud.LoadBalancer.CreatePool(context.Background(), opt)

	if sdkerr != nil {
		t.Fatalf("Expect nil but got %+v", sdkerr)
	}

	if pool == nil {
		t.Fatalf("Expect not nil but got nil")
	}

	t.Log("Result: ", pool)
	t.Log("PASS")
}

func TestCreatePoolWithMembersSuccess(t *testing.T) {
	vngcloud := validSdkConfig()
	opt := lbv2.NewCreatePoolRequest("test-pool-3", lbv2.PoolProtocolTCP)
	opt.LoadBalancerID = "lb-8bd4ea07-ab40-483d-8387-124ed2f2cecb"
	opt.Members = append(opt.Members, lbv2.NewMember("test-member-1", "10.84.0.32", 80, 80))
	opt.HealthMonitor = lbv2.NewHealthMonitor(lbv2.HealthCheckProtocolTCP)

	pool, sdkerr := vngcloud.LoadBalancer.CreatePool(context.Background(), opt)

	if sdkerr != nil {
		t.Fatalf("Expect nil but got %+v", sdkerr)
	}

	if pool == nil {
		t.Fatalf("Expect not nil but got nil")
	}

	t.Log("Result: ", pool)
	t.Log("PASS")
}

func TestCreateListenerSuccess(t *testing.T) {
	vngcloud := validSdkConfig()
	opt := lbv2.NewCreateListenerRequest("test-listener-100", lbv2.ListenerProtocolHTTP, 8081)
	opt.LoadBalancerID = "lb-a02759ad-4661-4555-b281-500fd268497e"
	headers := []entity.ListenerInsertHeader{
		{HeaderName: "X-Forwarded-For", HeaderValue: "true"},
		{HeaderName: "X-Forwarded-Proto", HeaderValue: "true"},
		{HeaderName: "Access-Control-Allow-Origin", HeaderValue: "https://*.example.com, https://*.example2.com"},
		{HeaderName: "Access-Control-Allow-Methods", HeaderValue: "GET, HEAD, PATCH, PROPFIND, REPORT"},
		{HeaderName: "Access-Control-Allow-Headers", HeaderValue: "X-Requested-With, X-CSRF-Token, X-PINGOTHER"},
		{HeaderName: "Access-Control-Max-Age", HeaderValue: "86400"},
		{HeaderName: "Access-Control-Allow-Credentials", HeaderValue: "false"},
		{HeaderName: "Access-Control-Expose-Headers", HeaderValue: "X-RateLimit-Limit, X-RateLimit-Remaining, Retry-After"},
	}
	opt.InsertHeaders = &headers

	listener, sdkerr := vngcloud.LoadBalancer.CreateListener(context.Background(), opt)
	if sdkerr != nil {
		t.Fatalf("Expect nil but got %+v", sdkerr)
	}

	if listener == nil {
		t.Fatalf("Expect not nil but got nil")
	}

	t.Log("Result: ", listener)
	t.Log("PASS")
}

func TestCreateListenerWithPoolIDSuccess(t *testing.T) {
	vngcloud := validSdkConfig()
	opt := lbv2.NewCreateListenerRequest("test-listener-14", lbv2.ListenerProtocolTCP, 8087)
	opt.LoadBalancerID = "lb-f7adf4ba-7734-45f3-8cb5-9b0c3850cd6f"
	poolID := "pool-82c3c670-6662-4087-bfc1-8098f25e84df"
	opt.DefaultPoolID = &poolID

	listener, sdkerr := vngcloud.LoadBalancer.CreateListener(context.Background(), opt)
	if sdkerr != nil {
		t.Fatalf("Expect nil but got %+v", sdkerr)
	}

	if listener == nil {
		t.Fatalf("Expect not nil but got nil")
	}

	t.Log("Result: ", listener)
	t.Log("PASS")
}

func TestUpdateListenerSuccess(t *testing.T) {
	vngcloud := validSdkConfig()

	opt := lbv2.NewUpdateListenerRequest(
		"lb-ab73cad3-1dd3-4f2c-9c4c-49702133c5c9",
		"lis-ed226fe5-65d2-4bb1-986e-7814deb3a55b")
	opt.TimeoutClient = 50
	opt.TimeoutConnection = 5
	opt.TimeoutMember = 50
	opt.AllowedCidrs = "0.0.0.0/0"
	opt.DefaultPoolID = "pool-a9239c24-9289-4641-a16b-2d71883d168b"

	sdkerr := vngcloud.LoadBalancer.UpdateListener(context.Background(), opt)
	if sdkerr != nil {
		t.Fatalf("Expect nil but got %+v", sdkerr)
	}

	t.Log("Result: ", sdkerr)
	t.Log("PASS")
}

func TestListListenersByLoadBalancerID(t *testing.T) {
	vngcloud := validSdkConfig()
	opt := lbv2.NewListListenersByLoadBalancerIDRequest("lb-1d3a92bb-6ebd-4b19-ad4b-5f47f5953144")
	listeners, sdkerr := vngcloud.LoadBalancer.ListListenersByLoadBalancerID(context.Background(), opt)
	if sdkerr != nil {
		t.Fatalf("Expect nil but got %+v", sdkerr)
	}

	if listeners == nil {
		t.Fatalf("Expect not nil but got nil")
	}

	t.Log("Result: ", listeners.At(0).DefaultPoolID)
	t.Log("PASS")
}

func TestListPoolsByLoadBalancerID(t *testing.T) {
	vngcloud := validSdkConfig()
	opt := lbv2.NewListPoolsByLoadBalancerIDRequest("lb-4cc1add7-677f-4130-b71a-206940dad28e")
	pools, sdkerr := vngcloud.LoadBalancer.ListPoolsByLoadBalancerID(context.Background(), opt)
	if sdkerr != nil {
		t.Fatalf("Expect nil but got %+v", sdkerr)
	}

	if pools == nil {
		t.Fatalf("Expect not nil but got nil")
	}

	t.Log("Result: ", pools)
	for _, pool := range pools.Items {
		t.Logf("Pool: %+v", pool)
	}
	t.Log("PASS")
}

func TestUpdatePoolMembersSuccess(t *testing.T) {
	vngcloud := validSdkConfig()
	opt := lbv2.NewUpdatePoolMembersRequest(
		"lb-f7adf4ba-7734-45f3-8cb5-9b0c3850cd6f",
		"pool-82c3c670-6662-4087-bfc1-8098f25e84df")
	opt.Members = append(opt.Members, lbv2.NewMember("test-member-50", "10.84.0.50", 80, 80))

	sdkerr := vngcloud.LoadBalancer.UpdatePoolMembers(context.Background(), opt)
	if sdkerr != nil {
		t.Fatalf("Expect nil but got %+v", sdkerr)
	}

	t.Log("Result: ", sdkerr)
	t.Log("PASS")
}

func TestListPoolMembersSuccess(t *testing.T) {
	vngcloud := validSdkConfig()
	opt := lbv2.NewListPoolMembersRequest(
		"lb-8bd4ea07-ab40-483d-8387-124ed2f2cecb",
		"pool-528261c5-9fb4-40bb-bd48-f47b79b272f3")
	members, sdkerr := vngcloud.LoadBalancer.ListPoolMembers(context.Background(), opt)
	if sdkerr != nil {
		t.Fatalf("Expect nil but got %+v", sdkerr)
	}

	if members == nil {
		t.Fatalf("Expect not nil but got nil")
	}

	t.Log("Result: ", members)
	t.Log("PASS")
}

func TestDeletePoolSuccess(t *testing.T) {
	vngcloud := validSdkConfig()
	opt := lbv2.NewDeletePoolByIDRequest(
		"lb-f7adf4ba-7734-45f3-8cb5-9b0c3850cd6f",
		"pool-82c3c670-6662-4087-bfc1-8098f25e84df")
	sdkerr := vngcloud.LoadBalancer.DeletePoolByID(context.Background(), opt)
	if sdkerr != nil {
		t.Fatalf("Expect nil but got %+v", sdkerr)
	}

	t.Log("Result: ", sdkerr)
	t.Log("PASS")
}

func TestDeleteListenterSuccess(t *testing.T) {
	vngcloud := validSdkConfig()
	opt := lbv2.NewDeleteListenerByIDRequest(
		"lb-f7adf4ba-7734-45f3-8cb5-9b0c3850cd6f",
		"lis-23655c30-e458-49ac-ba55-49dfcd104db8")
	sdkerr := vngcloud.LoadBalancer.DeleteListenerByID(context.Background(), opt)
	if sdkerr != nil {
		t.Fatalf("Expect nil but got %+v", sdkerr)
	}

	t.Log("Result: ", sdkerr)
	t.Log("PASS")
}

func TestDeleteLoadBalancer(t *testing.T) {
	vngcloud := validHcm3bSdkConfig()
	opt := lbv2.NewDeleteLoadBalancerByIDRequest("lb-50b72305-02a5-4235-8422-42d6638c7845")
	sdkerr := vngcloud.LoadBalancer.DeleteLoadBalancerByID(context.Background(), opt)
	if sdkerr != nil {
		t.Fatalf("Expect nil but got %+v", sdkerr)
	}

	t.Log("Result: ", sdkerr)
	t.Log("PASS")
}

func TestListTagsSuccess(t *testing.T) {
	vngcloud := validSuperSdkConfig2()
	opt := lbv2.NewListTagsRequest("lb-b1153f05-fd44-4861-8b66-d8b811597faf")
	tags, sdkErr := vngcloud.LoadBalancer.ListTags(context.Background(), opt)
	if sdkErr != nil {
		t.Fatalf("Expect nil but got %+v", sdkErr)
	}

	if tags == nil {
		t.Fatalf("Expect not nil but got nil")
	}

	t.Log("Result: ", tags)
	t.Log("PASS")
}

func TestCreateTagsSuccess(t *testing.T) {
	vngcloud := validSdkConfig()
	opt := lbv2.NewCreateTagsRequest("lb-3b53db2e-357a-406b-9c56-499f1c21a48c")
	opt.TagRequestList = common.NewTags("vks-owned-cluster2", "none")
	sdkErr := vngcloud.LoadBalancer.CreateTags(context.Background(), opt)
	if sdkErr != nil {
		t.Fatalf("Expect nil but got %+v", sdkErr)
	}

	t.Log("Result: ", sdkErr)
	t.Log("PASS")
}

func TestUpdateTagsSuccess(t *testing.T) {
	vngcloud := validSdkConfig()
	opt := lbv2.NewUpdateTagsRequest("lb-39e1750b-7141-455e-a668-a03d53b0328b")
	opt.TagRequestList = common.NewTags("vks-user", "test-user")
	sdkErr := vngcloud.LoadBalancer.UpdateTags(context.Background(), opt)
	if sdkErr != nil {
		t.Fatalf("Expect nil but got %+v", sdkErr)
	}

	t.Log("Result: ", sdkErr)
	t.Log("PASS")
}

func TestUpdatePoolSuccess(t *testing.T) {
	vngcloud := validSdkConfig()

	opt := lbv2.NewUpdatePoolRequest(
		"lb-2af92b71-1da8-4ba3-87bc-b32bb6ab3267",
		"pool-e31c3e31-e285-493e-8ebd-0f0d2cf541b2")
	opt.Algorithm = lbv2.PoolAlgorithmLeastConn
	opt.HealthMonitor = &lbv2.HealthMonitor{
		HealthCheckProtocol: lbv2.HealthCheckProtocolPINGUDP,
		Timeout:             6,
		UnhealthyThreshold:  4,
		HealthyThreshold:    7,
		Interval:            29,
	}

	sdkerr := vngcloud.LoadBalancer.UpdatePool(context.Background(), opt)
	if sdkerr != nil {
		t.Fatalf("Expect nil but got %+v", sdkerr)
	}

	t.Log("Result: ", sdkerr)
	t.Log("PASS")
}

func TestGetPoolHealthMonitorSuccess(t *testing.T) {
	vngcloud := validSdkConfig()
	opt := lbv2.NewGetPoolHealthMonitorByIDRequest(
		"lb-d5501a8c-d40e-4e3d-b86a-3e4041c629f7",
		"pool-1c5dfb52-922a-4dac-9dc0-970980637199")
	hm, sdkerr := vngcloud.LoadBalancer.GetPoolHealthMonitorByID(context.Background(), opt)
	if sdkerr != nil {
		t.Fatalf("Expect nil but got %+v", sdkerr)
	}

	if hm == nil {
		t.Fatalf("Expect not nil but got nil")
	}

	t.Log("Result: ", hm)
	t.Log("PASS")
}

func TestListPoliciesSuccess(t *testing.T) {
	vngcloud := validSdkConfig()
	opt := lbv2.NewListPoliciesRequest(
		"lb-eb9f558a-4724-4d0b-a197-60fd642236f4",
		"lis-b38a9abc-2979-444f-afce-da824e32ea75")
	policies, sdkerr := vngcloud.LoadBalancer.ListPolicies(context.Background(), opt)
	if sdkerr != nil {
		t.Fatalf("Expect nil but got %+v", sdkerr)
	}

	if policies == nil {
		t.Fatalf("Expect not nil but got nil")
	}

	t.Log("Result: ", policies)
	for _, policy := range policies.Items {
		t.Logf("Policy: %+v", policy)
	}
	t.Log("PASS")
}

func TestCreatePolicySuccess(t *testing.T) {
	vngcloud := validSdkConfig()
	opt := lbv2.NewCreatePolicyRequest(
		"lb-eb9f558a-4724-4d0b-a197-60fd642236f4",
		"lis-b38a9abc-2979-444f-afce-da824e32ea75")
	opt.Name = "test-policy-1"
	opt.Action = lbv2.PolicyActionREJECT
	opt.Rules = []lbv2.L7RuleRequest{
		{
			CompareType: lbv2.PolicyCompareTypeCONTAINS,
			RuleType:    lbv2.PolicyRuleTypeHOSTNAME,
			RuleValue:   "vngcloud.vn",
		},
	}

	policy, sdkerr := vngcloud.LoadBalancer.CreatePolicy(context.Background(), opt)
	if sdkerr != nil {
		t.Fatalf("Expect nil but got %+v", sdkerr)
	}

	if policy == nil {
		t.Fatalf("Expect not nil but got nil")
	}

	t.Log("Result: ", policy)
	t.Log("PASS")
}

func TestGetPolicyByIDSuccess(t *testing.T) {
	vngcloud := validSdkConfig()
	opt := lbv2.NewGetPolicyByIDRequest(
		"lb-eb9f558a-4724-4d0b-a197-60fd642236f4",
		"lis-b38a9abc-2979-444f-afce-da824e32ea75",
		"policy-dea6106b-dd41-4fc1-bddc-61acc034787b")
	policy, sdkerr := vngcloud.LoadBalancer.GetPolicyByID(context.Background(), opt)
	if sdkerr != nil {
		t.Fatalf("Expect nil but got %+v", sdkerr)
	}

	if policy == nil {
		t.Fatalf("Expect not nil but got nil")
	}

	t.Log("Result: ", policy)
	t.Log("PASS")
}

func TestUpdatePolicySuccess(t *testing.T) {
	vngcloud := validSdkConfig()
	opt := lbv2.NewUpdatePolicyRequest(
		"lb-eb9f558a-4724-4d0b-a197-60fd642236f4",
		"lis-b38a9abc-2979-444f-afce-da824e32ea75",
		"policy-dea6106b-dd41-4fc1-bddc-61acc034787b")
	opt.Action = lbv2.PolicyActionREDIRECTTOURL
	opt.RedirectURL = "https://vngcloud.vn"
	opt.RedirectHTTPCode = 301
	opt.KeepQueryString = true
	opt.Rules = []lbv2.L7RuleRequest{
		{
			CompareType: lbv2.PolicyCompareTypeCONTAINS,
			RuleType:    lbv2.PolicyRuleTypeHOSTNAME,
			RuleValue:   "vngcloud.com.vn",
		},
	}

	sdkerr := vngcloud.LoadBalancer.UpdatePolicy(context.Background(), opt)
	if sdkerr != nil {
		t.Fatalf("Expect nil but got %+v", sdkerr)
	}

	t.Log("Result: ", sdkerr)
	t.Log("PASS")
}

func TestDeletePolicySuccess(t *testing.T) {
	vngcloud := validSdkConfig()
	opt := lbv2.NewDeletePolicyByIDRequest(
		"lb-eb9f558a-4724-4d0b-a197-60fd642236f4",
		"lis-b38a9abc-2979-444f-afce-da824e32ea75",
		"policy-5cf4bacb-93b6-4078-bbf7-cb5d0d701828")
	sdkerr := vngcloud.LoadBalancer.DeletePolicyByID(context.Background(), opt)
	if sdkerr != nil {
		t.Fatalf("Expect nil but got %+v", sdkerr)
	}

	t.Log("Result: ", sdkerr)
	t.Log("PASS")
}

func TestReorderPoliciesSucces(t *testing.T) {
	vngcloud := validSdkConfig()
	opt := lbv2.NewReorderPoliciesRequest(
		"lb-d08b5093-b923-4064-b38c-828add7d439a",
		"lis-654e4105-b729-4bd1-9a33-61ddacbe3430")
	opt.PolicyPositions = []lbv2.PolicyPosition{
		{Position: 1, PolicyID: "policy-57b9e7d3-7ae6-4cb3-a649-6aa35f3ae26d"},
		{Position: 2, PolicyID: "policy-1d29aa49-e9da-4551-a349-39d3338cfc4a"},
		{Position: 3, PolicyID: "policy-bdee1abb-a4b6-4331-92b2-d1a8ddd51904"},
		{Position: 4, PolicyID: "policy-f6cfc6ec-3a4c-4cb0-a56e-16c9f9a2ac74"},
	}

	sdkerr := vngcloud.LoadBalancer.ReorderPolicies(context.Background(), opt)
	if sdkerr != nil {
		t.Fatalf("Expect nil but got %+v", sdkerr)
	}

	t.Log("Result: ", sdkerr)
	t.Log("PASS")
}

func TestScaleLoadBalancerSuccess(t *testing.T) {
	vngcloud := validSdkConfig()
	opt := lbv2.NewScaleLoadBalancerRequest("lb-7dd1f328-c7d6-4cbd-b194-7d8502de3fc8")
	opt.Scaling = &lbv2.ScalingConfig{
		MinNodes: 2,
		MaxNodes: 5,
	}
	opt.Networking = &lbv2.NetworkingConfig{
		Subnets: []string{
			"sub-403b36d2-39fc-47c4-b40b-8df0ecb71045",
			"sub-70b263e5-094b-44d4-9861-834a8ad190ce",
			// "sub-5fc1ee76-b754-490e-afa7-c7d963848481",
		},
	}

	lb, sdkerr := vngcloud.LoadBalancer.ScaleLoadBalancer(context.Background(), opt)
	if sdkerr != nil {
		t.Fatalf("Expect nil but got %+v", sdkerr)
	}

	if lb == nil {
		t.Fatalf("Expect not nil but got nil")
	}

	t.Log("Result: ", lb)
	t.Log("PASS")
}
