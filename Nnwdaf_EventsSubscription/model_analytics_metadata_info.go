/*
 * Nnwdaf_EventsSubscription
 *
 * Nnwdaf_EventsSubscription Service API.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Nnwdaf_EventsSubscription

// AnalyticsMetadataInfo - Contains analytics metadata information required for analytics aggregation.
type AnalyticsMetadataInfo struct {

	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible.
	NumSamples int32 `json:"numSamples,omitempty"`

	DataWindow TimeWindow `json:"dataWindow,omitempty"`

	DataStatProps []DatasetStatisticalProperty `json:"dataStatProps,omitempty"`

	Strategy OutputStrategy `json:"strategy,omitempty"`

	Accuracy Accuracy `json:"accuracy,omitempty"`
}
