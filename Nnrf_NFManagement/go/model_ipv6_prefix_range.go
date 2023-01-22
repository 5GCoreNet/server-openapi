/*
 * NRF NFManagement Service
 *
 * NRF NFManagement Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.3.0-alpha.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Nnrf_NFManagement

// Ipv6PrefixRange - Range of IPv6 prefixes
type Ipv6PrefixRange struct {

	Start Ipv6Prefix `json:"start,omitempty"`

	End Ipv6Prefix `json:"end,omitempty"`
}
