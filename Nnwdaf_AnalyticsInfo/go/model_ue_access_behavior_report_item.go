/*
 * Nnwdaf_AnalyticsInfo
 *
 * Nnwdaf_AnalyticsInfo Service API.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Nnwdaf_AnalyticsInfo

// UeAccessBehaviorReportItem - Report Item for UE Access Behavior Trends event.
type UeAccessBehaviorReportItem struct {

	StateTransitionType AccessStateTransitionType `json:"stateTransitionType"`

	// indicating a time in seconds.
	Spacing int32 `json:"spacing"`

	// indicating a time in seconds.
	Duration int32 `json:"duration"`
}
