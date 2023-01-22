/*
 * CAPIF_Security_API
 *
 * API for CAPIF security management.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package CAPIF_Security_API

// AccessTokenClaims - Represents the claims data structure for the access token.
type AccessTokenClaims struct {

	Iss string `json:"iss"`

	Scope string `json:"scope"`

	// Unsigned integer identifying a period of time in units of seconds.
	Exp int32 `json:"exp"`
}
