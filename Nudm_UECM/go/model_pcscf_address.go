/*
 * Nudm_UECM
 *
 * Nudm Context Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Nudm_UECM

// PcscfAddress - Contains the addressing information (IP addresses and/or FQDN) of the P-CSCF
type PcscfAddress struct {

	Ipv4Addrs []string `json:"ipv4Addrs,omitempty"`

	Ipv6Addrs []Ipv6Addr `json:"ipv6Addrs,omitempty"`

	// Fully Qualified Domain Name
	Fqdn string `json:"fqdn,omitempty"`
}
