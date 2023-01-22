/*
 * Nudm_EE
 *
 * Nudm Event Exposure Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Nudm_EE

import (
	"time"
)

type ReachabilityForSmsReport struct {

	SmsfAccessType AccessType `json:"smsfAccessType"`

	// string with format 'date-time' as defined in OpenAPI.
	MaxAvailabilityTime time.Time `json:"maxAvailabilityTime,omitempty"`
}
