/*
 * 3gpp-analyticsexposure
 *
 * API for Analytics Exposure.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.1.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_AnalyticsExposure

// AnalyticsFailureEventInfo - Represents an event for which the subscription request was not successful and including the associated failure reason. 
type AnalyticsFailureEventInfo struct {

	Event AnalyticsEvent `json:"event"`

	FailureCode AnalyticsFailureCode `json:"failureCode"`
}