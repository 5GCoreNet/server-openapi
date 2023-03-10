/*
 * M1_ContentHostingProvisioning
 *
 * 5GMS AF M1 Content Hosting Provisioning API © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 
 *
 * API version: 2.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_M1_ContentHostingProvisioning

// IngestConfiguration - A configuration for content ingest.
type IngestConfiguration struct {

	Path string `json:"path,omitempty"`

	Pull bool `json:"pull,omitempty"`

	// String providing an URI formatted according to RFC 3986.
	Protocol string `json:"protocol,omitempty"`

	// Uniform Resource Locator, comforming with the URI Generic Syntax specified in IETF RFC 3986.
	EntryPoint string `json:"entryPoint,omitempty"`
}
