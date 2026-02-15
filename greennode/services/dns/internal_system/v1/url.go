package v1

import "github.com/dannyota/greennode-community-sdk/v2/greennode/client"

func getHostedZoneByIDURL(sc *client.ServiceClient, opts *GetHostedZoneByIDRequest) string {
	return sc.ServiceURL(
		"dns",
		"hosted-zone",
		opts.GetHostedZoneID())
}

func listHostedZonesURL(sc *client.ServiceClient, opts *ListHostedZonesRequest) string {
	url := sc.ServiceURL("dns", "hosted-zone")
	if opts.GetName() != "" {
		url += "?name=" + opts.GetName()
	}
	return url
}

func listRecordsURL(sc *client.ServiceClient, opts *ListRecordsRequest) string {
	url := sc.ServiceURL("dns", "hosted-zone", opts.GetHostedZoneID(), "record")
	if opts.GetName() != "" {
		url += "?name=" + opts.GetName()
	}
	return url
}

func createHostedZoneURL(sc *client.ServiceClient) string {
	return sc.ServiceURL("dns", "hosted-zone")
}

func deleteHostedZoneURL(sc *client.ServiceClient, opts *DeleteHostedZoneRequest) string {
	return sc.ServiceURL("dns", "hosted-zone", opts.GetHostedZoneID())
}

func updateHostedZoneURL(sc *client.ServiceClient, opts *UpdateHostedZoneRequest) string {
	return sc.ServiceURL("dns", "hosted-zone", opts.GetHostedZoneID())
}

func getRecordURL(sc *client.ServiceClient, opts *GetRecordRequest) string {
	return sc.ServiceURL("dns", "hosted-zone", opts.GetHostedZoneID(), "record", opts.GetRecordID())
}

func updateRecordURL(sc *client.ServiceClient, opts *UpdateRecordRequest) string {
	return sc.ServiceURL("dns", "hosted-zone", opts.GetHostedZoneID(), "record", opts.GetRecordID())
}

func deleteRecordURL(sc *client.ServiceClient, opts *DeleteRecordRequest) string {
	return sc.ServiceURL("dns", "hosted-zone", opts.GetHostedZoneID(), "record", opts.GetRecordID())
}

func createDnsRecordURL(sc *client.ServiceClient, opts *CreateDnsRecordRequest) string {
	return sc.ServiceURL("dns", "hosted-zone", opts.GetHostedZoneID(), "record")
}
