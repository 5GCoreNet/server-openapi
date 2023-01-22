/*
 * Unified Data Repository Service API file for subscription data
 *
 * Unified Data Repository Service (subscription data).   The API version is defined in 3GPP TS 29.504.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: -
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Subscription_Data

type TraceDepthAnyOf string

// List of TraceDepthAnyOf
const (
	MINIMUM TraceDepthAnyOf = "MINIMUM"
	MEDIUM TraceDepthAnyOf = "MEDIUM"
	MAXIMUM TraceDepthAnyOf = "MAXIMUM"
	MINIMUM_WO_VENDOR_EXTENSION TraceDepthAnyOf = "MINIMUM_WO_VENDOR_EXTENSION"
	MEDIUM_WO_VENDOR_EXTENSION TraceDepthAnyOf = "MEDIUM_WO_VENDOR_EXTENSION"
	MAXIMUM_WO_VENDOR_EXTENSION TraceDepthAnyOf = "MAXIMUM_WO_VENDOR_EXTENSION"
)
