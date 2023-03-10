/*
 * CAPIF_Routing_Info_API
 *
 * API for Routing information.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_CAPIF_Routing_Info_API

// Ipv4AddressRange - Range of IPv4 addresses
type Ipv4AddressRange struct {

	// String identifying a IPv4 address formatted in the 'dotted decimal' notation as defined in RFC 1166. 
	Start string `json:"start,omitempty"`

	// String identifying a IPv4 address formatted in the 'dotted decimal' notation as defined in RFC 1166. 
	End string `json:"end,omitempty"`
}
