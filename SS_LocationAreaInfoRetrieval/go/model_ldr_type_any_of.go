/*
 * SS_LocationAreaInfoRetrieval
 *
 * API for SEAL Location Area Info Retrieval.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package SS_LocationAreaInfoRetrieval

type LdrTypeAnyOf string

// List of LdrTypeAnyOf
const (
	UE_AVAILABLE LdrTypeAnyOf = "UE_AVAILABLE"
	PERIODIC LdrTypeAnyOf = "PERIODIC"
	ENTERING_INTO_AREA LdrTypeAnyOf = "ENTERING_INTO_AREA"
	LEAVING_FROM_AREA LdrTypeAnyOf = "LEAVING_FROM_AREA"
	BEING_INSIDE_AREA LdrTypeAnyOf = "BEING_INSIDE_AREA"
	MOTION LdrTypeAnyOf = "MOTION"
)
