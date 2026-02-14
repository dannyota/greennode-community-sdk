package v1

import "github.com/dannyota/greennode-community-sdk/v2/greennode/client"

func getHostedZoneByIdUrl(psc client.IServiceClient, popts IGetHostedZoneByIdRequest) string {
	return psc.ServiceURL(
		"dns",
		"hosted-zone",
		popts.GetHostedZoneId())
}

func listHostedZonesUrl(psc client.IServiceClient, popts IListHostedZonesRequest) string {
	url := psc.ServiceURL("dns", "hosted-zone")
	if popts.GetName() != "" {
		url += "?name=" + popts.GetName()
	}
	return url
}

func listRecordsUrl(psc client.IServiceClient, popts IListRecordsRequest) string {
	url := psc.ServiceURL("dns", "hosted-zone", popts.GetHostedZoneId(), "record")
	if popts.GetName() != "" {
		url += "?name=" + popts.GetName()
	}
	return url
}

func createHostedZoneUrl(psc client.IServiceClient) string {
	return psc.ServiceURL("dns", "hosted-zone")
}

func deleteHostedZoneUrl(psc client.IServiceClient, popts IDeleteHostedZoneRequest) string {
	return psc.ServiceURL("dns", "hosted-zone", popts.GetHostedZoneId())
}

func updateHostedZoneUrl(psc client.IServiceClient, popts IUpdateHostedZoneRequest) string {
	return psc.ServiceURL("dns", "hosted-zone", popts.GetHostedZoneId())
}

func getRecordUrl(psc client.IServiceClient, popts IGetRecordRequest) string {
	return psc.ServiceURL("dns", "hosted-zone", popts.GetHostedZoneId(), "record", popts.GetRecordId())
}

func updateRecordUrl(psc client.IServiceClient, popts IUpdateRecordRequest) string {
	return psc.ServiceURL("dns", "hosted-zone", popts.GetHostedZoneId(), "record", popts.GetRecordId())
}

func deleteRecordUrl(psc client.IServiceClient, popts IDeleteRecordRequest) string {
	return psc.ServiceURL("dns", "hosted-zone", popts.GetHostedZoneId(), "record", popts.GetRecordId())
}

func createDnsRecordUrl(psc client.IServiceClient, popts ICreateDnsRecordRequest) string {
	return psc.ServiceURL("dns", "hosted-zone", popts.GetHostedZoneId(), "record")
}
