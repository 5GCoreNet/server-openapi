/*
 * TS 28.532 Streaming data reporting service
 *
 * OAS 3.0.1 specification for the Streaming data reporting service (Streaming MnS) © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.
 *
 * API version: 17.1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_StreamingDataMnS

// PerformanceInfoType - Information specific to performance data reporting
type PerformanceInfoType struct {

	MeasObjDn MeasObjDnType `json:"measObjDn"`

	// an ordered list of performance metric names (see clause 4.4.1 of 3GPP TS 28.622[11]) whose values are to be reported by the Performance Data Stream Units (see Annex C of TS 28.550 [42]) via this stream. Performance metrics include measurement and KPI
	PerformanceMetrics []string `json:"performanceMetrics"`

	JobId string `json:"jobId,omitempty"`
}