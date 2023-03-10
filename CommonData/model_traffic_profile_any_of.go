/*
 * Common Data Types
 *
 * Common Data Types for Service Based Interfaces.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.   
 *
 * API version: 1.4.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_CommonData

type TrafficProfileAnyOf string

// List of TrafficProfileAnyOf
const (
	SINGLE_TRANS_UL TrafficProfileAnyOf = "SINGLE_TRANS_UL"
	SINGLE_TRANS_DL TrafficProfileAnyOf = "SINGLE_TRANS_DL"
	DUAL_TRANS_UL_FIRST TrafficProfileAnyOf = "DUAL_TRANS_UL_FIRST"
	DUAL_TRANS_DL_FIRST TrafficProfileAnyOf = "DUAL_TRANS_DL_FIRST"
	MULTI_TRANS TrafficProfileAnyOf = "MULTI_TRANS"
)
