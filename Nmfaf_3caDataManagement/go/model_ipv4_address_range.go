/*
 * Nmfaf_3caDataManagement
 *
 * MFAF 3GPP Consumer Adaptor (3CA) Data Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Nmfaf_3caDataManagement

// Ipv4AddressRange - Range of IPv4 addresses
type Ipv4AddressRange struct {

	// String identifying a IPv4 address formatted in the 'dotted decimal' notation as defined in RFC 1166. 
	Start string `json:"start,omitempty"`

	// String identifying a IPv4 address formatted in the 'dotted decimal' notation as defined in RFC 1166. 
	End string `json:"end,omitempty"`
}
