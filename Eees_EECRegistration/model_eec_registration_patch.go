/*
 * Eees_EECRegistration
 *
 * API for EEC registration. © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Eees_EECRegistration

import (
	"time"
)

// EecRegistrationPatch - Describes the parameters to perform EEC Registration update.
type EecRegistrationPatch struct {

	// Profiles of ACs for which the EEC provides edge enabling services.
	AcProfs []AcProfile `json:"acProfs,omitempty"`

	// string with format \"date-time\" as defined in OpenAPI.
	ExpTime time.Time `json:"expTime,omitempty"`

	UnfulfilledAcProfs UnfulfilledAcProfile `json:"unfulfilledAcProfs,omitempty"`
}
