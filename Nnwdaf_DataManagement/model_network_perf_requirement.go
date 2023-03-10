/*
 * Nnwdaf_DataManagement
 *
 * Nnwdaf_DataManagement API Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Nnwdaf_DataManagement

// NetworkPerfRequirement - Represents a network performance requirement.
type NetworkPerfRequirement struct {

	NwPerfType NetworkPerfType `json:"nwPerfType"`

	// Unsigned integer indicating Sampling Ratio (see clauses 4.15.1 of 3GPP TS 23.502), expressed in percent.  
	RelativeRatio int32 `json:"relativeRatio,omitempty"`

	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible.
	AbsoluteNum int32 `json:"absoluteNum,omitempty"`
}
