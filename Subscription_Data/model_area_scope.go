/*
 * Unified Data Repository Service API file for subscription data
 *
 * Unified Data Repository Service (subscription data).   The API version is defined in 3GPP TS 29.504.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: -
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Subscription_Data

// AreaScope - Contain the area based on Cells or Tracking Areas.
type AreaScope struct {

	EutraCellIdList []string `json:"eutraCellIdList,omitempty"`

	NrCellIdList []string `json:"nrCellIdList,omitempty"`

	TacList []string `json:"tacList,omitempty"`

	// A map (list of key-value pairs) where PlmnId converted to a string serves as key 
	TacInfoPerPlmn map[string]TacInfo `json:"tacInfoPerPlmn,omitempty"`
}
