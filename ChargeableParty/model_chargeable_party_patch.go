/*
 * 3gpp-chargeable-party
 *
 * API for Chargeable Party management.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_ChargeableParty

// ChargeablePartyPatch - Represents a modification request of a chargeable party resource.
type ChargeablePartyPatch struct {

	// Describes the IP flows.
	FlowInfo []FlowInfo `json:"flowInfo,omitempty"`

	// Identifies the external Application Identifier.
	ExterAppId string `json:"exterAppId,omitempty"`

	// Identifies Ethernet packet flows.
	EthFlowInfo []EthFlowDescription `json:"ethFlowInfo,omitempty"`

	// Indicates whether the sponsoring data connectivity is enabled (true) or not (false). 
	SponsoringEnabled bool `json:"sponsoringEnabled,omitempty"`

	// string identifying a BDT Reference ID as defined in clause 5.3.3 of 3GPP TS 29.154.
	ReferenceId string `json:"referenceId,omitempty"`

	UsageThreshold *UsageThresholdRm `json:"usageThreshold,omitempty"`

	// string formatted according to IETF RFC 3986 identifying a referenced resource.
	NotificationDestination string `json:"notificationDestination,omitempty"`

	// Represents the list of event(s) to which the SCS/AS requests to subscribe to.
	Events []Event `json:"events,omitempty"`
}
