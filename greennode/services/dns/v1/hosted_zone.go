package v1

import (
	"github.com/dannyota/greennode-community-sdk/v2/greennode/client"
	"github.com/dannyota/greennode-community-sdk/v2/greennode/entity"
	sdkerror "github.com/dannyota/greennode-community-sdk/v2/greennode/sdkerror"
)

func (s *VDnsServiceV1) GetHostedZoneByID(opts *GetHostedZoneByIDRequest) (*entity.HostedZone, sdkerror.Error) {
	url := getHostedZoneByIDURL(s.DnsClient, opts)
	resp := new(GetHostedZoneByIDResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.NetworkGatewayErrorType)
	req := client.NewRequest().
		WithHeader("User-Agent", opts.ParseUserAgent()).
		WithOkCodes(200).
		WithJSONResponse(resp).
		WithJSONError(errResp)

	if _, sdkErr := s.DnsClient.Get(url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp).
			WithKVparameters(
				"hostedZoneId", opts.GetHostedZoneID()).
			WithErrorCategories(sdkerror.ErrCatProductVdns)
	}

	return resp.ToEntityHostedZone(), nil
}

func (s *VDnsServiceV1) ListHostedZones(opts *ListHostedZonesRequest) (*entity.ListHostedZone, sdkerror.Error) {
	url := listHostedZonesURL(s.DnsClient, opts)
	resp := new(ListHostedZonesResponse)
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

	return resp.ToEntityListHostedZones(), nil
}

func (s *VDnsServiceV1) CreateHostedZone(opts *CreateHostedZoneRequest) (*entity.HostedZone, sdkerror.Error) {
	url := createHostedZoneURL(s.DnsClient)
	resp := new(CreateHostedZoneResponse)
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

	return resp.ToEntityHostedZone(), nil
}

func (s *VDnsServiceV1) DeleteHostedZone(opts *DeleteHostedZoneRequest) sdkerror.Error {
	url := deleteHostedZoneURL(s.DnsClient, opts)
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

func (s *VDnsServiceV1) UpdateHostedZone(opts *UpdateHostedZoneRequest) sdkerror.Error {
	url := updateHostedZoneURL(s.DnsClient, opts)
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
