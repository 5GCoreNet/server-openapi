/*
 * Namf_MT
 *
 * AMF Mobile Terminated Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Namf_MT

// AccessTokenErr - Error returned in the access token response message
type AccessTokenErr struct {

	Error string `json:"error"`

	ErrorDescription string `json:"error_description,omitempty"`

	ErrorUri string `json:"error_uri,omitempty"`
}
