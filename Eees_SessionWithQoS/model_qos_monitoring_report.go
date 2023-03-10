/*
 * EES Session with QoS API
 *
 * API for EES Session with Qos service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Eees_SessionWithQoS

// QosMonitoringReport - Represents a QoS monitoring report.
type QosMonitoringReport struct {

	UlDelays []int32 `json:"ulDelays,omitempty"`

	DlDelays []int32 `json:"dlDelays,omitempty"`

	RtDelays []int32 `json:"rtDelays,omitempty"`
}
