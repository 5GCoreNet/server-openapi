/*
 * LMF Location
 *
 * LMF Location Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Nlmf_Location

// PeriodicEventInfo - Indicates the information of periodic event reporting.
type PeriodicEventInfo struct {

	// Number of required periodic event reports.
	ReportingAmount int32 `json:"reportingAmount"`

	// Event reporting periodic interval.
	ReportingInterval int32 `json:"reportingInterval"`
}
