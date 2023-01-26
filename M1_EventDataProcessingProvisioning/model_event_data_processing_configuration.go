/*
 * M1_EventDataProcessingProvisioning
 *
 * 5GMS AF M1 Event Data Processing Provisioning API © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 
 *
 * API version: 2.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_M1_EventDataProcessingProvisioning

// EventDataProcessingConfiguration - A representation of an Event Data Processing Configuration resource.
type EventDataProcessingConfiguration struct {

	// String chosen by the 5GMS AF to serve as an identifier in a resource URI.
	EventDataProcessingConfigurationId string `json:"eventDataProcessingConfigurationId"`

	EventId AfEvent `json:"eventId"`

	// Uniform Resource Locator, comforming with the URI Generic Syntax specified in IETF RFC 3986.
	AuthorizationUrl string `json:"authorizationUrl,omitempty"`

	DataAccessProfiles []DataAccessProfile `json:"dataAccessProfiles"`
}