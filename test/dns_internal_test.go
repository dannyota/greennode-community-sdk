package test

import (
	"context"
	"testing"

	dnsinternalv1 "github.com/dannyota/greennode-community-sdk/v2/greennode/services/dns/internal_system/v1"
)

func TestDnsServiceInternal_ListHostedZonesDefault(t *testing.T) {
	vngcloud := validSdkConfig()
	portalUserID := "53461"

	req := dnsinternalv1.NewListHostedZonesRequest().WithName("")

	resp, sdkErr := vngcloud.VDnsGateway().Internal().DnsService().ListHostedZones(context.Background(), req, portalUserID)
	if sdkErr != nil {
		t.Fatalf("Failed to list hosted zones: %v", sdkErr)
	}

	if resp == nil {
		t.Fatal("Response should not be nil")
	}

	t.Logf("Successfully listed hosted zones. Count: %d", len(resp.ListData))
	for _, zone := range resp.ListData {
		t.Logf("Hosted Zone: %+v", zone)
	}
}

func TestDnsServiceInternal_ListRecordsDefault(t *testing.T) {
	vngcloud := validSdkConfig()
	portalUserID := "53461"
	hostedZoneID := "hosted-zone-5ba95110-e8d1-40f5-92b2-50eefa33bde2"

	req := dnsinternalv1.NewListRecordsRequest(hostedZoneID)

	resp, sdkErr := vngcloud.VDnsGateway().Internal().DnsService().ListRecords(context.Background(), req, portalUserID)
	if sdkErr != nil {
		t.Fatalf("Failed to list record: %v", sdkErr)
	}

	if resp == nil {
		t.Fatal("Response should not be nil")
	}

	t.Logf("Successfully listed record. Count: %d", len(resp.ListData))
	for _, record := range resp.ListData {
		t.Logf("Record: %+v", record)
	}
}

func TestDnsServiceInternal_DeleteRecord(t *testing.T) {
	vngcloud := validSdkConfig()
	portalUserID := "53461"
	hostedZoneID := "hosted-zone-a6acbf48-9f7b-455d-b5be-efa9081722f9"
	recordID := "record-test-id"

	req := dnsinternalv1.NewDeleteRecordRequest(hostedZoneID, recordID)

	sdkErr := vngcloud.VDnsGateway().Internal().DnsService().DeleteRecord(context.Background(), req, portalUserID)
	if sdkErr != nil {
		t.Fatalf("Failed to delete record: %v", sdkErr)
	}

	t.Logf("Successfully deleted record with ID: %s", recordID)
}

func TestDnsServiceInternal_DeleteHostedZone(t *testing.T) {
	vngcloud := validSdkConfig()
	portalUserID := "53461"
	hostedZoneID := "hosted-zone-243d64ba-c955-4aa1-8640-95d3463446d8"

	req := dnsinternalv1.NewDeleteHostedZoneRequest(hostedZoneID)

	sdkErr := vngcloud.VDnsGateway().Internal().DnsService().DeleteHostedZone(context.Background(), req, portalUserID)
	if sdkErr != nil {
		t.Fatalf("Failed to delete hosted zone: %v", sdkErr)
	}

	t.Logf("Successfully deleted hosted zone with ID: %s", hostedZoneID)
}

func TestDnsServiceInternal_ListHostedZonesByVpc(t *testing.T) {
	vngcloud := validSdkConfig()
	portalUserID := "53461"
	targetVpcID := "net-5ed4bdc1-99d9-4d20-aea8-ce4049d9261d" // Replace with actual VPC ID to test

	// List all hosted zones
	req := dnsinternalv1.NewListHostedZonesRequest()
	resp, sdkErr := vngcloud.VDnsGateway().Internal().DnsService().ListHostedZones(context.Background(), req, portalUserID)
	if sdkErr != nil {
		t.Fatalf("Failed to list hosted zones: %v", sdkErr)
	}

	if resp == nil {
		t.Fatal("Response should not be nil")
	}

	t.Logf("Total hosted zones found: %d", len(resp.ListData))

	// Filter hosted zones that contain the target VPC ID
	var matchingZones []map[string]any
	for _, zone := range resp.ListData {
		for _, vpcID := range zone.AssocVpcIDs {
			if vpcID == targetVpcID {
				matchingZones = append(matchingZones, map[string]any{
					"hostedZoneId": zone.HostedZoneID,
					"domainName":   zone.DomainName,
					"type":         zone.Type,
					"status":       zone.Status,
					"assocVpcIds":  zone.AssocVpcIDs,
				})
				break
			}
		}
	}

	// Print results
	t.Logf("Found %d hosted zones associated with VPC ID: %s", len(matchingZones), targetVpcID)
	for i, zone := range matchingZones {
		t.Logf("Zone %d:", i+1)
		t.Logf("  - Hosted Zone ID: %s", zone["hostedZoneId"])
		t.Logf("  - Domain Name: %s", zone["domainName"])
		t.Logf("  - Type: %s", zone["type"])
		t.Logf("  - Status: %s", zone["status"])
		t.Logf("  - Associated VPCs: %v", zone["assocVpcIds"])
	}

	if len(matchingZones) == 0 {
		t.Logf("No hosted zones found for VPC ID: %s", targetVpcID)
	}
}

func TestDnsServiceInternal_ListAndDeleteAllRecords(t *testing.T) {
	vngcloud := validSdkConfig()
	portalUserID := "53461"
	hostedZoneID := "hosted-zone-243d64ba-c955-4aa1-8640-95d3463446d8"

	// First, list all records in the hosted zone
	listReq := dnsinternalv1.NewListRecordsRequest(hostedZoneID)
	listResp, sdkErr := vngcloud.VDnsGateway().Internal().DnsService().ListRecords(context.Background(), listReq, portalUserID)
	if sdkErr != nil {
		t.Fatalf("Failed to list records: %v", sdkErr)
	}

	if listResp == nil {
		t.Fatal("List response should not be nil")
	}

	t.Logf("Found %d records in hosted zone %s", len(listResp.ListData), hostedZoneID)

	// Delete each record (skip NS records as they cannot be deleted)
	allowedTypes := map[string]bool{
		"A":     true,
		"CNAME": true,
		"MX":    true,
		"PTR":   true,
		"TXT":   true,
		"SRV":   true,
	}

	for _, record := range listResp.ListData {
		t.Logf("Found record: ID=%s, SubDomain=%s, Type=%s", record.RecordID, record.SubDomain, record.Type)

		// Skip NS records as they cannot be deleted
		if !allowedTypes[record.Type] {
			t.Logf("Skipping record %s of type %s (cannot be deleted)", record.RecordID, record.Type)
			continue
		}

		deleteReq := dnsinternalv1.NewDeleteRecordRequest(hostedZoneID, record.RecordID)
		sdkErr := vngcloud.VDnsGateway().Internal().DnsService().DeleteRecord(context.Background(), deleteReq, portalUserID)
		if sdkErr != nil {
			t.Logf("Failed to delete record %s: %v", record.RecordID, sdkErr)
			continue
		}

		t.Logf("Successfully deleted record: %s", record.RecordID)
	}

	// req := dnsinternalv1.NewDeleteHostedZoneRequest(hostedZoneId)

	// sdkErr = vngcloud.VDnsGateway().Internal().DnsService().DeleteHostedZone(req, portalUserId)
	// if sdkErr != nil {
	// 	t.Fatalf("Failed to delete hosted zone: %v", sdkErr)
	// }

	t.Logf("Completed deletion of all records in hosted zone %s", hostedZoneID)
}
