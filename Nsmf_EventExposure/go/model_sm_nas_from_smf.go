/*
 * Nsmf_EventExposure
 *
 * Session Management Event Exposure Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Nsmf_EventExposure

import (
	"time"
)

// SmNasFromSmf - Represents information on the SM congestion control applied SM NAS messages that SMF sends to UE for PDU Session. 
type SmNasFromSmf struct {

	SmNasType string `json:"smNasType"`

	// string with format 'date-time' as defined in OpenAPI.
	TimeStamp time.Time `json:"timeStamp"`

	// indicating a time in seconds.
	BackoffTimer int32 `json:"backoffTimer"`

	AppliedSmccType AppliedSmccType `json:"appliedSmccType"`
}