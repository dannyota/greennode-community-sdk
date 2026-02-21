package v1

import (
	"context"

	"danny.vn/greennode/client"
	sdkerror "danny.vn/greennode/sdkerror"
	"danny.vn/greennode/services/common"
)

func (s *VDnsServiceInternal) GetHostedZoneByID(ctx context.Context, opts *GetHostedZoneByIDRequest, portalUserID string) (*HostedZone, error) {
	url := getHostedZoneByIDURL(s.Client, opts)
	resp := new(GetHostedZoneByIDResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.NetworkGatewayErrorType)
	req := client.NewRequest().
		WithUserID(portalUserID).
		WithOKCodes(200).
		WithJSONResponse(resp).
		WithJSONError(errResp)

	if _, sdkErr := s.Client.Get(ctx, url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp).
			WithKVparameters(
				"hostedZoneId", opts.HostedZoneID).
			AppendCategories(sdkerror.ErrCatProductVdns)
	}

	return resp.ToEntityHostedZone(), nil
}

func (s *VDnsServiceInternal) ListHostedZones(ctx context.Context, opts *ListHostedZonesRequest, portalUserID string) (*ListHostedZones, error) {
	url := listHostedZonesURL(s.Client, opts)
	resp := new(ListHostedZonesResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.NetworkGatewayErrorType)
	req := client.NewRequest().
		WithUserID(portalUserID).
		WithOKCodes(200).
		WithJSONResponse(resp).
		WithJSONError(errResp)

	if _, sdkErr := s.Client.Get(ctx, url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp).
			WithParameters(common.StructToMap(opts)).
			AppendCategories(sdkerror.ErrCatProductVdns)
	}

	return resp.ToEntityListHostedZones(), nil
}

func (s *VDnsServiceInternal) CreateHostedZone(ctx context.Context, opts *CreateHostedZoneRequest, portalUserID string) (*HostedZone, error) {
	url := createHostedZoneURL(s.Client)
	resp := new(CreateHostedZoneResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.NetworkGatewayErrorType)
	req := client.NewRequest().
		WithUserID(portalUserID).
		WithOKCodes(200).
		WithJSONBody(opts.ToRequestBody(s.Client)).
		WithJSONResponse(resp).
		WithJSONError(errResp)

	if _, sdkErr := s.Client.Post(ctx, url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp).
			WithParameters(common.StructToMap(opts)).
			AppendCategories(sdkerror.ErrCatProductVdns)
	}

	return resp.ToEntityHostedZone(), nil
}

func (s *VDnsServiceInternal) DeleteHostedZone(ctx context.Context, opts *DeleteHostedZoneRequest, portalUserID string) error {
	url := deleteHostedZoneURL(s.Client, opts)
	errResp := sdkerror.NewErrorResponse(sdkerror.NetworkGatewayErrorType)
	req := client.NewRequest().
		WithUserID(portalUserID).
		WithOKCodes(204).
		WithJSONError(errResp)

	if _, sdkErr := s.Client.Delete(ctx, url, req); sdkErr != nil {
		return sdkerror.SdkErrorHandler(sdkErr, errResp).
			WithParameters(common.StructToMap(opts)).
			AppendCategories(sdkerror.ErrCatProductVdns)
	}

	return nil
}

func (s *VDnsServiceInternal) UpdateHostedZone(ctx context.Context, opts *UpdateHostedZoneRequest, portalUserID string) error {
	url := updateHostedZoneURL(s.Client, opts)
	errResp := sdkerror.NewErrorResponse(sdkerror.NetworkGatewayErrorType)
	req := client.NewRequest().
		WithUserID(portalUserID).
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

func (s *VDnsServiceInternal) ListRecords(ctx context.Context, opts *ListRecordsRequest, portalUserID string) (*ListDnsRecords, error) {
	url := listRecordsURL(s.Client, opts)
	resp := new(ListRecordsResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.NetworkGatewayErrorType)
	req := client.NewRequest().
		WithUserID(portalUserID).
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

func (s *VDnsServiceInternal) GetRecord(ctx context.Context, opts *GetRecordRequest, portalUserID string) (*DnsRecord, error) {
	url := getRecordURL(s.Client, opts)
	resp := new(GetRecordResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.NetworkGatewayErrorType)
	req := client.NewRequest().
		WithUserID(portalUserID).
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

func (s *VDnsServiceInternal) UpdateRecord(ctx context.Context, opts *UpdateRecordRequest, portalUserID string) error {
	url := updateRecordURL(s.Client, opts)
	errResp := sdkerror.NewErrorResponse(sdkerror.NetworkGatewayErrorType)
	req := client.NewRequest().
		WithUserID(portalUserID).
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

func (s *VDnsServiceInternal) DeleteRecord(ctx context.Context, opts *DeleteRecordRequest, portalUserID string) error {
	url := deleteRecordURL(s.Client, opts)
	errResp := sdkerror.NewErrorResponse(sdkerror.NetworkGatewayErrorType)
	req := client.NewRequest().
		WithUserID(portalUserID).
		WithOKCodes(204).
		WithJSONError(errResp)

	if _, sdkErr := s.Client.Delete(ctx, url, req); sdkErr != nil {
		return sdkerror.SdkErrorHandler(sdkErr, errResp).
			WithParameters(common.StructToMap(opts)).
			AppendCategories(sdkerror.ErrCatProductVdns)
	}

	return nil
}

func (s *VDnsServiceInternal) CreateDnsRecord(ctx context.Context, opts *CreateDnsRecordRequest, portalUserID string) (*DnsRecord, error) {
	url := createDnsRecordURL(s.Client, opts)
	resp := new(CreateDnsRecordResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.NetworkGatewayErrorType)
	req := client.NewRequest().
		WithUserID(portalUserID).
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
