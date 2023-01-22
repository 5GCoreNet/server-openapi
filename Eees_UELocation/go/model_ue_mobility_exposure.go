/*
 * EES UE Location Information_API
 *
 * API for EES UE Location Information.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Eees_UELocation

import (
	"time"
)

// UeMobilityExposure - Represents a UE mobility information.
type UeMobilityExposure struct {

	// string with format \"date-time\" as defined in OpenAPI.
	Ts time.Time `json:"ts,omitempty"`

	RecurringTime ScheduledCommunicationTime `json:"recurringTime,omitempty"`

	// Unsigned integer identifying a period of time in units of seconds.
	Duration int32 `json:"duration"`

	// string with format 'float' as defined in OpenAPI.
	DurationVariance float32 `json:"durationVariance,omitempty"`

	LocInfo []UeLocationInfo `json:"locInfo"`
}
