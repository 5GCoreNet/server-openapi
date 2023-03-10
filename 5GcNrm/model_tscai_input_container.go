/*
 * 3GPP 5GC NRM
 *
 * OAS 3.0.1 specification of the 5GC NRM © 2020, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.
 *
 * API version: 17.7.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_5GcNrm

import (
	"time"
)

// TscaiInputContainer - Indicates TSC Traffic pattern.
type TscaiInputContainer struct {

	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible.
	Periodicity int32 `json:"periodicity,omitempty"`

	// string with format 'date-time' as defined in OpenAPI.
	BurstArrivalTime time.Time `json:"burstArrivalTime,omitempty"`

	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible.
	SurTimeInNumMsg int32 `json:"surTimeInNumMsg,omitempty"`

	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible.
	SurTimeInTime int32 `json:"surTimeInTime,omitempty"`
}
