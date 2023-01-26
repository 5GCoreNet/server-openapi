/*
 * M1_ContentHostingProvisioning
 *
 * 5GMS AF M1 Content Hosting Provisioning API © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 
 *
 * API version: 2.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_M1_ContentHostingProvisioning

// PathRewriteRule - A rule to manipulate URL paths.
type PathRewriteRule struct {

	RequestPattern string `json:"requestPattern"`

	MappedPath string `json:"mappedPath"`
}