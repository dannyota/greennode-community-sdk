package v1

import (
	"github.com/dannyota/greennode-community-sdk/v2/greennode/client"
	"github.com/dannyota/greennode-community-sdk/v2/greennode/entity"
	sdkerror "github.com/dannyota/greennode-community-sdk/v2/greennode/sdkerror"
	"github.com/dannyota/greennode-community-sdk/v2/greennode/services/common"
)

func (s *VDnsServiceInternal) GetHostedZoneByID(opts *GetHostedZoneByIDRequest, portalUserID string) (*entity.HostedZone, error) {
	url := getHostedZoneByIDURL(s.DnsClient, opts)
	resp := new(GetHostedZoneByIDResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.NetworkGatewayErrorType)
	req := client.NewRequest().
		WithHeader("User-Agent", opts.ParseUserAgent()).
		WithUserID(portalUserID).
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

func (s *VDnsServiceInternal) ListHostedZones(opts *ListHostedZonesRequest, portalUserID string) (*entity.ListHostedZone, error) {
	url := listHostedZonesURL(s.DnsClient, opts)
	resp := new(ListHostedZonesResponse)
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

	return resp.ToEntityListHostedZones(), nil
}

func (s *VDnsServiceInternal) CreateHostedZone(opts *CreateHostedZoneRequest, portalUserID string) (*entity.HostedZone, error) {
	url := createHostedZoneURL(s.DnsClient)
	resp := new(CreateHostedZoneResponse)
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

	return resp.ToEntityHostedZone(), nil
}

func (s *VDnsServiceInternal) DeleteHostedZone(opts *DeleteHostedZoneRequest, portalUserID string) error {
	url := deleteHostedZoneURL(s.DnsClient, opts)
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

func (s *VDnsServiceInternal) UpdateHostedZone(opts *UpdateHostedZoneRequest, portalUserID string) error {
	url := updateHostedZoneURL(s.DnsClient, opts)
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
