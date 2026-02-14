package v2

import (
	"github.com/dannyota/greennode-community-sdk/v2/greennode/client"
	"github.com/dannyota/greennode-community-sdk/v2/greennode/entity"
	sdkerror "github.com/dannyota/greennode-community-sdk/v2/greennode/sdkerror"
)

func (s *NetworkServiceV2) GetAllAddressPairByVirtualSubnetId(opts IGetAllAddressPairByVirtualSubnetIdRequest) ([]*entity.AddressPair, sdkerror.Error) {
	url := getAllAddressPairByVirtualSubnetIdUrl(s.VserverClient, opts)
	resp := new(GetAllAddressPairByVirtualSubnetIdResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithHeader("User-Agent", opts.ParseUserAgent()).
		WithOkCodes(200).
		WithJsonResponse(resp).
		WithJsonError(errResp)

	if _, sdkErr := s.VserverClient.Get(url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp)
		// sdkerror.WithErrorSubnetNotBelongNetwork(sdkErr),
		// sdkerror.WithErrorSubnetNotFound(errResp)).
		// WithKVparameters(
		// 	"virtualSubnetId", popts.GetVirtualSubnetId(),
		// 	"projectId", s.getProjectId(),

	}
	return resp.ToListAddressPair(), nil
}

func (s *NetworkServiceV2) SetAddressPairInVirtualSubnet(opts ISetAddressPairInVirtualSubnetRequest) (*entity.AddressPair, sdkerror.Error) {
	url := setAddressPairInVirtualSubnetUrl(s.VserverClient, opts)
	resp := new(SetAddressPairInVirtualSubnetResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithHeader("User-Agent", opts.ParseUserAgent()).
		WithOkCodes(200, 201, 202, 203, 204).
		WithJsonBody(opts.ToRequestBody()).
		WithJsonResponse(resp).
		WithJsonError(errResp)

	if _, sdkErr := s.VserverClient.Post(url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp)
	}
	return resp.ToAddressPair(), nil
}

func (s *NetworkServiceV2) DeleteAddressPair(opts IDeleteAddressPairRequest) sdkerror.Error {
	url := deleteAddressPairUrl(s.VserverClient, opts)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithHeader("User-Agent", opts.ParseUserAgent()).
		WithOkCodes(200, 201, 202, 203, 204).
		WithJsonError(errResp)

	if _, sdkErr := s.VserverClient.Delete(url, req); sdkErr != nil {
		return sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.WithErrorAddressPairNotFound(errResp)).
			WithKVparameters("addressPairId", opts.GetAddressPairID())
	}
	return nil
}

func (s *NetworkServiceV2) CreateAddressPair(opts ICreateAddressPairRequest) (*entity.AddressPair, sdkerror.Error) {
	url := createAddressPairUrl(s.VserverClient, opts)
	resp := new(CreateAddressPairResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithHeader("User-Agent", opts.ParseUserAgent()).
		WithOkCodes(201).
		WithJsonBody(opts.ToRequestBody()).
		WithJsonResponse(resp).
		WithJsonError(errResp)

	if _, sdkErr := s.VserverClient.Post(url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.WithErrorInternalNetworkInterfaceNotFound(errResp),
			sdkerror.WithErrorAddressPairExisted(errResp)).
			WithErrorCategories(sdkerror.ErrCatVServer, sdkerror.ErrCatVirtualAddress)
	}

	return resp.ToAddressPair(), nil
}
