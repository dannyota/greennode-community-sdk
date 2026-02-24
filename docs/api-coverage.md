# API Coverage

What the SDK supports today vs. what GreenNode exposes. Based on
[hotpot feature inventory](/home/danny/src/hotpot/docs/features/GREENNODE.md).

| Service | ✅ | Missing | Total |
|---------|:---:|:-------:|:-----:|
| Portal | 6 | 0 | 6 |
| Compute | 14 | 8 | 22 |
| Network | 18 | 14 | 32 |
| vNetwork | 4 | 1 | 5 |
| Volume | 14 | 2 | 16 |
| Load Balancer | 30 | 0 | 30 |
| Global LB | 21 | 0 | 21 |
| DNS | 10 | 0 | 10 |
| IAM | 1 | 1 | 2 |
| Kubernetes | 0 | 2 | 2 |
| Database | 0 | 7 | 7 |
| Object Storage | 0 | 3 | 3 |
| Monitoring | 0 | 4 | 4 |
| Billing | 0 | 1 | 1 |
| **Total** | **118** | **43** | **161** |

---

## Portal (`Client.Portal`, `Client.PortalV1`)

| Resource | SDK Method | REST Endpoint | |
|----------|-----------|---------------|:---:|
| Regions | `Portal.ListRegions()` | `GET /v2/{projectId}/region` | ✅ |
| Quotas | `Portal.ListAllQuotaUsed()` | `GET /v2/{projectId}/quotas/quotaUsed` | ✅ |
| Quota by Name | `Portal.GetQuotaByName()` | (wraps ListAllQuotaUsed) | ✅ |
| Portal Info | `PortalV1.GetPortalInfo()` | `GET /v1/projects/{id}/detail` | ✅ |
| Projects | `PortalV1.ListProjects()` | `GET /v1/projects` | ✅ |
| Zones | `PortalV1.ListZones()` | `GET /v1/{projectId}/zones` | ✅ |

**6/6**

## Compute (`Client.Compute`)

| Resource | SDK Method | REST Endpoint | |
|----------|-----------|---------------|:---:|
| List Servers | `ListServers()` | `GET /v2/{projectId}/servers` | ✅ |
| Get Server | `GetServerByID()` | `GET /v2/{projectId}/servers/{id}` | ✅ |
| Create Server | `CreateServer()` | `POST /v2/{projectId}/servers` | ✅ |
| Delete Server | `DeleteServerByID()` | `DELETE /v2/{projectId}/servers/{id}` | ✅ |
| Update Server Secgroups | `UpdateServerSecgroupsByServerID()` | `PUT /v2/{projectId}/servers/{id}/update-sec-group` | ✅ |
| Attach Volume | `AttachBlockVolume()` | `PUT /v2/{projectId}/volumes/{id}/servers/{id}/attach` | ✅ |
| Detach Volume | `DetachBlockVolume()` | `PUT /v2/{projectId}/volumes/{id}/servers/{id}/detach` | ✅ |
| Attach Floating IP | `AttachFloatingIp()` | `PUT /v2/{projectId}/servers/{id}/wan-ips/auto/attach` | ✅ |
| Detach Floating IP | `DetachFloatingIp()` | `PUT /v2/{projectId}/servers/{id}/wan-ips/{id}/detach` | ✅ |
| List SSH Keys | `ListSSHKeys()` | `GET /v2/{projectId}/sshKeys` | ✅ |
| List Server Groups | `ListServerGroups()` | `GET /v2/{projectId}/serverGroups` | ✅ |
| Create Server Group | `CreateServerGroup()` | `POST /v2/{projectId}/serverGroups` | ✅ |
| Delete Server Group | `DeleteServerGroupByID()` | `DELETE /v2/{projectId}/serverGroups/{id}` | ✅ |
| List Server Group Policies | `ListServerGroupPolicies()` | `GET /v2/{projectId}/serverGroups/policies` | ✅ |
| OS Images | — | `GET /v1/{projectId}/images/os` | |
| GPU Images | — | `GET /v1/{projectId}/images/gpu` | |
| User Images | — | `GET /v2/{projectId}/user-images` | |
| Flavors | — | `GET /v1/{projectId}/flavors/families/{family}/platforms/{code}` | |
| Flavor Zones | — | `GET /v1/{projectId}/flavor_zones/product` | |
| Flavor Families | — | `GET /v1/{projectId}/flavor_zones/families` | |
| Tags | — | `GET /v2/{projectId}/tag` | |
| Tag Keys | — | `GET /v2/{projectId}/tag/tag-key` | |

**14/22**

## Network (`Client.Network`)

| Resource | SDK Method | REST Endpoint | |
|----------|-----------|---------------|:---:|
| Get Network | `GetNetworkByID()` | `GET /v2/{projectId}/networks/{id}` | ✅ |
| Get Subnet | `GetSubnetByID()` | `GET /v2/{projectId}/networks/{id}/subnets/{id}` | ✅ |
| Update Subnet | `UpdateSubnetByID()` | `PATCH /v2/{projectId}/networks/{id}/subnets/{id}` | ✅ |
| List Secgroups | `ListSecgroup()` | `GET /v2/{projectId}/secgroups` | ✅ |
| Get Secgroup | `GetSecgroupByID()` | `GET /v2/{projectId}/secgroups/{id}` | ✅ |
| Create Secgroup | `CreateSecgroup()` | `POST /v2/{projectId}/secgroups` | ✅ |
| Delete Secgroup | `DeleteSecgroupByID()` | `DELETE /v2/{projectId}/secgroups/{id}` | ✅ |
| List Secgroup Rules | `ListSecgroupRulesBySecgroupID()` | `GET /v2/{projectId}/secgroups/{id}/secGroupRules` | ✅ |
| Create Secgroup Rule | `CreateSecgroupRule()` | `POST /v2/{projectId}/secgroups/{id}/secgroupRules` | ✅ |
| Delete Secgroup Rule | `DeleteSecgroupRuleByID()` | `DELETE /v2/{projectId}/secgroups/{id}/secgroupRules/{id}` | ✅ |
| Create Virtual Address | `CreateVirtualAddressCrossProject()` | `POST /v2/{projectId}/virtualIpAddress` | ✅ |
| Get Virtual Address | `GetVirtualAddressByID()` | `GET /v2/{projectId}/virtualIpAddress/{id}` | ✅ |
| Delete Virtual Address | `DeleteVirtualAddressByID()` | `DELETE /v2/{projectId}/virtualIpAddress/{id}` | ✅ |
| List Address Pairs | `ListAddressPairsByVirtualAddressID()` | `GET /v2/{projectId}/virtualIpAddress/{id}/addressPairs` | ✅ |
| Create Address Pair | `CreateAddressPair()` | `POST /v2/{projectId}/virtualIpAddress/{id}/addressPairs` | ✅ |
| Delete Address Pair | `DeleteAddressPair()` | `DELETE /v2/{projectId}/virtual-subnets/addressPairs/{id}` | ✅ |
| Get All Address Pairs by Subnet | `GetAllAddressPairByVirtualSubnetID()` | `GET /v2/{projectId}/virtual-subnets/{id}/addressPairs` | ✅ |
| Set Address Pair in Subnet | `SetAddressPairInVirtualSubnet()` | `POST /v2/{projectId}/virtual-subnets/{id}/addressPairs` | ✅ |
| List Networks (VPCs) | — | `GET /v2/{projectId}/networks` | |
| List Subnets | — | `GET /v2/{projectId}/networks/{id}/subnets` | |
| Network ACLs | — | `GET /v2/{projectId}/network-acl/list` | |
| Network ACL Rules | — | `GET /v2/{projectId}/network-acl/{id}/rules` | |
| Route Tables | — | `GET /v2/{projectId}/route-table` | |
| Route Table Routes | — | `GET /v2/{projectId}/route-table/route/{id}` | |
| DHCP Options | — | `GET /v2/{projectId}/dhcp_option` | |
| Elastic IPs | — | `GET /v2/{projectId}/elastic-ips` | |
| Network Interfaces | — | `GET /v2/{projectId}/network-interfaces-elastic` | |
| Virtual IPs | — | `GET /v2/{projectId}/virtualIpAddress` | |
| Public VIPs | — | `GET /v2/{projectId}/public-vips/externalNetworkInterfaces` | |
| Peering | — | `GET /v2/{projectId}/peering` | |
| Interconnects | — | `GET /v2/{projectId}/interconnects` | |
| Interconnect Connections | — | `GET /v2/{projectId}/interconnects/{id}/connections` | |
| WAN IPs | — | `GET /v2/{projectId}/wanIps` | |

**18/32**

## vNetwork (`Client.NetworkV1`)

| Resource | SDK Method | REST Endpoint | |
|----------|-----------|---------------|:---:|
| List Endpoints | `ListEndpoints()` | `GET /vnetwork/v1/{regionId}/{projectId}/endpoints` | ✅ |
| Get Endpoint | `GetEndpointByID()` | `GET /vnetwork/v1/{regionId}/{projectId}/endpoints/{id}` | ✅ |
| Create Endpoint | `CreateEndpoint()` | `POST /vnetwork/v1/{regionId}/{projectId}/endpoints` | ✅ |
| Delete Endpoint | `DeleteEndpointByID()` | `DELETE /vnetwork/v1/{regionId}/{projectId}/endpoints/{id}` | ✅ |
| Regions | — | `GET /vnetwork/v1/regions` | |

**4/5**

## Volume (`Client.Volume`, `Client.VolumeV1`)

| Resource | SDK Method | REST Endpoint | |
|----------|-----------|---------------|:---:|
| List Volumes | `Volume.ListBlockVolumes()` | `GET /v2/{projectId}/volumes` | ✅ |
| Get Volume | `Volume.GetBlockVolumeByID()` | `GET /v2/{projectId}/volumes/{id}` | ✅ |
| Create Volume | `Volume.CreateBlockVolume()` | `POST /v2/{projectId}/volumes` | ✅ |
| Delete Volume | `Volume.DeleteBlockVolumeByID()` | `DELETE /v2/{projectId}/volumes/{id}` | ✅ |
| Resize Volume | `Volume.ResizeBlockVolumeByID()` | `PUT /v2/{projectId}/volumes/{id}/resize` | ✅ |
| Get Volume Mapping | `Volume.GetUnderBlockVolumeByID()` | `GET /v2/{projectId}/volumes/{id}/mapping` | ✅ |
| Migrate Volume | `Volume.MigrateBlockVolumeByID()` | `PUT /v2/{projectId}/volumes/{id}/change-device-type` | ✅ |
| List Snapshots | `Volume.ListSnapshotsByBlockVolumeID()` | `GET /v2/{projectId}/volumes/{id}/snapshots` | ✅ |
| Create Snapshot | `Volume.CreateSnapshotByBlockVolumeID()` | `POST /v2/{projectId}/volumes/{id}/snapshots` | ✅ |
| Delete Snapshot | `Volume.DeleteSnapshotByID()` | `DELETE /v2/{projectId}/volumes/{id}/snapshots/{id}` | ✅ |
| Volume Types | `VolumeV1.GetListVolumeTypes()` | `GET /v1/{volumeTypeZoneId}/volume_types` | ✅ |
| Volume Type by ID | `VolumeV1.GetVolumeTypeByID()` | `GET /v1/{projectId}/volume_types/{id}` | ✅ |
| Default Volume Type | `VolumeV1.GetDefaultVolumeType()` | `GET /v1/{projectId}/volume_default_id` | ✅ |
| Volume Type Zones | `VolumeV1.GetVolumeTypeZones()` | `GET /v1/{projectId}/volume_type_zones` | ✅ |
| Persistent Volumes | — | `GET /v2/{projectId}/persistent-volumes` | |
| Encryption Types | — | `GET /v1/{projectId}/volumes/encryption_types` | |

**14/16**

## Load Balancer (`Client.LoadBalancer`)

| Resource | SDK Method | REST Endpoint | |
|----------|-----------|---------------|:---:|
| List LBs | `ListLoadBalancers()` | `GET /v2/{projectId}/loadBalancers` | ✅ |
| Get LB | `GetLoadBalancerByID()` | `GET /v2/{projectId}/loadBalancers/{id}` | ✅ |
| Create LB | `CreateLoadBalancer()` | `POST /v2/{projectId}/loadBalancers` | ✅ |
| Delete LB | `DeleteLoadBalancerByID()` | `DELETE /v2/{projectId}/loadBalancers/{id}` | ✅ |
| Resize LB | `ResizeLoadBalancer()` | `PUT /v2/{projectId}/loadBalancers/{id}/resize` | ✅ |
| Scale LB | `ScaleLoadBalancer()` | `PUT /v2/{projectId}/loadBalancers/{id}/rebalancing` | ✅ |
| LB Packages | `ListLoadBalancerPackages()` | `GET /v2/{projectId}/loadBalancers/packages` | ✅ |
| List Listeners | `ListListenersByLoadBalancerID()` | `GET /v2/{projectId}/loadBalancers/{id}/listeners` | ✅ |
| Get Listener | `GetListenerByID()` | `GET /v2/{projectId}/loadBalancers/{id}/listeners/{id}` | ✅ |
| Create Listener | `CreateListener()` | `POST /v2/{projectId}/loadBalancers/{id}/listeners` | ✅ |
| Update Listener | `UpdateListener()` | `PUT /v2/{projectId}/loadBalancers/{id}/listeners/{id}` | ✅ |
| Delete Listener | `DeleteListenerByID()` | `DELETE /v2/{projectId}/loadBalancers/{id}/listeners/{id}` | ✅ |
| List Pools | `ListPoolsByLoadBalancerID()` | `GET /v2/{projectId}/loadBalancers/{id}/pools` | ✅ |
| Get Pool | `GetPoolByID()` | `GET /v2/{projectId}/loadBalancers/{id}/pools/{id}` | ✅ |
| Create Pool | `CreatePool()` | `POST /v2/{projectId}/loadBalancers/{id}/pools` | ✅ |
| Update Pool | `UpdatePool()` | `PUT /v2/{projectId}/loadBalancers/{id}/pools/{id}` | ✅ |
| Delete Pool | `DeletePoolByID()` | `DELETE /v2/{projectId}/loadBalancers/{id}/pools/{id}` | ✅ |
| Pool Health Monitor | `GetPoolHealthMonitorByID()` | `GET /v2/{projectId}/loadBalancers/{id}/pools/{id}/healthMonitor` | ✅ |
| List Pool Members | `ListPoolMembers()` | `GET /v2/{projectId}/loadBalancers/{id}/pools/{id}/members` | ✅ |
| Update Pool Members | `UpdatePoolMembers()` | `PUT /v2/{projectId}/loadBalancers/{id}/pools/{id}/members` | ✅ |
| List L7 Policies | `ListPolicies()` | `GET /v2/{projectId}/loadBalancers/{id}/listeners/{id}/l7policies` | ✅ |
| Get L7 Policy | `GetPolicyByID()` | `GET /v2/{projectId}/.../{id}/l7policies/{id}` | ✅ |
| Create L7 Policy | `CreatePolicy()` | `POST /v2/{projectId}/.../{id}/l7policies` | ✅ |
| Update L7 Policy | `UpdatePolicy()` | `PUT /v2/{projectId}/.../{id}/l7policies/{id}` | ✅ |
| Delete L7 Policy | `DeletePolicyByID()` | `DELETE /v2/{projectId}/.../{id}/l7policies/{id}` | ✅ |
| Reorder L7 Policies | `ReorderPolicies()` | `PUT /v2/{projectId}/.../{id}/reorderL7Policies` | ✅ |
| List Certificates | `ListCertificates()` | `GET /v2/{projectId}/cas` | ✅ |
| Get Certificate | `GetCertificateByID()` | `GET /v2/{projectId}/cas/{id}` | ✅ |
| Create Certificate | `CreateCertificate()` | `POST /v2/{projectId}/cas` | ✅ |
| Delete Certificate | `DeleteCertificateByID()` | `DELETE /v2/{projectId}/cas/{id}` | ✅ |

**30/30**

## Global Load Balancer (`Client.GLB`)

| Resource | SDK Method | REST Endpoint | |
|----------|-----------|---------------|:---:|
| List GLBs | `ListGlobalLoadBalancers()` | `GET /v1/global-load-balancers` | ✅ |
| Get GLB | `GetGlobalLoadBalancerByID()` | `GET /v1/global-load-balancers/{id}` | ✅ |
| Create GLB | `CreateGlobalLoadBalancer()` | `POST /v1/global-load-balancers` | ✅ |
| Delete GLB | `DeleteGlobalLoadBalancer()` | `DELETE /v1/global-load-balancers/{id}` | ✅ |
| Usage Histories | `GetGlobalLoadBalancerUsageHistories()` | `GET /v1/global-load-balancers/{id}/usage-histories` | ✅ |
| List Listeners | `ListGlobalListeners()` | `GET /v1/global-load-balancers/{id}/global-listeners` | ✅ |
| Get Listener | `GetGlobalListener()` | `GET /v1/global-load-balancers/{id}/global-listeners/{id}` | ✅ |
| Create Listener | `CreateGlobalListener()` | `POST /v1/global-load-balancers/{id}/global-listeners` | ✅ |
| Update Listener | `UpdateGlobalListener()` | `PUT /v1/global-load-balancers/{id}/global-listeners/{id}` | ✅ |
| Delete Listener | `DeleteGlobalListener()` | `DELETE /v1/global-load-balancers/{id}/global-listeners/{id}` | ✅ |
| List Pools | `ListGlobalPools()` | `GET /v1/global-load-balancers/{id}/global-pools` | ✅ |
| Create Pool | `CreateGlobalPool()` | `POST /v1/global-load-balancers/{id}/global-pools` | ✅ |
| Update Pool | `UpdateGlobalPool()` | `PUT /v1/global-load-balancers/{id}/global-pools/{id}` | ✅ |
| Delete Pool | `DeleteGlobalPool()` | `DELETE /v1/global-load-balancers/{id}/global-pools/{id}` | ✅ |
| List Pool Members | `ListGlobalPoolMembers()` | `GET /v1/.../global-pools/{id}/pool-members` | ✅ |
| Get Pool Member | `GetGlobalPoolMember()` | `GET /v1/.../pool-members/{id}` | ✅ |
| Update Pool Member | `UpdateGlobalPoolMember()` | `PUT /v1/.../pool-members/{id}` | ✅ |
| Delete Pool Member | `DeleteGlobalPoolMember()` | `DELETE /v1/.../pool-members/{id}` | ✅ |
| Patch Pool Members | `PatchGlobalPoolMembers()` | `PATCH /v1/.../pool-members` | ✅ |
| Packages | `ListGlobalPackages()` | `GET /v1/packages` | ✅ |
| Regions | `ListGlobalRegions()` | `GET /v1/regions` | ✅ |

**21/21**

## DNS (`Client.DNS`)

| Resource | SDK Method | REST Endpoint | |
|----------|-----------|---------------|:---:|
| List Hosted Zones | `ListHostedZones()` | `GET /v1/{projectId}/dns/hosted-zone` | ✅ |
| Get Hosted Zone | `GetHostedZoneByID()` | `GET /v1/{projectId}/dns/hosted-zone/{id}` | ✅ |
| Create Hosted Zone | `CreateHostedZone()` | `POST /v1/{projectId}/dns/hosted-zone` | ✅ |
| Delete Hosted Zone | `DeleteHostedZone()` | `DELETE /v1/{projectId}/dns/hosted-zone/{id}` | ✅ |
| Update Hosted Zone | `UpdateHostedZone()` | `PUT /v1/{projectId}/dns/hosted-zone/{id}` | ✅ |
| List Records | `ListRecords()` | `GET /v1/{projectId}/dns/hosted-zone/{id}/record` | ✅ |
| Get Record | `GetRecord()` | `GET /v1/{projectId}/dns/hosted-zone/{id}/record/{id}` | ✅ |
| Create Record | `CreateDnsRecord()` | `POST /v1/{projectId}/dns/hosted-zone/{id}/record` | ✅ |
| Update Record | `UpdateRecord()` | `PUT /v1/{projectId}/dns/hosted-zone/{id}/record/{id}` | ✅ |
| Delete Record | `DeleteRecord()` | `DELETE /v1/{projectId}/dns/hosted-zone/{id}/record/{id}` | ✅ |

**10/10**

## IAM (`Client.Identity`)

| Resource | SDK Method | REST Endpoint | |
|----------|-----------|---------------|:---:|
| Access Token | `GetAccessToken()` | `POST /accounts-api/v2/auth/token` | ✅ |
| User Info | — | `GET /accounts-api/v1/auth/userinfo` | |

**1/2**

## Kubernetes (vKS) — No SDK client

| Resource | REST Endpoint | |
|----------|---------------|:---:|
| Clusters | `GET /v2/{projectId}/clusters` | |
| Cluster Node Groups | `GET /v2/{projectId}/clusters/{id}/nodeGroups` | |

**0/2**

## Database (vDB) — No SDK client

| Resource | REST Endpoint | |
|----------|---------------|:---:|
| Relational DBs | `GET /v1/{projectId}/relational/databases` | |
| Relational Backups | `GET /v1/{projectId}/relational/backups` | |
| Relational Config Groups | `GET /v1/{projectId}/relational/config-groups` | |
| Memstore DBs | `GET /v1/{projectId}/memstore/databases` | |
| Memstore Backups | `GET /v1/{projectId}/memstore/backups` | |
| Memstore Config Groups | `GET /v1/{projectId}/memstore/config-groups` | |
| DB Packages | `GET /v1/{projectId}/packages` | |

**0/7**

## Object Storage (vStorage) — No SDK client

S3/Swift compatible — may not need SDK wrappers.

| Resource | REST Endpoint | |
|----------|---------------|:---:|
| Containers | Swift: `GET /v1/{account}` | |
| Objects | Swift: `GET /v1/{account}/{container}` | |
| Buckets | S3: `GET /` | |

**0/3**

## Monitoring (vMonitor) — No SDK client

Host: `vmonitor.console.vngcloud.vn`. Note: uses `lstData` wrapper, not `listData`.

| Resource | REST Endpoint | |
|----------|---------------|:---:|
| Dashboards | `GET /vmonitor-api/api/v1/dashboards` | |
| Configurations | `GET /vmonitor-api/api/v1/configurations/key/{key}` | |
| User Info | `GET /user-api/v1/userinfo` | |
| Quota Status | `POST /billing-api/v1/introspect-quotas` | |

**0/4**

## Billing — No SDK client

| Resource | REST Endpoint | |
|----------|---------------|:---:|
| User Info | `GET /v1/users/info` | |

**0/1**

## Internal/System Services

Exist in the SDK for internal use, not public API coverage:

| Service | Client Field | Methods |
|---------|-------------|---------|
| DNS Internal | `Client.DNSInternal` | 10 methods (mirrors DNS V1) |
| Network Internal | `Client.NetworkInternal` | 5 methods (endpoint tags) |
| LB Internal | `Client.LoadBalancerInternal` | 1 method (CreateLoadBalancer) |
| Server Internal | `Client.ServerInternal` | 1 method (CreateSystemTags) |

---

### Next priorities

1. **Compute gaps** — images, flavors, tags (needed for full server provisioning)
2. **Network gaps** — list networks, list subnets, elastic IPs, route tables, ACLs
3. **IAM user info** — single endpoint, quick win
4. **Kubernetes** — high demand, only 2 endpoints
5. **Volume gaps** — persistent volumes, encryption types
6. **vNetwork regions** — single endpoint, quick win
