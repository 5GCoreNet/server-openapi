/*
 * Nsmf_PDUSession
 *
 * SMF PDU Session Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.3.0-alpha.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Nsmf_PDUSession

// StatusNotification - Data within Notify Status Request
type StatusNotification struct {

	StatusInfo StatusInfo `json:"statusInfo"`

	SmallDataRateStatus SmallDataRateStatus `json:"smallDataRateStatus,omitempty"`

	ApnRateStatus ApnRateStatus `json:"apnRateStatus,omitempty"`

	TargetDnaiInfo TargetDnaiInfo `json:"targetDnaiInfo,omitempty"`

	// String providing an URI formatted according to RFC 3986.
	OldPduSessionRef string `json:"oldPduSessionRef,omitempty"`

	// String uniquely identifying a NF instance. The format of the NF Instance ID shall be a  Universally Unique Identifier (UUID) version 4, as described in IETF RFC 4122.  
	NewSmfId string `json:"newSmfId,omitempty"`

	EpsPdnCnxInfo EpsPdnCnxInfo `json:"epsPdnCnxInfo,omitempty"`

	// String providing an URI formatted according to RFC 3986.
	InterPlmnApiRoot string `json:"interPlmnApiRoot,omitempty"`

	// String providing an URI formatted according to RFC 3986.
	IntraPlmnApiRoot string `json:"intraPlmnApiRoot,omitempty"`
}
