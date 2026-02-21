//go:build integration

package test

import (
	"context"
	"testing"

	v1 "danny.vn/greennode/services/dns/v1"
)

func TestDnsServiceV1_GetHostedZoneByID(t *testing.T) {

	vngcloud := validSdkConfig()
	opt := v1.NewGetHostedZoneByIDRequest(
		"hosted-zone-32a21aa3-99a3-4d03-9045-37aa701fa03a",
	)
	hostedZone, sdkerr := vngcloud.DNS.GetHostedZoneByID(context.Background(), opt)
	if sdkerr != nil {
		t.Fatalf("Expect nil but got %+v", sdkerr)
	}

	if hostedZone == nil {
		t.Fatalf("Expect not nil but got nil")
	}

	t.Logf("Result: %+v", hostedZone)
	t.Log("PASS")
}

func TestDnsServiceV1_ListHostedZonesDefault(t *testing.T) {
	vngcloud := validSdkConfig()
	opt := v1.NewListHostedZonesRequest()
	listHostedZones, sdkerr := vngcloud.DNS.ListHostedZones(context.Background(), opt)
	if sdkerr != nil {
		t.Fatalf("Expect nil but got %+v", sdkerr)
	}

	if listHostedZones == nil {
		t.Fatalf("Expect not nil but got nil")
	}

	for _, hostedZone := range listHostedZones.ListData {
		t.Logf("HostedZone: %+v", hostedZone)
	}
	t.Log("PASS")
}

func TestDnsServiceV1_ListHostedZonesWithFilter(t *testing.T) {
	vngcloud := validSdkConfig()
	opt := v1.NewListHostedZonesRequest()
	opt.Name = "test-zone"
	listHostedZones, sdkerr := vngcloud.DNS.ListHostedZones(context.Background(), opt)
	if sdkerr != nil {
		t.Fatalf("Expect nil but got %+v", sdkerr)
	}

	if listHostedZones == nil {
		t.Fatalf("Expect not nil but got nil")
	}

	for _, hostedZone := range listHostedZones.ListData {
		t.Logf("HostedZone: %+v", hostedZone)
	}
	t.Log("PASS")
}

func TestDnsServiceV1_ListRecordsDefault(t *testing.T) {
	vngcloud := validSdkConfig()
	opt := v1.NewListRecordsRequest("hosted-zone-32a21aa3-99a3-4d03-9045-37aa701fa03a")
	listRecords, sdkerr := vngcloud.DNS.ListRecords(context.Background(), opt)
	if sdkerr != nil {
		t.Fatalf("Expect nil but got %+v", sdkerr)
	}

	if listRecords == nil {
		t.Fatalf("Expect not nil but got nil")
	}

	for _, record := range listRecords.ListData {
		t.Logf("Record: %+v", record)
	}
	t.Log("PASS")
}

func TestDnsServiceV1_ListRecordsWithFilter(t *testing.T) {
	vngcloud := validSdkConfig()
	opt := v1.NewListRecordsRequest("hosted-zone-32a21aa3-99a3-4d03-9045-37aa701fa03a")
	opt.Name = "k8s"
	listRecords, sdkerr := vngcloud.DNS.ListRecords(context.Background(), opt)
	if sdkerr != nil {
		t.Fatalf("Expect nil but got %+v", sdkerr)
	}

	if listRecords == nil {
		t.Fatalf("Expect not nil but got nil")
	}

	for _, record := range listRecords.ListData {
		t.Logf("Record: %+v", record)
	}
	t.Log("PASS")
}

func TestDnsServiceV1_CreateHostedZone(t *testing.T) {
	vngcloud := validSdkConfig()
	opt := v1.NewCreateHostedZoneRequest(
		"test-sdk.example.com",
		[]string{"net-dc14bb60-d500-40b5-945f-218540990187"},
		v1.HostedZoneTypePrivate,
	)
	opt.Description = "Test hosted zone created by SDK"

	hostedZone, sdkerr := vngcloud.DNS.CreateHostedZone(context.Background(), opt)
	if sdkerr != nil {
		t.Fatalf("Expect nil but got %+v", sdkerr)
	}

	if hostedZone == nil {
		t.Fatalf("Expect not nil but got nil")
	}

	t.Logf("Created HostedZone: %+v", hostedZone)
	t.Log("PASS")
}

func TestDnsServiceV1_DeleteHostedZone(t *testing.T) {
	vngcloud := validSdkConfig()
	opt := v1.NewDeleteHostedZoneRequest("hosted-zone-8d556e58-e84c-4dff-aeda-dc246b296f32")

	sdkerr := vngcloud.DNS.DeleteHostedZone(context.Background(), opt)
	if sdkerr != nil {
		t.Fatalf("Expect nil but got %+v", sdkerr)
	}

	t.Log("HostedZone deleted successfully")
	t.Log("PASS")
}

func TestDnsServiceV1_UpdateHostedZone(t *testing.T) {
	vngcloud := validSdkConfig()
	opt := v1.NewUpdateHostedZoneRequest("hosted-zone-32a21aa3-99a3-4d03-9045-37aa701fa03a")
	opt.AssocVpcIDs = []string{"net-dc14bb60-d500-40b5-945f-218540990187"}
	opt.Description = "Updated description for hosted zone."

	sdkerr := vngcloud.DNS.UpdateHostedZone(context.Background(), opt)
	if sdkerr != nil {
		t.Fatalf("Expect nil but got %+v", sdkerr)
	}

	t.Log("HostedZone updated successfully")
	t.Log("PASS")
}

func TestDnsServiceV1_CreateDnsRecord(t *testing.T) {
	vngcloud := validSdkConfig()
	hostedZoneID := "hosted-zone-32a21aa3-99a3-4d03-9045-37aa701fa03a"

	// Test 1: A record with multiple IP addresses
	t.Run("CreateARecord", func(t *testing.T) {
		weight1, weight2 := 10, 20
		aRecordValues := []v1.RecordValueRequest{
			v1.NewRecordValueRequest("10.0.0.1", nil, &weight1),
			v1.NewRecordValueRequest("10.0.0.2", nil, &weight2),
		}

		opt := v1.NewCreateDnsRecordRequest(
			hostedZoneID,
			"test-a-record",
			300,
			v1.DnsRecordTypeA,
			v1.RoutingPolicyWeighted,
			aRecordValues,
		)

		dnsRecord, sdkerr := vngcloud.DNS.CreateDnsRecord(context.Background(), opt)
		if sdkerr != nil {
			t.Fatalf("Expect nil but got %+v", sdkerr)
		}
		t.Logf("Created A Record: %+v", dnsRecord)
	})

	// Test 2: CNAME record
	t.Run("CreateCNAMERecord", func(t *testing.T) {
		cnameValues := []v1.RecordValueRequest{
			v1.NewRecordValueRequest("www.example.com", nil, nil),
		}

		stickySession := false
		opt := v1.NewCreateDnsRecordRequest(
			hostedZoneID,
			"test-cname",
			300,
			v1.DnsRecordTypeCNAME,
			v1.RoutingPolicySimple,
			cnameValues,
		)
		opt.EnableStickySession = &stickySession

		dnsRecord, sdkerr := vngcloud.DNS.CreateDnsRecord(context.Background(), opt)
		if sdkerr != nil {
			t.Fatalf("Expect nil but got %+v", sdkerr)
		}
		t.Logf("Created CNAME Record: %+v", dnsRecord)
	})

	// Test 3: TXT record with multiple values
	t.Run("CreateTXTRecord", func(t *testing.T) {
		txtValues := []v1.RecordValueRequest{
			v1.NewRecordValueRequest("v=spf1 include:_spf.google.com ~all", nil, nil),
			v1.NewRecordValueRequest("google-site-verification=abcd1234", nil, nil),
		}

		opt := v1.NewCreateDnsRecordRequest(
			hostedZoneID,
			"test-txt",
			600,
			v1.DnsRecordTypeTXT,
			v1.RoutingPolicySimple,
			txtValues,
		)

		dnsRecord, sdkerr := vngcloud.DNS.CreateDnsRecord(context.Background(), opt)
		if sdkerr != nil {
			t.Fatalf("Expect nil but got %+v", sdkerr)
		}
		t.Logf("Created TXT Record: %+v", dnsRecord)
	})

	// Test 4: MX record with multiple mail servers
	t.Run("CreateMXRecord", func(t *testing.T) {
		mxValues := []v1.RecordValueRequest{
			v1.NewRecordValueRequest("10 mail1.example.com", nil, nil),
			v1.NewRecordValueRequest("20 mail2.example.com", nil, nil),
			v1.NewRecordValueRequest("30 mail3.example.com", nil, nil),
		}

		opt := v1.NewCreateDnsRecordRequest(
			hostedZoneID,
			"test-mx",
			3600,
			v1.DnsRecordTypeMX,
			v1.RoutingPolicySimple,
			mxValues,
		)

		dnsRecord, sdkerr := vngcloud.DNS.CreateDnsRecord(context.Background(), opt)
		if sdkerr != nil {
			t.Fatalf("Expect nil but got %+v", sdkerr)
		}
		t.Logf("Created MX Record: %+v", dnsRecord)
	})

	// Test 5: SRV record with multiple services
	t.Run("CreateSRVRecord", func(t *testing.T) {
		weight1, weight2 := 5, 10
		srvValues := []v1.RecordValueRequest{
			v1.NewRecordValueRequest("10 5 443 target1.example.com", nil, &weight1),
			v1.NewRecordValueRequest("20 10 443 target2.example.com", nil, &weight2),
		}

		opt := v1.NewCreateDnsRecordRequest(
			hostedZoneID,
			"test-srv",
			300,
			v1.DnsRecordTypeSRV,
			v1.RoutingPolicyWeighted,
			srvValues,
		)

		dnsRecord, sdkerr := vngcloud.DNS.CreateDnsRecord(context.Background(), opt)
		if sdkerr != nil {
			t.Fatalf("Expect nil but got %+v", sdkerr)
		}
		t.Logf("Created SRV Record: %+v", dnsRecord)
	})

	t.Log("All DNS Record types created successfully")
}

func TestDnsServiceV1_GetRecord(t *testing.T) {
	vngcloud := validSdkConfig()

	// Use one of the existing records for testing
	hostedZoneID := "hosted-zone-32a21aa3-99a3-4d03-9045-37aa701fa03a"

	// First list records to get a record ID
	listOpt := v1.NewListRecordsRequest(hostedZoneID)
	listRecords, sdkerr := vngcloud.DNS.ListRecords(context.Background(), listOpt)
	if sdkerr != nil {
		t.Fatalf("Failed to list records: %+v", sdkerr)
	}

	if len(listRecords.ListData) == 0 {
		t.Skip("No records found to test GetRecord")
	}

	recordID := listRecords.ListData[0].RecordID
	t.Logf("Testing GetRecord with recordId: %s", recordID)

	// Test GetRecord
	opt := v1.NewGetRecordRequest(hostedZoneID, recordID)
	record, sdkerr := vngcloud.DNS.GetRecord(context.Background(), opt)
	if sdkerr != nil {
		t.Fatalf("Expect nil but got %+v", sdkerr)
	}

	if record == nil {
		t.Fatalf("Expect not nil but got nil")
	}

	if record.RecordID != recordID {
		t.Fatalf("Expected record ID %s but got %s", recordID, record.RecordID)
	}

	t.Logf("Retrieved Record: %+v", record)
	t.Log("PASS")
}

func TestDnsServiceV1_UpdateRecord(t *testing.T) {
	vngcloud := validSdkConfig()

	hostedZoneID := "hosted-zone-32a21aa3-99a3-4d03-9045-37aa701fa03a"

	// First list records to get a record ID
	listOpt := v1.NewListRecordsRequest(hostedZoneID)
	listRecords, sdkerr := vngcloud.DNS.ListRecords(context.Background(), listOpt)
	if sdkerr != nil {
		t.Fatalf("Failed to list records: %+v", sdkerr)
	}

	if len(listRecords.ListData) == 0 {
		t.Skip("No records found to test UpdateRecord")
	}

	// Find a record that's not NS or SOA (system records that can't be modified)
	var targetRecord *v1.DnsRecord
	for _, record := range listRecords.ListData {
		if record.Type != "NS" && record.Type != "SOA" {
			targetRecord = record
			break
		}
	}

	if targetRecord == nil {
		t.Skip("No modifiable records found to test UpdateRecord")
	}

	recordID := targetRecord.RecordID
	t.Logf("Testing UpdateRecord with recordId: %s", recordID)

	// Update the record with new TTL
	newTTL := 3600
	values := []v1.RecordValueRequest{
		v1.NewRecordValueRequest("updated.example.com", nil, nil),
	}

	opt := v1.NewUpdateRecordRequest(hostedZoneID, recordID)
	opt.SubDomain = "updated-test"
	opt.TTL = newTTL
	opt.Type = v1.DnsRecordTypeCNAME
	opt.RoutingPolicy = v1.RoutingPolicySimple
	opt.Value = values

	sdkerr = vngcloud.DNS.UpdateRecord(context.Background(), opt)
	if sdkerr != nil {
		t.Fatalf("Expect nil but got %+v", sdkerr)
	}

	t.Log("Record updated successfully")
	t.Log("PASS")
}

func TestDnsServiceV1_DeleteRecord(t *testing.T) {
	vngcloud := validSdkConfig()

	hostedZoneID := "hosted-zone-32a21aa3-99a3-4d03-9045-37aa701fa03a"
	recordID := "record-63ecb31a-2477-4afa-a46b-12ca3f2b1060"

	deleteOpt := v1.NewDeleteRecordRequest(hostedZoneID, recordID)
	sdkerr := vngcloud.DNS.DeleteRecord(context.Background(), deleteOpt)
	if sdkerr != nil {
		t.Fatalf("Expect nil but got %+v", sdkerr)
	}

	t.Log("Record deleted successfully")
	t.Log("PASS")
}
