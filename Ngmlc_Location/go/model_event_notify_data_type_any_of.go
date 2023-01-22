/*
 * Ngmlc_Location
 *
 * GMLC Location Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.1.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Ngmlc_Location

type EventNotifyDataTypeAnyOf string

// List of EventNotifyDataTypeAnyOf
const (
	UE_AVAILABLE EventNotifyDataTypeAnyOf = "UE_AVAILABLE"
	PERIODIC EventNotifyDataTypeAnyOf = "PERIODIC"
	ENTERING_INTO_AREA EventNotifyDataTypeAnyOf = "ENTERING_INTO_AREA"
	LEAVING_FROM_AREA EventNotifyDataTypeAnyOf = "LEAVING_FROM_AREA"
	BEING_INSIDE_AREA EventNotifyDataTypeAnyOf = "BEING_INSIDE_AREA"
	MOTION EventNotifyDataTypeAnyOf = "MOTION"
	MAXIMUM_INTERVAL_EXPIRATION_EVENT EventNotifyDataTypeAnyOf = "MAXIMUM_INTERVAL_EXPIRATION_EVENT"
	LOCATION_CANCELLATION_EVENT EventNotifyDataTypeAnyOf = "LOCATION_CANCELLATION_EVENT"
	ACTIVATION_OF_DEFERRED_LOCATION EventNotifyDataTypeAnyOf = "ACTIVATION_OF_DEFERRED_LOCATION"
	UE_MOBILITY_FOR_DEFERRED_LOCATION EventNotifyDataTypeAnyOf = "UE_MOBILITY_FOR_DEFERRED_LOCATION"
	_5_GC_MT_LR EventNotifyDataTypeAnyOf = "5GC_MT_LR"
)
