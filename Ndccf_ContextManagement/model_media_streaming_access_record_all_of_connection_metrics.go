/*
 * Ndccf_ContextManagement
 *
 * DCCF Context Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Ndccf_ContextManagement

type MediaStreamingAccessRecordAllOfConnectionMetrics struct {

	// string with format 'float' as defined in OpenAPI.
	MeanNetworkRoundTripTime float32 `json:"meanNetworkRoundTripTime"`

	// string with format 'float' as defined in OpenAPI.
	NetworkRoundTripTimeVariation float32 `json:"networkRoundTripTimeVariation"`

	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible.
	CongestionWindowSize int32 `json:"congestionWindowSize"`
}
