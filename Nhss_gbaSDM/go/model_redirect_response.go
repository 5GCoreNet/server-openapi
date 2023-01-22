/*
 * Nhss_gbaSDM
 *
 * Nhss Subscriber Data Management Service for GBA.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Nhss_gbaSDM

// RedirectResponse - The response shall include a Location header field containing a different URI  (pointing to a different URI of an other service instance), or the same URI if a request  is redirected to the same target resource via a different SCP. 
type RedirectResponse struct {

	Cause string `json:"cause,omitempty"`

	// String providing an URI formatted according to RFC 3986.
	TargetScp string `json:"targetScp,omitempty"`

	// String providing an URI formatted according to RFC 3986.
	TargetSepp string `json:"targetSepp,omitempty"`
}
