/*
 * CAPIF_Discover_Service_API
 *
 * API for discovering service APIs.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_CAPIF_Discover_Service_API

// CustomOperation - Represents the description of a custom operation.
type CustomOperation struct {

	CommType CommunicationType `json:"commType"`

	// it is set as {custOpName} part of the URI structure for a custom operation without resource association as defined in clause 5.2.4 of 3GPP TS 29.122. 
	CustOpName string `json:"custOpName"`

	// Supported HTTP methods for the API resource. Only applicable when the protocol in AefProfile indicates HTTP. 
	Operations []Operation `json:"operations,omitempty"`

	// Text description of the custom operation
	Description string `json:"description,omitempty"`
}
