/*
 * SS_NetworkResourceAdaptation
 *
 * SS Network Resource Adaptation Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package SS_NetworkResourceAdaptation

// LocalMbmsInfo - Contains the local MBMS information.
type LocalMbmsInfo struct {

	// String identifying a IPv4 address formatted in the 'dotted decimal' notation as defined in RFC 1166. 
	MbmsEnbIpv4MulAddr string `json:"mbmsEnbIpv4MulAddr,omitempty"`

	MbmsEnbIpv6MulAddr Ipv6Prefix `json:"mbmsEnbIpv6MulAddr,omitempty"`

	// String identifying a IPv4 address formatted in the 'dotted decimal' notation as defined in RFC 1166. 
	MbmsGwIpv4SsmAddr string `json:"mbmsGwIpv4SsmAddr,omitempty"`

	MbmsGwIpv6SsmAddr Ipv6Addr `json:"mbmsGwIpv6SsmAddr,omitempty"`

	Cteid string `json:"cteid,omitempty"`

	// String identifying a IPv4 address formatted in the 'dotted decimal' notation as defined in RFC 1166. 
	BmscIpv4Addr string `json:"bmscIpv4Addr,omitempty"`

	BmscIpv6Addr Ipv6Addr `json:"bmscIpv6Addr,omitempty"`

	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible.
	BmscPort int32 `json:"bmscPort,omitempty"`
}
