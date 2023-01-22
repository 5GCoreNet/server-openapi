/*
 * NSSF NSSAI Availability
 *
 * NSSF NSSAI Availability Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.3.0-alpha.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Nnssf_NSSAIAvailability

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
