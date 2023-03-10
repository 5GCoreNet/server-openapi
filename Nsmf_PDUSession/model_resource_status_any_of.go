/*
 * Nsmf_PDUSession
 *
 * SMF PDU Session Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.3.0-alpha.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Nsmf_PDUSession

type ResourceStatusAnyOf string

// List of ResourceStatusAnyOf
const (
	RELEASED ResourceStatusAnyOf = "RELEASED"
	UNCHANGED ResourceStatusAnyOf = "UNCHANGED"
	TRANSFERRED ResourceStatusAnyOf = "TRANSFERRED"
	UPDATED ResourceStatusAnyOf = "UPDATED"
	ALT_ANCHOR_SMF ResourceStatusAnyOf = "ALT_ANCHOR_SMF"
)
