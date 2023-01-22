/*
 * Nnwdaf_DataManagement
 *
 * Nnwdaf_DataManagement API Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Nnwdaf_DataManagement

import (
	"time"
)

// RedundantTransmissionExpPerTs - The redundant transmission experience per Time Slot.
type RedundantTransmissionExpPerTs struct {

	// string with format 'date-time' as defined in OpenAPI.
	TsStart time.Time `json:"tsStart"`

	// indicating a time in seconds.
	TsDuration int32 `json:"tsDuration"`

	ObsvRedTransExp ObservedRedundantTransExp `json:"obsvRedTransExp"`

	// Redundant Transmission Status. Set to \"true\" if redundant transmission was activated, otherwise set to \"false\". Default value is \"false\" if omitted. 
	RedTransStatus bool `json:"redTransStatus,omitempty"`

	// Unsigned integer indicating Sampling Ratio (see clauses 4.15.1 of 3GPP TS 23.502), expressed in percent.  
	UeRatio int32 `json:"ueRatio,omitempty"`

	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible.
	Confidence int32 `json:"confidence,omitempty"`
}
