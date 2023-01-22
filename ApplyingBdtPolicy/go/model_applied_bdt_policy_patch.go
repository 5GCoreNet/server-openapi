/*
 * 3gpp-applying-bdt-policy
 *
 * API for applying BDT policy   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.1.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package ApplyingBdtPolicy

// AppliedBdtPolicyPatch - Represents the parameters to request the modification of a subscription to applied BDT policy. 
type AppliedBdtPolicyPatch struct {

	// string identifying a BDT Reference ID as defined in clause 5.3.3 of 3GPP TS 29.154.
	BdtRefId string `json:"bdtRefId"`
}
