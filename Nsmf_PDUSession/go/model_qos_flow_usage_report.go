/*
 * Nsmf_PDUSession
 *
 * SMF PDU Session Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.3.0-alpha.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Nsmf_PDUSession

import (
	"time"
)

// QosFlowUsageReport - Contains QoS flows usage data information.
type QosFlowUsageReport struct {

	// Unsigned integer identifying a QoS flow, within the range 0 to 63.
	Qfi int32 `json:"qfi"`

	// string with format 'date-time' as defined in OpenAPI.
	StartTimeStamp time.Time `json:"startTimeStamp"`

	// string with format 'date-time' as defined in OpenAPI.
	EndTimeStamp time.Time `json:"endTimeStamp"`

	// string with format 'int64' as defined in OpenAPI.
	DownlinkVolume int64 `json:"downlinkVolume"`

	// string with format 'int64' as defined in OpenAPI.
	UplinkVolume int64 `json:"uplinkVolume"`
}
