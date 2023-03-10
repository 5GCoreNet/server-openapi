/*
 * CAPIF_Events_API
 *
 * API for event subscription management.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_CAPIF_Events_API

// CapifEventDetail - Represents a CAPIF event details.
type CapifEventDetail struct {

	// Description of the service API as published by the APF.
	ServiceAPIDescriptions []ServiceApiDescription `json:"serviceAPIDescriptions,omitempty"`

	// Identifier of the service API
	ApiIds []string `json:"apiIds,omitempty"`

	// Identity of the API invoker
	ApiInvokerIds []string `json:"apiInvokerIds,omitempty"`

	AccCtrlPolList AccessControlPolicyListExt `json:"accCtrlPolList,omitempty"`

	// Invocation logs.
	InvocationLogs []InvocationLog `json:"invocationLogs,omitempty"`

	ApiTopoHide TopologyHiding `json:"apiTopoHide,omitempty"`
}
