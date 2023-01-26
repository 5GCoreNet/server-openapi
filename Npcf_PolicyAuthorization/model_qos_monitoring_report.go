/*
 * Npcf_PolicyAuthorization Service API
 *
 * PCF Policy Authorization Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Npcf_PolicyAuthorization

// QosMonitoringReport - QoS Monitoring reporting information.
type QosMonitoringReport struct {

	Flows []Flows `json:"flows,omitempty"`

	UlDelays []int32 `json:"ulDelays,omitempty"`

	DlDelays []int32 `json:"dlDelays,omitempty"`

	RtDelays []int32 `json:"rtDelays,omitempty"`
}