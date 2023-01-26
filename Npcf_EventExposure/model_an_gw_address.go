/*
 * Npcf_EventExposure
 *
 * PCF Event Exposure Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Npcf_EventExposure

// AnGwAddress - Describes the address of the access network gateway control node.
type AnGwAddress struct {

	// String identifying a IPv4 address formatted in the 'dotted decimal' notation as defined in RFC 1166. 
	AnGwIpv4Addr string `json:"anGwIpv4Addr,omitempty"`

	AnGwIpv6Addr Ipv6Addr `json:"anGwIpv6Addr,omitempty"`
}