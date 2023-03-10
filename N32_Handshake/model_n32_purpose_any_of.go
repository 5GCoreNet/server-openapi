/*
 * N32 Handshake API
 *
 * N32-c Handshake Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.3.0-alpha.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_N32_Handshake

type N32PurposeAnyOf string

// List of N32PurposeAnyOf
const (
	ROAMING N32PurposeAnyOf = "ROAMING"
	INTER_PLMN_MOBILITY N32PurposeAnyOf = "INTER_PLMN_MOBILITY"
	SMS_INTERCONNECT N32PurposeAnyOf = "SMS_INTERCONNECT"
	ROAMING_TEST N32PurposeAnyOf = "ROAMING_TEST"
	INTER_PLMN_MOBILITY_TEST N32PurposeAnyOf = "INTER_PLMN_MOBILITY_TEST"
	SMS_INTERCONNECT_TEST N32PurposeAnyOf = "SMS_INTERCONNECT_TEST"
	SNPN_INTERCONNECT N32PurposeAnyOf = "SNPN_INTERCONNECT"
	SNPN_INTERCONNECT_TEST N32PurposeAnyOf = "SNPN_INTERCONNECT_TEST"
	DISASTER_ROAMING N32PurposeAnyOf = "DISASTER_ROAMING"
	DISASTER_ROAMING_TEST N32PurposeAnyOf = "DISASTER_ROAMING_TEST"
)
