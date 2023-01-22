/*
 * 3gpp-mbs-session
 *
 * API for MBS Session Management.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package MBSSession

type PatchOperationAnyOf string

// List of PatchOperationAnyOf
const (
	ADD PatchOperationAnyOf = "add"
	COPY PatchOperationAnyOf = "copy"
	MOVE PatchOperationAnyOf = "move"
	REMOVE PatchOperationAnyOf = "remove"
	REPLACE PatchOperationAnyOf = "replace"
	TEST PatchOperationAnyOf = "test"
)
