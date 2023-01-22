/*
 * Nadrf_DataManagement
 *
 * ADRF Data Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Nadrf_DataManagement

// LmfInfo - Information of an LMF NF Instance
type LmfInfo struct {

	ServingClientTypes []ExternalClientType `json:"servingClientTypes,omitempty"`

	// LMF identification.
	LmfId string `json:"lmfId,omitempty"`

	ServingAccessTypes []AccessType `json:"servingAccessTypes,omitempty"`

	ServingAnNodeTypes []AnNodeType `json:"servingAnNodeTypes,omitempty"`

	ServingRatTypes []RatType `json:"servingRatTypes,omitempty"`

	TaiList []Tai `json:"taiList,omitempty"`

	TaiRangeList []TaiRange `json:"taiRangeList,omitempty"`

	SupportedGADShapes []SupportedGadShapes `json:"supportedGADShapes,omitempty"`
}
