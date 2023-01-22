/*
 * Nsmf_PDUSession
 *
 * SMF PDU Session Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.3.0-alpha.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Nsmf_PDUSession

// QosFlowReleaseRequestItem - Individual QoS flow requested to be released
type QosFlowReleaseRequestItem struct {

	// Unsigned integer identifying a QoS flow, within the range 0 to 63.
	Qfi int32 `json:"qfi"`

	// string with format 'bytes' as defined in OpenAPI
	QosRules string `json:"qosRules,omitempty"`

	// string with format 'bytes' as defined in OpenAPI
	QosFlowDescription string `json:"qosFlowDescription,omitempty"`
}