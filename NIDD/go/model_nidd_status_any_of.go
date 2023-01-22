/*
 * 3gpp-nidd
 *
 * API for non IP data delivery.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package NIDD

type NiddStatusAnyOf string

// List of NiddStatusAnyOf
const (
	ACTIVE NiddStatusAnyOf = "ACTIVE"
	TERMINATED_UE_NOT_AUTHORIZED NiddStatusAnyOf = "TERMINATED_UE_NOT_AUTHORIZED"
	TERMINATED NiddStatusAnyOf = "TERMINATED"
	RDS_PORT_UNKNOWN NiddStatusAnyOf = "RDS_PORT_UNKNOWN"
)
