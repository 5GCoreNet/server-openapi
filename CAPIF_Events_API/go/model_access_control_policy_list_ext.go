/*
 * CAPIF_Events_API
 *
 * API for event subscription management.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package CAPIF_Events_API

// AccessControlPolicyListExt - Represents the extension for access control policies.
type AccessControlPolicyListExt struct {

	// Policy of each API invoker.
	ApiInvokerPolicies []ApiInvokerPolicy `json:"apiInvokerPolicies,omitempty"`

	ApiId string `json:"apiId"`
}
