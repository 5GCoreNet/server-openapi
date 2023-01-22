/*
 * 3gpp-analyticsexposure
 *
 * API for Analytics Exposure.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.1.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package AnalyticsExposure

import (
	"time"
)

// CircumstanceDescription - Contains the description of a circumstance.
type CircumstanceDescription struct {

	// string with format 'float' as defined in OpenAPI.
	Freq float32 `json:"freq,omitempty"`

	// string with format 'date-time' as defined in OpenAPI.
	Tm time.Time `json:"tm,omitempty"`

	LocArea NetworkAreaInfo `json:"locArea,omitempty"`

	// Unsigned integer identifying a volume in units of bytes.
	Vol int64 `json:"vol,omitempty"`
}
