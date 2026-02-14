package v1

import "github.com/dannyota/greennode-community-sdk/v2/greennode/client"

func getHostedZoneByIdUrl(sc client.ServiceClient, opts IGetHostedZoneByIdRequest) string {
	return sc.ServiceURL(
		"dns",
		"hosted-zone",
		opts.GetHostedZoneId())
}

func listHostedZonesUrl(sc client.ServiceClient, opts IListHostedZonesRequest) string {
	url := sc.ServiceURL("dns", "hosted-zone")
	if opts.GetName() != "" {
		url += "?name=" + opts.GetName()
	}
	return url
}

func listRecordsUrl(sc client.ServiceClient, opts IListRecordsRequest) string {
	url := sc.ServiceURL("dns", "hosted-zone", opts.GetHostedZoneId(), "record")
	if opts.GetName() != "" {
		url += "?name=" + opts.GetName()
	}
	return url
}

func createHostedZoneUrl(sc client.ServiceClient) string {
	return sc.ServiceURL("dns", "hosted-zone")
}

func deleteHostedZoneUrl(sc client.ServiceClient, opts IDeleteHostedZoneRequest) string {
	return sc.ServiceURL("dns", "hosted-zone", opts.GetHostedZoneId())
}

func updateHostedZoneUrl(sc client.ServiceClient, opts IUpdateHostedZoneRequest) string {
	return sc.ServiceURL("dns", "hosted-zone", opts.GetHostedZoneId())
}

func getRecordUrl(sc client.ServiceClient, opts IGetRecordRequest) string {
	return sc.ServiceURL("dns", "hosted-zone", opts.GetHostedZoneId(), "record", opts.GetRecordId())
}

func updateRecordUrl(sc client.ServiceClient, opts IUpdateRecordRequest) string {
	return sc.ServiceURL("dns", "hosted-zone", opts.GetHostedZoneId(), "record", opts.GetRecordId())
}

func deleteRecordUrl(sc client.ServiceClient, opts IDeleteRecordRequest) string {
	return sc.ServiceURL("dns", "hosted-zone", opts.GetHostedZoneId(), "record", opts.GetRecordId())
}

func createDnsRecordUrl(sc client.ServiceClient, opts ICreateDnsRecordRequest) string {
	return sc.ServiceURL("dns", "hosted-zone", opts.GetHostedZoneId(), "record")
}
