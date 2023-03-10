/*
 * Slice NRM
 *
 * OAS 3.0.1 specification of the Slice NRM @ 2020, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.
 *
 * API version: 17.7.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_SliceNrm

type SubNetworkSingleAllOf1 struct {

	SubNetwork []SubNetworkSingle `json:"SubNetwork,omitempty"`

	NetworkSlice []NetworkSliceSingle `json:"NetworkSlice,omitempty"`

	NetworkSliceSubnet []NetworkSliceSubnetSingle `json:"NetworkSliceSubnet,omitempty"`

	EPTransport []EpTransportSingle `json:"EP_Transport,omitempty"`

	NetworkSliceSubnetProviderCapabilities []NetworkSliceSubnetProviderCapabilitiesSingle `json:"NetworkSliceSubnetProviderCapabilities,omitempty"`

	FeasibilityCheckAndReservationJob []FeasibilityCheckAndReservationJobSingle `json:"FeasibilityCheckAndReservationJob,omitempty"`
}
