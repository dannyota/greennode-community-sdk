package v1

import (
	"github.com/dannyota/greennode-community-sdk/v2/greennode/client"
	"github.com/dannyota/greennode-community-sdk/v2/greennode/entity"
	sdkerror "github.com/dannyota/greennode-community-sdk/v2/greennode/sdkerror"
	"github.com/dannyota/greennode-community-sdk/v2/greennode/services/common"
)

func (s *VDnsServiceInternal) ListRecords(opts *ListRecordsRequest, portalUserID string) (*entity.ListDnsRecords, error) {
	url := listRecordsURL(s.DnsClient, opts)
	resp := new(ListRecordsResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.NetworkGatewayErrorType)
	req := client.NewRequest().
		WithHeader("User-Agent", opts.ParseUserAgent()).
		WithUserID(portalUserID).
		WithOkCodes(200).
		WithJSONResponse(resp).
		WithJSONError(errResp)

	if _, sdkErr := s.DnsClient.Get(url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp).
			WithParameters(common.StructToMap(opts)).
			WithErrorCategories(sdkerror.ErrCatProductVdns)
	}

	return resp.ToEntityListRecords(), nil
}

func (s *VDnsServiceInternal) GetRecord(opts *GetRecordRequest, portalUserID string) (*entity.DnsRecord, error) {
	url := getRecordURL(s.DnsClient, opts)
	resp := new(GetRecordResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.NetworkGatewayErrorType)
	req := client.NewRequest().
		WithHeader("User-Agent", opts.ParseUserAgent()).
		WithUserID(portalUserID).
		WithOkCodes(200).
		WithJSONResponse(resp).
		WithJSONError(errResp)

	if _, sdkErr := s.DnsClient.Get(url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp).
			WithParameters(common.StructToMap(opts)).
			WithErrorCategories(sdkerror.ErrCatProductVdns)
	}

	return resp.ToEntityDnsRecord(), nil
}

func (s *VDnsServiceInternal) UpdateRecord(opts *UpdateRecordRequest, portalUserID string) error {
	url := updateRecordURL(s.DnsClient, opts)
	errResp := sdkerror.NewErrorResponse(sdkerror.NetworkGatewayErrorType)
	req := client.NewRequest().
		WithHeader("User-Agent", opts.ParseUserAgent()).
		WithUserID(portalUserID).
		WithOkCodes(204).
		WithJSONBody(opts.ToRequestBody(s.DnsClient)).
		WithJSONError(errResp)

	if _, sdkErr := s.DnsClient.Put(url, req); sdkErr != nil {
		return sdkerror.SdkErrorHandler(sdkErr, errResp).
			WithParameters(common.StructToMap(opts)).
			WithErrorCategories(sdkerror.ErrCatProductVdns)
	}

	return nil
}

func (s *VDnsServiceInternal) DeleteRecord(opts *DeleteRecordRequest, portalUserID string) error {
	url := deleteRecordURL(s.DnsClient, opts)
	errResp := sdkerror.NewErrorResponse(sdkerror.NetworkGatewayErrorType)
	req := client.NewRequest().
		WithHeader("User-Agent", opts.ParseUserAgent()).
		WithUserID(portalUserID).
		WithOkCodes(204).
		WithJSONError(errResp)

	if _, sdkErr := s.DnsClient.Delete(url, req); sdkErr != nil {
		return sdkerror.SdkErrorHandler(sdkErr, errResp).
			WithParameters(common.StructToMap(opts)).
			WithErrorCategories(sdkerror.ErrCatProductVdns)
	}

	return nil
}

func (s *VDnsServiceInternal) CreateDnsRecord(opts *CreateDnsRecordRequest, portalUserID string) (*entity.DnsRecord, error) {
	url := createDnsRecordURL(s.DnsClient, opts)
	resp := new(CreateDnsRecordResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.NetworkGatewayErrorType)
	req := client.NewRequest().
		WithHeader("User-Agent", opts.ParseUserAgent()).
		WithUserID(portalUserID).
		WithOkCodes(200).
		WithJSONBody(opts.ToRequestBody(s.DnsClient)).
		WithJSONResponse(resp).
		WithJSONError(errResp)

	if _, sdkErr := s.DnsClient.Post(url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp).
			WithParameters(common.StructToMap(opts)).
			WithErrorCategories(sdkerror.ErrCatProductVdns)
	}

	return resp.ToEntityDnsRecord(), nil
}
