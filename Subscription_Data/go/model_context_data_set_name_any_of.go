/*
 * Unified Data Repository Service API file for subscription data
 *
 * Unified Data Repository Service (subscription data).   The API version is defined in 3GPP TS 29.504.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: -
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Subscription_Data

type ContextDataSetNameAnyOf string

// List of ContextDataSetNameAnyOf
const (
	AMF_3_GPP ContextDataSetNameAnyOf = "AMF_3GPP"
	AMF_NON_3_GPP ContextDataSetNameAnyOf = "AMF_NON_3GPP"
	SDM_SUBSCRIPTIONS ContextDataSetNameAnyOf = "SDM_SUBSCRIPTIONS"
	EE_SUBSCRIPTIONS ContextDataSetNameAnyOf = "EE_SUBSCRIPTIONS"
	SMSF_3_GPP ContextDataSetNameAnyOf = "SMSF_3GPP"
	SMSF_NON_3_GPP ContextDataSetNameAnyOf = "SMSF_NON_3GPP"
	SUBS_TO_NOTIFY ContextDataSetNameAnyOf = "SUBS_TO_NOTIFY"
	SMF_REG ContextDataSetNameAnyOf = "SMF_REG"
	IP_SM_GW ContextDataSetNameAnyOf = "IP_SM_GW"
	ROAMING_INFO ContextDataSetNameAnyOf = "ROAMING_INFO"
	PEI_INFO ContextDataSetNameAnyOf = "PEI_INFO"
)
