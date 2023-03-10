/*
 * CAPIF_Security_API
 *
 * API for CAPIF security management.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_CAPIF_Security_API

type SecurityMethodAnyOf string

// List of SecurityMethodAnyOf
const (
	PSK SecurityMethodAnyOf = "PSK"
	PKI SecurityMethodAnyOf = "PKI"
	OAUTH SecurityMethodAnyOf = "OAUTH"
)
