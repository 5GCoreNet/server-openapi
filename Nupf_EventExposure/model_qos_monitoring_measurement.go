/*
 * UPF Event Exposure Service
 *
 * UPF Event Exposure Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Nupf_EventExposure

// QosMonitoringMeasurement - QoS Monitoring Measurement information
type QosMonitoringMeasurement struct {

	// Integer where the allowed values correspond to the value range of an unsigned 32-bit integer. 
	DlPacketDelay int32 `json:"dlPacketDelay,omitempty"`

	// Integer where the allowed values correspond to the value range of an unsigned 32-bit integer. 
	UlPacketDelay int32 `json:"ulPacketDelay,omitempty"`

	// Integer where the allowed values correspond to the value range of an unsigned 32-bit integer. 
	RtrPacketDelay int32 `json:"rtrPacketDelay,omitempty"`

	MeasureFailure bool `json:"measureFailure,omitempty"`
}
