package v1

import (
	"github.com/dannyota/greennode-community-sdk/v2/greennode/client"
	"github.com/dannyota/greennode-community-sdk/v2/greennode/entity"
	sdkerror "github.com/dannyota/greennode-community-sdk/v2/greennode/sdkerror"
)

func (s *VDnsServiceV1) ListRecords(opts IListRecordsRequest) (*entity.ListDnsRecords, sdkerror.Error) {
	url := listRecordsURL(s.DnsClient, opts)
	resp := new(ListRecordsResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.NetworkGatewayErrorType)
	req := client.NewRequest().
		WithHeader("User-Agent", opts.ParseUserAgent()).
		WithOkCodes(200).
		WithJSONResponse(resp).
		WithJSONError(errResp)

	if _, sdkErr := s.DnsClient.Get(url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp).
			WithParameters(opts.ToMap()).
			WithErrorCategories(sdkerror.ErrCatProductVdns)
	}

	return resp.ToEntityListRecords(), nil
}

func (s *VDnsServiceV1) GetRecord(opts IGetRecordRequest) (*entity.DnsRecord, sdkerror.Error) {
	url := getRecordURL(s.DnsClient, opts)
	resp := new(GetRecordResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.NetworkGatewayErrorType)
	req := client.NewRequest().
		WithHeader("User-Agent", opts.ParseUserAgent()).
		WithOkCodes(200).
		WithJSONResponse(resp).
		WithJSONError(errResp)

	if _, sdkErr := s.DnsClient.Get(url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp).
			WithParameters(opts.ToMap()).
			WithErrorCategories(sdkerror.ErrCatProductVdns)
	}

	return resp.ToEntityDnsRecord(), nil
}

func (s *VDnsServiceV1) UpdateRecord(opts IUpdateRecordRequest) sdkerror.Error {
	url := updateRecordURL(s.DnsClient, opts)
	errResp := sdkerror.NewErrorResponse(sdkerror.NetworkGatewayErrorType)
	req := client.NewRequest().
		WithHeader("User-Agent", opts.ParseUserAgent()).
		WithOkCodes(204).
		WithJSONBody(opts.ToRequestBody(s.DnsClient)).
		WithJSONError(errResp)

	if _, sdkErr := s.DnsClient.Put(url, req); sdkErr != nil {
		return sdkerror.SdkErrorHandler(sdkErr, errResp).
			WithParameters(opts.ToMap()).
			WithErrorCategories(sdkerror.ErrCatProductVdns)
	}

	return nil
}

func (s *VDnsServiceV1) DeleteRecord(opts IDeleteRecordRequest) sdkerror.Error {
	url := deleteRecordURL(s.DnsClient, opts)
	errResp := sdkerror.NewErrorResponse(sdkerror.NetworkGatewayErrorType)
	req := client.NewRequest().
		WithHeader("User-Agent", opts.ParseUserAgent()).
		WithOkCodes(204).
		WithJSONError(errResp)

	if _, sdkErr := s.DnsClient.Delete(url, req); sdkErr != nil {
		return sdkerror.SdkErrorHandler(sdkErr, errResp).
			WithParameters(opts.ToMap()).
			WithErrorCategories(sdkerror.ErrCatProductVdns)
	}

	return nil
}

func (s *VDnsServiceV1) CreateDnsRecord(opts ICreateDnsRecordRequest) (*entity.DnsRecord, sdkerror.Error) {
	url := createDnsRecordURL(s.DnsClient, opts)
	resp := new(CreateDnsRecordResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.NetworkGatewayErrorType)
	req := client.NewRequest().
		WithHeader("User-Agent", opts.ParseUserAgent()).
		WithOkCodes(200).
		WithJSONBody(opts.ToRequestBody(s.DnsClient)).
		WithJSONResponse(resp).
		WithJSONError(errResp)

	if _, sdkErr := s.DnsClient.Post(url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp).
			WithParameters(opts.ToMap()).
			WithErrorCategories(sdkerror.ErrCatProductVdns)
	}

	return resp.ToEntityDnsRecord(), nil
}
