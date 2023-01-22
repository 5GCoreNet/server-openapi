/*
 * Nbsf_Management
 *
 * Binding Support Management Service API.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.3.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Nbsf_Management

// BindingResp - Contains the binding information for a PCF for a PDU Session.
type BindingResp struct {

	// Fully Qualified Domain Name
	PcfSmFqdn string `json:"pcfSmFqdn,omitempty"`

	// IP end points of the PCF hosting the Npcf_SMPolicyControl service.
	PcfSmIpEndPoints []IpEndPoint `json:"pcfSmIpEndPoints,omitempty"`
}
