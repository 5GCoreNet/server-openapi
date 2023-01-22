/*
 * EES EAS Registration_API
 *
 * API for EAS Registration.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Eees_EASRegistration

import (
	"time"
)

// EasRegistrationPatch - Represents partial update request of individual EAS registration information.
type EasRegistrationPatch struct {

	EasProf EasProfile `json:"easProf,omitempty"`

	// string with format 'date-time' as defined in OpenAPI with 'nullable:true' property.  
	ExpTime *time.Time `json:"expTime,omitempty"`
}
