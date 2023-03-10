/*
 * Nsmf_PDUSession
 *
 * SMF PDU Session Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.3.0-alpha.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Nsmf_PDUSession

// DnaiInformation - DNAI Information
type DnaiInformation struct {

	// DNAI (Data network access identifier), see clause 5.6.7 of 3GPP TS 23.501.
	Dnai string `json:"dnai"`

	NoDnaiChangeInd bool `json:"noDnaiChangeInd,omitempty"`

	NoLocalPsaChangeInd bool `json:"noLocalPsaChangeInd,omitempty"`
}
