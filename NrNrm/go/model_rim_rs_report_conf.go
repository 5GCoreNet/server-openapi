/*
 * NR NRM
 *
 * OAS 3.0.1 specification of the NR NRM © 2020, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.
 *
 * API version: 17.7.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package NrNrm

type RimRsReportConf struct {

	ReportIndicator string `json:"reportIndicator,omitempty"`

	ReportInterval int32 `json:"reportInterval,omitempty"`

	NrofRIMRSReportInfo int32 `json:"nrofRIMRSReportInfo,omitempty"`

	MaxPropagationDelay int32 `json:"maxPropagationDelay,omitempty"`

	RimRSReportInfoList []RimRsReportInfo `json:"rimRSReportInfoList,omitempty"`
}