/*
 * Nnwdaf_DataManagement
 *
 * Nnwdaf_DataManagement API Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Nnwdaf_DataManagement

// TargetArea - TA list or TAI range list or any TA
type TargetArea struct {

	TaList []Tai `json:"taList,omitempty"`

	TaiRangeList []TaiRange `json:"taiRangeList,omitempty"`

	AnyTa bool `json:"anyTa,omitempty"`
}
