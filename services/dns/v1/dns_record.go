package v1

import (
	"context"

	"danny.vn/greennode/client"
	sdkerror "danny.vn/greennode/sdkerror"
	"danny.vn/greennode/services/common"
)

func (s *VDnsServiceV1) ListRecords(ctx context.Context, opts *ListRecordsRequest) (*ListDnsRecords, error) {
	url := listRecordsURL(s.Client, opts)
	resp := new(ListRecordsResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.NetworkGatewayErrorType)
	req := client.NewRequest().
		WithOKCodes(200).
		WithJSONResponse(resp).
		WithJSONError(errResp)

	if _, sdkErr := s.Client.Get(ctx, url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp).
			WithParameters(common.StructToMap(opts)).
			AppendCategories(sdkerror.ErrCatProductVdns)
	}

	return resp.ToEntityListRecords(), nil
}

func (s *VDnsServiceV1) GetRecord(ctx context.Context, opts *GetRecordRequest) (*DnsRecord, error) {
	url := getRecordURL(s.Client, opts)
	resp := new(GetRecordResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.NetworkGatewayErrorType)
	req := client.NewRequest().
		WithOKCodes(200).
		WithJSONResponse(resp).
		WithJSONError(errResp)

	if _, sdkErr := s.Client.Get(ctx, url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp).
			WithParameters(common.StructToMap(opts)).
			AppendCategories(sdkerror.ErrCatProductVdns)
	}

	return resp.ToEntityDnsRecord(), nil
}

func (s *VDnsServiceV1) UpdateRecord(ctx context.Context, opts *UpdateRecordRequest) error {
	url := updateRecordURL(s.Client, opts)
	errResp := sdkerror.NewErrorResponse(sdkerror.NetworkGatewayErrorType)
	req := client.NewRequest().
		WithOKCodes(204).
		WithJSONBody(opts.ToRequestBody(s.Client)).
		WithJSONError(errResp)

	if _, sdkErr := s.Client.Put(ctx, url, req); sdkErr != nil {
		return sdkerror.SdkErrorHandler(sdkErr, errResp).
			WithParameters(common.StructToMap(opts)).
			AppendCategories(sdkerror.ErrCatProductVdns)
	}

	return nil
}

func (s *VDnsServiceV1) DeleteRecord(ctx context.Context, opts *DeleteRecordRequest) error {
	url := deleteRecordURL(s.Client, opts)
	errResp := sdkerror.NewErrorResponse(sdkerror.NetworkGatewayErrorType)
	req := client.NewRequest().
		WithOKCodes(204).
		WithJSONError(errResp)

	if _, sdkErr := s.Client.Delete(ctx, url, req); sdkErr != nil {
		return sdkerror.SdkErrorHandler(sdkErr, errResp).
			WithParameters(common.StructToMap(opts)).
			AppendCategories(sdkerror.ErrCatProductVdns)
	}

	return nil
}

func (s *VDnsServiceV1) CreateDnsRecord(ctx context.Context, opts *CreateDnsRecordRequest) (*DnsRecord, error) {
	url := createDnsRecordURL(s.Client, opts)
	resp := new(CreateDnsRecordResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.NetworkGatewayErrorType)
	req := client.NewRequest().
		WithOKCodes(200).
		WithJSONBody(opts.ToRequestBody(s.Client)).
		WithJSONResponse(resp).
		WithJSONError(errResp)

	if _, sdkErr := s.Client.Post(ctx, url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp).
			WithParameters(common.StructToMap(opts)).
			AppendCategories(sdkerror.ErrCatProductVdns)
	}

	return resp.ToEntityDnsRecord(), nil
}
