/*
 * Nchf_SpendingLimitControl
 *
 * Nchf Spending Limit Control Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Nchf_SpendingLimitControl

// PolicyCounterInfo - Represents the data structure presenting the policy counter status.
type PolicyCounterInfo struct {

	// Identifies a policy counter.
	PolicyCounterId string `json:"policyCounterId"`

	// Identifies the policy counter status applicable for a specific policy counter identified by the policyCounterId. The values (e.g. valid, invalid or any other status) are not specified. The interpretation and actions related to the defined values are out of scope of 3GPP. 
	CurrentStatus string `json:"currentStatus"`

	// Provides the pending policy counter status.
	PenPolCounterStatuses []PendingPolicyCounterStatus `json:"penPolCounterStatuses,omitempty"`
}