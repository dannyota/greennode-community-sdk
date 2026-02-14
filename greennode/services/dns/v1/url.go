package v1

import "github.com/dannyota/greennode-community-sdk/v2/greennode/client"

func getHostedZoneByIdUrl(psc client.ServiceClient, popts IGetHostedZoneByIdRequest) string {
	return psc.ServiceURL(
		"dns",
		"hosted-zone",
		popts.GetHostedZoneId())
}

func listHostedZonesUrl(psc client.ServiceClient, popts IListHostedZonesRequest) string {
	url := psc.ServiceURL("dns", "hosted-zone")
	if popts.GetName() != "" {
		url += "?name=" + popts.GetName()
	}
	return url
}

func listRecordsUrl(psc client.ServiceClient, popts IListRecordsRequest) string {
	url := psc.ServiceURL("dns", "hosted-zone", popts.GetHostedZoneId(), "record")
	if popts.GetName() != "" {
		url += "?name=" + popts.GetName()
	}
	return url
}

func createHostedZoneUrl(psc client.ServiceClient) string {
	return psc.ServiceURL("dns", "hosted-zone")
}

func deleteHostedZoneUrl(psc client.ServiceClient, popts IDeleteHostedZoneRequest) string {
	return psc.ServiceURL("dns", "hosted-zone", popts.GetHostedZoneId())
}

func updateHostedZoneUrl(psc client.ServiceClient, popts IUpdateHostedZoneRequest) string {
	return psc.ServiceURL("dns", "hosted-zone", popts.GetHostedZoneId())
}

func getRecordUrl(psc client.ServiceClient, popts IGetRecordRequest) string {
	return psc.ServiceURL("dns", "hosted-zone", popts.GetHostedZoneId(), "record", popts.GetRecordId())
}

func updateRecordUrl(psc client.ServiceClient, popts IUpdateRecordRequest) string {
	return psc.ServiceURL("dns", "hosted-zone", popts.GetHostedZoneId(), "record", popts.GetRecordId())
}

func deleteRecordUrl(psc client.ServiceClient, popts IDeleteRecordRequest) string {
	return psc.ServiceURL("dns", "hosted-zone", popts.GetHostedZoneId(), "record", popts.GetRecordId())
}

func createDnsRecordUrl(psc client.ServiceClient, popts ICreateDnsRecordRequest) string {
	return psc.ServiceURL("dns", "hosted-zone", popts.GetHostedZoneId(), "record")
}
