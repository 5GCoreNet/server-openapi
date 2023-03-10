/*
 * 3gpp-time-sync-exposure
 *
 * API for time synchronization exposure.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_TimeSyncExposure

type InstanceTypeAnyOf string

// List of InstanceTypeAnyOf
const (
	BOUNDARY_CLOCK InstanceTypeAnyOf = "BOUNDARY_CLOCK"
	E2_E_TRANS_CLOCK InstanceTypeAnyOf = "E2E_TRANS_CLOCK"
	P2_P_TRANS_CLOCK InstanceTypeAnyOf = "P2P_TRANS_CLOCK"
	P2_P_RELAY_INSTANCE InstanceTypeAnyOf = "P2P_RELAY_INSTANCE"
)
