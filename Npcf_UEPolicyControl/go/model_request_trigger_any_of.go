/*
 * Npcf_UEPolicyControl
 *
 * UE Policy Control Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Npcf_UEPolicyControl

type RequestTriggerAnyOf string

// List of RequestTriggerAnyOf
const (
	LOC_CH RequestTriggerAnyOf = "LOC_CH"
	PRA_CH RequestTriggerAnyOf = "PRA_CH"
	UE_POLICY RequestTriggerAnyOf = "UE_POLICY"
	PLMN_CH RequestTriggerAnyOf = "PLMN_CH"
	CON_STATE_CH RequestTriggerAnyOf = "CON_STATE_CH"
	GROUP_ID_LIST_CHG RequestTriggerAnyOf = "GROUP_ID_LIST_CHG"
	UE_CAP_CH RequestTriggerAnyOf = "UE_CAP_CH"
)
