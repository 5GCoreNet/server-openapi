/*
 * Ndcaf_DataReporting
 *
 * Data Collection AF: Data Collection and Reporting Configuration API and Data Reporting API © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 
 *
 * API version: 1.1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Ndcaf_DataReporting

// ReportingCondition - A condition that triggers data reporting by a data collection client to the Data Collection AF.
type ReportingCondition struct {

	Type ReportingConditionType `json:"type"`

	// indicating a time in seconds.
	Period int32 `json:"period,omitempty"`

	Parameter string `json:"parameter,omitempty"`

	Threshold ReportingConditionThreshold `json:"threshold,omitempty"`

	ReportWhenBelow bool `json:"reportWhenBelow,omitempty"`

	EventTrigger ReportingEventTrigger `json:"eventTrigger,omitempty"`
}
