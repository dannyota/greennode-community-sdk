package v1

import (
	"github.com/dannyota/greennode-community-sdk/v2/greennode/client"
	"github.com/dannyota/greennode-community-sdk/v2/greennode/entity"
	sdkerror "github.com/dannyota/greennode-community-sdk/v2/greennode/sdkerror"
)

func (s *VDnsServiceInternal) GetHostedZoneById(popts IGetHostedZoneByIdRequest, portalUserId string) (*entity.HostedZone, sdkerror.IError) {
	url := getHostedZoneByIdUrl(s.DnsClient, popts)
	resp := new(GetHostedZoneByIdResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.NetworkGatewayErrorType)
	req := client.NewRequest().
		WithHeader("User-Agent", popts.ParseUserAgent()).
		WithUserId(portalUserId).
		WithOkCodes(200).
		WithJsonResponse(resp).
		WithJsonError(errResp)

	if _, sdkErr := s.DnsClient.Get(url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp).
			WithKVparameters(
				"hostedZoneId", popts.GetHostedZoneId()).
			WithErrorCategories(sdkerror.ErrCatProductVdns)
	}

	return resp.ToEntityHostedZone(), nil
}

func (s *VDnsServiceInternal) ListHostedZones(popts IListHostedZonesRequest, portalUserId string) (*entity.ListHostedZone, sdkerror.IError) {
	url := listHostedZonesUrl(s.DnsClient, popts)
	resp := new(ListHostedZonesResponse)
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

	return resp.ToEntityListHostedZones(), nil
}

func (s *VDnsServiceInternal) CreateHostedZone(popts ICreateHostedZoneRequest, portalUserId string) (*entity.HostedZone, sdkerror.IError) {
	url := createHostedZoneUrl(s.DnsClient)
	resp := new(CreateHostedZoneResponse)
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

	return resp.ToEntityHostedZone(), nil
}

func (s *VDnsServiceInternal) DeleteHostedZone(popts IDeleteHostedZoneRequest, portalUserId string) sdkerror.IError {
	url := deleteHostedZoneUrl(s.DnsClient, popts)
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

func (s *VDnsServiceInternal) UpdateHostedZone(popts IUpdateHostedZoneRequest, portalUserId string) sdkerror.IError {
	url := updateHostedZoneUrl(s.DnsClient, popts)
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
