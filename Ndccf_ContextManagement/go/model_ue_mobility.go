/*
 * Ndccf_ContextManagement
 *
 * DCCF Context Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Ndccf_ContextManagement

import (
	"time"
)

// UeMobility - Represents UE mobility information.
type UeMobility struct {

	// string with format 'date-time' as defined in OpenAPI.
	Ts time.Time `json:"ts,omitempty"`

	RecurringTime ScheduledCommunicationTime `json:"recurringTime,omitempty"`

	// indicating a time in seconds.
	Duration int32 `json:"duration,omitempty"`

	// string with format 'float' as defined in OpenAPI.
	DurationVariance float32 `json:"durationVariance,omitempty"`

	LocInfos []LocationInfo `json:"locInfos,omitempty"`
}
