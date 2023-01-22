/*
 * CAPIF_Routing_Info_API
 *
 * API for Routing information.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package CAPIF_Routing_Info_API

// Resource - Represents the API resource data.
type Resource struct {

	// Resource name
	ResourceName string `json:"resourceName"`

	CommType CommunicationType `json:"commType"`

	// Relative URI of the API resource, it is set as {apiSpecificSuffixes} part of the URI structure as defined in clause 5.2.4 of 3GPP TS 29.122. 
	Uri string `json:"uri"`

	// it is set as {custOpName} part of the URI structure for a custom operation associated with a resource as defined in clause 5.2.4 of 3GPP TS 29.122. 
	CustOpName string `json:"custOpName,omitempty"`

	// Supported HTTP methods for the API resource. Only applicable when the protocol in AefProfile indicates HTTP. 
	Operations []Operation `json:"operations,omitempty"`

	// Text description of the API resource
	Description string `json:"description,omitempty"`
}
