/*
 * 3gpp-monitoring-event
 *
 * API for Monitoring Event.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_MonitoringEvent

// MinorLocationQoS - Contain Minor Location QoS.
type MinorLocationQoS struct {

	// Indicates value of accuracy.
	HAccuracy float32 `json:"hAccuracy,omitempty"`

	// Indicates value of accuracy.
	VAccuracy float32 `json:"vAccuracy,omitempty"`
}