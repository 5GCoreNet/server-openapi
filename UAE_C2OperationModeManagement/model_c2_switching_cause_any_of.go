/*
 * UAE Server C2 Operation Mode Management Service
 *
 * UAE Server C2 Operation Mode Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_UAE_C2OperationModeManagement

type C2SwitchingCauseAnyOf string

// List of C2SwitchingCauseAnyOf
const (
	DIRECT_LINK_QUALITY_DEGRADATION C2SwitchingCauseAnyOf = "DIRECT_LINK_QUALITY_DEGRADATION"
	DIRECT_LINK_AVAILABLE C2SwitchingCauseAnyOf = "DIRECT_LINK_AVAILABLE"
	MOVING_BVLOS C2SwitchingCauseAnyOf = "MOVING_BVLOS"
	LOCATION_CHANGE C2SwitchingCauseAnyOf = "LOCATION_CHANGE"
	TRAFFIC_CONTROL_NEEDED C2SwitchingCauseAnyOf = "TRAFFIC_CONTROL_NEEDED"
	SECURITY_REASONS C2SwitchingCauseAnyOf = "SECURITY_REASONS"
	OTHER_REASONS C2SwitchingCauseAnyOf = "OTHER_REASONS"
)
