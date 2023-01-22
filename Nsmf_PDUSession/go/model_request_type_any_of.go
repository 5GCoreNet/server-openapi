/*
 * Nsmf_PDUSession
 *
 * SMF PDU Session Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.3.0-alpha.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Nsmf_PDUSession

type RequestTypeAnyOf string

// List of RequestTypeAnyOf
const (
	INITIAL_REQUEST RequestTypeAnyOf = "INITIAL_REQUEST"
	EXISTING_PDU_SESSION RequestTypeAnyOf = "EXISTING_PDU_SESSION"
	INITIAL_EMERGENCY_REQUEST RequestTypeAnyOf = "INITIAL_EMERGENCY_REQUEST"
	EXISTING_EMERGENCY_PDU_SESSION RequestTypeAnyOf = "EXISTING_EMERGENCY_PDU_SESSION"
)
