/*
 * Ndcaf_DataReportingProvisioning
 *
 * Data Collection AF: Provisioning Sessions API © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 
 *
 * API version: 1.1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Ndcaf_DataReportingProvisioning

// DataReportingConfigurationPatch - A JSON patch for a Data Reporting Configuration.
type DataReportingConfigurationPatch struct {

	// Uniform Resource Locator, comforming with the URI Generic Syntax specified in IETF RFC 3986.
	AuthorizationURL string `json:"authorizationURL,omitempty"`

	DataAccessProfiles []DataAccessProfile `json:"dataAccessProfiles,omitempty"`
}
