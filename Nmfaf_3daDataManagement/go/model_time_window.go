/*
 * Nmfaf_3daDataManagement
 *
 * MFAF 3GPP DCCF Adaptor (3DA) Data Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Nmfaf_3daDataManagement

import (
	"time"
)

// TimeWindow - Represents a time window identified by a start time and a stop time.
type TimeWindow struct {

	// string with format \"date-time\" as defined in OpenAPI.
	StartTime time.Time `json:"startTime"`

	// string with format \"date-time\" as defined in OpenAPI.
	StopTime time.Time `json:"stopTime"`
}