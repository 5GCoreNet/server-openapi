/*
 * Unified Data Repository Service API file for Application Data
 *
 * The API version is defined in 3GPP TS 29.504   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: -
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Application_Data

// PfdContent - Represents the content of a PFD for an application identifier.
type PfdContent struct {

	// Identifies a PDF of an application identifier.
	PfdId string `json:"pfdId,omitempty"`

	// Represents a 3-tuple with protocol, server ip and server port for UL/DL application traffic. 
	FlowDescriptions []string `json:"flowDescriptions,omitempty"`

	// Indicates a URL or a regular expression which is used to match the significant parts of the URL. 
	Urls []string `json:"urls,omitempty"`

	// Indicates an FQDN or a regular expression as a domain name matching criteria.
	DomainNames []string `json:"domainNames,omitempty"`

	DnProtocol DomainNameProtocol `json:"dnProtocol,omitempty"`
}