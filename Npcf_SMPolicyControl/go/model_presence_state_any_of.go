/*
 * Npcf_SMPolicyControl API
 *
 * Session Management Policy Control Service   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Npcf_SMPolicyControl

type PresenceStateAnyOf string

// List of PresenceStateAnyOf
const (
	IN_AREA PresenceStateAnyOf = "IN_AREA"
	OUT_OF_AREA PresenceStateAnyOf = "OUT_OF_AREA"
	UNKNOWN PresenceStateAnyOf = "UNKNOWN"
	INACTIVE PresenceStateAnyOf = "INACTIVE"
)
