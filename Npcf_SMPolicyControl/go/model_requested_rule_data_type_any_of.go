/*
 * Npcf_SMPolicyControl API
 *
 * Session Management Policy Control Service   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Npcf_SMPolicyControl

type RequestedRuleDataTypeAnyOf string

// List of RequestedRuleDataTypeAnyOf
const (
	CH_ID RequestedRuleDataTypeAnyOf = "CH_ID"
	MS_TIME_ZONE RequestedRuleDataTypeAnyOf = "MS_TIME_ZONE"
	USER_LOC_INFO RequestedRuleDataTypeAnyOf = "USER_LOC_INFO"
	RES_RELEASE RequestedRuleDataTypeAnyOf = "RES_RELEASE"
	SUCC_RES_ALLO RequestedRuleDataTypeAnyOf = "SUCC_RES_ALLO"
	EPS_FALLBACK RequestedRuleDataTypeAnyOf = "EPS_FALLBACK"
)