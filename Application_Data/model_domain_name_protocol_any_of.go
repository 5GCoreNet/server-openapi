/*
 * Unified Data Repository Service API file for Application Data
 *
 * The API version is defined in 3GPP TS 29.504   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: -
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Application_Data

type DomainNameProtocolAnyOf string

// List of DomainNameProtocolAnyOf
const (
	DNS_QNAME DomainNameProtocolAnyOf = "DNS_QNAME"
	TLS_SNI DomainNameProtocolAnyOf = "TLS_SNI"
	TLS_SAN DomainNameProtocolAnyOf = "TLS_SAN"
	TSL_SCN DomainNameProtocolAnyOf = "TSL_SCN"
)
