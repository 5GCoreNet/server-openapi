/*
 * CAPIF_Events_API
 *
 * API for event subscription management.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_CAPIF_Events_API

// AefProfile - Represents the AEF profile data.
type AefProfile struct {

	// Identifier of the API exposing function
	AefId string `json:"aefId"`

	// API version
	Versions []Version `json:"versions"`

	Protocol Protocol `json:"protocol,omitempty"`

	DataFormat DataFormat `json:"dataFormat,omitempty"`

	// Security methods supported by the AEF
	SecurityMethods []SecurityMethod `json:"securityMethods,omitempty"`

	// Domain to which API belongs to
	DomainName string `json:"domainName,omitempty"`

	// Interface details
	InterfaceDescriptions []InterfaceDescription `json:"interfaceDescriptions,omitempty"`

	AefLocation AefLocation `json:"aefLocation,omitempty"`
}
