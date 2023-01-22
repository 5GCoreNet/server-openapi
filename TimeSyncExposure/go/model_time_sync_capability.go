/*
 * 3gpp-time-sync-exposure
 *
 * API for time synchronization exposure.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package TimeSyncExposure

// TimeSyncCapability - Contains time synchronization capability.
type TimeSyncCapability struct {

	// Integer where the allowed values correspond to the value range of an unsigned 64-bit integer. 
	UpNodeId int32 `json:"upNodeId"`

	GmCapables []GmCapable `json:"gmCapables,omitempty"`

	AsTimeRes AsTimeResource `json:"asTimeRes,omitempty"`

	// Contains the PTP capabilities supported by each of the UE(s). The key of the map is the gpsi. 
	PtpCapForUes map[string]PtpCapabilitiesPerUe `json:"ptpCapForUes,omitempty"`
}