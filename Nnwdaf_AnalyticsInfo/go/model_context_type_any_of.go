/*
 * Nnwdaf_AnalyticsInfo
 *
 * Nnwdaf_AnalyticsInfo Service API.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Nnwdaf_AnalyticsInfo

type ContextTypeAnyOf string

// List of ContextTypeAnyOf
const (
	PENDING_ANALYTICS ContextTypeAnyOf = "PENDING_ANALYTICS"
	HISTORICAL_ANALYTICS ContextTypeAnyOf = "HISTORICAL_ANALYTICS"
	AGGR_SUBS ContextTypeAnyOf = "AGGR_SUBS"
	DATA ContextTypeAnyOf = "DATA"
	AGGR_INFO ContextTypeAnyOf = "AGGR_INFO"
	ML_MODELS ContextTypeAnyOf = "ML_MODELS"
)
