/*
 * N32 Handshake API
 *
 * N32-c Handshake Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.3.0-alpha.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_N32_Handshake

type N32fErrorTypeAnyOf string

// List of N32fErrorTypeAnyOf
const (
	INTEGRITY_CHECK_FAILED N32fErrorTypeAnyOf = "INTEGRITY_CHECK_FAILED"
	INTEGRITY_CHECK_ON_MODIFICATIONS_FAILED N32fErrorTypeAnyOf = "INTEGRITY_CHECK_ON_MODIFICATIONS_FAILED"
	MODIFICATIONS_INSTRUCTIONS_FAILED N32fErrorTypeAnyOf = "MODIFICATIONS_INSTRUCTIONS_FAILED"
	DECIPHERING_FAILED N32fErrorTypeAnyOf = "DECIPHERING_FAILED"
	MESSAGE_RECONSTRUCTION_FAILED N32fErrorTypeAnyOf = "MESSAGE_RECONSTRUCTION_FAILED"
	CONTEXT_NOT_FOUND N32fErrorTypeAnyOf = "CONTEXT_NOT_FOUND"
	INTEGRITY_KEY_EXPIRED N32fErrorTypeAnyOf = "INTEGRITY_KEY_EXPIRED"
	ENCRYPTION_KEY_EXPIRED N32fErrorTypeAnyOf = "ENCRYPTION_KEY_EXPIRED"
	POLICY_MISMATCH N32fErrorTypeAnyOf = "POLICY_MISMATCH"
)