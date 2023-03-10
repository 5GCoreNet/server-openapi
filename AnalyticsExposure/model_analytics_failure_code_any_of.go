/*
 * 3gpp-analyticsexposure
 *
 * API for Analytics Exposure.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.1.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_AnalyticsExposure

type AnalyticsFailureCodeAnyOf string

// List of AnalyticsFailureCodeAnyOf
const (
	UNAVAILABLE_DATA AnalyticsFailureCodeAnyOf = "UNAVAILABLE_DATA"
	BOTH_STAT_PRED_NOT_ALLOWED AnalyticsFailureCodeAnyOf = "BOTH_STAT_PRED_NOT_ALLOWED"
	UNSATISFIED_REQUESTED_ANALYTICS_TIME AnalyticsFailureCodeAnyOf = "UNSATISFIED_REQUESTED_ANALYTICS_TIME"
	OTHER AnalyticsFailureCodeAnyOf = "OTHER"
)
