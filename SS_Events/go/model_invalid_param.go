/*
 * SS_Events
 *
 * API for SEAL Events management.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.1.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package SS_Events

// InvalidParam - Represents the description of invalid parameters, for a request rejected due to invalid parameters.
type InvalidParam struct {

	// Attribute's name encoded as a JSON Pointer, or header's name.
	Param string `json:"param"`

	// A human-readable reason, e.g. \"must be a positive integer\".
	Reason string `json:"reason,omitempty"`
}
