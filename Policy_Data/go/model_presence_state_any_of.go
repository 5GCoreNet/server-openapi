/*
 * Unified Data Repository Service API file for policy data
 *
 * The API version is defined in 3GPP TS 29.504   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: -
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Policy_Data

type PresenceStateAnyOf string

// List of PresenceStateAnyOf
const (
	IN_AREA PresenceStateAnyOf = "IN_AREA"
	OUT_OF_AREA PresenceStateAnyOf = "OUT_OF_AREA"
	UNKNOWN PresenceStateAnyOf = "UNKNOWN"
	INACTIVE PresenceStateAnyOf = "INACTIVE"
)