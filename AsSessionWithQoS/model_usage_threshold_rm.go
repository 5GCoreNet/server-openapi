/*
 * 3gpp-as-session-with-qos
 *
 * API for setting us an AS session with required QoS.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_AsSessionWithQoS

// UsageThresholdRm - Represents the same as the UsageThreshold data type but with the nullable:true property.
type UsageThresholdRm struct {

	// Unsigned integer identifying a period of time in units of seconds with \"nullable=true\" property.
	Duration *int32 `json:"duration,omitempty"`

	// Unsigned integer identifying a volume in units of bytes with \"nullable=true\" property.
	TotalVolume *int64 `json:"totalVolume,omitempty"`

	// Unsigned integer identifying a volume in units of bytes with \"nullable=true\" property.
	DownlinkVolume *int64 `json:"downlinkVolume,omitempty"`

	// Unsigned integer identifying a volume in units of bytes with \"nullable=true\" property.
	UplinkVolume *int64 `json:"uplinkVolume,omitempty"`
}
