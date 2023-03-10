/*
 * 3gpp-am-policyauthorization
 *
 * API for AM policy authorization.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_AMPolicyAuthorization

// AsTimeDistributionParam - Contains the 5G acess stratum time distribution parameters.
type AsTimeDistributionParam struct {

	AsTimeDistInd bool `json:"asTimeDistInd,omitempty"`

	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible with the OpenAPI 'nullable: true' property. 
	UuErrorBudget *int32 `json:"uuErrorBudget,omitempty"`
}
