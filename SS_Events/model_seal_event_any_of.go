/*
 * SS_Events
 *
 * API for SEAL Events management.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.1.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_SS_Events

type SealEventAnyOf string

// List of SealEventAnyOf
const (
	LM_LOCATION_INFO_CHANGE SealEventAnyOf = "LM_LOCATION_INFO_CHANGE"
	GM_GROUP_INFO_CHANGE SealEventAnyOf = "GM_GROUP_INFO_CHANGE"
	CM_USER_PROFILE_CHANGE SealEventAnyOf = "CM_USER_PROFILE_CHANGE"
	GM_GROUP_CREATE SealEventAnyOf = "GM_GROUP_CREATE"
	NRM_MONITOR_UE_USER_EVENTS SealEventAnyOf = "NRM_MONITOR_UE_USER_EVENTS"
	LM_LOCATION_DEVIATION_MONITOR SealEventAnyOf = "LM_LOCATION_DEVIATION_MONITOR"
	GM_TEMP_GROUP_FORMATION SealEventAnyOf = "GM_TEMP_GROUP_FORMATION"
	LM_LOCATION_AREA_MONITOR SealEventAnyOf = "LM_LOCATION_AREA_MONITOR"
)
