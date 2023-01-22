/*
 * Ndccf_DataManagement
 *
 * DCCF Data Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Ndccf_DataManagement

type TrafficProfileAnyOf string

// List of TrafficProfileAnyOf
const (
	SINGLE_TRANS_UL TrafficProfileAnyOf = "SINGLE_TRANS_UL"
	SINGLE_TRANS_DL TrafficProfileAnyOf = "SINGLE_TRANS_DL"
	DUAL_TRANS_UL_FIRST TrafficProfileAnyOf = "DUAL_TRANS_UL_FIRST"
	DUAL_TRANS_DL_FIRST TrafficProfileAnyOf = "DUAL_TRANS_DL_FIRST"
	MULTI_TRANS TrafficProfileAnyOf = "MULTI_TRANS"
)
