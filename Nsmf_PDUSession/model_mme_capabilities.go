/*
 * Nsmf_PDUSession
 *
 * SMF PDU Session Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.3.0-alpha.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Nsmf_PDUSession

// MmeCapabilities - MME capabilities
type MmeCapabilities struct {

	NonIpSupported bool `json:"nonIpSupported,omitempty"`

	EthernetSupported bool `json:"ethernetSupported,omitempty"`

	UpipSupported bool `json:"upipSupported,omitempty"`
}
