/*
 * EES EEC Context Relocation API
 *
 * API for EEC Context Relocation.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Eees_EECContextRelocation

// SessionContexts - Represents the list of service session contexts information.
type SessionContexts struct {

	// List of service session contexts information.
	SessCntxs []IndividualSessionContext `json:"sessCntxs"`
}
