/*
 * Ndccf_DataManagement
 *
 * DCCF Data Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Ndccf_DataManagement

// V2xCapability - Indicate the supported V2X Capability by the PCF.
type V2xCapability struct {

	LteV2x bool `json:"lteV2x,omitempty"`

	NrV2x bool `json:"nrV2x,omitempty"`
}
