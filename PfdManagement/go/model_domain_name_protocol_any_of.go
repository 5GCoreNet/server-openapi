/*
 * 3gpp-pfd-management
 *
 * API for PFD management.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package PfdManagement

type DomainNameProtocolAnyOf string

// List of DomainNameProtocolAnyOf
const (
	DNS_QNAME DomainNameProtocolAnyOf = "DNS_QNAME"
	TLS_SNI DomainNameProtocolAnyOf = "TLS_SNI"
	TLS_SAN DomainNameProtocolAnyOf = "TLS_SAN"
	TSL_SCN DomainNameProtocolAnyOf = "TSL_SCN"
)
