package v1

import "github.com/dannyota/greennode-community-sdk/v2/greennode/client"

func getHostedZoneByIDURL(sc client.ServiceClient, opts IGetHostedZoneByIDRequest) string {
	return sc.ServiceURL(
		"dns",
		"hosted-zone",
		opts.GetHostedZoneID())
}

func listHostedZonesURL(sc client.ServiceClient, opts IListHostedZonesRequest) string {
	url := sc.ServiceURL("dns", "hosted-zone")
	if opts.GetName() != "" {
		url += "?name=" + opts.GetName()
	}
	return url
}

func listRecordsURL(sc client.ServiceClient, opts IListRecordsRequest) string {
	url := sc.ServiceURL("dns", "hosted-zone", opts.GetHostedZoneID(), "record")
	if opts.GetName() != "" {
		url += "?name=" + opts.GetName()
	}
	return url
}

func createHostedZoneURL(sc client.ServiceClient) string {
	return sc.ServiceURL("dns", "hosted-zone")
}

func deleteHostedZoneURL(sc client.ServiceClient, opts IDeleteHostedZoneRequest) string {
	return sc.ServiceURL("dns", "hosted-zone", opts.GetHostedZoneID())
}

func updateHostedZoneURL(sc client.ServiceClient, opts IUpdateHostedZoneRequest) string {
	return sc.ServiceURL("dns", "hosted-zone", opts.GetHostedZoneID())
}

func getRecordURL(sc client.ServiceClient, opts IGetRecordRequest) string {
	return sc.ServiceURL("dns", "hosted-zone", opts.GetHostedZoneID(), "record", opts.GetRecordID())
}

func updateRecordURL(sc client.ServiceClient, opts IUpdateRecordRequest) string {
	return sc.ServiceURL("dns", "hosted-zone", opts.GetHostedZoneID(), "record", opts.GetRecordID())
}

func deleteRecordURL(sc client.ServiceClient, opts IDeleteRecordRequest) string {
	return sc.ServiceURL("dns", "hosted-zone", opts.GetHostedZoneID(), "record", opts.GetRecordID())
}

func createDnsRecordURL(sc client.ServiceClient, opts ICreateDnsRecordRequest) string {
	return sc.ServiceURL("dns", "hosted-zone", opts.GetHostedZoneID(), "record")
}
