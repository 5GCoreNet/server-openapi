/*
 * Nadrf_DataManagement
 *
 * ADRF Data Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Nadrf_DataManagement

// NfGroupCond - Subscription to a set of NFs based on their Group Id
type NfGroupCond struct {

	NfType string `json:"nfType"`

	// Identifier of a group of NFs.
	NfGroupId string `json:"nfGroupId"`
}
