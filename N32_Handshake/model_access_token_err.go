/*
 * N32 Handshake API
 *
 * N32-c Handshake Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.3.0-alpha.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_N32_Handshake

// AccessTokenErr - Error returned in the access token response message
type AccessTokenErr struct {

	Error string `json:"error"`

	ErrorDescription string `json:"error_description,omitempty"`

	ErrorUri string `json:"error_uri,omitempty"`
}
