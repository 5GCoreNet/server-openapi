/*
 * Ndccf_ContextManagement
 *
 * DCCF Context Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Ndccf_ContextManagement

// DynamicPolicy - A representation of a Dynamic Policy resource.
type DynamicPolicy struct {

	// String chosen by the 5GMS AF to serve as an identifier in a resource URI.
	DynamicPolicyId string `json:"dynamicPolicyId"`

	// String chosen by the 5GMS AF to serve as an identifier in a resource URI.
	PolicyTemplateId string `json:"policyTemplateId"`

	ServiceDataFlowDescriptions []ServiceDataFlowDescription `json:"serviceDataFlowDescriptions"`

	// String chosen by the 5GMS AF to serve as an identifier in a resource URI.
	ProvisioningSessionId string `json:"provisioningSessionId"`

	QosSpecification M5QoSSpecification `json:"qosSpecification,omitempty"`

	EnforcementMethod string `json:"enforcementMethod,omitempty"`

	EnforcementBitRate int32 `json:"enforcementBitRate,omitempty"`
}
