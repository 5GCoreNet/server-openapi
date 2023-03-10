/*
 * MSGS_ASRegistration
 *
 * API for MSGS AS Registration Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_MSGS_ASRegistration

// AsRegistrationAck - AS registration response data
type AsRegistrationAck struct {

	AsSvcId string `json:"asSvcId"`

	Result ProblemDetails `json:"result"`
}
