/*
 * Nnwdaf_DataManagement
 *
 * Nnwdaf_DataManagement API Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Nnwdaf_DataManagement

// NetworkAssistanceSession - A representation of a Network Assistance Session resource.
type NetworkAssistanceSession struct {

	// String chosen by the 5GMS AF to serve as an identifier in a resource URI.
	NaSessionId string `json:"naSessionId"`

	ServiceDataFlowDescription []ServiceDataFlowDescription `json:"serviceDataFlowDescription,omitempty"`

	// String chosen by the 5GMS AF to serve as an identifier in a resource URI.
	PolicyTemplateId string `json:"policyTemplateId,omitempty"`

	RequestedQoS M5QoSSpecification `json:"requestedQoS,omitempty"`

	RecommendedQoS M5QoSSpecification `json:"recommendedQoS,omitempty"`

	// Uniform Resource Locator, comforming with the URI Generic Syntax specified in IETF RFC 3986.
	NotficationURL string `json:"notficationURL,omitempty"`
}
