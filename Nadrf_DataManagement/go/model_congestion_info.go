/*
 * Nadrf_DataManagement
 *
 * ADRF Data Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Nadrf_DataManagement

// CongestionInfo - Represents the congestion information.
type CongestionInfo struct {

	CongType CongestionType `json:"congType"`

	TimeIntev TimeWindow `json:"timeIntev"`

	Nsi ThresholdLevel `json:"nsi"`

	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible.
	Confidence int32 `json:"confidence,omitempty"`

	TopAppListUl []TopApplication `json:"topAppListUl,omitempty"`

	TopAppListDl []TopApplication `json:"topAppListDl,omitempty"`
}
