/*
 * Nadrf_DataManagement
 *
 * ADRF Data Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Nadrf_DataManagement

type ReachabilityForDataConfiguration struct {

	ReportCfg ReachabilityForDataReportConfig `json:"reportCfg"`

	// indicating a time in seconds.
	MinInterval int32 `json:"minInterval,omitempty"`
}
