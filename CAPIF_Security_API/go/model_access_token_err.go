/*
 * CAPIF_Security_API
 *
 * API for CAPIF security management.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package CAPIF_Security_API

// AccessTokenErr - Represents an error in the access token request.
type AccessTokenErr struct {

	Error string `json:"error"`

	ErrorDescription string `json:"error_description,omitempty"`

	ErrorUri string `json:"error_uri,omitempty"`
}
