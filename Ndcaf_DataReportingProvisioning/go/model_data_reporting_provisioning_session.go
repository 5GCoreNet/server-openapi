/*
 * Ndcaf_DataReportingProvisioning
 *
 * Data Collection AF: Provisioning Sessions API © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 
 *
 * API version: 1.1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Ndcaf_DataReportingProvisioning

// DataReportingProvisioningSession - A representation of a Data Reporting Provisioning Session.
type DataReportingProvisioningSession struct {

	// String chosen by the 5GMS AF to serve as an identifier in a resource URI.
	ProvisioningSessionId string `json:"provisioningSessionId"`

	// Contains an identity of an application service provider.
	AspId string `json:"aspId"`

	// String providing an application identifier.
	ExternalApplicationId string `json:"externalApplicationId"`

	// String providing an application identifier.
	InternalApplicationId string `json:"internalApplicationId,omitempty"`

	EventId AfEvent `json:"eventId"`

	DataReportingConfigurationIds []string `json:"dataReportingConfigurationIds"`
}