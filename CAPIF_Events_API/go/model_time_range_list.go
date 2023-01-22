/*
 * CAPIF_Events_API
 *
 * API for event subscription management.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package CAPIF_Events_API

import (
	"time"
)

// TimeRangeList - Represents the time range during which the invocation of a service API is allowed by the API invoker. 
type TimeRangeList struct {

	// string with format \"date-time\" as defined in OpenAPI.
	StartTime time.Time `json:"startTime,omitempty"`

	// string with format \"date-time\" as defined in OpenAPI.
	StopTime time.Time `json:"stopTime,omitempty"`
}