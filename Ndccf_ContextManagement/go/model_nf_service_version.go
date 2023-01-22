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

// NfServiceVersion - Contains the version details of an NF service
type NfServiceVersion struct {

	ApiVersionInUri string `json:"apiVersionInUri"`

	ApiFullVersion string `json:"apiFullVersion"`

	// string with format 'date-time' as defined in OpenAPI.
	Expiry time.Time `json:"expiry,omitempty"`
}
