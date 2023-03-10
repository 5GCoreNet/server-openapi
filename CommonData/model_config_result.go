/*
 * TS 29.122 Common Data Types
 *
 * Data types applicable to several APIs.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_CommonData

// ConfigResult - Represents one configuration processing result for a group's members.
type ConfigResult struct {

	// Each element indicates an external identifier of the UE.
	ExternalIds []string `json:"externalIds,omitempty"`

	// Each element identifies the MS internal PSTN/ISDN number allocated for the UE.
	Msisdns []string `json:"msisdns,omitempty"`

	ResultReason ResultReason `json:"resultReason"`
}
