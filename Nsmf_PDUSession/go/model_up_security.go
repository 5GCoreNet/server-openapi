/*
 * Nsmf_PDUSession
 *
 * SMF PDU Session Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.3.0-alpha.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Nsmf_PDUSession

// UpSecurity - Contains Userplain security information.
type UpSecurity struct {

	UpIntegr UpIntegrity `json:"upIntegr"`

	UpConfid UpConfidentiality `json:"upConfid"`
}