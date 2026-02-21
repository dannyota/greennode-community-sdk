//go:build integration

package test

import (
	"context"
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/dannyota/greennode-community-sdk/greennode/sdkerror"
	"github.com/dannyota/greennode-community-sdk/greennode/services/common"
	v1 "github.com/dannyota/greennode-community-sdk/greennode/services/glb/v1"
)

func TestGetGlobalListenerSuccess(t *testing.T) {
	vngcloud := validSdkConfig()
	opt := v1.NewGetGlobalListenerRequest(
		"glb-a9799830-f7ef-40a8-ad05-ba7f81a8bb8d",
		"glis-6e0b0a21-0b52-40b7-90e5-f59a83e577c2",
	)
	listener, sdkerr := vngcloud.GLB.GetGlobalListener(context.Background(), opt)
	if sdkerr != nil {
		t.Fatalf("Expect nil but got %+v", sdkerr)
	}

	if listener == nil {
		t.Fatalf("Expect not nil but got nil")
	}

	t.Logf("Result: %+v", listener)
	t.Log("PASS")
}

func TestGetGlobalPoolMemberSuccess(t *testing.T) {
	vngcloud := validSdkConfig()
	opt := v1.NewGetGlobalPoolMemberRequest(
		"glb-a9799830-f7ef-40a8-ad05-ba7f81a8bb8d",
		"gpool-e5de4670-27e6-45cf-bc68-ec3803ed6849",
		"gpool-mem-4b3a819d-a83f-4964-8336-da6cb8edf529",
	)
	poolMember, sdkerr := vngcloud.GLB.GetGlobalPoolMember(context.Background(), opt)
	if sdkerr != nil {
		t.Fatalf("Expect nil but got %+v", sdkerr)
	}

	if poolMember == nil {
		t.Fatalf("Expect not nil but got nil")
	}

	t.Logf("Pool Member: %+v", poolMember)
	if poolMember.Members != nil && len(poolMember.Members.Items) > 0 {
		t.Logf("Members count: %d", len(poolMember.Members.Items))
		for i, member := range poolMember.Members.Items {
			t.Logf("Member[%d]: %+v", i, member)
		}
	} else {
		t.Log("No members found")
	}
	t.Log("PASS")
}

func TestDeleteGlobalPoolMemberSuccess(t *testing.T) {
	vngcloud := validSdkConfig()
	opt := v1.NewDeleteGlobalPoolMemberRequest(
		"glb-a9799830-f7ef-40a8-ad05-ba7f81a8bb8d",
		"gpool-e5de4670-27e6-45cf-bc68-ec3803ed6849",
		"gpool-mem-4b3a819d-a83f-4964-8336-da6cb8edf529",
	)

	sdkerr := vngcloud.GLB.DeleteGlobalPoolMember(context.Background(), opt)
	if sdkerr != nil {
		t.Fatalf("Expect nil but got %+v", sdkerr)
	}

	t.Log("PASS")
}

func TestUpdateGlobalPoolMemberSuccess(t *testing.T) {
	vngcloud := validSdkConfig()
	opt := v1.NewUpdateGlobalPoolMemberRequest(
		"glb-a9799830-f7ef-40a8-ad05-ba7f81a8bb8d",
		"gpool-e5de4670-27e6-45cf-bc68-ec3803ed6849",
		"gpool-mem-4b3a819d-a83f-4964-8336-da6cb8edf529",
		100,
	)
	opt.Members = append(opt.Members,
		v1.NewGlobalMemberRequest("updated-member", "10.0.0.9", "sub-e208484a-69cd-4a70-a7dd-f60bbfd4b04d", 80, 80, 1, false),
		v1.NewGlobalMemberRequest("updated-member", "10.0.0.10", "sub-e208484a-69cd-4a70-a7dd-f60bbfd4b04d", 80, 80, 1, false),
	)

	poolMember, sdkerr := vngcloud.GLB.UpdateGlobalPoolMember(context.Background(), opt)
	if sdkerr != nil {
		t.Fatalf("Expect nil but got %+v", sdkerr)
	}

	if poolMember == nil {
		t.Fatalf("Expect not nil but got nil")
	}

	t.Logf("Updated Pool Member: %+v", poolMember)
	t.Log("PASS")
}

func TestListGlobalPackagesSuccess(t *testing.T) {
	vngcloud := validSdkConfig()
	opt := v1.NewListGlobalPackagesRequest()
	packages, sdkerr := vngcloud.GLB.ListGlobalPackages(context.Background(), opt)
	if sdkerr != nil {
		t.Fatalf("Expect nil but got %+v", sdkerr)
	}

	if packages == nil {
		t.Fatalf("Expect not nil but got nil")
	}

	t.Logf("Packages: %+v", packages)
	if len(packages.Items) > 0 {
		t.Logf("Packages count: %d", len(packages.Items))
		for i, pkg := range packages.Items {
			t.Logf("Package[%d]: ID=%s, Name=%s, Description=%s", i, pkg.ID, pkg.Name, pkg.Description)
		}
	} else {
		t.Log("No packages found")
	}
	t.Log("PASS")
}

func TestListGlobalRegionsSuccess(t *testing.T) {
	vngcloud := validSdkConfig()
	opt := v1.NewListGlobalRegionsRequest()
	regions, sdkerr := vngcloud.GLB.ListGlobalRegions(context.Background(), opt)
	if sdkerr != nil {
		t.Fatalf("Expect nil but got %+v", sdkerr)
	}

	if regions == nil {
		t.Fatalf("Expect not nil but got nil")
	}

	t.Logf("Regions: %+v", regions)
	if len(regions.Items) > 0 {
		t.Logf("Regions count: %d", len(regions.Items))
		for i, region := range regions.Items {
			t.Logf("Region[%d]: %+v", i, region)
		}
	} else {
		t.Log("No regions found")
	}
	t.Log("PASS")
}

func TestGetGlobalLoadBalancerUsageHistoriesSuccess(t *testing.T) {
	vngcloud := validSdkConfig()
	// Using timestamps from 1 day ago to current moment
	now := time.Now().Unix()
	oneDayAgo := now - 86400 // 86400 seconds = 24 hours
	from := fmt.Sprintf("%d", oneDayAgo)
	to := fmt.Sprintf("%d", now)
	usageType := "connection_rate"
	opt := v1.NewGetGlobalLoadBalancerUsageHistoriesRequest("glb-a9799830-f7ef-40a8-ad05-ba7f81a8bb8d", from, to, usageType)
	histories, sdkerr := vngcloud.GLB.GetGlobalLoadBalancerUsageHistories(context.Background(), opt)
	if sdkerr != nil {
		t.Fatalf("Expect nil but got %+v", sdkerr)
	}

	if histories == nil {
		t.Fatalf("Expect not nil but got nil")
	}

	t.Logf("Usage Histories: %+v", histories)
	if len(histories.Items) > 0 {
		t.Logf("Usage histories count: %d", len(histories.Items))
		for i, history := range histories.Items {
			t.Logf("History[%d]: %+v", i, history)
		}
	} else {
		t.Log("No usage histories found")
	}
	t.Log("PASS")
}

func TestListGlobalPoolsSuccess(t *testing.T) {
	vngcloud := validSdkConfig()
	opt := v1.NewListGlobalPoolsRequest("glb-2e550a10-8a9e-4e0e-9086-80d8297ca3f7")
	pools, sdkerr := vngcloud.GLB.ListGlobalPools(context.Background(), opt)
	if sdkerr != nil {
		t.Fatalf("Expect nil but got %+v", sdkerr)
	}

	if pools == nil {
		t.Fatalf("Expect not nil but got nil")
	}

	t.Logf("Result: %+v", pools)
	for _, pool := range pools.Items {
		t.Logf("Pool: %+v", pool)
		t.Logf("Health: %+v", pool.Health)
	}
	t.Log("PASS")
}

func TestCreateGlobalPoolSuccess(t *testing.T) {
	vngcloud := validSdkConfig()
	member := v1.NewGlobalMemberRequest("p_name", "10.105.0.4", "sub-8aa727dd-9857-472f-8766-ece41282d437", 80, 80, 1, false)
	poolMember := v1.NewGlobalPoolMemberRequest("p_name", "hcm", "net-80a4eb74-c7d9-46b4-9705-ffed0e2bc3c2", 100, v1.GlobalPoolMemberTypePublic)
	poolMember.Members = append(poolMember.Members, member)
	opt := v1.NewCreateGlobalPoolRequest("test-pool-4", v1.GlobalPoolProtocolTCP)
	opt.LoadBalancerID = "glb-2e550a10-8a9e-4e0e-9086-80d8297ca3f7"
	opt.HealthMonitor = v1.NewGlobalHealthMonitor(v1.GlobalPoolHealthCheckProtocolTCP)
	opt.GlobalPoolMembers = append(opt.GlobalPoolMembers, poolMember)
	pool, sdkerr := vngcloud.GLB.CreateGlobalPool(context.Background(), opt)

	if sdkerr != nil {
		var sdkErr *sdkerror.SdkError
		if errors.As(sdkerr, &sdkErr) {
			t.Log(sdkErr.Err())
			t.Log(sdkErr.ErrorCode())
			t.Log(sdkErr.GetMessage())
			t.Log(sdkErr.ErrorCategories())
		}
		t.Fatalf("Expect nil but got %+v", sdkerr)
	}

	if pool == nil {
		t.Fatalf("Expect not nil but got nil")
	}

	t.Log("Result: ", pool)
	t.Log("PASS")
}

func TestCreateGlobalPoolHTTPSSuccess(t *testing.T) {
	vngcloud := validSdkConfig()
	member := v1.NewGlobalMemberRequest("p_name", "10.105.0.4", "sub-8aa727dd-9857-472f-8766-ece41282d437", 80, 80, 1, false)
	poolMember := v1.NewGlobalPoolMemberRequest("p_name", "hcm", "net-80a4eb74-c7d9-46b4-9705-ffed0e2bc3c2", 100, v1.GlobalPoolMemberTypePrivate)
	poolMember.Members = append(poolMember.Members, member)
	healthMonitor := v1.NewGlobalHealthMonitor(v1.GlobalPoolHealthCheckProtocolHTTPS)
	healthMonitor.HTTPMethod = common.Ptr(v1.GlobalPoolHealthCheckMethodGET)
	healthMonitor.Path = common.Ptr("/sfdsaf")
	healthMonitor.HTTPVersion = common.Ptr(v1.GlobalPoolHealthCheckHTTPVersionHTTP1Minor1)
	healthMonitor.SuccessCode = common.Ptr("200")
	healthMonitor.DomainName = common.Ptr("example.com")
	opt := v1.NewCreateGlobalPoolRequest("test-pool-5", v1.GlobalPoolProtocolTCP)
	opt.LoadBalancerID = "glb-2e550a10-8a9e-4e0e-9086-80d8297ca3f7"
	opt.HealthMonitor = healthMonitor
	opt.GlobalPoolMembers = append(opt.GlobalPoolMembers, poolMember)
	pool, sdkerr := vngcloud.GLB.CreateGlobalPool(context.Background(), opt)

	if sdkerr != nil {
		var sdkErr *sdkerror.SdkError
		if errors.As(sdkerr, &sdkErr) {
			t.Log(sdkErr.Err())
			t.Log(sdkErr.ErrorCode())
			t.Log(sdkErr.GetMessage())
			t.Log(sdkErr.ErrorCategories())
		}
		t.Fatalf("Expect nil but got %+v", sdkerr)
	}

	if pool == nil {
		t.Fatalf("Expect not nil but got nil")
	}

	t.Log("Result: ", pool)
	t.Log("PASS")
}

func TestUpdateGlobalPoolHTTPSSuccess(t *testing.T) {
	vngcloud := validSdkConfig()
	httpMonitor := v1.NewGlobalHealthMonitor(v1.GlobalPoolHealthCheckProtocolHTTPS)
	httpMonitor.DomainName = common.Ptr("exampleee.com")
	httpMonitor.HTTPMethod = common.Ptr(v1.GlobalPoolHealthCheckMethodPOST)
	httpMonitor.Path = common.Ptr("/hghjgj")
	httpMonitor.HTTPVersion = common.Ptr(v1.GlobalPoolHealthCheckHTTPVersionHTTP1Minor1)
	opt := v1.NewUpdateGlobalPoolRequest("glb-2e550a10-8a9e-4e0e-9086-80d8297ca3f7", "gpool-30c2a387-7912-4be7-8e3b-448ef16548ab")
	opt.HealthMonitor = httpMonitor

	pool, sdkerr := vngcloud.GLB.UpdateGlobalPool(context.Background(), opt)

	if sdkerr != nil {
		var sdkErr *sdkerror.SdkError
		if errors.As(sdkerr, &sdkErr) {
			t.Log(sdkErr.Err())
			t.Log(sdkErr.ErrorCode())
			t.Log(sdkErr.GetMessage())
			t.Log(sdkErr.ErrorCategories())
		}
		t.Fatalf("Expect nil but got %+v", sdkerr)
	}

	if pool == nil {
		t.Fatalf("Expect not nil but got nil")
	}

	t.Log("Result: ", pool)
	t.Log("PASS")
}

func TestDeleteGlobalPoolSuccess(t *testing.T) {
	vngcloud := validSdkConfig()
	opt := v1.NewDeleteGlobalPoolRequest("glb-2e550a10-8a9e-4e0e-9086-80d8297ca3f7", "gpool-1ffbe2f4-0bb5-4272-afe9-0dfa8a4365df")
	sdkerr := vngcloud.GLB.DeleteGlobalPool(context.Background(), opt)

	if sdkerr != nil {
		var sdkErr *sdkerror.SdkError
		if errors.As(sdkerr, &sdkErr) {
			t.Log(sdkErr.Err())
			t.Log(sdkErr.ErrorCode())
			t.Log(sdkErr.GetMessage())
			t.Log(sdkErr.ErrorCategories())
		}
		t.Fatalf("Expect nil but got %+v", sdkerr)
	}

	t.Log("PASS")
}

func TestListGlobalPoolMembersSuccess(t *testing.T) {
	vngcloud := validSdkConfig()
	opt := v1.NewListGlobalPoolMembersRequest("glb-2e550a10-8a9e-4e0e-9086-80d8297ca3f7", "gpool-0f4ba08b-e09d-4a1c-b953-523179cea006")
	members, sdkerr := vngcloud.GLB.ListGlobalPoolMembers(context.Background(), opt)
	if sdkerr != nil {
		t.Fatalf("Expect nil but got %+v", sdkerr)
	}

	if members == nil {
		t.Fatalf("Expect not nil but got nil")
	}

	t.Logf("Result: %+v", members)
	for _, member := range members.Items {
		t.Logf("Member: %+v", member)
		for _, m := range member.Members.Items {
			t.Logf("  - Member: %+v", m)
		}
	}
	t.Log("PASS")
}

func TestPatchGlobalPoolMemberSuccess(t *testing.T) {
	vngcloud := validSdkConfig()
	patchMember := v1.NewGlobalPoolMemberRequest("patch_name", "hcm", "net-86b7c84a-b3dd-4e6a-b66b-d28f36f3fc5f", 100, v1.GlobalPoolMemberTypePublic)
	patchMember.Members = append(patchMember.Members,
		v1.NewGlobalMemberRequest("patch_name_4", "10.105.0.4", "sub-7ceeed28-2cad-4bcd-9a4a-a0041c6d6304", 80, 80, 1, false),
		v1.NewGlobalMemberRequest("patch_name_3", "10.105.0.3", "sub-a7fceae7-5ab5-4768-993f-8e6465f75050", 80, 80, 1, false),
	)
	createAction := v1.NewPatchGlobalPoolCreateBulkActionRequest(patchMember)
	updateMember := v1.NewUpdateGlobalPoolMemberRequest("", "", "", 100)
	updateMember.Members = append(updateMember.Members,
		v1.NewGlobalMemberRequest("patch_name_44", "10.105.0.44", "sub-7ceeed28-2cad-4bcd-9a4a-a0041c6d6304", 80, 80, 1, false),
		v1.NewGlobalMemberRequest("patch_name_33", "10.105.0.33", "sub-a7fceae7-5ab5-4768-993f-8e6465f75050", 80, 80, 1, false),
	)
	updateAction := v1.NewPatchGlobalPoolUpdateBulkActionRequest("gpool-mem-4568b7da-e82b-4417-b991-ac040967c0c1", updateMember)
	// deleteAction := v1.NewPatchGlobalPoolDeleteBulkActionRequest("gpool-mem-e4a56d03-baf8-448b-98ab-404219fddede")
	opt := v1.NewPatchGlobalPoolMembersRequest("glb-2e550a10-8a9e-4e0e-9086-80d8297ca3f7", "gpool-0f4ba08b-e09d-4a1c-b953-523179cea006")
	opt.BulkActions = []any{
		createAction,
		// deleteAction,
		updateAction,
	}
	sdkerr := vngcloud.GLB.PatchGlobalPoolMembers(context.Background(), opt)
	if sdkerr != nil {
		t.Fatalf("Expect nil but got %+v", sdkerr)
	}

	t.Log("PASS")
}

func TestListGlobalListenersSuccess(t *testing.T) {
	vngcloud := validSdkConfig()
	opt := v1.NewListGlobalListenersRequest("glb-2e550a10-8a9e-4e0e-9086-80d8297ca3f7")
	listeners, sdkerr := vngcloud.GLB.ListGlobalListeners(context.Background(), opt)
	if sdkerr != nil {
		t.Fatalf("Expect nil but got %+v", sdkerr)
	}

	if listeners == nil {
		t.Fatalf("Expect not nil but got nil")
	}

	t.Logf("Result: %+v", listeners)
	for _, listener := range listeners.Items {
		t.Logf("Listener: %+v", listener)
	}
	t.Log("PASS")
}

func TestCreateGlobalListenerSuccess(t *testing.T) {
	vngcloud := validSdkConfig()
	opt := v1.NewCreateGlobalListenerRequest("glb-2e550a10-8a9e-4e0e-9086-80d8297ca3f7", "test-listener")
	opt.Description = "hihi"
	opt.Port = 85
	opt.TimeoutClient = 50
	opt.TimeoutConnection = 5
	opt.TimeoutMember = 50
	opt.GlobalPoolID = "gpool-7000d491-b441-40a0-af01-8039baa8e346"
	listener, sdkerr := vngcloud.GLB.CreateGlobalListener(context.Background(), opt)
	if sdkerr != nil {
		t.Fatalf("Expect nil but got %+v", sdkerr)
	}

	if listener == nil {
		t.Fatalf("Expect not nil but got nil")
	}

	t.Logf("Result: %+v", listener)
	t.Log("PASS")
}

func TestUpdateGlobalListenerSuccess(t *testing.T) {
	vngcloud := validSdkConfig()
	opt := v1.NewUpdateGlobalListenerRequest("glb-2e550a10-8a9e-4e0e-9086-80d8297ca3f7", "glis-7ffc4f19-7218-4d38-8016-e3ad2401e3bd")
	opt.TimeoutClient = 60
	opt.TimeoutConnection = 6
	opt.TimeoutMember = 60
	opt.GlobalPoolID = "gpool-7000d491-b441-40a0-af01-8039baa8e346"
	listener, sdkerr := vngcloud.GLB.UpdateGlobalListener(context.Background(), opt)
	if sdkerr != nil {
		t.Fatalf("Expect nil but got %+v", sdkerr)
	}

	if listener == nil {
		t.Fatalf("Expect not nil but got nil")
	}

	t.Logf("Result: %+v", listener)
	t.Log("PASS")
}

func TestDeleteGlobalListenerSuccess(t *testing.T) {
	vngcloud := validSdkConfig()
	opt := v1.NewDeleteGlobalListenerRequest("glb-2e550a10-8a9e-4e0e-9086-80d8297ca3f7", "glis-7ffc4f19-7218-4d38-8016-e3ad2401e3bd")
	sdkerr := vngcloud.GLB.DeleteGlobalListener(context.Background(), opt)
	if sdkerr != nil {
		t.Fatalf("Expect nil but got %+v", sdkerr)
	}

	t.Log("PASS")
}

func TestListGlobalLoadBalancerSuccess(t *testing.T) {
	vngcloud := validSdkConfig()
	opt := v1.NewListGlobalLoadBalancersRequest(0, 10)
	lbs, sdkerr := vngcloud.GLB.ListGlobalLoadBalancers(context.Background(), opt)
	if sdkerr != nil {
		t.Fatalf("Expect nil but got %+v", sdkerr)
	}

	if lbs == nil {
		t.Fatalf("Expect not nil but got nil")
	}

	t.Logf("Result: %+v", lbs)
	for _, lb := range lbs.Items {
		t.Logf("LB: %+v", lb)
		for _, vip := range lb.Vips {
			t.Logf("  - VIP: %+v", vip)
		}
		for _, domain := range lb.Domains {
			t.Logf("  - Domain: %+v", domain)
		}
	}
	t.Log("PASS")
}

func TestCreateGlobalLoadBalancerSuccess(t *testing.T) {
	healthMonitor := v1.NewGlobalHealthMonitor(v1.GlobalPoolHealthCheckProtocolHTTPS)
	healthMonitor.HTTPMethod = common.Ptr(v1.GlobalPoolHealthCheckMethodGET)
	healthMonitor.Path = common.Ptr("/sfdsaf")
	healthMonitor.HTTPVersion = common.Ptr(v1.GlobalPoolHealthCheckHTTPVersionHTTP1Minor1)
	healthMonitor.SuccessCode = common.Ptr("200")
	healthMonitor.DomainName = common.Ptr("example.com")
	poolMember := v1.NewGlobalPoolMemberRequest(
		"p_name",
		"hcm",
		"net-80a4eb74-c7d9-46b4-9705-ffed0e2bc3c2",
		100,
		v1.GlobalPoolMemberTypePrivate)
	poolMember.Members = append(poolMember.Members,
		v1.NewGlobalMemberRequest("p_name", "10.105.0.4", "sub-8aa727dd-9857-472f-8766-ece41282d437", 80, 80, 1, false),
	)
	pool := v1.NewCreateGlobalPoolRequest("test-pool-5", v1.GlobalPoolProtocolTCP)
	pool.LoadBalancerID = "glb-2e550a10-8a9e-4e0e-9086-80d8297ca3f7"
	pool.HealthMonitor = healthMonitor
	pool.GlobalPoolMembers = append(pool.GlobalPoolMembers, poolMember)
	listener := v1.NewCreateGlobalListenerRequest("glb-2e550a10-8a9e-4e0e-9086-80d8297ca3f7", "test-listener")
	listener.Description = "hihi"
	listener.Port = 85
	listener.TimeoutClient = 50
	listener.TimeoutConnection = 5
	listener.TimeoutMember = 50
	listener.GlobalPoolID = "gpool-7000d491-b441-40a0-af01-8039baa8e346"
	vngcloud := validSdkConfig()
	opt := v1.NewCreateGlobalLoadBalancerRequest("test-glb")
	opt.Description = "hihi"
	opt.GlobalListener = listener
	opt.GlobalPool = pool
	opt.Package = "pkg-b02e62ab-a282-4faf-8732-a172ef497a7b"

	lb, sdkerr := vngcloud.GLB.CreateGlobalLoadBalancer(context.Background(), opt)
	if sdkerr != nil {
		t.Fatalf("Expect nil but got %+v", sdkerr)
	}

	if lb == nil {
		t.Fatalf("Expect not nil but got nil")
	}

	t.Logf("Result: %+v", lb)
	t.Log("PASS")
}

func TestDeleteGlobalLoadBalancerSuccess(t *testing.T) {
	vngcloud := validSdkConfig()
	opt := v1.NewDeleteGlobalLoadBalancerRequest("glb-3fd57a7e-7bb3-4152-a329-adba6d779c4a")
	sdkerr := vngcloud.GLB.DeleteGlobalLoadBalancer(context.Background(), opt)
	if sdkerr != nil {
		t.Fatalf("Expect nil but got %+v", sdkerr)
	}

	t.Log("PASS")
}

func TestGetGlobalLoadBalancerSuccess(t *testing.T) {
	vngcloud := validSdkConfig()
	opt := v1.NewGetGlobalLoadBalancerByIDRequest("glb-2e550a10-8a9e-4e0e-9086-80d8297ca3f7")
	lb, sdkerr := vngcloud.GLB.GetGlobalLoadBalancerByID(context.Background(), opt)
	if sdkerr != nil {
		t.Fatalf("Expect nil but got %+v", sdkerr)
	}

	if lb == nil {
		t.Fatalf("Expect not nil but got nil")
	}

	t.Logf("Result: %+v", lb)
	for _, vip := range lb.Vips {
		t.Logf("  - VIP: %+v", vip)
	}
	for _, domain := range lb.Domains {
		t.Logf("  - Domain: %+v", domain)
	}
	t.Log("PASS")
}
