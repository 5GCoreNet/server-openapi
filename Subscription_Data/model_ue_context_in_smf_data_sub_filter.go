/*
 * Unified Data Repository Service API file for subscription data
 *
 * Unified Data Repository Service (subscription data).   The API version is defined in 3GPP TS 29.504.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: -
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Subscription_Data

// UeContextInSmfDataSubFilter - UE Context In Smf Data Subscription Filter.
type UeContextInSmfDataSubFilter struct {

	DnnList []string `json:"dnnList,omitempty"`

	SnssaiList []Snssai `json:"snssaiList,omitempty"`

	EmergencyInd bool `json:"emergencyInd,omitempty"`
}
