/*
 * 3gpp-racs-parameter-provisioning
 *
 * API for provisioning UE radio capability parameters.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_RacsParameterProvisioning

// RacsConfiguration - Represents a single UE radio capability configuration data.
type RacsConfiguration struct {

	// The UE radio capability ID provided by the SCS/AS to identify the UE radio capability data. See 3GPP TS 23.003 for the encoding.
	RacsId string `json:"racsId"`

	// The UE radio capability data in EPS.
	RacsParamEps string `json:"racsParamEps,omitempty"`

	// The UE radio capability data in 5GS.
	RacsParam5Gs string `json:"racsParam5Gs,omitempty"`

	// Related UE model's IMEI-TAC values.
	ImeiTacs []string `json:"imeiTacs"`
}
