/*
 * Nhss_imsUECM
 *
 * Nhss UE Context Management Service for IMS.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Nhss_imsUECM

// ScscfCapabilityList - Information about mandatory and optional S-CSCF capabilities
type ScscfCapabilityList struct {

	// A list of capabilities of the S-CSCF
	MandatoryCapabilityList []int32 `json:"mandatoryCapabilityList,omitempty"`

	// A list of capabilities of the S-CSCF
	OptionalCapabilityList []int32 `json:"optionalCapabilityList,omitempty"`
}
