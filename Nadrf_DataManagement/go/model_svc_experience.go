/*
 * Nadrf_DataManagement
 *
 * ADRF Data Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Nadrf_DataManagement

// SvcExperience - Contains a mean opinion score with the customized range.
type SvcExperience struct {

	// string with format 'float' as defined in OpenAPI.
	Mos float32 `json:"mos,omitempty"`

	// string with format 'float' as defined in OpenAPI.
	UpperRange float32 `json:"upperRange,omitempty"`

	// string with format 'float' as defined in OpenAPI.
	LowerRange float32 `json:"lowerRange,omitempty"`
}
