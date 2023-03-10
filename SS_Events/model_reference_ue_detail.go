/*
 * SS_Events
 *
 * API for SEAL Events management.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.1.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_SS_Events

// ReferenceUeDetail - Reference UE details, where the UEs moving in and out to be monitored.
type ReferenceUeDetail struct {

	ValTgtUe ValTargetUe `json:"valTgtUe"`

	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible.
	ProxRange int32 `json:"proxRange"`

	// string with format 'float' as defined in OpenAPI.
	ProxRangeFrac float32 `json:"proxRangeFrac,omitempty"`
}
