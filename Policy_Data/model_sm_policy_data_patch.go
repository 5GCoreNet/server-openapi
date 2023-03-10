/*
 * Unified Data Repository Service API file for policy data
 *
 * The API version is defined in 3GPP TS 29.504   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: -
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Policy_Data

// SmPolicyDataPatch - Contains the SM policy data for a given subscriber.
type SmPolicyDataPatch struct {

	// Contains the remaining allowed usage data associated with the subscriber. The value of the limit identifier is used as the key of the map. 
	UmData *map[string]UsageMonData `json:"umData,omitempty"`

	// Modifiable Session Management Policy data per S-NSSAI for all the SNSSAIs of the subscriber. The key of the map is the S-NSSAI. 
	SmPolicySnssaiData map[string]SmPolicySnssaiDataPatch `json:"smPolicySnssaiData,omitempty"`
}
