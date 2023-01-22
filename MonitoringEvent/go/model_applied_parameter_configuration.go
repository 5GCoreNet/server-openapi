/*
 * 3gpp-monitoring-event
 *
 * API for Monitoring Event.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package MonitoringEvent

// AppliedParameterConfiguration - Represents the parameter configuration applied in the network.
type AppliedParameterConfiguration struct {

	// Each element uniquely identifies a user.
	ExternalIds []string `json:"externalIds,omitempty"`

	// Each element identifies the MS internal PSTN/ISDN number allocated for a UE.
	Msisdns []string `json:"msisdns,omitempty"`

	// Unsigned integer identifying a period of time in units of seconds.
	MaximumLatency int32 `json:"maximumLatency,omitempty"`

	// Unsigned integer identifying a period of time in units of seconds.
	MaximumResponseTime int32 `json:"maximumResponseTime,omitempty"`

	// Unsigned integer identifying a period of time in units of seconds.
	MaximumDetectionTime int32 `json:"maximumDetectionTime,omitempty"`
}
