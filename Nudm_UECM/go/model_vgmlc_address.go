/*
 * Nudm_UECM
 *
 * Nudm Context Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Nudm_UECM

type VgmlcAddress struct {

	// String identifying a IPv4 address formatted in the 'dotted decimal' notation as defined in RFC 1166. 
	VgmlcAddressIpv4 string `json:"vgmlcAddressIpv4,omitempty"`

	VgmlcAddressIpv6 Ipv6Addr `json:"vgmlcAddressIpv6,omitempty"`

	// Fully Qualified Domain Name
	VgmlcFqdn string `json:"vgmlcFqdn,omitempty"`
}
