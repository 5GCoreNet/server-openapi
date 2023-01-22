/*
 * Nchf_OfflineOnlyCharging
 *
 * OfflineOnlyCharging Service © 2021, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Nchf_OfflineOnlyCharging

import (
	"time"
)

type QosFlowsUsageReport struct {

	// Unsigned integer identifying a QoS flow, within the range 0 to 63.
	QFI int32 `json:"qFI,omitempty"`

	// string with format 'date-time' as defined in OpenAPI.
	StartTimestamp time.Time `json:"startTimestamp,omitempty"`

	// string with format 'date-time' as defined in OpenAPI.
	EndTimestamp time.Time `json:"endTimestamp,omitempty"`

	// Integer where the allowed values correspond to the value range of an unsigned 64-bit integer. 
	UplinkVolume int32 `json:"uplinkVolume,omitempty"`

	// Integer where the allowed values correspond to the value range of an unsigned 64-bit integer. 
	DownlinkVolume int32 `json:"downlinkVolume,omitempty"`
}
