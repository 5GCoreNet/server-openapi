/*
 * CAPIF_Publish_Service_API
 *
 * API for publishing service APIs.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package CAPIF_Publish_Service_API

type SecurityMethodAnyOf string

// List of SecurityMethodAnyOf
const (
	PSK SecurityMethodAnyOf = "PSK"
	PKI SecurityMethodAnyOf = "PKI"
	OAUTH SecurityMethodAnyOf = "OAUTH"
)
