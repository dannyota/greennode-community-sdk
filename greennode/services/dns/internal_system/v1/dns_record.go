package v1

import (
	"github.com/dannyota/greennode-community-sdk/v2/greennode/client"
	"github.com/dannyota/greennode-community-sdk/v2/greennode/entity"
	sdkerror "github.com/dannyota/greennode-community-sdk/v2/greennode/sdkerror"
)

func (s *VDnsServiceInternal) ListRecords(popts IListRecordsRequest, portalUserId string) (*entity.ListDnsRecords, sdkerror.IError) {
	url := listRecordsUrl(s.DnsClient, popts)
	resp := new(ListRecordsResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.NetworkGatewayErrorType)
	req := client.NewRequest().
		WithHeader("User-Agent", popts.ParseUserAgent()).
		WithUserId(portalUserId).
		WithOkCodes(200).
		WithJsonResponse(resp).
		WithJsonError(errResp)

	if _, sdkErr := s.DnsClient.Get(url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp).
			WithParameters(popts.ToMap()).
			WithErrorCategories(sdkerror.ErrCatProductVdns)
	}

	return resp.ToEntityListRecords(), nil
}

func (s *VDnsServiceInternal) GetRecord(popts IGetRecordRequest, portalUserId string) (*entity.DnsRecord, sdkerror.IError) {
	url := getRecordUrl(s.DnsClient, popts)
	resp := new(GetRecordResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.NetworkGatewayErrorType)
	req := client.NewRequest().
		WithHeader("User-Agent", popts.ParseUserAgent()).
		WithUserId(portalUserId).
		WithOkCodes(200).
		WithJsonResponse(resp).
		WithJsonError(errResp)

	if _, sdkErr := s.DnsClient.Get(url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp).
			WithParameters(popts.ToMap()).
			WithErrorCategories(sdkerror.ErrCatProductVdns)
	}

	return resp.ToEntityDnsRecord(), nil
}

func (s *VDnsServiceInternal) UpdateRecord(popts IUpdateRecordRequest, portalUserId string) sdkerror.IError {
	url := updateRecordUrl(s.DnsClient, popts)
	errResp := sdkerror.NewErrorResponse(sdkerror.NetworkGatewayErrorType)
	req := client.NewRequest().
		WithHeader("User-Agent", popts.ParseUserAgent()).
		WithUserId(portalUserId).
		WithOkCodes(204).
		WithJsonBody(popts.ToRequestBody(s.DnsClient)).
		WithJsonError(errResp)

	if _, sdkErr := s.DnsClient.Put(url, req); sdkErr != nil {
		return sdkerror.SdkErrorHandler(sdkErr, errResp).
			WithParameters(popts.ToMap()).
			WithErrorCategories(sdkerror.ErrCatProductVdns)
	}

	return nil
}

func (s *VDnsServiceInternal) DeleteRecord(popts IDeleteRecordRequest, portalUserId string) sdkerror.IError {
	url := deleteRecordUrl(s.DnsClient, popts)
	errResp := sdkerror.NewErrorResponse(sdkerror.NetworkGatewayErrorType)
	req := client.NewRequest().
		WithHeader("User-Agent", popts.ParseUserAgent()).
		WithUserId(portalUserId).
		WithOkCodes(204).
		WithJsonError(errResp)

	if _, sdkErr := s.DnsClient.Delete(url, req); sdkErr != nil {
		return sdkerror.SdkErrorHandler(sdkErr, errResp).
			WithParameters(popts.ToMap()).
			WithErrorCategories(sdkerror.ErrCatProductVdns)
	}

	return nil
}

func (s *VDnsServiceInternal) CreateDnsRecord(popts ICreateDnsRecordRequest, portalUserId string) (*entity.DnsRecord, sdkerror.IError) {
	url := createDnsRecordUrl(s.DnsClient, popts)
	resp := new(CreateDnsRecordResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.NetworkGatewayErrorType)
	req := client.NewRequest().
		WithHeader("User-Agent", popts.ParseUserAgent()).
		WithUserId(portalUserId).
		WithOkCodes(200).
		WithJsonBody(popts.ToRequestBody(s.DnsClient)).
		WithJsonResponse(resp).
		WithJsonError(errResp)

	if _, sdkErr := s.DnsClient.Post(url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp).
			WithParameters(popts.ToMap()).
			WithErrorCategories(sdkerror.ErrCatProductVdns)
	}

	return resp.ToEntityDnsRecord(), nil
}
