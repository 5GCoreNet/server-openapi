/*
 * Nhss_EE
 *
 * HSS Event Exposure.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Nhss_EE

import (
	"time"
)

// ReachabilityForDataReport - Contains data for a Monitoring Event Report, specific to the 'Reachability For Data' event type
type ReachabilityForDataReport struct {

	ReachabilityDataStatus bool `json:"reachabilityDataStatus"`

	// string with format 'date-time' as defined in OpenAPI.
	MaxAvailabilityTime time.Time `json:"maxAvailabilityTime,omitempty"`
}
